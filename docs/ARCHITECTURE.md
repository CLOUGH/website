# Architecture Overview

This document provides a comprehensive overview of the website's architecture, design patterns, and technical decisions.

## System Architecture

### High-Level Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Web Browser   │    │   Load Balancer │    │   Application   │
│                 │◄──►│    (Optional)   │◄──►│     Server      │
│  HTMX + CSS     │    │                 │    │   (Go/Gin)     │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                                        │
                                                        ▼
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Admin Panel   │    │    Firebase     │    │   Google OAuth  │
│                 │◄──►│   Firestore     │    │                 │
│  HTMX Dashboard │    │   Database      │    │ Authentication  │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

### Technology Stack

**Backend:**
- **Language**: Go 1.21+
- **Framework**: Gin Web Framework
- **Database**: Google Firestore (NoSQL)
- **Authentication**: Firebase Auth with Google OAuth
- **Deployment**: Docker, Google Cloud Run, App Engine

**Frontend:**
- **Core**: Server-side rendered HTML templates
- **Styling**: TailwindCSS
- **Interactivity**: HTMX for dynamic behavior
- **Icons**: Devicons, Font Awesome
- **No JavaScript framework** - HTMX handles all dynamic interactions

**Infrastructure:**
- **Hosting**: Google Cloud Platform
- **Database**: Firestore with automatic scaling
- **CDN**: Google Cloud CDN / Cloudflare
- **Monitoring**: Google Cloud Monitoring
- **SSL**: Automatic with Google-managed certificates

## Application Structure

### Directory Layout

```
├── internal/                 # Private application code
│   ├── config/              # Configuration management
│   │   └── config.go        # Environment and app configuration
│   ├── handlers/            # HTTP request handlers
│   │   ├── home.go          # Public website handlers
│   │   └── admin.go         # Admin interface handlers
│   ├── middleware/          # HTTP middleware components
│   │   ├── auth.go          # Authentication middleware
│   │   └── middleware.go    # CORS, security headers, etc.
│   ├── models/              # Data models and structures
│   │   ├── models.go        # Legacy models (for migration)
│   │   └── content.go       # Core CMS models (Page, Section, User)
│   └── services/            # Business logic layer
│       ├── content.go       # Content management service
│       └── firebase.go      # Firebase/Firestore service
├── web/                     # Web assets and templates
│   ├── static/              # Static files (CSS, JS, images)
│   │   ├── css/            # Custom CSS files
│   │   ├── js/             # Custom JavaScript files
│   │   └── images/         # Image assets
│   └── templates/           # HTML templates
│       ├── base.html        # Base layout template
│       ├── home.html        # Homepage template
│       ├── error.html       # Error page template
│       └── admin/           # Admin interface templates
│           ├── dashboard.html
│           ├── page-form.html
│           └── section-form.html
├── scripts/                 # Utility scripts
│   ├── seed.go             # Database seeding script
│   └── test-data.go        # Data testing utilities
├── docs/                    # Documentation
├── main.go                  # Application entry point
├── go.mod                   # Go module definition
├── Dockerfile              # Container definition
├── firebase.json           # Firebase configuration
├── firestore.rules         # Database security rules
└── firestore.indexes.json  # Database indexes
```

## Core Components

### 1. Content Management System (CMS)

The application implements a generic CMS based on two primary entities:

#### Pages
- Represent top-level content areas (home, blog, about, etc.)
- Contains metadata for SEO (title, description, keywords)
- Can be enabled/disabled independently
- Support custom routing patterns

#### Sections
- Content blocks within pages
- Ordered display sequence
- Flexible JSON data storage for type-specific content
- Support for different templates and custom CSS/JS
- Types: hero, about, technologies, experience, projects, contact

```go
type Page struct {
    ID          string    `firestore:"id"`
    Slug        string    `firestore:"slug"`
    Title       string    `firestore:"title"`
    Description string    `firestore:"description"`
    Metadata    PageMeta  `firestore:"metadata"`
    Visible     bool      `firestore:"visible"`
    CreatedAt   time.Time `firestore:"created_at"`
    UpdatedAt   time.Time `firestore:"updated_at"`
    UpdatedBy   string    `firestore:"updated_by"`
}

type Section struct {
    ID        string                 `firestore:"id"`
    PageSlug  string                 `firestore:"page_slug"`
    Slug      string                 `firestore:"slug"`
    Type      string                 `firestore:"type"`
    Title     string                 `firestore:"title"`
    Subtitle  string                 `firestore:"subtitle"`
    Content   string                 `firestore:"content"`
    Data      map[string]interface{} `firestore:"data"`
    Order     int                    `firestore:"order"`
    Visible   bool                   `firestore:"visible"`
    Template  string                 `firestore:"template"`
    CreatedAt time.Time              `firestore:"created_at"`
    UpdatedAt time.Time              `firestore:"updated_at"`
    UpdatedBy string                 `firestore:"updated_by"`
}
```

