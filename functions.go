package main

import (
	"context"
	"log"
	"net/http"

	"warrenclough.com/internal/config"
	"warrenclough.com/internal/handlers"
	"warrenclough.com/internal/middleware"
	"warrenclough.com/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Firebase Cloud Function entry point
func WarrenCloughWebsite(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("Failed to initialize Firebase: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer firebaseService.Close()

	// Initialize services
	blogService := services.NewBlogService(firebaseService)
	portfolioService := services.NewPortfolioService(firebaseService)

	// Initialize handlers
	homeHandler := handlers.NewHomeHandler()
	blogHandler := handlers.NewBlogHandler(blogService)
	portfolioHandler := handlers.NewPortfolioHandler(portfolioService)
	resumeHandler := handlers.NewResumeHandler()

	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)

	// Create Gin router
	router := gin.New()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())
	router.Use(middleware.SecurityHeaders())

	// Static files
	router.Static("/static", "./web/static")
	router.LoadHTMLGlob("web/templates/*")

	// Routes
	router.GET("/", homeHandler.Home)
	router.GET("/portfolio", portfolioHandler.Portfolio)
	router.GET("/blog", blogHandler.Blog)
	router.GET("/blog/:slug", blogHandler.BlogPost)
	router.GET("/resume", resumeHandler.Resume)
	router.GET("/contact", homeHandler.Contact)

	// API routes for HTMX
	api := router.Group("/api")
	{
		api.GET("/blog/posts", blogHandler.GetPosts)
		api.GET("/portfolio/projects", portfolioHandler.GetProjects)
		api.POST("/contact", homeHandler.ContactForm)
		api.POST("/newsletter/subscribe", homeHandler.NewsletterSubscribe)
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// Serve the request
	router.ServeHTTP(w, r)
}
