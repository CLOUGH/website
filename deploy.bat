@echo off
REM Warren Clough Website Deployment Script for Windows
REM This script deploys the website to Firebase

echo 🚀 Starting deployment of Warren Clough website...

REM Check if Firebase CLI is installed
firebase --version >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ Firebase CLI is not installed. Please install it first:
    echo npm install -g firebase-tools
    pause
    exit /b 1
)

REM Check if user is logged in to Firebase
firebase projects:list >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ Not logged in to Firebase. Please run:
    echo firebase login
    pause
    exit /b 1
)

REM Install Go dependencies
echo 📦 Installing Go dependencies...
go mod tidy

REM Build the function
echo 🔨 Building Go function...
go build -o functions.exe functions.go

REM Deploy to Firebase
echo 🚀 Deploying to Firebase...
firebase deploy

echo ✅ Deployment completed successfully!
echo 🌐 Your website should be available at your Firebase hosting URL
pause
