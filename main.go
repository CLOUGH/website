package main

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"warrenclough.com/internal/config"
	"warrenclough.com/internal/handlers"
	"warrenclough.com/internal/middleware"
	"warrenclough.com/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
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

	// Initialize services
	contentService := services.NewContentService(firebaseService, cfg)
	storageService := services.NewStorageService(firebaseService)

	// Initialize handlers
	homeHandler := handlers.NewHomeHandler(contentService)
	adminHandler := handlers.NewAdminHandler(contentService, cfg)
	storageHandler := handlers.NewStorageHandler(storageService)

	// Initialize authentication middleware
	authMiddleware := middleware.NewAuthMiddleware(firebaseService.GetAuthClient(), contentService, cfg)

	// Set up Gin router
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Initialize rate limiters
	// General: 100 requests per second per IP, burst of 200
	generalLimiter := middleware.NewRateLimiter(100, 200)
	// Strict: 10 requests per second per IP for API endpoints, burst of 20
	strictLimiter := middleware.NewRateLimiter(10, 20)
	// Contact form: Allow reasonable testing - 1 request per 30 seconds, burst of 5
	contactLimiter := middleware.NewRateLimiter(rate.Limit(1.0/30.0), 5)

	// Initialize cache (5 minute TTL for public pages)
	pageCache := middleware.NewCache(5 * time.Minute)

	// Middleware
	router.Use(middleware.CORS())
	router.Use(middleware.SecurityHeaders())
	router.Use(generalLimiter.Middleware()) // Global rate limit

	// Static files from Firebase Storage
	// Keep manifest.json local as it needs to be served from same origin
	router.StaticFile("/manifest.json", "./web/static/manifest.json")

	// Proxy to Firebase Storage for other static files
	router.GET("/static/:filename", storageHandler.GetFile)
	router.GET("/resume", storageHandler.GetResume)

	// Set template functions and load templates
	router.SetFuncMap(template.FuncMap{
		"json": func(v interface{}) string {
			if v == nil {
				return "{}"
			}
			jsonBytes, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				return "{}"
			}
			return string(jsonBytes)
		},
		"html": func(s string) template.HTML {
			return template.HTML(s)
		},
		"has": func(slice interface{}, item string) bool {
			switch s := slice.(type) {
			case []string:
				for _, v := range s {
					if v == item {
						return true
					}
				}
			case []interface{}:
				for _, v := range s {
					if str, ok := v.(string); ok && str == item {
						return true
					}
				}
			}
			return false
		},
		"sequence": func(start, end int) []int {
			if start > end {
				return []int{}
			}
			seq := make([]int, end-start+1)
			for i := range seq {
				seq[i] = start + i
			}
			return seq
		},
	})

	// Load all templates from templates directory and subdirectories
	router.LoadHTMLGlob("web/templates/**/*.html")

	// Public Routes (with caching for GET requests)
	publicGroup := router.Group("/")
	publicGroup.Use(pageCache.Middleware())
	{
		publicGroup.GET("/", homeHandler.Home)
		publicGroup.GET("/contact", homeHandler.Contact)
	}

	// POST routes without cache, with strict rate limiting
	router.POST("/contact", contactLimiter.StrictMiddleware(), homeHandler.ContactForm)
	router.POST("/newsletter", strictLimiter.StrictMiddleware(), homeHandler.NewsletterSubscribe)

	// API Routes (with rate limiting and caching)
	apiRoutes := router.Group("/api")
	{
		// Generic page and section API (with cache)
		apiRoutes.GET("/page/:pageSlug", homeHandler.GetPage)
		apiRoutes.GET("/page/:pageSlug/sections", homeHandler.GetPageSections)
		apiRoutes.GET("/page/:pageSlug/sections/:sectionSlug", homeHandler.GetPageSection)

		// Contact form API (with contact-specific rate limit)
		apiRoutes.POST("/contact", contactLimiter.Middleware(), handlers.HandleContact)
	}

	// Admin Routes (with moderate rate limiting)
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(strictLimiter.Middleware()) // Protect admin endpoints
	{
		// Login routes (no auth required)
		adminRoutes.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin/login.html", gin.H{
				"FirebaseConfig": cfg.FirebaseConfig,
			})
		})
		adminRoutes.POST("/auth", authMiddleware.Authenticate)

		// Dashboard can be accessed without auth, but will check auth client-side
		adminRoutes.GET("/dashboard", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin/dashboard-new.html", gin.H{
				"title":          "Admin Dashboard",
				"FirebaseConfig": cfg.FirebaseConfig,
			})
		})

		// Protected admin routes - API endpoints
		protected := adminRoutes.Group("", authMiddleware.RequireAdmin())
		{
			// Generic Page & Section Management API
			protected.GET("/pages", adminHandler.GetAllPages)
			protected.POST("/pages", adminHandler.CreatePage)
			protected.PUT("/pages/:id", adminHandler.UpdatePage)
			protected.DELETE("/pages/:id", adminHandler.DeletePage)

			// HTMX-friendly section endpoints
			protected.GET("/pages/:pageSlug/sections", adminHandler.GetPageSectionsHTML)
			protected.POST("/pages/:pageSlug/sections", adminHandler.CreateSection)
			protected.PUT("/sections/:id", adminHandler.UpdateSection)
			protected.DELETE("/sections/:id", adminHandler.DeleteSection)
			protected.PUT("/sections/:id/reorder", adminHandler.UpdateSectionOrder)
			protected.PUT("/sections/:id/toggle-visibility", adminHandler.ToggleSectionVisibility)

			// Section forms and components
			protected.GET("/sections/new", adminHandler.RenderSectionForm)
			protected.GET("/sections/:id/edit", adminHandler.RenderSectionEditForm)

		}
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// Start server with timeouts for DDoS protection
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,  // Maximum time to read request
		WriteTimeout:   10 * time.Second,  // Maximum time to write response
		IdleTimeout:    120 * time.Second, // Maximum time for idle connections
		MaxHeaderBytes: 1 << 20,           // 1 MB max header size
	}

	log.Printf("Server starting on port %s with DDoS protection enabled", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
