#!/bin/bash

# Warren Clough Website Deployment Script
# This script deploys the website to Google Cloud Run and Firebase Hosting

set -e

echo "🚀 Starting deployment of Warren Clough website..."

# Set project variables
PROJECT_ID="${FIREBASE_PROJECT_ID:-my-website-a3970}"
SERVICE_NAME="warren-clough-website"
REGION="us-central1"
IMAGE_NAME="us-docker.pkg.dev/$PROJECT_ID/gcr.io/$SERVICE_NAME"

# Check if gcloud CLI is installed
if ! command -v gcloud &> /dev/null; then
    echo "❌ Google Cloud SDK is not installed. Please install it first:"
    echo "https://cloud.google.com/sdk/docs/install"
    exit 1
fi

# Check if Firebase CLI is installed
if ! command -v firebase &> /dev/null; then
    echo "❌ Firebase CLI is not installed. Please install it first:"
    echo "npm install -g firebase-tools"
    exit 1
fi

# Check if user is authenticated
if ! gcloud auth list --filter=status:ACTIVE --format="value(account)" | grep -q .; then
    echo "❌ Not authenticated with gcloud. Please run:"
    echo "gcloud auth login"
    exit 1
fi

# Set the project
echo "🔧 Setting up Google Cloud project..."
gcloud config set project $PROJECT_ID

# Enable required APIs
echo "⚙️ Enabling required APIs..."
gcloud services enable cloudbuild.googleapis.com
gcloud services enable run.googleapis.com
gcloud services enable containerregistry.googleapis.com

# Install Go dependencies
echo "📦 Installing Go dependencies..."
go mod tidy

# Build and push Docker image
echo "Building and pushing Docker image..."
gcloud builds submit --tag us-docker.pkg.dev/$PROJECT_ID/gcr.io/$SERVICE_NAME

# Deploy to Cloud Run
echo "🚀 Deploying to Cloud Run..."
gcloud run deploy $SERVICE_NAME \
    --image $IMAGE_NAME \
    --platform managed \
    --region $REGION \
    --allow-unauthenticated \
    --port 8080 \
    --memory=512Mi \
    --concurrency=1000 \
    --timeout=300 \
    --set-env-vars="GIN_MODE=release,FIREBASE_PROJECT_ID=$PROJECT_ID,SMTP_HOST=smtp.gmail.com,SMTP_PORT=587,ADMIN_EMAILS=clough.warren@gmail.com" \
    --set-secrets="SMTP_USERNAME=smtp-username:latest,SMTP_PASSWORD=smtp-password:latest,GOOGLE_APPLICATION_CREDENTIALS=firebase-service-account:latest"

# Deploy Firebase Hosting (this will connect to Cloud Run)
echo "🌐 Deploying Firebase Hosting..."
firebase deploy --only hosting

echo "✅ Deployment completed successfully!"
echo "🌐 Your website should be available at:"
echo "   - Cloud Run: $(gcloud run services describe $SERVICE_NAME --region=$REGION --format='value(status.url)')"
echo "   - Firebase Hosting: https://$PROJECT_ID.web.app"
