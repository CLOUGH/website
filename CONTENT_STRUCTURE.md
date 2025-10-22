# Content Structure & Database Schema

## Overview
This document describes how the content from your website has been structured in Firestore and how it maps to the dynamic CMS system.

## Database Collections

### Pages Collection
- **home**: Main landing page
  - Title: "Warren Clough - Senior Software Engineer"
  - Description: Enterprise-grade solutions specialist
  - SEO metadata included

### Sections Collection
All sections belong to the "home" page and are ordered sequentially.

---

## Section Breakdown

### 1. Hero Section (`hero`)
**Order:** 1  
**Type:** `hero`  
**Title:** "Hi, I'm Warren Clough"

**Data Structure:**
```json
{
  "name": "Warren Clough",
  "title": "Senior Software Engineer",
  "tagline": "I build exceptional digital experiences...",
  "github_url": "https://github.com/clough",
  "linkedin_url": "https://linkedin.com/in/warrenclough"
}
```

**Content from Page:**
- Name and title display
- Professional tagline
- Social links

---

### 2. About Section (`about`)
**Order:** 2  
**Type:** `about`  
**Title:** "About Me"

**Data Structure:**
```json
{
  "paragraphs": [
    "I'm a Full-Stack Software Engineer...",
    "With over 10 years of professional experience...",
    "Throughout my career, I've had the privilege...",
    "Beyond coding, I'm passionate about mentoring..."
  ]
}
```

**Content from Page:**
- 4 paragraphs describing background, experience, team collaboration, and personal interests
- Journey in software development
- Professional expertise in Java/Spring Boot and Laravel/PHP
- Passion for Go and continuous learning

---

### 3. Technologies Section (`technologies`)
**Order:** 3  
**Type:** `technologies`  
**Title:** "Technologies I Work With"  
**Subtitle:** "I believe in choosing the right tool for the job..."

**Data Structure:**
```json
{
  "categories": {
    "Languages & Frameworks": [
      "Java & Spring Boot",
      "PHP & Laravel",
      "Go (Golang)",
      // ... 13 total items
    ],
    "Cloud & DevOps": [
      "Kubernetes",
      "Docker",
      "Google Cloud Platform",
      // ... 8 total items
    ],
    "Databases & Tools": [
      "PostgreSQL",
      "Oracle SQL",
      "MySQL",
      // ... 14 total items
    ]
  }
}
```

**Content from Page:**
- 3 categories of technologies
- 35+ different technologies listed
- Organized by type (Languages, Cloud, Databases)

---

### 4. Experience Section (`experience`)
**Order:** 4  
**Type:** `experience`  
**Title:** "Experience"

**Data Structure:**
```json
{
  "experiences": [
    {
      "company": "National Commercial Bank of Jamaica",
      "position": "Senior Developer Analyst",
      "location": "Jamaica",
      "employment": "Full-time",
      "start_date": "2018",
      "end_date": "Present",
      "description": [
        "Technical Lead for NCB's Consumer Loans & Credit Card team...",
        "Design and build Java microservices with Kubernetes...",
        // ... 6 bullet points total
      ],
      "technologies": ["Java", "React", "Kubernetes", "PostgreSQL", "Oracle", "GCP", "Microservices"]
    },
    // ... 3 total positions
  ]
}
```

**Content from Page:**
- 3 work positions (NCB, Epic Technologies x2)
- Date ranges (2018-Present, 2014-2018, 2010-2014)
- Detailed bullet points for each role
- Technology stacks for each position

---

### 5. Projects Section (`projects`)
**Order:** 5  
**Type:** `projects`  
**Title:** "Projects"

**Data Structure:**
```json
{
  "projects": [
    {
      "id": "poe-recipe",
      "title": "Path of Exile Recipe Helper",
      "description": "Fun gaming side project...",
      "full_description": "Desktop application that integrates...",
      "technologies": ["Electron", "Angular", "TypeScript", "REST API"],
      "screenshots": ["https://github.com/..."],
      "github_url": "https://github.com/clough/poe-recipe",
      "featured": true,
      "order": 1
    },
    // ... 5 total projects
  ],
  "confidential_notice": {
    "title": "Confidential Professional Work",
    "content": "The majority of my professional projects..."
  }
}
```

