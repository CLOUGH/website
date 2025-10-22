# Admin Access Configuration

This guide explains how to configure admin access for your website's admin dashboard using environment variables.

## Overview

Admin access is controlled by the `ADMIN_EMAILS` environment variable, which contains a comma-separated list of email addresses that have admin privileges.

## Configuration

### Local Development

Update your `.env` file:

```env
# Admin Configuration
ADMIN_EMAILS=clough.warren@gmail.com
```

For multiple admins:
```env
# Multiple admins (comma-separated, no spaces around commas)
ADMIN_EMAILS=admin1@example.com,admin2@example.com,admin3@example.com
```

### Cloud Run Deployment

Admin emails are set as environment variables during deployment:

**Option 1: Update deploy.sh**
```bash
--set-env-vars="...,ADMIN_EMAILS=clough.warren@gmail.com"
```

**Option 2: Update cloudbuild.yaml**
```yaml
- '--set-env-vars=...,ADMIN_EMAILS=clough.warren@gmail.com'
```

### Adding Multiple Admins

To add multiple admin emails, use a comma-separated list (no spaces):

**Local (.env)**:
```env
ADMIN_EMAILS=warren@example.com,admin@company.com,manager@company.com
```

**Cloud Run (deploy.sh)**:
```bash
--set-env-vars="...,ADMIN_EMAILS=warren@example.com,admin@company.com,manager@company.com"
```

## How It Works

1. **Environment Variable Loading**: The application loads `ADMIN_EMAILS` from environment variables
2. **Email Parsing**: The config system splits the comma-separated string into an array
3. **Authentication Check**: When a user tries to access admin routes, their email is checked against this list
4. **Access Control**: Only emails in the list can access `/admin/*` routes

## Code Structure

The admin email configuration flows through:

1. **Config Loading** (`internal/config/config.go`):
   ```go
   AdminEmails: getAdminEmails()
   ```

2. **Content Service** (`internal/services/content.go`):
   ```go
   func (cs *ContentService) IsUserAdmin(ctx context.Context, email string) (bool, error) {
       adminEmails := cs.config.AdminEmails
       // Check if email is in admin list
   }
   ```

3. **Auth Middleware** (`internal/middleware/auth.go`):
   ```go
   isAdmin, err := am.contentService.IsUserAdmin(c.Request.Context(), userEmail)
   ```

## Testing Admin Access

### 1. Local Testing

Start your local server:
```bash
go run main.go
```

Navigate to admin login:
```
http://localhost:8080/admin/login
```

### 2. Cloud Run Testing

After deployment, test admin access:
```bash
CLOUD_RUN_URL=$(gcloud run services describe warren-clough-website --region=us-central1 --format='value(status.url)')
echo "Admin login: $CLOUD_RUN_URL/admin/login"
```

## Security Notes

- **Email Verification**: Only verified Gmail/Google accounts can access admin routes
- **Firebase Auth**: Admin status is checked after Firebase authentication
- **Case Sensitive**: Email addresses are case-sensitive in the comparison
- **No Wildcards**: Each admin email must be explicitly listed

## Troubleshooting

### Admin Access Denied

1. **Check Email Format**: Ensure the email matches exactly (case-sensitive)
2. **Verify Environment Variable**: 
   ```bash
   # Local
   echo $ADMIN_EMAILS
   
   # Cloud Run
   gcloud run services describe warren-clough-website --region=us-central1 --format='value(spec.template.spec.template.spec.containers[0].env[?(@.name=="ADMIN_EMAILS")].value)'
   ```

3. **Check Deployment**: Ensure environment variable was set during deployment
4. **Verify Firebase Auth**: Make sure the user is authenticated with Firebase

### Multiple Admins Not Working

1. **No Spaces**: Ensure no spaces around commas: `admin1@example.com,admin2@example.com`
2. **Proper Escaping**: In shell scripts, quote the entire value: `"admin1@example.com,admin2@example.com"`

### Environment Variable Not Set

If `ADMIN_EMAILS` is not set, it defaults to `clough.warren@gmail.com`. You can verify this in the logs:

```bash
# Check application logs
gcloud logs read --service=warren-clough-website --limit=10
```

## Updating Admin Emails

### For Cloud Run

1. **Update deployment script**:
   ```bash
   # Edit deploy.sh or cloudbuild.yaml
   vim deploy.sh
   ```

2. **Redeploy**:
   ```bash
   ./deploy.sh
   ```

### For Local Development

1. **Update .env file**:
   ```bash
   echo "ADMIN_EMAILS=new@example.com,admin@example.com" >> .env
   ```

2. **Restart server**:
   ```bash
   go run main.go
   ```

## Best Practices

1. **Least Privilege**: Only add necessary admin emails
2. **Regular Review**: Periodically review and remove unused admin access
3. **Secure Emails**: Use corporate/verified email addresses for admins
4. **Backup Access**: Ensure at least 2 admins to prevent lockout
5. **Documentation**: Keep a record of who has admin access and why

## Example Configurations

### Single Admin
```env
ADMIN_EMAILS=warren@mycompany.com
```

### Team Access
```env
ADMIN_EMAILS=warren@mycompany.com,alice@mycompany.com,bob@mycompany.com
```

### Development vs Production
```env
# Development
ADMIN_EMAILS=warren@mycompany.com,dev-team@mycompany.com

# Production  
ADMIN_EMAILS=warren@mycompany.com,admin@mycompany.com
```