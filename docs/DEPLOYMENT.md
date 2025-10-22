# Deployment Guide

This guide covers various deployment options for your personal website.

## Prerequisites

- Go 1.21 or later
- Firebase project with Firestore enabled
- Service account key from Firebase Console
- Domain name (optional, for custom domains)

## Environment Setup

Create the following environment variables for all deployment methods:

```bash
export FIREBASE_PROJECT_ID="your-firebase-project-id"
export GOOGLE_APPLICATION_CREDENTIALS="./path/to/service-account-key.json"
export PORT="8080"
export GIN_MODE="release"
```

## Deployment Options

### 1. Google Cloud Platform

#### Cloud Run (Recommended)

Cloud Run provides automatic scaling and is cost-effective for personal websites.

**Step 1: Prepare the application**
```bash
# Create a Dockerfile if it doesn't exist
cat > Dockerfile << EOF
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/web ./web
COPY --from=builder /app/*.json .
EXPOSE 8080
CMD ["./main"]
EOF
```

**Step 2: Build and deploy**
```bash
# Set project ID
gcloud config set project YOUR-PROJECT-ID

# Build the container
gcloud builds submit --tag gcr.io/YOUR-PROJECT-ID/my-website

# Deploy to Cloud Run
gcloud run deploy my-website \
  --image gcr.io/YOUR-PROJECT-ID/my-website \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars="FIREBASE_PROJECT_ID=YOUR-PROJECT-ID,GIN_MODE=release" \
  --port 8080
```

**Step 3: Set up custom domain (optional)**
```bash
# Map custom domain
gcloud run domain-mappings create \
  --service my-website \
  --domain yourdomain.com \
  --region us-central1
```

#### App Engine

App Engine provides a fully managed platform with automatic scaling.

**Step 1: Create app.yaml**
```yaml
runtime: go121

env_variables:
  FIREBASE_PROJECT_ID: "your-project-id"
  GIN_MODE: "release"

automatic_scaling:
  min_instances: 0
  max_instances: 10
  target_cpu_utilization: 0.6
```

**Step 2: Deploy**
```bash
gcloud app deploy
```

### 2. Firebase Hosting + Cloud Functions

This option uses Firebase Hosting for static assets and Cloud Functions for the backend.

**Step 1: Install Firebase CLI**
```bash
npm install -g firebase-tools
firebase login
```

**Step 2: Initialize Firebase**
```bash
firebase init hosting
firebase init functions
```

**Step 3: Create Cloud Function**
Create `functions/index.js`:
```javascript
const functions = require('firebase-functions');
const { spawn } = require('child_process');

exports.app = functions.https.onRequest((req, res) => {
  // Proxy to Go application
  // Implementation depends on your specific needs
});
```

**Step 4: Deploy**
```bash
firebase deploy
```

### 3. Docker Deployment

#### Local Docker
```bash
# Build the image
docker build -t my-website .

# Run locally
docker run -p 8080:8080 \
  -e FIREBASE_PROJECT_ID=your-project-id \
  -e GOOGLE_APPLICATION_CREDENTIALS=/app/credentials.json \
  -v /path/to/credentials.json:/app/credentials.json \
  my-website
```

#### Docker Compose
Create `docker-compose.yml`:
```yaml
version: '3.8'
services:
  website:
    build: .
    ports:
      - "8080:8080"
    environment:
      - FIREBASE_PROJECT_ID=your-project-id
      - GOOGLE_APPLICATION_CREDENTIALS=/app/credentials.json
      - GIN_MODE=release
    volumes:
      - ./credentials.json:/app/credentials.json:ro
    restart: unless-stopped
```

Run with:
```bash
docker-compose up -d
```

### 4. Traditional VPS/Server

#### Systemd Service

**Step 1: Build and deploy**
```bash
# Build for Linux (if cross-compiling)
GOOS=linux GOARCH=amd64 go build -o my-website

# Copy to server
scp my-website user@server:/opt/my-website/
scp -r web user@server:/opt/my-website/
scp credentials.json user@server:/opt/my-website/
```

**Step 2: Create systemd service**
```ini
# /etc/systemd/system/my-website.service
[Unit]
Description=My Personal Website
After=network.target

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/my-website
ExecStart=/opt/my-website/my-website
Environment=FIREBASE_PROJECT_ID=your-project-id
Environment=GOOGLE_APPLICATION_CREDENTIALS=/opt/my-website/credentials.json
Environment=GIN_MODE=release
Environment=PORT=8080
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
```

