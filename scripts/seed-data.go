package main

import (
	"context"
	"log"
	"time"

	"warrenclough.com/internal/config"
	"warrenclough.com/internal/models"
	"warrenclough.com/internal/services"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize configuration
	cfg := config.Load()

	// Initialize Firebase
	ctx := context.Background()
	firebaseService, err := services.NewFirebaseService(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}
	defer firebaseService.Close()

	client := firebaseService.GetClient()

	// Seed blog posts
	log.Println("Seeding blog posts...")
	blogPosts := []models.BlogPost{
		{
			Title:       "Building Scalable Go Applications",
			Slug:        "building-scalable-go-applications",
			Content:     "<p>Go is an excellent language for building scalable applications...</p>",
			Excerpt:     "Learn how to structure and scale Go applications for production environments.",
			Author:      "Warren Clough",
			PublishedAt: time.Now().AddDate(0, 0, -5),
			UpdatedAt:   time.Now().AddDate(0, 0, -5),
			Tags:        []string{"go", "backend", "scalability"},
			Featured:    true,
			ImageURL:    "",
		},
		{
			Title:       "Modern Web Development with HTMX",
			Slug:        "modern-web-development-htmx",
			Content:     "<p>HTMX is revolutionizing how we think about web development...</p>",
			Excerpt:     "Exploring how HTMX can simplify modern web development without complex JavaScript frameworks.",
			Author:      "Warren Clough",
			PublishedAt: time.Now().AddDate(0, 0, -10),
			UpdatedAt:   time.Now().AddDate(0, 0, -10),
			Tags:        []string{"htmx", "frontend", "web-development"},
			Featured:    true,
			ImageURL:    "",
		},
		{
			Title:       "Firebase Best Practices",
			Slug:        "firebase-best-practices",
			Content:     "<p>Firebase provides a powerful platform for building modern applications...</p>",
			Excerpt:     "Tips and tricks for building efficient applications with Firebase and Firestore.",
			Author:      "Warren Clough",
			PublishedAt: time.Now().AddDate(0, 0, -15),
			UpdatedAt:   time.Now().AddDate(0, 0, -15),
			Tags:        []string{"firebase", "cloud", "best-practices"},
			Featured:    false,
			ImageURL:    "",
		},
	}

	for _, post := range blogPosts {
		_, _, err := client.Collection("blog_posts").Add(ctx, post)
		if err != nil {
			log.Printf("Error adding blog post %s: %v", post.Title, err)
		} else {
			log.Printf("Added blog post: %s", post.Title)
		}
	}

	// Seed portfolio projects
	log.Println("Seeding portfolio projects...")
	projects := []models.PortfolioProject{
		{
			Title:            "E-Commerce Platform",
			Description:      "A full-stack e-commerce solution built with Go, React, and Firebase.",
			LongDescription:  "A comprehensive e-commerce platform featuring user authentication, product management, shopping cart, payment processing, and order management. Built with modern technologies and best practices.",
			ImageURL:         "",
			LiveURL:          "https://example-ecommerce.com",
			GitHubURL:        "https://github.com/warrenclough/ecommerce-platform",
			Technologies:     []string{"Go", "React", "Firebase", "Stripe", "Docker"},
			Featured:         true,
			CreatedAt:        time.Now().AddDate(0, 0, -30),
			UpdatedAt:        time.Now().AddDate(0, 0, -30),
		},
		{
			Title:            "Task Management App",
			Description:      "A collaborative task management application with real-time updates.",
			LongDescription:  "A modern task management application that allows teams to collaborate effectively. Features include real-time updates, drag-and-drop interface, file attachments, and team management.",
			ImageURL:         "",
			LiveURL:          "https://example-tasks.com",
			GitHubURL:        "https://github.com/warrenclough/task-manager",
			Technologies:     []string{"Go", "HTMX", "TailwindCSS", "WebSockets", "PostgreSQL"},
			Featured:         true,
			CreatedAt:        time.Now().AddDate(0, 0, -45),
			UpdatedAt:        time.Now().AddDate(0, 0, -45),
		},
		{
			Title:            "Analytics Dashboard",
			Description:      "A comprehensive analytics dashboard with real-time data visualization.",
			LongDescription:  "An advanced analytics dashboard that provides insights into user behavior, system performance, and business metrics. Features interactive charts, real-time data updates, and customizable widgets.",
			ImageURL:         "",
			LiveURL:          "https://example-analytics.com",
			GitHubURL:        "https://github.com/warrenclough/analytics-dashboard",
			Technologies:     []string{"Go", "React", "D3.js", "Redis", "Kubernetes"},
			Featured:         false,
			CreatedAt:        time.Now().AddDate(0, 0, -60),
			UpdatedAt:        time.Now().AddDate(0, 0, -60),
		},
	}

	for _, project := range projects {
		_, _, err := client.Collection("portfolio_projects").Add(ctx, project)
		if err != nil {
			log.Printf("Error adding project %s: %v", project.Title, err)
		} else {
			log.Printf("Added project: %s", project.Title)
		}
	}

	// Seed resume data
	log.Println("Seeding resume data...")
	resume := models.Resume{
		Name:     "Warren Clough",
		Title:    "Software Developer & Designer",
		Email:    "warren@warrenclough.com",
		Phone:    "+1 (555) 123-4567",
		Location: "Remote",
		Summary:  "Passionate software developer with expertise in modern web technologies, cloud architecture, and user experience design. I love building scalable applications and creating meaningful digital experiences.",
		Experience: []models.Experience{
			{
				Company:      "Tech Solutions Inc.",
				Position:     "Senior Software Developer",
				Location:     "Remote",
				StartDate:    time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				EndDate:      nil, // Current position
				Description:  "Leading development of scalable web applications using modern technologies. Mentoring junior developers and implementing best practices.",
				Technologies: []string{"Go", "React", "Firebase", "Docker", "Kubernetes"},
			},
			{
				Company:      "StartupXYZ",
				Position:     "Full Stack Developer",
				Location:     "San Francisco, CA",
				StartDate:    time.Date(2018, 6, 1, 0, 0, 0, 0, time.UTC),
				EndDate:      &[]time.Time{time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC)}[0],
				Description:  "Developed and maintained web applications using various technologies. Collaborated with cross-functional teams to deliver high-quality software.",
				Technologies: []string{"JavaScript", "Node.js", "React", "MongoDB", "AWS"},
			},
		},
		Education: []models.Education{
			{
				Institution: "University of Technology",
				Degree:      "Bachelor of Science",
				Field:       "Computer Science",
				StartDate:   time.Date(2014, 9, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2018, 5, 31, 0, 0, 0, 0, time.UTC),
				GPA:         "3.8",
			},
		},
		Skills: []models.Skill{
			{Name: "Go", Category: "Backend", Proficiency: 5},
			{Name: "JavaScript/TypeScript", Category: "Frontend", Proficiency: 5},
			{Name: "React", Category: "Frontend", Proficiency: 4},
			{Name: "Firebase", Category: "Cloud", Proficiency: 4},
			{Name: "Docker", Category: "DevOps", Proficiency: 4},
			{Name: "Kubernetes", Category: "DevOps", Proficiency: 3},
			{Name: "AWS", Category: "Cloud", Proficiency: 4},
			{Name: "PostgreSQL", Category: "Database", Proficiency: 4},
		},
		UpdatedAt: time.Now(),
	}

	_, err = client.Collection("resume").Doc("main").Set(ctx, resume)
	if err != nil {
		log.Printf("Error adding resume: %v", err)
	} else {
		log.Printf("Added resume data")
	}

	log.Println("✅ Data seeding completed successfully!")
}