### 2. Authentication System

**Google OAuth Integration:**
- Users authenticate via Google OAuth 2.0
- Firebase Auth handles token validation
- Session management through secure HTTP cookies
- Admin access controlled through configurable user whitelist

**Authentication Flow:**
```
1. User visits /admin/login
2. Click "Sign in with Google"
3. Redirect to Google OAuth
4. Google returns ID token
5. Validate token with Firebase Auth
6. Check user admin permissions
7. Create session and redirect to dashboard
```

### 3. API Layer

**RESTful API Design:**
- Generic endpoints: `/api/page/{pageSlug}/sections`
- Consistent JSON responses
- Proper HTTP status codes
- Error handling with structured responses

**Admin API:**
- CRUD operations for pages and sections
- Real-time updates via HTMX
- Drag-and-drop section reordering
- Inline editing capabilities

### 4. Template System

**Server-Side Rendering:**
- Go's `html/template` package
- Component-based template architecture
- Template inheritance with base layout
- Context-aware rendering based on user authentication

**Template Structure:**
```
base.html                    # Main layout with navigation, footer
├── home.html               # Homepage content
├── error.html              # Error pages (404, 500, etc.)
└── admin/
    ├── dashboard.html      # Admin dashboard
    ├── page-form.html      # Page creation/editing
    └── section-form.html   # Section creation/editing
```

## Design Patterns

### 1. Repository Pattern

Content service acts as a repository layer:
```go
type ContentService interface {
    GetPage(ctx context.Context, slug string) (*models.Page, error)
    GetPageSections(ctx context.Context, pageSlug string) ([]models.Section, error)
    CreateSection(ctx context.Context, section *models.Section) error
    UpdateSection(ctx context.Context, section *models.Section) error
    DeleteSection(ctx context.Context, sectionID string) error
}
```

### 2. Middleware Chain

HTTP middleware for cross-cutting concerns:
```go
// Middleware stack
router.Use(gin.Logger())
router.Use(gin.Recovery())
router.Use(CORS())
router.Use(SecurityHeaders())

// Admin routes with authentication
admin := router.Group("/admin")
admin.Use(authMiddleware.RequireAuth())
```

### 3. Dependency Injection

Services are injected into handlers:
```go
func main() {
    firebaseService := services.NewFirebaseService(ctx, config)
    contentService := services.NewContentService(firebaseService)
    
    homeHandler := handlers.NewHomeHandler(contentService)
    adminHandler := handlers.NewAdminHandler(contentService)
}
```

### 4. Configuration Pattern

Environment-based configuration:
```go
type Config struct {
    FirebaseProjectID string
    Port              string
    Environment       string
}

func Load() *Config {
    return &Config{
        FirebaseProjectID: getEnv("FIREBASE_PROJECT_ID", "default"),
        Port:              getEnv("PORT", "8080"),
        Environment:       getEnv("ENVIRONMENT", "development"),
    }
}
```

## Data Flow

### 1. Public Content Rendering

```
1. User requests page (e.g., /)
2. Home handler retrieves page metadata
3. Handler fetches ordered sections from Firestore
4. Template engine renders sections using type-specific templates
5. Complete HTML page returned to user
```

### 2. Admin Content Management

```
1. Admin user authenticated via Google OAuth
2. Admin requests section edit (/admin/sections/123/edit)
3. Handler retrieves section data from Firestore
4. HTMX renders edit form inline
5. Admin submits changes via HTMX
6. Handler validates and updates Firestore
7. HTMX swaps updated content without page reload
```

### 3. API Data Access

```
1. API request for sections (/api/page/home/sections)
2. Handler validates page exists
3. Content service queries Firestore with composite index
4. Results serialized to JSON
5. Response returned with appropriate caching headers
```

## Security Architecture

### 1. Authentication & Authorization

- **OAuth 2.0**: Google OAuth for admin authentication
- **Session Management**: Secure HTTP-only cookies
- **CSRF Protection**: Built-in Gin CSRF middleware
- **Admin Whitelist**: Configurable admin user authorization

### 2. Data Security

