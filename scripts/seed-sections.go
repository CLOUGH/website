package main

import (
	"context"
	"log"

	"warrenclough.com/internal/config"
	"warrenclough.com/internal/models"
	"warrenclough.com/internal/services"
)

func seedData() {
	ctx := context.Background()

	// Load configuration
	cfg := config.Load()

	// Initialize Firebase service
	firebaseService, err := services.NewFirebaseService(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase service: %v", err)
	}
	defer firebaseService.Close()

	// Initialize content service
	contentService := services.NewContentService(firebaseService)

	// Create Home page
	homePage := &models.Page{
		Slug:        "home",
		Title:       "Warren Clough - Software Engineer",
		Description: "Personal portfolio and blog of Warren Clough, Full-Stack Software Engineer",
		Visible:     true,
	}

	err = contentService.CreatePage(ctx, homePage)
	if err != nil {
		log.Printf("Failed to create home page: %v", err)
	} else {
		log.Println("Created home page")
	}

	// Create sample sections for home page
	sections := []models.Section{
		{
			Slug:     "hero",
			Type:     "hero",
			Title:    "Warren Clough",
			Subtitle: "Full Stack Software Engineer",
			Content:  "Passionate about building robust, scalable applications with modern technologies.",
			Order:    1,
			Visible:  true,
			PageSlug: "home",
			Data: map[string]interface{}{
				"background_image": "/static/hero-bg.jpg",
				"cta_text":         "View My Work",
				"cta_link":         "#projects",
			},
		},
		{
			Slug:     "about",
			Type:     "about",
			Title:    "About Me",
			Subtitle: "Get to know me better",
			Content:  "I'm a passionate software engineer with expertise in full-stack development, cloud technologies, and modern frameworks. I love solving complex problems and building applications that make a difference.",
			Order:    2,
			Visible:  true,
			PageSlug: "home",
			Data: map[string]interface{}{
				"image":  "/static/profile.jpg",
				"skills": []string{"Go", "JavaScript", "Python", "React", "Node.js", "Docker", "AWS"},
			},
		},
		{
			Slug:     "technologies",
			Type:     "technologies",
			Title:    "Technologies",
			Subtitle: "Tools and technologies I work with",
			Content:  "I have experience with a wide range of technologies and frameworks.",
			Order:    3,
			Visible:  true,
			PageSlug: "home",
			Data: map[string]interface{}{
				"categories": map[string][]map[string]interface{}{
					"Backend": {
						{"name": "Go", "level": 90, "icon": "devicon-go-original-wordmark"},
						{"name": "Node.js", "level": 85, "icon": "devicon-nodejs-plain"},
						{"name": "Python", "level": 80, "icon": "devicon-python-plain"},
					},
					"Frontend": {
						{"name": "React", "level": 85, "icon": "devicon-react-original"},
						{"name": "JavaScript", "level": 90, "icon": "devicon-javascript-plain"},
						{"name": "TypeScript", "level": 80, "icon": "devicon-typescript-plain"},
					},
					"Cloud": {
						{"name": "AWS", "level": 75, "icon": "devicon-amazonwebservices-original"},
						{"name": "Docker", "level": 80, "icon": "devicon-docker-plain"},
						{"name": "Firebase", "level": 70, "icon": "devicon-firebase-plain"},
					},
				},
			},
		},
		{
			Slug:     "experience",
			Type:     "experience",
			Title:    "Work Experience",
			Subtitle: "My professional journey",
			Content:  "Here's an overview of my professional experience and key achievements.",
			Order:    4,
			Visible:  true,
			PageSlug: "home",
			Data: map[string]interface{}{
				"experiences": []map[string]interface{}{
					{
						"company":     "Tech Corp",
						"position":    "Senior Software Engineer",
						"duration":    "2022 - Present",
						"description": "Led development of microservices architecture serving 1M+ users",
						"achievements": []string{
							"Reduced system latency by 40%",
							"Implemented CI/CD pipeline",
							"Mentored 3 junior developers",
						},
					},
					{
						"company":     "StartupXYZ",
						"position":    "Full Stack Developer",
						"duration":    "2020 - 2022",
						"description": "Built scalable web applications from ground up",
						"achievements": []string{
							"Developed MVP that acquired 10K users",
							"Implemented real-time chat system",
							"Optimized database queries by 60%",
						},
					},
				},
			},
		},
		{
			Slug:     "projects",
			Type:     "projects",
			Title:    "Featured Projects",
			Subtitle: "Some of my recent work",
			Content:  "Here are some projects I've worked on that showcase my skills and expertise.",
			Order:    5,
			Visible:  true,
			PageSlug: "home",
			Data: map[string]interface{}{
				"projects": []map[string]interface{}{
					{
						"title":        "E-commerce Platform",
						"description":  "Full-stack e-commerce solution with React frontend and Go backend",
						"image":        "/static/project1.jpg",
						"technologies": []string{"Go", "React", "PostgreSQL", "Docker"},
						"github":       "https://github.com/warren/ecommerce",
						"demo":         "https://demo.example.com",
						"featured":     true,
					},
					{
						"title":        "Task Management App",
						"description":  "Real-time collaborative task management with WebSocket integration",
						"image":        "/static/project2.jpg",
						"technologies": []string{"Node.js", "Vue.js", "MongoDB", "Socket.io"},
						"github":       "https://github.com/warren/taskapp",
						"demo":         "https://tasks.example.com",
						"featured":     true,
					},
					{
						"title":        "API Gateway",
						"description":  "High-performance API gateway built with Go and Redis",
						"image":        "/static/project3.jpg",
						"technologies": []string{"Go", "Redis", "Docker", "Kubernetes"},
						"github":       "https://github.com/warren/gateway",
						"featured":     false,
					},
				},
			},
		},
		{
			Slug:     "contact",
			Type:     "contact",
			Title:    "Get In Touch",
			Subtitle: "Let's work together",
			Content:  "I'm always interested in hearing about new opportunities and interesting projects.",
			Order:    6,
			Visible:  true,
			PageSlug: "home",
			Data: map[string]interface{}{
				"email":    "warren@example.com",
				"phone":    "+1 (555) 123-4567",
				"location": "San Francisco, CA",
				"social": map[string]string{
					"github":   "https://github.com/warren",
					"linkedin": "https://linkedin.com/in/warren",
					"twitter":  "https://twitter.com/warren",
				},
				"availability": "Available for full-time opportunities",
			},
		},
	}

	for _, section := range sections {
		// Data is already properly structured for the model

		err = contentService.CreateSection(ctx, &section)
		if err != nil {
			log.Printf("Failed to create section %s: %v", section.Slug, err)
		} else {
			log.Printf("Created section: %s", section.Title)
		}
	}

	log.Println("Seed data created successfully!")
}

func main() {
	seedData()
}
