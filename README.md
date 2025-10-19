# Warren Clough Personal Website

A modern, responsive personal website built with Go, HTMX, TailwindCSS, and Firebase. This project is designed to run as a serverless solution on Google Cloud Platform using Firebase Cloud Functions.

## Features

- **Modern Tech Stack**: Go backend with HTMX for dynamic interactions and TailwindCSS for styling
- **Serverless Architecture**: Deployed as Firebase Cloud Functions for scalability and cost-effectiveness
- **Firebase Integration**: Uses Firestore for database and Firebase Hosting for static assets
- **Responsive Design**: Clean, modern design that works on all devices
- **Performance Optimized**: Fast loading times with optimized assets and caching
- **SEO Friendly**: Proper meta tags and structured content

## Project Structure

```
├── internal/
│   ├── config/          # Configuration management
│   ├── handlers/        # HTTP request handlers
│   ├── middleware/      # HTTP middleware
│   ├── models/          # Data models
│   └── services/        # Business logic services
├── web/
│   ├── static/          # Static assets (CSS, JS, images)
│   └── templates/       # HTML templates
├── firebase.json        # Firebase configuration
├── firestore.rules      # Firestore security rules
├── firestore.indexes.json # Firestore indexes
├── functions.go         # Firebase Cloud Function entry point
├── main.go             # Local development server
└── go.mod              # Go module dependencies
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
   cd warrenclough-website
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   ```bash
   cp env.example .env
   # Edit .env with your Firebase project details
   ```

4. **Set up Firebase**
   ```bash
   firebase login
   firebase use --add
   ```

5. **Run locally**
   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8080`

### Firebase Setup

1. **Create a Firebase project** at [Firebase Console](https://console.firebase.google.com/)

2. **Enable Firestore Database**
   - Go to Firestore Database in your Firebase console
   - Create database in production mode
   - Deploy the security rules: `firebase deploy --only firestore:rules`

3. **Set up Firebase Hosting**
   - Enable Firebase Hosting in your project
   - Deploy: `firebase deploy --only hosting`

4. **Deploy Cloud Functions**
   ```bash
   firebase deploy --only functions
   ```

## Deployment

### Deploy Everything
```bash
firebase deploy
```

### Deploy Specific Services
```bash
# Deploy only functions
firebase deploy --only functions

# Deploy only hosting
firebase deploy --only hosting

# Deploy only Firestore rules
firebase deploy --only firestore:rules
```

## Configuration

### Environment Variables

- `FIREBASE_PROJECT_ID`: Your Firebase project ID
- `GOOGLE_APPLICATION_CREDENTIALS`: Path to service account key (for local development)
- `PORT`: Server port (default: 8080)
- `ENVIRONMENT`: Environment (development/production)

### Firebase Configuration

The `firebase.json` file contains the configuration for:
- Cloud Functions
- Firebase Hosting
- Firestore rules and indexes

## Content Management

### Adding Blog Posts

Blog posts are stored in the `blog_posts` Firestore collection with the following structure:

```json
{
  "title": "Post Title",
  "slug": "post-slug",
  "content": "Post content in HTML",
  "excerpt": "Short description",
  "author": "Warren Clough",
  "published_at": "2023-12-15T10:00:00Z",
  "updated_at": "2023-12-15T10:00:00Z",
  "tags": ["go", "web-development"],
  "featured": true,
  "image_url": "https://example.com/image.jpg"
}
```

### Adding Portfolio Projects

Portfolio projects are stored in the `portfolio_projects` collection:

```json
{
  "title": "Project Title",
  "description": "Short description",
  "long_description": "Detailed description",
  "image_url": "https://example.com/project.jpg",
  "live_url": "https://project.com",
  "github_url": "https://github.com/user/project",
  "technologies": ["Go", "React", "Firebase"],
  "featured": true,
  "created_at": "2023-12-15T10:00:00Z",
  "updated_at": "2023-12-15T10:00:00Z"
}
```

## Customization

### Styling

The website uses TailwindCSS with a custom color scheme defined in the base template. You can customize colors by modifying the Tailwind config in `web/templates/base.html`.

### Templates

HTML templates are located in `web/templates/`. The base template provides the common layout, and individual pages extend it with their specific content.

### HTMX Integration

HTMX is used for dynamic content loading without full page refreshes. Examples include:
- Loading more blog posts
- Loading more portfolio projects
- Contact form submissions
- Newsletter subscriptions

## Performance

- **Static Assets**: Served through Firebase Hosting with CDN
- **Caching**: Proper cache headers for static assets
- **Compression**: Automatic compression for text assets
- **Image Optimization**: Responsive images with proper sizing

## Security

- **Firestore Rules**: Secure read/write rules for different collections
- **Security Headers**: CORS, XSS protection, and other security headers
- **Input Validation**: Proper validation for all form inputs
- **HTTPS**: All traffic served over HTTPS

## Monitoring

- **Health Check**: `/health` endpoint for monitoring
- **Error Handling**: Proper error pages and logging
- **Analytics**: Ready for Google Analytics integration

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test locally
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

For questions or support, please contact [warren@warrenclough.com](mailto:warren@warrenclough.com)