- **Firestore Rules**: Server-side security rules
- **Input Validation**: Comprehensive request validation
- **SQL Injection**: Not applicable (NoSQL database)
- **XSS Protection**: Template auto-escaping

### 3. Infrastructure Security

- **TLS**: HTTPS everywhere with automatic certificates
- **Security Headers**: HSTS, CSP, X-Frame-Options
- **CORS**: Configurable cross-origin policies
- **Rate Limiting**: API endpoint protection

## Performance Considerations

### 1. Database Optimization

**Firestore Indexing:**
```json
{
  "indexes": [
    {
      "collectionGroup": "sections",
      "queryScope": "COLLECTION",
      "fields": [
        {"fieldPath": "page_slug", "order": "ASCENDING"},
        {"fieldPath": "order", "order": "ASCENDING"}
      ]
    }
  ]
}
```

**Query Optimization:**
- Composite indexes for complex queries
- Pagination for large result sets
- Selective field projection
- Connection pooling

### 2. Caching Strategy

**Application Level:**
- Template compilation caching
- Firestore connection pooling
- Static asset caching

**Infrastructure Level:**
- CDN for static assets
- HTTP caching headers
- Browser caching strategies

### 3. Rendering Performance

- Server-side rendering (no client-side framework overhead)
- HTMX for minimal JavaScript
- Optimized template compilation
- Lazy loading of non-critical content

## Scalability Architecture

### 1. Horizontal Scaling

**Stateless Application:**
- No server-side sessions (cookie-based auth)
- Shared database state
- Container-friendly architecture

**Load Balancing:**
- Multiple application instances
- Database connection pooling
- Session affinity not required

### 2. Database Scaling

**Firestore Benefits:**
- Automatic scaling
- Global distribution
- Real-time synchronization
- Built-in backup and recovery

### 3. Deployment Scaling

**Container Orchestration:**
```yaml
# Kubernetes deployment example
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-website
spec:
  replicas: 3
  selector:
    matchLabels:
      app: my-website
  template:
    metadata:
      labels:
        app: my-website
    spec:
      containers:
      - name: website
        image: gcr.io/project/my-website:latest
        ports:
        - containerPort: 8080
        env:
        - name: FIREBASE_PROJECT_ID
          value: "your-project"
```

## Monitoring and Observability

### 1. Application Monitoring

**Health Checks:**
```go
router.GET("/health", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "status":    "ok",
        "timestamp": time.Now(),
        "version":   version,
    })
})
```

**Logging:**
- Structured logging with JSON format
- Request/response logging
- Error tracking with stack traces
- Performance metrics

### 2. Infrastructure Monitoring

- **Uptime Monitoring**: External health check monitoring
- **Performance Metrics**: Response time, throughput
- **Error Rates**: 4xx/5xx response tracking
- **Resource Usage**: CPU, memory, database connections

### 3. User Analytics

- **Google Analytics**: User behavior tracking
- **Performance**: Core Web Vitals monitoring
- **Admin Usage**: Admin interface usage patterns

## Future Architecture Considerations

### 1. Microservices Evolution

Current monolithic structure could evolve to:
- **Content Service**: Page and section management
- **Auth Service**: User authentication and authorization
- **Media Service**: Image and file management
- **Analytics Service**: Usage and performance tracking

### 2. Enhanced Caching

- **Redis Integration**: Application-level caching
- **CDN Enhancement**: Advanced caching strategies
- **Edge Computing**: Cloudflare Workers for edge rendering

### 3. Real-time Features

- **WebSocket Integration**: Real-time admin updates
- **Live Preview**: Real-time content preview
- **Collaborative Editing**: Multiple admin users

### 4. API Evolution

- **GraphQL**: More flexible data querying
- **API Versioning**: Backward compatibility
- **Rate Limiting**: Advanced throttling strategies
- **API Documentation**: Auto-generated documentation

## Testing Strategy

### 1. Unit Testing
```go
func TestContentService_GetPageSections(t *testing.T) {
    // Mock Firestore client
    // Test section retrieval
    // Verify ordering and filtering
}
```

### 2. Integration Testing
```go
func TestAdminHandler_CreateSection(t *testing.T) {
    // Test with real Firestore emulator
    // Test authentication flow
    // Verify database state changes
}
```

### 3. End-to-End Testing
- Automated browser testing with tools like Playwright
- Admin workflow testing
- Content management scenarios
- Authentication flows

This architecture provides a solid foundation for a personal website with CMS capabilities, emphasizing simplicity, security, and maintainability while leveraging modern cloud-native technologies.