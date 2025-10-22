# Cloud Run SMTP Configuration Guide

This guide explains how to configure SMTP settings for your Cloud Run deployment using Google Secret Manager for secure credential storage.

## Overview

Your application uses Gmail SMTP for sending contact form emails. In production, credentials are stored securely in Google Secret Manager and injected as environment variables into Cloud Run.

## Setup Process

### 1. Store Secrets in Google Secret Manager

Run the setup script to create secrets:

```bash
./setup-secrets.sh
```

This script will:
- Create secrets for SMTP username and password
- Upload your Firebase service account key
- Grant permissions to Cloud Run service account

### 2. Update SMTP Password

After running the setup script, update your Gmail app password:

```bash
# Replace with your actual Gmail app password
echo -n 'your-actual-app-password' | gcloud secrets versions add smtp-password --data-file=-
```

### 3. Deploy to Cloud Run

Deploy using the deploy script:

```bash
./deploy.sh
```

Or using Cloud Build:

```bash
gcloud builds submit --config cloudbuild.yaml
```

## Environment Variables in Cloud Run

The following environment variables are configured automatically:

### Regular Environment Variables
- `GIN_MODE=release` - Sets Gin to production mode
- `FIREBASE_PROJECT_ID` - Your Firebase project ID
- `SMTP_HOST=smtp.gmail.com` - Gmail SMTP server
- `SMTP_PORT=587` - Gmail SMTP port
- `ADMIN_EMAILS` - Comma-separated list of admin email addresses

### Secret Environment Variables
- `SMTP_USERNAME` - Your Gmail address (from Secret Manager)
- `SMTP_PASSWORD` - Your Gmail app password (from Secret Manager)
- `GOOGLE_APPLICATION_CREDENTIALS` - Firebase service account key (from Secret Manager)

## Gmail App Password Setup

1. **Enable 2-Factor Authentication** on your Gmail account
2. **Generate App Password**:
   - Go to [Google Account Settings](https://myaccount.google.com/)
   - Security → 2-Step Verification → App passwords
   - Generate password for "Mail"
3. **Update Secret**:
   ```bash
   echo -n 'generated-app-password' | gcloud secrets versions add smtp-password --data-file=-
   ```

## Verifying Configuration

After deployment, test the contact form:

```bash
CLOUD_RUN_URL=$(gcloud run services describe warren-clough-website --region=us-central1 --format='value(status.url)')

curl -X POST $CLOUD_RUN_URL/api/contact \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com", 
    "message": "Test message from Cloud Run"
  }'
```

## Security Notes

- Secrets are never stored in code or environment files
- Secret Manager provides automatic encryption and access logging
- Cloud Run service account has minimal required permissions
- Secrets are injected at runtime as environment variables

## Local Development

For local development, the application falls back to the `.env` file:

```env
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your-email@gmail.com
SMTP_PASSWORD=your-app-password
```

## Troubleshooting

### Secret Access Issues
```bash
# Check if secrets exist
gcloud secrets list

# Check secret permissions
gcloud secrets get-iam-policy smtp-password
```

### Cloud Run Logs
```bash
# View application logs
gcloud logs read --service=warren-clough-website --limit=50
```

### Email Delivery Issues
- Verify Gmail app password is correct
- Check Gmail security settings
- Verify sender email is authorized
- Check Cloud Run service logs for SMTP errors

## Manual Secret Management

If you need to manage secrets manually:

```bash
# Create secret
echo -n 'secret-value' | gcloud secrets create secret-name --data-file=-

# Update secret
echo -n 'new-value' | gcloud secrets versions add secret-name --data-file=-

# Grant access to Cloud Run service account
PROJECT_ID=$(gcloud config get-value project)
COMPUTE_SA="${PROJECT_ID}-compute@developer.gserviceaccount.com"

gcloud secrets add-iam-policy-binding secret-name \
  --member="serviceAccount:${COMPUTE_SA}" \
  --role="roles/secretmanager.secretAccessor"
```