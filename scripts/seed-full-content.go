package main

import (
	"context"
	"log"

	"warrenclough.com/internal/config"
	"warrenclough.com/internal/models"
	"warrenclough.com/internal/services"
)

func main() {
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
		Title:       "Warren Clough - Senior Software Engineer",
		Description: "Senior Software Engineer specializing in enterprise-grade solutions. Building scalable, high-performance applications using Java, Spring Boot, Laravel, React, and cloud-native architectures.",
		Metadata: models.PageMeta{
			Keywords:    "Warren Clough, Software Engineer, Java, Spring Boot, Laravel, React, Go, Full Stack Developer, Enterprise Software",
			Description: "Senior Software Engineer with 10+ years of experience in building enterprise-grade applications. Specializing in Java/Spring Boot, Laravel, React, and cloud-native architectures.",
		},
		Visible: true,
	}

	err = contentService.CreatePage(ctx, homePage)
	if err != nil {
		log.Printf("Failed to create home page: %v", err)
	} else {
		log.Println("Created home page")
	}

	// Create sections for home page
	sections := []models.Section{
		// Hero Section
		{
			Slug:     "hero",
			Type:     "hero",
			Title:    "Hi, I'm Warren Clough",
			Subtitle: "Senior Software Engineer",
			Content:  "A passionate Senior Software Engineer specializing in enterprise-grade solutions. I build scalable, high-performance applications using Java, Spring Boot, Laravel, React, and cloud-native architectures. With professional experience in insurance, human resources, and fintech sectors, I transform complex business requirements into elegant, mission-critical systems.",
			Order:    1,
			Visible:  true,
			PageSlug: "home",
			Data: map[string]interface{}{
				"name":         "Warren Clough",
				"title":        "Senior Software Engineer",
				"tagline":      "I build exceptional digital experiences that blend beautiful design with powerful functionality.",
				"github_url":   "https://github.com/clough",
				"linkedin_url": "https://linkedin.com/in/warrenclough",
			},
		},
		// About Section
		{
			Slug:     "about",
			Type:     "about",
			Title:    "About Me",
			Subtitle: "",
			Content:  "",
			Order:    2,
			Visible:  true,
			PageSlug: "home",
			Data: map[string]interface{}{
				"paragraphs": []string{
					"I'm a Full-Stack Software Engineer with a passion for building elegant, scalable solutions that make a real impact. My journey in software development has been driven by an insatiable curiosity for how things work and a desire to create technology that enhances people's lives. I thrive at the intersection of beautiful design and robust engineering, where pixel-perfect user interfaces meet bulletproof backend systems.",
					"With over 10 years of professional experience across insurance, human resources, and fintech industries, I've specialized in architecting and delivering enterprise-grade applications serving thousands of users daily. My professional expertise centers on Java/Spring Boot and Laravel/PHP ecosystems, building mission-critical systems that meet the rigorous demands of enterprise environments—high availability, security, scalability, and maintainability. I'm passionate about simple, elegant solutions that solve complex problems, which draws me to explore Go for its minimalist philosophy and powerful concurrency model in side projects.",
					"Throughout my career, I've had the privilege of working with diverse teams across startups, digital agencies, and established tech companies. These experiences have taught me the importance of clear communication, collaborative problem-solving, and writing code that not only works today but remains maintainable tomorrow. I believe great software is built by great teams, and I'm committed to fostering environments where innovation thrives.",
					"Beyond coding, I'm passionate about mentoring aspiring developers and sharing knowledge with my team. When I'm not at my keyboard, you'll find me gaming, 3D printing custom projects, or tinkering with electronics and Arduino boards. I love the intersection of software and hardware, and I'm always exploring emerging technologies and experimenting with new frameworks. I believe continuous learning isn't just important—it's essential in our rapidly evolving field.",
				},
			},
		},
		// Technologies Section
		{
			Slug:     "technologies",
			Type:     "technologies",
			Title:    "Technologies I Work With",
			Subtitle: "I believe in choosing the right tool for the job. Here's my technical toolkit:",
			Content:  "",
			Order:    3,
			Visible:  true,
			PageSlug: "home",
			Data: map[string]interface{}{
				"categories": map[string][]string{
					"Languages & Frameworks": {
						"Java & Spring Boot",
						"PHP & Laravel",
						"Go (Golang)",
						"Python",
						"React",
						"Angular & AngularJS",
						"TypeScript & JavaScript",
						"C++ & OpenGL",
						"Unreal Engine",
						"HTMX",
						"4D Language",
						"Hibernate & JPA",
						"Electron",
					},
					"Cloud & DevOps": {
						"Kubernetes",
						"Docker",
						"Google Cloud Platform",
						"Firebase",
						"Apache Kafka",
						"GitLab CI/CD",
						"Octopus Deploy",
						"GitHub Actions",
					},
					"Databases & Tools": {
						"PostgreSQL",
						"Oracle SQL",
						"MySQL",
						"MongoDB",
						"Redis",
						"D3.js Visualization",
						"REST APIs",
						"Microservices",
						"Maven & Gradle",
						"Git & GitLab",
						"Jira",
						"Tailwind CSS",
						"JUnit & Mockito",
						"GLSL Shaders",
					},
				},
			},
		},
		// Experience Section
		{
			Slug:     "experience",
			Type:     "experience",
			Title:    "Experience",
			Subtitle: "",
			Content:  "",
			Order:    4,
			Visible:  true,
			PageSlug: "home",
			Data: map[string]interface{}{
				"experiences": []map[string]interface{}{
					{
						"company":    "National Commercial Bank of Jamaica",
						"position":   "Senior Developer Analyst",
						"location":   "Jamaica",
						"employment": "Full-time",
						"start_date": "2018",
						"end_date":   "Present",
						"description": []string{
							"Technical Lead for NCB's Consumer Loans & Credit Card team, delivering enterprise solutions for loan and credit card applications processing millions in transactions daily",
							"Design and build Java microservices with Kubernetes, implementing scalable cloud-native architectures for critical banking operations",
							"Develop modern front-end applications with React, creating intuitive user experiences for banking customers and internal staff",
							"Perform root cause analysis and resolution for production issues, ensuring 99.9% uptime for mission-critical financial systems",
							"Mentor development team members on best practices, code quality, and architectural patterns for enterprise banking applications",
							"Support and enhance core banking systems, working with PostgreSQL and integrating with legacy financial platforms",
						},
						"technologies": []string{
							"Java",
							"React",
							"Kubernetes",
							"PostgreSQL",
							"Oracle",
							"GCP",
							"Microservices",
						},
					},
					{
						"company":    "Epic Technologies",
						"position":   "Senior Software Developer",
						"location":   "Jamaica",
						"employment": "Full-time",
						"start_date": "2014",
						"end_date":   "2018",
						"description": []string{
							"Technical Lead for the company's Human Resource Management system and insurance product data migration tools, managing full development lifecycle",
							"Designed and built enterprise web applications using Laravel/PHP and Angular, creating scalable solutions for HR and payroll management serving multiple organizations",
							"Developed desktop applications with 4D platform, delivering robust tools for complex business operations and data processing",
							"Mentored junior developers on coding best practices, architecture patterns, and professional software development methodologies",
							"Provided technical support and maintenance for the company's payroll software product, ensuring system reliability and data integrity",
							"Implemented data migration strategies for insurance products, safely transferring critical business data with zero data loss",
						},
						"technologies": []string{
							"PHP",
							"Laravel",
							"Angular",
							"4D",
							"JavaScript",
							"PostgreSQL",
						},
					},
					{
						"company":    "Epic Technologies & Alteroo Consulting Group",
						"position":   "Software Developer / Web Developer",
						"location":   "Jamaica",
						"employment": "Part-time / Contract",
						"start_date": "2010",
						"end_date":   "2014",
						"description": []string{
							"Resolved bugs and extended features for Epic Technologies' insurance software product, ensuring system stability and enhanced functionality",
							"Customized Plone CMS to client specifications and designs at Alteroo Consulting Group, delivering tailored web solutions",
							"Gained hands-on experience with enterprise insurance systems, including policy management, claims processing, and business workflows",
							"Developed proficiency in web technologies and CMS platforms while working on client projects and internal tools",
							"Collaborated with senior developers to understand insurance domain requirements and implement technical solutions",
						},
						"technologies": []string{
							"PHP",
							"Python",
							"Plone CMS",
							"JavaScript",
							"PostgreSQL",
						},
					},
				},
			},
		},
		// Projects Section
		{
			Slug:     "projects",
			Type:     "projects",
			Title:    "Projects",
			Subtitle: "",
			Content:  "",
			Order:    5,
			Visible:  true,
			PageSlug: "home",
			Data: map[string]interface{}{
				"projects": []map[string]interface{}{
					{
						"id":                "poe-recipe",
						"title":             "Path of Exile Recipe Helper",
						"short_description": "Fun gaming side project built for personal use as a game add-on for Path of Exile.",
						"description":       "Desktop application that integrates with Path of Exile's public API to automatically scan player inventory, generate a visual copy with intelligent highlighting of items eligible for chaos recipes and other crafting recipes. Built with Electron for cross-platform compatibility and Angular for a responsive, game-like UI.",
						"full_description":  "Desktop application that integrates with Path of Exile's public API to automatically scan player inventory, generate a visual copy with intelligent highlighting of items eligible for chaos recipes and other crafting recipes. Built with Electron for cross-platform compatibility and Angular for a responsive, game-like UI. Features real-time inventory syncing, recipe pattern matching, and visual item categorization. Implements efficient caching strategies and parallel API requests for optimal performance.",
						"technologies": []string{
							"Electron",
							"Angular",
							"TypeScript",
							"REST API",
						},
						"screenshots": []string{
							"https://github.com/CLOUGH/poe-recipe/raw/main/src/assets/screenshots/Screenshot-2021-01-19-214751.png",
						},
						"github_url": "https://github.com/clough/poe-recipe",
						"featured":   true,
						"order":      1,
					},
					{
						"id":                "go-trader",
						"title":             "Go Trader - Jamaica Stock Exchange",
						"short_description": "Side project built to learn Go and HTMX while exploring financial data visualization.",
						"description":       "Real-time stock market application that pulls data from the Jamaica Stock Exchange API and visualizes it with custom-built interactive charts using D3.js. Features include live stock data fetching, historical price tracking, custom candlestick and line charts, search functionality, and responsive design.",
						"full_description":  "Real-time stock market application that pulls data from the Jamaica Stock Exchange API and visualizes it with custom-built interactive charts using D3.js. Built as a learning project to master Go's concurrency patterns, HTMX for dynamic updates, and D3.js for data visualization. Features include live stock data fetching, historical price tracking, custom candlestick and line charts, search functionality, and responsive design. Demonstrates modern web development patterns with minimal JavaScript and server-side rendering. Implements efficient data caching and real-time updates using Go channels.",
						"technologies": []string{
							"Go",
							"HTMX",
							"D3.js",
							"REST API",
						},
						"screenshots": []string{
							"https://storage.googleapis.com/my-website-a3970.appspot.com/static/go-trader-home-page.png",
							"https://storage.googleapis.com/my-website-a3970.appspot.com/static/go-trader-stock-page.png",
							"https://storage.googleapis.com/my-website-a3970.appspot.com/static/go-trader-stock-detail-page.png",
						},
						"github_url": "https://github.com/clough/go-trader",
						"featured":   true,
						"order":      2,
					},
					{
						"id":                "hrmnext-recruitment",
						"title":             "HRMNext Recruitment Platform",
						"short_description": "Public-facing SaaS recruitment solution developed at Epic Technologies for modern hiring workflows.",
						"description":       "Enterprise platform that enables recruiters to create job openings, specify requirements, setup and conduct interviews with real-time messaging, manage interview stages, and send invitations. Features seamless integration with Epic's HRMNext 4D Solution for comprehensive HR management.",
						"full_description":  "Enterprise platform that enables recruiters to create job openings, specify requirements, setup and conduct interviews with real-time messaging, manage interview stages, and send invitations. Features seamless integration with Epic's HRMNext 4D Solution for comprehensive HR management. Built as a multi-tenant SaaS application with role-based access control and automated workflow management. Implements advanced scheduling algorithms, candidate tracking dashboards, and comprehensive analytics for hiring teams.",
						"technologies": []string{
							"AngularJS",
							"Laravel",
							"4D",
							"SaaS",
						},
						"screenshots": []string{
							"https://storage.googleapis.com/my-website-a3970.appspot.com/static/hrmnext-home.png",
						},
						"featured": true,
						"order":    3,
					},
					{
						"id":                "ncb-credit-card",
						"title":             "NCB Credit Card Application Portal",
						"short_description": "Public-facing web application for NCB (National Commercial Bank) credit card applications.",
						"description":       "Enterprise-grade portal that streamlines the credit card application process with a secure, user-friendly interface. Built with modern microservices architecture using React and Spring/Java backend, deployed on Kubernetes for high availability and scalability. Features real-time form validation, document upload, and application tracking.",
						"full_description":  "Enterprise-grade credit card application portal built for NCB to streamline the credit card application process. Customers can apply for credit cards online with a secure, user-friendly interface that guides them through the application steps. Built with modern microservices architecture using React for the frontend and Spring/Java backend, deployed on Kubernetes for high availability and scalability. Features include real-time form validation, document upload, application tracking, and integration with NCB's core banking systems. Implements strict security measures for handling sensitive financial data and PCI compliance. Deployed across multiple availability zones with auto-scaling capabilities.",
						"technologies": []string{
							"React",
							"Spring Boot",
							"Java",
							"Kubernetes",
						},
						"screenshots": []string{
							"https://storage.googleapis.com/my-website-a3970.appspot.com/static/credit-card-web-portal.png",
						},
						"live_url": "https://creditcard.jncb.com/",
						"featured": true,
						"order":    4,
					},
					{
						"id":                "opengl-snake",
						"title":             "OpenGL Snake Game",
						"short_description": "One of my first public projects from university where I learned OpenGL and computer graphics fundamentals.",
						"description":       "Built a classic Snake game from scratch using modern OpenGL, implementing custom rendering pipelines, collision detection, and game state management. Features smooth animations, score tracking, and increasing difficulty levels.",
						"full_description":  "Built a classic Snake game from scratch using modern OpenGL, implementing custom rendering pipelines, collision detection, and game state management. This project taught me low-level graphics programming, shader development, and game loop architecture. Features smooth animations, score tracking, and increasing difficulty levels. Implemented using C++ with OpenGL 3.3+ core profile, custom GLSL shaders for rendering, and efficient vertex buffer management. Includes custom physics calculations and optimized frame timing.",
						"technologies": []string{
							"OpenGL",
							"C++",
							"GLSL Shaders",
							"Game Development",
						},
						"video_url":       "https://www.youtube.com/watch?v=voOTQlxe6oo",
						"video_embed_url": "https://www.youtube.com/embed/voOTQlxe6oo?si=asDBgtiRSfk9znNA",
						"github_url":      "https://github.com/clough/snake",
						"featured":        false,
						"order":           5,
					},
				},
				"confidential_notice": map[string]string{
					"title":   "Confidential Professional Work",
					"content": "The majority of my professional projects involve private, confidential, and sensitive systems developed for enterprise clients in the financial services, insurance, and human resources sectors.\n\nDue to non-disclosure agreements (NDAs) and client confidentiality requirements, I am unable to showcase these projects publicly or discuss their technical implementations in detail.\n\nThe projects displayed above represent a selection of public-facing work and personal projects that I'm able to share. For discussions about my confidential professional experience, please reach out directly.",
				},
			},
		},
		// Contact Section
		{
			Slug:     "contact",
			Type:     "contact",
			Title:    "Let's Build Something Amazing",
			Subtitle: "",
			Content:  "",
			Order:    6,
			Visible:  true,
			PageSlug: "home",
			Data: map[string]interface{}{
				"intro": []string{
					"I'm always excited to connect with fellow developers, discuss innovative projects, or explore new opportunities. Whether you need technical consultation, want to discuss a project opportunity, or just want to chat about the latest in web development—I'd love to hear from you!",
					"Currently open to new opportunities including full-time positions, contract work, and technical advisory roles. I'm particularly interested in projects involving Java/Spring Boot, Laravel, Angular, and cloud-native architectures. With strong experience in Angular, I'm passionate about building modern, reactive enterprise applications. I'm also eager to work on Go projects professionally, drawn by its elegant simplicity and philosophy of building clean, maintainable systems.",
				},
				"call_to_action": "Drop me an email at clough.warren@gmail.com or connect with me on LinkedIn and GitHub. I typically respond within 24 hours!",
				"email":          "clough.warren@gmail.com",
				"github_url":     "https://github.com/clough",
				"linkedin_url":   "https://linkedin.com/in/warrenclough",
				"buttons": []map[string]string{
					{
						"text": "View Resume",
						"url":  "https://storage.googleapis.com/my-website-a3970.appspot.com/static/MyResume-20250830-compressed.pdf",
						"type": "primary",
					},
					{
						"text": "Send Message",
						"url":  "/contact",
						"type": "secondary",
					},
				},
			},
		},
	}

	for _, section := range sections {
		err = contentService.CreateSection(ctx, &section)
		if err != nil {
			log.Printf("Failed to create section %s: %v", section.Slug, err)
		} else {
			log.Printf("Created section: %s (%s)", section.Title, section.Slug)
		}
	}

	log.Println("\n✅ Seed data created successfully!")
	log.Println("📄 Created 1 page (home)")
	log.Printf("📦 Created %d sections\n", len(sections))
}