**Projects Included:**
1. **Path of Exile Recipe Helper** (Electron, Angular)
   - Gaming side project
   - GitHub link
   
2. **Go Trader - Jamaica Stock Exchange** (Go, HTMX, D3.js)
   - Stock market visualization
   - 3 screenshots
   - GitHub link

3. **HRMNext Recruitment Platform** (AngularJS, Laravel, 4D)
   - SaaS recruitment solution
   - Professional work

4. **NCB Credit Card Application Portal** (React, Spring Boot, Java, Kubernetes)
   - Enterprise banking portal
   - Live URL: https://creditcard.jncb.com/

5. **OpenGL Snake Game** (OpenGL, C++, GLSL)
   - University project
   - YouTube video link
   - GitHub link

**Additional Content:**
- Confidential work disclaimer
- NDA notice for enterprise projects

---

### 6. Contact Section (`contact`)
**Order:** 6  
**Type:** `contact`  
**Title:** "Let's Build Something Amazing"

**Data Structure:**
```json
{
  "intro": [
    "I'm always excited to connect with fellow developers...",
    "Currently open to new opportunities..."
  ],
  "call_to_action": "Drop me an email at clough.warren@gmail.com...",
  "email": "clough.warren@gmail.com",
  "github_url": "https://github.com/clough",
  "linkedin_url": "https://linkedin.com/in/warrenclough",
  "buttons": [
    {"text": "View Resume", "url": "/resume", "type": "primary"},
    {"text": "Send Message", "url": "/contact", "type": "secondary"}
  ]
}
```

**Content from Page:**
- Professional availability statement
- Interest areas (Java/Spring Boot, Laravel, Angular, Go)
- Contact information
- Call-to-action buttons

---

## How to Use This Data

### Editing Content via Admin Dashboard

1. **Navigate to:** `/admin/dashboard`
2. **Login** with your Google account (clough.warren@gmail.com)
3. **Select a section** to edit
4. **Modify the JSON data** in the expanded editor

### Data Access in Templates

The sections are available in your Go templates through the content service:

```go
sections, err := contentService.GetPageSections(ctx, "home")
```

Each section has:
- `Title`, `Subtitle`, `Content` (text fields)
- `Data` (flexible JSON for complex structures)
- `Order`, `Visible` (display control)
- `Type` (section type identifier)

### Section Types Supported

- `hero` - Hero/landing section
- `about` - About me biographical content
- `technologies` - Tech stack listings
- `experience` - Work history
- `projects` - Portfolio projects
- `contact` - Contact information and CTAs

---

## Rendering Dynamic Content

### Example: Rendering Projects

```go
// In your handler
sections, _ := contentService.GetPageSections(ctx, "home")

for _, section := range sections {
    if section.Type == "projects" {
        if projectsData, ok := section.Data["projects"].([]interface{}); ok {
            for _, proj := range projectsData {
                project := proj.(map[string]interface{})
                // Access: project["title"], project["technologies"], etc.
            }
        }
    }
}
```

### Example: Rendering Technologies

```go
for _, section := range sections {
    if section.Type == "technologies" {
        if categories, ok := section.Data["categories"].(map[string]interface{}); ok {
            for categoryName, techs := range categories {
                techList := techs.([]interface{})
                // Render each technology in the category
            }
        }
    }
}
```

---

## Seed Script

The content was populated using: `/scripts/seed-full-content.go`

**To re-seed:**
```bash
GOOGLE_APPLICATION_CREDENTIALS=credentials/firebase-service-account.json \
  go run scripts/seed-full-content.go
```

**To update specific sections:**
Use the admin dashboard at `/admin/dashboard`

---

## Next Steps

1. ✅ Database seeded with all current content
2. 🔄 Update home page template to pull from Firestore
3. 🔄 Create dynamic section rendering based on `type`
4. ✅ Admin dashboard ready for content editing
5. 🔄 Add new section types as needed

---

## Content Preservation

All original content from your static page has been preserved:
- ✅ Hero text and branding
- ✅ About me (4 paragraphs)
- ✅ 35+ technologies across 3 categories
- ✅ 3 work experiences with detailed descriptions
- ✅ 5 projects with full details and media
- ✅ Contact information and CTAs
- ✅ Professional disclaimers for confidential work