**Step 3: Start service**
```bash
sudo systemctl daemon-reload
sudo systemctl enable my-website
sudo systemctl start my-website
```

#### Nginx Reverse Proxy

Create `/etc/nginx/sites-available/my-website`:
```nginx
server {
    listen 80;
    server_name yourdomain.com www.yourdomain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # Optional: serve static files directly
    location /static/ {
        alias /opt/my-website/web/static/;
        expires 1y;
        add_header Cache-Control "public, immutable";
    }
}
```

Enable the site:
```bash
sudo ln -s /etc/nginx/sites-available/my-website /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

### 5. Heroku

**Step 1: Create Procfile**
```
web: ./main
```

**Step 2: Deploy**
```bash
# Login to Heroku
heroku login

# Create app
heroku create my-website-app

# Set environment variables
heroku config:set FIREBASE_PROJECT_ID=your-project-id
heroku config:set GIN_MODE=release

# Add credentials as config var (base64 encoded)
heroku config:set GOOGLE_APPLICATION_CREDENTIALS_JSON="$(base64 -i credentials.json)"

# Deploy
git push heroku main
```

## SSL/TLS Configuration

### Let's Encrypt (for VPS deployments)

```bash
# Install certbot
sudo apt install certbot python3-certbot-nginx

# Get certificate
sudo certbot --nginx -d yourdomain.com -d www.yourdomain.com

# Auto-renewal
sudo crontab -e
# Add: 0 12 * * * /usr/bin/certbot renew --quiet
```

### Cloudflare (DNS + SSL)

1. Add your domain to Cloudflare
2. Update nameservers at your registrar
3. Enable SSL/TLS encryption (Full mode)
4. Configure Page Rules for caching

## Monitoring and Maintenance

### Health Checks

Set up monitoring for the `/health` endpoint:

```bash
# Simple monitoring script
#!/bin/bash
URL="https://yourdomain.com/health"
if ! curl -f $URL > /dev/null 2>&1; then
    echo "Website is down!" | mail -s "Website Alert" your-email@example.com
fi
```

### Log Management

For VPS deployments, set up log rotation:
```bash
# /etc/logrotate.d/my-website
/var/log/my-website/*.log {
    daily
    missingok
    rotate 52
    compress
    delaycompress
    notifempty
    create 644 www-data www-data
    postrotate
        systemctl reload my-website
    endscript
}
```

### Backup Strategy

**Firestore backups:**
```bash
# Schedule regular exports
gcloud firestore export gs://your-backup-bucket/$(date +%Y-%m-%d)
```

**Application backups:**
```bash
# Backup application files and configs
tar -czf website-backup-$(date +%Y-%m-%d).tar.gz \
  /opt/my-website \
  /etc/nginx/sites-available/my-website \
  /etc/systemd/system/my-website.service
```

## Performance Optimization

### Caching

**Redis caching (optional):**
```bash
# Install Redis
sudo apt install redis-server

# Configure in your application
# Add caching middleware for static content and API responses
```

**CDN setup:**
- Configure Cloudflare or CloudFront
- Set appropriate cache headers
- Optimize image delivery

### Database Optimization

**Firestore best practices:**
- Use composite indexes for complex queries
- Implement pagination for large datasets
- Cache frequently accessed data
- Monitor usage in Firebase Console

## Troubleshooting

### Common Deployment Issues

**Port binding errors:**
```bash
# Check what's using the port
sudo lsof -i :8080
sudo netstat -tulpn | grep :8080
```

**Permission errors:**
```bash
# Fix file permissions
sudo chown -R www-data:www-data /opt/my-website
sudo chmod +x /opt/my-website/my-website
```

**Firestore connection issues:**
- Verify service account key path
- Check Firebase project ID
- Ensure Firestore is enabled
- Verify network connectivity to Firebase

### Rollback Procedures

**Cloud Run rollback:**
```bash
# List revisions
gcloud run revisions list --service=my-website

# Rollback to previous revision
gcloud run services update-traffic my-website \
  --to-revisions=REVISION-NAME=100
```

**VPS rollback:**
```bash
# Keep backup of working version
cp my-website my-website.backup

# Rollback if needed
mv my-website.backup my-website
sudo systemctl restart my-website
```

## Security Considerations

- Keep dependencies updated (`go mod tidy`)
- Use HTTPS everywhere
- Implement rate limiting for API endpoints
- Regularly rotate service account keys
- Monitor access logs for suspicious activity
- Use environment variables for all secrets
- Implement proper CORS policies
- Regular security audits of Firestore rules