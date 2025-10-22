# API Documentation

This document describes the API endpoints available in the website.

## Base URL

- **Local Development**: `http://localhost:3000`
- **Production**: `https://yourdomain.com`

## Authentication

### Admin Routes

Admin routes require Google OAuth authentication. Users must sign in through the admin interface.

**Admin Login Flow:**
1. Visit `/admin/login`
2. Click "Sign in with Google"
3. Authenticate with Google
4. System checks if user is authorized admin
5. Redirect to admin dashboard

## Public API Endpoints

### Health Check

Check if the application is running.

**Endpoint:** `GET /health`

**Response:**
```json
{
  "status": "ok",
  "timestamp": "2023-12-15T10:30:00Z"
}
```

### Pages

#### Get Page Information

Retrieve information about a specific page.

**Endpoint:** `GET /api/page/{pageSlug}`

**Parameters:**
- `pageSlug` (string): The page identifier (e.g., "home", "blog", "about")

**Response:**
```json
{
  "id": "page-id-123",
  "slug": "home",
  "title": "Home Page",
  "description": "Main landing page",
  "metadata": {
    "keywords": "portfolio, developer",
    "description": "SEO description",
    "og_image": "/static/og-image.jpg",
    "canonical": "https://yourdomain.com"
  },
  "visible": true,
  "created_at": "2023-12-15T10:00:00Z",
  "updated_at": "2023-12-15T10:30:00Z",
  "updated_by": "admin@example.com"
}
```

**Error Responses:**
- `404 Not Found`: Page does not exist
```json
{
  "error": "Page not found"
}
```

### Sections

#### Get Page Sections

Retrieve all sections for a specific page, ordered by their display order.

**Endpoint:** `GET /api/page/{pageSlug}/sections`

**Parameters:**
- `pageSlug` (string): The page identifier

**Query Parameters:**
- `visible` (boolean, optional): Filter by visibility (default: true)

**Response:**
```json
[
  {
    "id": "section-id-456",
    "page_slug": "home",
    "slug": "hero",
    "type": "hero",
    "title": "Welcome to My Website",
    "subtitle": "Full Stack Developer",
    "content": "I create amazing web applications...",
    "data": {
      "background_image": "/static/hero-bg.jpg",
      "cta_text": "View My Work",
      "cta_link": "#projects"
    },
    "order": 1,
    "visible": true,
    "template": "hero",
    "created_at": "2023-12-15T10:00:00Z",
    "updated_at": "2023-12-15T10:30:00Z",
    "updated_by": "admin@example.com"
  },
  {
    "id": "section-id-789",
    "page_slug": "home",
    "slug": "about",
    "type": "about",
    "title": "About Me",
    "subtitle": "Get to know me better",
    "content": "I'm a passionate developer...",
    "data": {
      "image": "/static/profile.jpg",
      "skills": ["Go", "JavaScript", "React"]
    },
    "order": 2,
    "visible": true,
    "created_at": "2023-12-15T10:00:00Z",
    "updated_at": "2023-12-15T10:30:00Z"
  }
]
```

**Error Responses:**
- `500 Internal Server Error`: Failed to fetch sections
```json
{
  "error": "Failed to fetch sections"
}
```

#### Get Specific Section

Retrieve a specific section by its ID.

**Endpoint:** `GET /api/page/{pageSlug}/sections/{sectionSlug}`

**Parameters:**
- `pageSlug` (string): The page identifier
- `sectionSlug` (string): The section identifier

**Response:**
```json
{
  "id": "section-id-456",
  "page_slug": "home",
  "slug": "hero",
  "type": "hero",
  "title": "Welcome to My Website",
  "subtitle": "Full Stack Developer",
  "content": "I create amazing web applications...",
  "data": {
    "background_image": "/static/hero-bg.jpg",
    "cta_text": "View My Work",
    "cta_link": "#projects"
  },
  "order": 1,
  "visible": true,
  "created_at": "2023-12-15T10:00:00Z",
  "updated_at": "2023-12-15T10:30:00Z"
}
```

**Error Responses:**
- `404 Not Found`: Section does not exist
```json
{
  "error": "Section not found"
}
```

## Admin API Endpoints

**Note:** All admin endpoints require authentication via Google OAuth.

### Authentication

#### Admin Login Page

Display the admin login page.

**Endpoint:** `GET /admin/login`

**Response:** HTML page with Google OAuth login button

#### Authenticate User

Process Google OAuth authentication.

**Endpoint:** `POST /admin/auth`

**Headers:**
- `Content-Type: application/json`

**Request Body:**
```json
{
  "id_token": "google-oauth-id-token"
}
```

**Response:**
- `302 Redirect`: Redirect to admin dashboard on success
- `401 Unauthorized`: Invalid or expired token
- `403 Forbidden`: User is not authorized as admin

### Dashboard

#### Admin Dashboard

Display the main admin dashboard.

