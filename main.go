package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"warrenclough.com/internal/config"
	"warrenclough.com/internal/handlers"
	"warrenclough.com/internal/middleware"
	"warrenclough.com/internal/services"

	"github.com/gin-gonic/gin"
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

	// Initialize services (for future use)
	_ = firebaseService // Keep firebase service for future features

	// Initialize handlers
	homeHandler := handlers.NewHomeHandler()

	// Set up Gin router
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Middleware
	router.Use(middleware.CORS())
	router.Use(middleware.SecurityHeaders())

	// Static files
	router.Static("/static", "./web/static")
	router.LoadHTMLGlob("web/templates/*")

	// Routes
	router.GET("/", homeHandler.Home)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
