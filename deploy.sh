#!/bin/bash

# Warren Clough Website Deployment Script
# This script deploys the website to Firebase

set -e

echo "🚀 Starting deployment of Warren Clough website..."

# Check if Firebase CLI is installed
if ! command -v firebase &> /dev/null; then
    echo "❌ Firebase CLI is not installed. Please install it first:"
    echo "npm install -g firebase-tools"
    exit 1
fi

# Check if user is logged in to Firebase
if ! firebase projects:list &> /dev/null; then
    echo "❌ Not logged in to Firebase. Please run:"
    echo "firebase login"
    exit 1
fi

# Install Go dependencies
echo "📦 Installing Go dependencies..."
go mod tidy

# Build the function
echo "🔨 Building Go function..."
go build -o functions functions.go

# Deploy to Firebase
echo "🚀 Deploying to Firebase..."
firebase deploy

echo "✅ Deployment completed successfully!"
echo "🌐 Your website should be available at your Firebase hosting URL"