**Endpoint:** `GET /admin/dashboard`

**Response:** HTML page with dashboard interface

### Page Management

#### Get All Pages

Retrieve all pages for admin management.

**Endpoint:** `GET /admin/pages`

**Response:**
```json
[
  {
    "id": "page-id-123",
    "slug": "home",
    "title": "Home Page",
    "description": "Main landing page",
    "visible": true,
    "created_at": "2023-12-15T10:00:00Z",
    "updated_at": "2023-12-15T10:30:00Z"
  }
]
```

#### Create Page

Create a new page.

**Endpoint:** `POST /admin/pages`

**Headers:**
- `Content-Type: application/json`

**Request Body:**
```json
{
  "slug": "about",
  "title": "About Page",
  "description": "About me page",
  "metadata": {
    "keywords": "about, developer",
    "description": "Learn more about me"
  },
  "visible": true
}
```

**Response:**
```json
{
  "id": "new-page-id",
  "slug": "about",
  "title": "About Page",
  "description": "About me page",
  "visible": true,
  "created_at": "2023-12-15T10:30:00Z",
  "updated_at": "2023-12-15T10:30:00Z"
}
```

#### Update Page

Update an existing page.

**Endpoint:** `PUT /admin/pages/{id}`

**Parameters:**
- `id` (string): The page ID

**Request Body:** Same as create page

**Response:** Updated page object

#### Delete Page

Delete a page and all its sections.

**Endpoint:** `DELETE /admin/pages/{id}`

**Parameters:**
- `id` (string): The page ID

**Response:**
```json
{
  "message": "Page deleted successfully"
}
```

### Section Management

#### Get Page Sections (Admin)

Get sections for admin interface with additional metadata.

**Endpoint:** `GET /admin/pages/{pageSlug}/sections`

**Response:** HTML fragment with section list for admin interface

#### Create Section

Create a new section within a page.

**Endpoint:** `POST /admin/pages/{pageSlug}/sections`

**Request Body:**
```json
{
  "slug": "new-section",
  "type": "about",
  "title": "New Section",
  "subtitle": "Section subtitle",
  "content": "Section content...",
  "data": {
    "custom_field": "custom_value"
  },
  "order": 3,
  "visible": true,
  "template": "about"
}
```

**Response:** Created section object

#### Update Section

Update an existing section.

**Endpoint:** `PUT /admin/sections/{id}`

**Parameters:**
- `id` (string): The section ID

**Request Body:** Same as create section

**Response:** Updated section object

#### Delete Section

Delete a section.

**Endpoint:** `DELETE /admin/sections/{id}`

**Parameters:**
- `id` (string): The section ID

**Response:**
```json
{
  "message": "Section deleted successfully"
}
```

#### Reorder Sections

Change the order of sections by swapping two sections.

**Endpoint:** `PUT /admin/sections/{id}/reorder`

**Parameters:**
- `id` (string): The source section ID

**Request Body:**
```json
{
  "target_id": "target-section-id"
}
```

**Response:**
```json
{
  "message": "Sections reordered successfully"
}
```

#### Toggle Section Visibility

Toggle the visibility of a section.

**Endpoint:** `PUT /admin/sections/{id}/toggle-visibility`

**Parameters:**
- `id` (string): The section ID

**Response:** HTML fragment with updated section

#### Section Form Rendering

Get HTML form for creating new sections.

**Endpoint:** `GET /admin/sections/new`

**Query Parameters:**
- `page_slug` (string): The page to create section for
- `type` (string, optional): Pre-select section type

**Response:** HTML form fragment

#### Section Edit Form

Get HTML form for editing existing sections.

**Endpoint:** `GET /admin/sections/{id}/edit`

**Parameters:**
- `id` (string): The section ID

**Response:** HTML form fragment with pre-filled data

## Contact and Newsletter

### Contact Form

Submit a contact form message.

**Endpoint:** `POST /contact`

