# My Personal Website

A modern, responsive personal website built with Go, Gin, HTMX, TailwindCSS, and Firebase Firestore. This project features a dynamic CMS system with admin interface for content management and can be deployed to various platforms including Google Cloud Platform.

## Features

- **Dynamic CMS**: Generic page and section management with flexible JSON data storage
- **Modern Tech Stack**: Go/Gin backend with HTMX for dynamic interactions and TailwindCSS for styling
- **Admin Interface**: HTMX-powered admin dashboard for content management with drag-and-drop ordering
- **Google OAuth**: Secure authentication for admin access using Firebase Auth
- **Generic API**: RESTful API with pattern `/api/page/{pageSlug}/sections` for flexible content access
- **Component-Based Templates**: Modular Go templates for different section types (hero, about, projects, etc.)
- **Firebase Integration**: Uses Firestore for database storage with real-time capabilities
- **Responsive Design**: Clean, modern design that works on all devices
- **Performance Optimized**: Fast loading times with optimized queries and caching

## Project Structure

```
├── internal/
│   ├── config/          # Configuration management
│   ├── handlers/        # HTTP request handlers (home, admin)
│   ├── middleware/      # HTTP middleware (auth, CORS, security)
│   ├── models/          # Data models (Page, Section, User)
│   └── services/        # Business logic services (content, firebase)
├── web/
│   ├── static/          # Static assets (CSS, JS, images)
│   └── templates/       # HTML templates (base, components, admin)
├── scripts/             # Utility scripts (seeding, testing)
├── google-cloud-sdk/    # Google Cloud SDK for deployment
├── firebase.json        # Firebase configuration
├── firestore.rules      # Firestore security rules
├── firestore.indexes.json # Firestore indexes
├── main.go             # Application entry point
├── go.mod              # Go module dependencies
└── *.json              # Firebase service account credentials
```

## Getting Started

### Prerequisites

- Go 1.21 or later
- Firebase CLI
- A Firebase project

### Local Development

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd website
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up Firebase credentials**
   - Download your Firebase service account key from the Firebase Console
   - Place it in the project root (e.g., `my-website-firebase-adminsdk.json`)

4. **Set up environment variables**
   ```bash
   export GOOGLE_APPLICATION_CREDENTIALS="./credentials/firebase-service-account.json"
   export FIREBASE_PROJECT_ID="your-firebase-project-id"
   export PORT="3000"
   ```

5. **Create Firestore indexes**
   - Visit the Firestore console and create required composite indexes
   - Or run: `firebase deploy --only firestore:indexes`

6. **Seed the database (optional)**
   ```bash
   go run scripts/seed.go
   ```

7. **Run locally**
   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:3000` (or your specified PORT)

### Firebase Setup

1. **Create a Firebase project** at [Firebase Console](https://console.firebase.google.com/)

2. **Enable required services**
   - Enable Firestore Database (Native mode)
   - Enable Authentication and configure Google provider
   - Generate a service account key for server-side access

3. **Configure Firestore**
   - Set up security rules from `firestore.rules`
   - Create composite indexes (see Firestore Index Requirements section)
   - Deploy rules: `firebase deploy --only firestore:rules`

4. **Set up Authentication**
   - Enable Google Sign-In in Authentication > Sign-in methods
   - Configure authorized domains for your deployment

5. **Download service account key**
   - Go to Project Settings > Service Accounts
   - Generate new private key and download the JSON file

## Deployment

### Deployment Options

This website can be deployed to various platforms:

#### 1. Google Cloud Platform (Recommended)

**Using Google Cloud Run:**
```bash
# Build and deploy to Cloud Run
gcloud builds submit --tag gcr.io/[PROJECT-ID]/my-website
gcloud run deploy --image gcr.io/[PROJECT-ID]/my-website --platform managed
```

**Using App Engine:**
```bash
# Deploy to App Engine
gcloud app deploy
```

#### 2. Firebase Hosting (Static + Functions)

```bash
# Deploy everything
firebase deploy

# Deploy specific services
firebase deploy --only functions
firebase deploy --only hosting
firebase deploy --only firestore:rules,firestore:indexes
```

#### 3. Docker Deployment

```bash
# Build Docker image
docker build -t my-website .

# Run locally
docker run -p 3000:3000 \
  -e GOOGLE_APPLICATION_CREDENTIALS=/app/credentials.json \
  -e FIREBASE_PROJECT_ID=your-project-id \
  -v /path/to/credentials.json:/app/credentials.json \
  my-website
```

#### 4. Traditional VPS/Server

```bash
# Build the application
go build -o website main.go

