#!/bin/bash

# Setup script for storing SMTP credentials in Google Secret Manager
# Run this once before deployment

set -e

PROJECT_ID="${FIREBASE_PROJECT_ID:-my-website-a3970}"

echo "🔐 Setting up SMTP secrets for project: $PROJECT_ID"

# Check if gcloud CLI is installed and authenticated
if ! command -v gcloud &> /dev/null; then
    echo "❌ Google Cloud SDK is not installed."
    exit 1
fi

# Set the project
gcloud config set project $PROJECT_ID

# Enable Secret Manager API
echo "⚙️ Enabling Secret Manager API..."
gcloud services enable secretmanager.googleapis.com

# Check if secrets already exist, if not create them
echo "📧 Setting up SMTP secrets..."

# SMTP Username
if ! gcloud secrets describe smtp-username &>/dev/null; then
    echo "Creating smtp-username secret..."
    echo -n "clough.warren@gmail.com" | gcloud secrets create smtp-username --data-file=-
else
    echo "✅ smtp-username secret already exists"
fi

# SMTP Password (you'll need to update this with your actual app password)
if ! gcloud secrets describe smtp-password &>/dev/null; then
    echo "Creating smtp-password secret..."
    echo -n "hstu gwkx lnkn grzq" | gcloud secrets create smtp-password --data-file=-
    echo "⚠️  Remember to update the smtp-password secret with your actual Gmail app password!"
else
    echo "✅ smtp-password secret already exists"
fi

# Firebase Service Account (copy from your credentials file)
if ! gcloud secrets describe firebase-service-account &>/dev/null; then
    echo "Creating firebase-service-account secret..."
    if [ -f "./credentials/firebase-service-account.json" ]; then
        gcloud secrets create firebase-service-account --data-file=./credentials/firebase-service-account.json
    else
        echo "❌ Firebase service account file not found at ./credentials/firebase-service-account.json"
        echo "Please place your service account file there and run this script again."
        exit 1
    fi
else
    echo "✅ firebase-service-account secret already exists"
fi

echo ""
echo "✅ All secrets are set up!"
echo ""
echo "📝 Next steps:"
echo "1. Update your Gmail app password if needed:"
echo "   echo -n 'your-actual-app-password' | gcloud secrets versions add smtp-password --data-file=-"
echo ""
echo "2. Update admin emails in deploy.sh or cloudbuild.yaml if needed"
echo "   Current admin email: clough.warren@gmail.com"
echo "   For multiple admins, use comma-separated list: email1@domain.com,email2@domain.com"
echo ""
echo "3. Deploy your application:"
echo "   ./deploy.sh"
echo ""

# Grant the default compute service account access to the secrets
COMPUTE_SA="${PROJECT_ID}-compute@developer.gserviceaccount.com"
echo "🔑 Granting compute service account access to secrets..."

gcloud secrets add-iam-policy-binding smtp-username \
    --member="serviceAccount:${COMPUTE_SA}" \
    --role="roles/secretmanager.secretAccessor" || true

gcloud secrets add-iam-policy-binding smtp-password \
    --member="serviceAccount:${COMPUTE_SA}" \
    --role="roles/secretmanager.secretAccessor" || true

gcloud secrets add-iam-policy-binding firebase-service-account \
    --member="serviceAccount:${COMPUTE_SA}" \
    --role="roles/secretmanager.secretAccessor" || true

echo "✅ Secret permissions configured!"