**Headers:**
- `Content-Type: application/json`

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "subject": "Project Inquiry",
  "message": "I'd like to discuss a project..."
}
```

**Response:**
```json
{
  "message": "Thank you for your message! I'll get back to you soon."
}
```

**Error Responses:**
- `400 Bad Request`: Invalid form data
```json
{
  "error": "Invalid form data"
}
```

### Newsletter Subscription

Subscribe to the newsletter.

**Endpoint:** `POST /newsletter`

**Headers:**
- `Content-Type: application/json`

**Request Body:**
```json
{
  "email": "subscriber@example.com"
}
```

**Response:**
```json
{
  "message": "Thank you for subscribing to my newsletter!"
}
```

**Error Responses:**
- `400 Bad Request`: Invalid email address
```json
{
  "error": "Invalid email address"
}
```

## Section Types and Data Structures

### Hero Section

**Type:** `hero`

**Data Structure:**
```json
{
  "background_image": "/static/hero-bg.jpg",
  "cta_text": "Get Started",
  "cta_link": "#contact",
  "video_background": "/static/hero-video.mp4",
  "overlay_opacity": 0.5
}
```

### About Section

**Type:** `about`

**Data Structure:**
```json
{
  "image": "/static/profile.jpg",
  "skills": ["Go", "JavaScript", "React", "Docker"],
  "resume_link": "/static/resume.pdf",
  "social_links": {
    "github": "https://github.com/username",
    "linkedin": "https://linkedin.com/in/username"
  }
}
```

### Technologies Section

**Type:** `technologies`

**Data Structure:**
```json
{
  "categories": {
    "Backend": [
      {
        "name": "Go",
        "level": 90,
        "icon": "devicon-go-original-wordmark"
      }
    ],
    "Frontend": [
      {
        "name": "React",
        "level": 85,
        "icon": "devicon-react-original"
      }
    ]
  }
}
```

### Experience Section

**Type:** `experience`

**Data Structure:**
```json
{
  "experiences": [
    {
      "company": "Tech Corp",
      "position": "Senior Developer",
      "duration": "2022 - Present",
      "description": "Led development of...",
      "achievements": [
        "Reduced latency by 40%",
        "Mentored junior developers"
      ],
      "technologies": ["Go", "React", "AWS"]
    }
  ]
}
```

### Projects Section

**Type:** `projects`

**Data Structure:**
```json
{
  "projects": [
    {
      "title": "E-commerce Platform",
      "description": "Full-stack solution...",
      "image": "/static/project1.jpg",
      "technologies": ["Go", "React", "PostgreSQL"],
      "github": "https://github.com/user/project",
      "demo": "https://demo.example.com",
      "featured": true
    }
  ]
}
```

### Contact Section

**Type:** `contact`

**Data Structure:**
```json
{
  "email": "contact@example.com",
  "phone": "+1 (555) 123-4567",
  "location": "City, Country",
  "social": {
    "github": "https://github.com/username",
    "linkedin": "https://linkedin.com/in/username",
    "twitter": "https://twitter.com/username"
  },
  "availability": "Available for freelance work",
  "contact_form_enabled": true
}
```

## Error Handling

### Standard Error Format

All API errors follow this format:

```json
{
  "error": "Error message describing what went wrong",
  "code": "ERROR_CODE",
  "details": {
    "field": "Additional details if applicable"
  }
}
```

### HTTP Status Codes

- `200 OK`: Request successful
- `201 Created`: Resource created successfully
- `400 Bad Request`: Invalid request data
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Access denied
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

### Common Error Codes

- `INVALID_REQUEST`: Request body is malformed
- `VALIDATION_ERROR`: Request data fails validation
- `NOT_FOUND`: Requested resource doesn't exist
- `UNAUTHORIZED`: Authentication token missing or invalid
- `FORBIDDEN`: User doesn't have permission
- `SERVER_ERROR`: Internal server error

## Rate Limiting

API endpoints may be rate limited to prevent abuse:

- **Public API**: 100 requests per minute per IP
- **Admin API**: 1000 requests per minute per authenticated user

Rate limit headers are included in responses:
```
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1640995200
```

## CORS Policy

Cross-Origin Resource Sharing (CORS) is configured to allow:

- **Origins**: `*` (development), specific domains (production)
- **Methods**: `GET, POST, PUT, DELETE, OPTIONS`
- **Headers**: `Content-Type, Authorization, X-Requested-With`
- **Credentials**: Allowed for admin endpoints

## SDK Examples

### JavaScript/Node.js

```javascript
// Fetch page sections
async function getPageSections(pageSlug) {
  const response = await fetch(`/api/page/${pageSlug}/sections`);
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
  return await response.json();
}

// Create new section (admin)
async function createSection(pageSlug, sectionData) {
  const response = await fetch(`/admin/pages/${pageSlug}/sections`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    body: JSON.stringify(sectionData)
  });
  return await response.json();
}
```

### Python

```python
import requests

# Fetch page sections
def get_page_sections(page_slug):
    response = requests.get(f'/api/page/{page_slug}/sections')
    response.raise_for_status()
    return response.json()

# Submit contact form
def submit_contact_form(form_data):
    response = requests.post('/contact', json=form_data)
    response.raise_for_status()
    return response.json()
```

### cURL Examples

```bash
# Get page sections
curl -X GET "http://localhost:3000/api/page/home/sections"

# Submit contact form
curl -X POST "http://localhost:3000/contact" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "subject": "Test",
    "message": "Hello!"
  }'

# Create section (admin)
curl -X POST "http://localhost:3000/admin/pages/home/sections" \
  -H "Content-Type: application/json" \
  -b "session-cookie" \
  -d '{
    "slug": "new-section",
    "type": "about",
    "title": "New Section",
    "content": "Content here...",
    "order": 3,
    "visible": true
  }'
```