# Run with systemd or supervisor
./website
```

## Configuration

### Environment Variables

- `FIREBASE_PROJECT_ID`: Your Firebase project ID (required)
- `GOOGLE_APPLICATION_CREDENTIALS`: Path to Firebase service account key JSON file (required for server)
- `PORT`: Server port (default: 8080)
- `ENVIRONMENT`: Environment mode (development/production)
- `GIN_MODE`: Gin framework mode (debug/release)

### Firestore Index Requirements

The application requires the following composite indexes in Firestore:

1. **Sections Collection:**
   - Fields: `page_slug` (Ascending), `order` (Ascending)
   - Required for: Fetching ordered sections for a page

**Create indexes automatically:**
```bash
firebase deploy --only firestore:indexes
```

**Or create manually:** Visit the URLs provided in error messages when running queries locally.

### Admin Access Configuration

To grant admin access to users:

1. Users must sign in with Google OAuth at `/admin/login`
2. Update the `IsUserAdmin` function in `internal/services/content.go` to check admin permissions
3. Consider storing admin emails in Firestore or environment variables

## Content Management

### Dynamic CMS System

The website uses a generic CMS system with Pages and Sections:

#### Pages
Pages represent top-level content areas (e.g., "home", "blog", "about"):

```json
{
  "slug": "home",
  "title": "Home Page",
  "description": "Main landing page",
  "visible": true,
  "metadata": {
    "keywords": "portfolio, developer",
    "description": "SEO description"
  }
}
```

#### Sections
Sections are content blocks within pages, ordered and flexible:

```json
{
  "page_slug": "home",
  "slug": "hero",
  "type": "hero",
  "title": "Welcome Section",
  "subtitle": "Subtitle text",
  "content": "Main content text",
  "order": 1,
  "visible": true,
  "data": {
    "background_image": "/static/hero.jpg",
    "cta_text": "Learn More",
    "skills": ["Go", "JavaScript", "React"]
  }
}
```

### Admin Interface

Access the admin interface at `/admin/login`:

- **Dashboard**: Overview of all pages and sections
- **Page Management**: Create, edit, and delete pages
- **Section Management**: Add, reorder, and modify sections
- **Drag & Drop**: Reorder sections with HTMX-powered interface
- **Inline Editing**: Quick edits without page refresh

### API Endpoints

- `GET /api/page/{pageSlug}` - Get page information
- `GET /api/page/{pageSlug}/sections` - Get all sections for a page
- `GET /api/page/{pageSlug}/sections/{sectionSlug}` - Get specific section

### Seeding Data

To populate the database with initial content:

```bash
# Run the seed script
GOOGLE_APPLICATION_CREDENTIALS=./credentials.json FIREBASE_PROJECT_ID=your-project-id go run scripts/seed.go
```

This creates:
- Home page with 6 sample sections
- Hero, About, Technologies, Experience, Projects, and Contact sections
- Structured data for each section type

## Customization

### Styling

The website uses TailwindCSS with custom components:

- **Base Template**: `web/templates/base.html` contains the main layout
- **Component Templates**: Individual templates for each section type
- **Admin Styles**: HTMX-powered admin interface with TailwindCSS
- **Custom CSS**: Additional styles in `/static/css/`

### Templates

Template structure:
```
web/templates/
├── base.html           # Main layout
├── home.html          # Home page
├── error.html         # Error pages
└── admin/
    ├── dashboard.html # Admin dashboard
    ├── page-form.html # Page creation/editing
    └── section-form.html # Section creation/editing
```

### Adding New Section Types

1. **Define the template component**:
   ```html
   {{define "my-section"}}
   <div class="my-section">
     <h2>{{.Title}}</h2>
     <p>{{.Content}}</p>
     <!-- Access custom data: {{.Data.custom_field}} -->
   </div>
   {{end}}
   ```

2. **Update the section rendering logic** in handlers

3. **Add admin form fields** for the new section type

### HTMX Integration

HTMX enables dynamic interactions:

- **Section Reordering**: Drag & drop with real-time updates
- **Inline Editing**: Edit content without page refresh
- **Dynamic Loading**: Load admin forms and content dynamically
- **Form Submissions**: Handle forms with partial page updates

Example HTMX usage:
```html
<button hx-post="/admin/sections/123/toggle-visibility" 
        hx-target="#section-123"
        hx-swap="outerHTML">
  Toggle Visibility
</button>
```

## Troubleshooting

### Common Issues

**1. Firestore Index Errors**
```
The query requires an index. You can create it here: [URL]
```
*Solution*: Visit the provided URL or run `firebase deploy --only firestore:indexes`

**2. Authentication Errors**
```
credentials: could not find default credentials
```
*Solution*: Set `GOOGLE_APPLICATION_CREDENTIALS` environment variable to your service account key path

**3. Port Already in Use**
```
listen tcp :3000: bind: address already in use
```
*Solution*: 
```bash
# Kill process using the port
lsof -ti:3000 | xargs kill -9
# Or use a different port
PORT=3001 go run main.go
```

**4. Template Errors**
```
template: function "json" not defined
```
*Solution*: Ensure all required template functions are registered in main.go

**5. Admin Access Issues**
- Verify Google OAuth is configured in Firebase Console
- Check that your email is configured as admin in the `IsUserAdmin` function
- Ensure Firebase Auth is enabled and properly configured

### Debug Mode

Run in debug mode for detailed logging:
```bash
GIN_MODE=debug go run main.go
```

### Health Check

Monitor application health:
```bash
curl http://localhost:3000/health
```

## Performance

- **Static Assets**: Served efficiently with proper caching headers
- **Database Queries**: Optimized Firestore queries with proper indexing
- **Template Caching**: Go templates are compiled and cached
- **HTMX**: Minimal JavaScript for dynamic interactions
- **Compression**: Gzip compression for text responses

## Security

- **Firestore Rules**: Secure read/write rules in `firestore.rules`
- **OAuth Authentication**: Google OAuth for admin access
- **Security Headers**: CORS, XSS protection, and security middleware
- **Input Validation**: Proper validation for all form inputs and API requests
- **HTTPS**: Force HTTPS in production environments

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices and formatting (`go fmt`, `go vet`)
- Write tests for new functionality
- Update documentation for new features
- Use meaningful commit messages
- Ensure all tests pass before submitting PR

## License

This project is licensed under the MIT License - see the LICENSE file for details.
