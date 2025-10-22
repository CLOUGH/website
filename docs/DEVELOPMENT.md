# Development Guide

This guide covers local development setup, coding standards, and contribution workflows for the website project.

## Quick Start

### Prerequisites

- **Go 1.21+**: [Download and install Go](https://golang.org/dl/)
- **Git**: For version control
- **Firebase Account**: [Create a Firebase project](https://console.firebase.google.com/)
- **Text Editor**: VS Code, GoLand, or similar with Go support

### Local Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd website
   ```

2. **Install Go dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up Firebase**
   - Create a new Firebase project
   - Enable Firestore Database
   - Enable Authentication with Google provider
   - Download service account key
   - Place the JSON file in project root

4. **Configure environment**
   ```bash
   export GOOGLE_APPLICATION_CREDENTIALS="./your-firebase-key.json"
   export FIREBASE_PROJECT_ID="your-project-id"
   export PORT="3000"
   export GIN_MODE="debug"
   ```

5. **Seed the database (optional)**
   ```bash
   go run scripts/seed.go
   ```

6. **Start the development server**
   ```bash
   go run main.go
   ```

   Visit `http://localhost:3000` to see your website.

## Development Workflow

### 1. Project Structure

Understanding the project layout is crucial for effective development:

```
internal/          # Private application code (not importable by other projects)
├── config/        # Configuration management
├── handlers/      # HTTP request handlers (controllers)
├── middleware/    # HTTP middleware components
├── models/        # Data structures and models
└── services/      # Business logic layer

web/              # Web assets and templates
├── static/       # Static files (CSS, JS, images)
└── templates/    # HTML templates

scripts/          # Utility scripts
docs/            # Documentation
```

### 2. Making Changes

**Backend Changes:**
1. Models: Define data structures in `internal/models/`
2. Services: Implement business logic in `internal/services/`
3. Handlers: Add HTTP endpoints in `internal/handlers/`
4. Middleware: Add cross-cutting concerns in `internal/middleware/`

**Frontend Changes:**
1. Templates: Modify HTML in `web/templates/`
2. Styles: Update CSS in `web/static/css/`
3. Scripts: Add JavaScript in `web/static/js/`
4. Assets: Add images in `web/static/images/`

**Database Changes:**
1. Update models in `internal/models/content.go`
2. Modify Firestore rules in `firestore.rules`
3. Update indexes in `firestore.indexes.json`
4. Create migration scripts if needed

### 3. Testing Your Changes

**Run the application:**
```bash
go run main.go
```

**Check for errors:**
```bash
go vet ./...
go fmt ./...
```

**Test database operations:**
```bash
go run scripts/test-data.go
```

**Test admin interface:**
1. Visit `http://localhost:3000/admin/login`
2. Sign in with Google
3. Test content management features

## Coding Standards

### 1. Go Code Style

**Follow standard Go conventions:**
- Use `gofmt` for formatting
- Use `go vet` for static analysis
- Follow effective Go practices
- Use meaningful variable and function names

**Example:**
```go
// Good
func (cs *ContentService) GetPageSections(ctx context.Context, pageSlug string) ([]models.Section, error) {
    var sections []models.Section
    
    iter := cs.fs.client.Collection("sections").
        Where("page_slug", "==", pageSlug).
        OrderBy("order", firestore.Asc).
        Documents(ctx)
    defer iter.Stop()
    
    for {
        doc, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            return nil, fmt.Errorf("failed to fetch sections: %w", err)
        }
        
        var section models.Section
        if err := doc.DataTo(&section); err != nil {
            return nil, fmt.Errorf("failed to parse section: %w", err)
        }
        section.ID = doc.Ref.ID
        sections = append(sections, section)
    }
    
    return sections, nil
}

// Bad
func getsections(slug string) []models.Section {
    // Missing error handling
    // No context
    // Poor naming
}
```

**Error Handling:**
```go
// Wrap errors with context
if err != nil {
    return fmt.Errorf("failed to create section: %w", err)
}

// Handle errors at appropriate levels
if err := contentService.CreateSection(ctx, section); err != nil {
    log.Printf("Error creating section: %v", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create section"})
    return
}
```

### 2. Database Patterns

**Use context everywhere:**
```go
func (cs *ContentService) GetPage(ctx context.Context, pageSlug string) (*models.Page, error) {
    // Always pass context to Firestore operations
    doc, err := cs.fs.client.Collection("pages").
        Where("slug", "==", pageSlug).
        Limit(1).
        Documents(ctx).Next()
}
```

**Consistent error handling:**
```go
// Service layer - return detailed errors
if err != nil {
    return nil, fmt.Errorf("failed to fetch page %s: %w", pageSlug, err)
}

// Handler layer - return user-friendly errors
if err != nil {
    log.Printf("Error fetching page: %v", err)
    c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
    return
}
```

**Timestamp management:**
```go
// Always set timestamps
section.CreatedAt = time.Now()
section.UpdatedAt = time.Now()
section.UpdatedBy = getCurrentUser(ctx)
```

### 3. Template Patterns

**Base template structure:**
```html
{{define "base"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.title}}</title>
    <!-- Common head elements -->
</head>
<body>
    {{template "content" .}}
</body>
</html>
{{end}}
```

**Component templates:**
```html
{{define "hero"}}
<section class="hero">
    <h1>{{.Title}}</h1>
    {{if .Subtitle}}<h2>{{.Subtitle}}</h2>{{end}}
    <p>{{.Content}}</p>
    {{with .Data.cta_text}}
        <a href="{{$.Data.cta_link}}" class="btn">{{.}}</a>
    {{end}}
</section>
{{end}}
```

**HTMX integration:**
```html
<!-- Inline editing -->
<div id="section-{{.ID}}">
    <h3>{{.Title}}</h3>
    <button hx-get="/admin/sections/{{.ID}}/edit" 
            hx-target="#section-{{.ID}}"
            hx-swap="outerHTML">
        Edit
    </button>
</div>

<!-- Form submission -->
<form hx-post="/admin/sections" 
      hx-target="#sections-list"
      hx-swap="beforeend">
    <!-- Form fields -->
</form>
```

### 4. API Design

**Consistent response format:**
```go
// Success response
c.JSON(http.StatusOK, gin.H{
    "data": result,
    "message": "Operation successful",
})

// Error response
c.JSON(http.StatusBadRequest, gin.H{
    "error": "Validation failed",
    "details": validationErrors,
})
```

**Input validation:**
```go
type CreateSectionRequest struct {
    Slug     string                 `json:"slug" binding:"required"`
    Type     string                 `json:"type" binding:"required"`
    Title    string                 `json:"title" binding:"required"`
    Content  string                 `json:"content"`
    Data     map[string]interface{} `json:"data"`
    Order    int                    `json:"order" binding:"min=1"`
    Visible  bool                   `json:"visible"`
}

func (h *AdminHandler) CreateSection(c *gin.Context) {
    var req CreateSectionRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Process request
}
```

## Common Development Tasks

### 1. Adding a New Section Type

**Step 1: Define the template**
```html
<!-- web/templates/sections/my-new-section.html -->
{{define "my-new-section"}}
<section class="my-new-section">
    <h2>{{.Title}}</h2>
    <p>{{.Content}}</p>
    <!-- Custom data access -->
    {{with .Data.custom_field}}
        <div class="custom-content">{{.}}</div>
    {{end}}
</section>
{{end}}
```

**Step 2: Update template loading**
```go
// main.go
router.LoadHTMLFiles(
    "web/templates/base.html",
    "web/templates/home.html",
    "web/templates/sections/my-new-section.html", // Add this
    // ... other templates
)
```

**Step 3: Add to admin form**
```html
<!-- web/templates/admin/section-form.html -->
<option value="my-new-section">My New Section</option>
```

### 2. Adding New API Endpoints

**Step 1: Define the handler**
```go
// internal/handlers/home.go
func (h *HomeHandler) GetMyData(c *gin.Context) {
    data, err := h.contentService.GetMyData(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
        return
    }
    c.JSON(http.StatusOK, data)
}
```

**Step 2: Add the route**
```go
// main.go
api := router.Group("/api")
{
    api.GET("/my-data", homeHandler.GetMyData)
}
```

**Step 3: Implement service method**
```go
// internal/services/content.go
func (cs *ContentService) GetMyData(ctx context.Context) ([]MyData, error) {
    // Implementation
}
```

### 3. Adding Database Migrations

For schema changes that require data migration:

**Step 1: Create migration script**
```go
// scripts/migrate-v2.go
package main

import (
    "context"
    "log"
    // imports
)

func main() {
    // Connect to Firebase
    // Perform data transformation
    // Update existing documents
}
```

**Step 2: Update models**
```go
// internal/models/content.go
type Section struct {
    // Existing fields
    NewField string `json:"new_field" firestore:"new_field"`
}
```

**Step 3: Update Firestore rules if needed**
```
// firestore.rules
match /sections/{sectionId} {
    allow read: if true;
    allow write: if request.auth != null 
        && request.auth.token.email in get(/databases/$(database)/documents/admin_users/list).data.emails
        && validateSectionData(request.resource.data);
}

function validateSectionData(data) {
    return data.keys().hasAll(['title', 'content', 'new_field']) // Add new field
        && data.title is string
        && data.content is string
        && data.new_field is string; // Validate new field
}
```

### 4. Debugging Common Issues

**Template errors:**
```bash
# Check template syntax
go run main.go 2>&1 | grep template

# Test template rendering
{{/* Add debug output in templates */}}
<!-- Debug: {{printf "%+v" .}} -->
```

**Database connection issues:**
```bash
# Test Firebase connection
go run scripts/test-data.go

# Check Firestore rules
firebase deploy --only firestore:rules
```

**Authentication problems:**
```bash
# Verify environment variables
echo $GOOGLE_APPLICATION_CREDENTIALS
echo $FIREBASE_PROJECT_ID

# Test admin access
# Visit /admin/login and check browser console for errors
```

**HTMX issues:**
```html
<!-- Add HTMX debug attributes -->
<div hx-get="/api/data" 
     hx-trigger="click"
     hx-indicator="#loading"
     hx-on="htmx:responseError: console.error('HTMX Error:', event.detail)">
</div>
```

## Performance Optimization

### 1. Database Optimization

**Use appropriate indexes:**
```json
// firestore.indexes.json
{
  "indexes": [
    {
      "collectionGroup": "sections",
      "queryScope": "COLLECTION",
      "fields": [
        {"fieldPath": "page_slug", "order": "ASCENDING"},
        {"fieldPath": "visible", "order": "ASCENDING"},
        {"fieldPath": "order", "order": "ASCENDING"}
      ]
    }
  ]
}
```

**Optimize queries:**
```go
// Good - use specific queries
iter := cs.fs.client.Collection("sections").
    Where("page_slug", "==", pageSlug).
    Where("visible", "==", true).
    OrderBy("order", firestore.Asc).
    Limit(10).
    Documents(ctx)

// Bad - fetch everything then filter
iter := cs.fs.client.Collection("sections").Documents(ctx)
// Then filter in code
```

### 2. Template Optimization

**Cache template compilation:**
```go
// Load templates once at startup
router.LoadHTMLGlob("web/templates/**/*")

// Don't reload templates on every request
```

**Minimize template complexity:**
```html
<!-- Good - simple logic -->
{{if .Visible}}
    {{template "section" .}}
{{end}}

<!-- Bad - complex logic in templates -->
{{if and .Visible (eq .Type "hero") (gt .Order 0)}}
    <!-- Complex template logic -->
{{end}}
```

### 3. Static Asset Optimization

**Serve static files efficiently:**
```go
// Use Gin's static file serving
router.Static("/static", "./web/static")

// Add caching headers
router.Use(func(c *gin.Context) {
    if strings.HasPrefix(c.Request.URL.Path, "/static") {
        c.Header("Cache-Control", "public, max-age=31536000")
    }
    c.Next()
})
```

## Testing Guidelines

### 1. Unit Testing

**Test service functions:**
```go
// internal/services/content_test.go
func TestContentService_GetPageSections(t *testing.T) {
    // Setup mock Firestore
    // Create test data
    // Call service method
    // Assert results
}
```

**Test handlers:**
```go
// internal/handlers/admin_test.go
func TestAdminHandler_CreateSection(t *testing.T) {
    // Setup test server
    // Make request
    // Assert response
}
```

### 2. Integration Testing

**Test with Firestore emulator:**
```bash
# Start emulator
firebase emulators:start --only firestore

# Run tests
FIRESTORE_EMULATOR_HOST=localhost:8080 go test ./...
```

### 3. Manual Testing

**Test checklist:**
- [ ] Homepage loads correctly
- [ ] All sections render properly
- [ ] Admin login works
- [ ] CRUD operations in admin work
- [ ] Drag and drop reordering works
- [ ] API endpoints return correct data
- [ ] Error pages display correctly
- [ ] Mobile responsiveness

## Deployment Preparation

### 1. Build Optimization

**Production build:**
```bash
# Set production mode
export GIN_MODE=release

# Build optimized binary
go build -ldflags="-s -w" -o website main.go
```

### 2. Security Checklist

- [ ] Update Firestore security rules
- [ ] Configure CORS for production domains
- [ ] Set up proper admin user authorization
- [ ] Enable HTTPS
- [ ] Configure security headers
- [ ] Review environment variables

### 3. Performance Checklist

- [ ] Test with production data volume
- [ ] Verify Firestore indexes are deployed
- [ ] Configure caching headers
- [ ] Optimize image sizes
- [ ] Test page load speeds
- [ ] Verify mobile performance

## Troubleshooting

### Common Issues

**"Template not found" errors:**
```bash
# Check template file paths
ls -la web/templates/
# Verify LoadHTMLFiles paths match actual files
```

**"Index not found" Firestore errors:**
```bash
# Deploy indexes
firebase deploy --only firestore:indexes
# Wait for indexes to build (can take several minutes)
```

**Authentication failures:**
```bash
# Check environment variables
env | grep FIREBASE
env | grep GOOGLE

# Verify service account permissions
# Check Firebase console for auth configuration
```

**HTMX not working:**
```html
<!-- Check HTMX is loaded -->
<script src="https://unpkg.com/htmx.org@1.9.6"></script>

<!-- Add debug attributes -->
<div hx-get="/api/test" hx-trigger="click" hx-swap="innerHTML">
    Click me
</div>
```

### Getting Help

1. **Check logs**: Look at application logs for error details
2. **Browser console**: Check for JavaScript errors with HTMX
3. **Firebase console**: Monitor Firestore usage and errors
4. **Documentation**: Refer to framework documentation
5. **Community**: Search Go, Gin, HTMX, and Firebase communities

This development guide should help you get started with the codebase and contribute effectively. Remember to follow the coding standards and test your changes thoroughly before deployment.