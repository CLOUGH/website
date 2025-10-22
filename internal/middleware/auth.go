package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"warrenclough.com/internal/config"
	"warrenclough.com/internal/models"
	"warrenclough.com/internal/services"
)

// AuthMiddleware handles Firebase authentication
type AuthMiddleware struct {
	authClient     *auth.Client
	contentService *services.ContentService
	config         *config.Config
}

// NewAuthMiddleware creates a new auth middleware
func NewAuthMiddleware(authClient *auth.Client, contentService *services.ContentService, cfg *config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		authClient:     authClient,
		contentService: contentService,
		config:         cfg,
	}
}

// RequireAuth middleware that requires authentication
func (am *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		idToken := tokenParts[1]

		// Verify the ID token
		token, err := am.authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			log.Printf("Error verifying ID token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Get or create user
		user := &models.User{
			Email:   token.Claims["email"].(string),
			Name:    token.Claims["name"].(string),
			Picture: token.Claims["picture"].(string),
		}

		err = am.contentService.CreateOrUpdateUser(c.Request.Context(), user)
		if err != nil {
			log.Printf("Error creating/updating user: %v", err)
		}

		// Store user info in context
		c.Set("user_email", user.Email)
		c.Set("user_name", user.Name)
		c.Set("user_picture", user.Picture)
		c.Set("firebase_claims", token.Claims)

		c.Next()
	}
}

// RequireAdmin middleware that requires admin privileges
func (am *AuthMiddleware) RequireAdmin() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// First require auth
		am.RequireAuth()(c)
		if c.IsAborted() {
			return
		}

		// Check if user is admin
		userEmail := c.GetString("user_email")
		isAdmin, err := am.contentService.IsUserAdmin(c.Request.Context(), userEmail)
		if err != nil {
			log.Printf("Error checking admin status: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking admin status"})
			c.Abort()
			return
		}

		if !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}

		c.Next()
	})
}

// OptionalAuth middleware that allows both authenticated and unauthenticated access
func (am *AuthMiddleware) OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.Next()
			return
		}

		idToken := tokenParts[1]
		token, err := am.authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.Next()
			return
		}

		// Store user info in context if authenticated
		c.Set("user_email", token.Claims["email"].(string))
		c.Set("user_name", token.Claims["name"].(string))
		c.Set("authenticated", true)

		c.Next()
	}
}

// GoogleOAuthConfig represents Google OAuth configuration
type GoogleOAuthConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
	Scope        string `json:"scope"`
}

// LoginHandler handles the Google OAuth login flow
func (am *AuthMiddleware) LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// This will be handled by Firebase Auth on the frontend
		// Return the Firebase config for client-side authentication
		c.HTML(http.StatusOK, "admin/login.html", gin.H{
			"title":          "Admin Login",
			"FirebaseConfig": am.config.FirebaseConfig,
		})
	}
}

// LogoutHandler handles logout
func (am *AuthMiddleware) LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Client-side logout with Firebase
		c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
	}
}

// AuthStatusHandler returns current authentication status
func (am *AuthMiddleware) AuthStatusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userEmail := c.GetString("user_email")
		if userEmail == "" {
			c.JSON(http.StatusOK, gin.H{
				"authenticated": false,
			})
			return
		}

		isAdmin, _ := am.contentService.IsUserAdmin(c.Request.Context(), userEmail)

		c.JSON(http.StatusOK, gin.H{
			"authenticated": true,
			"email":         userEmail,
			"name":          c.GetString("user_name"),
			"picture":       c.GetString("user_picture"),
			"is_admin":      isAdmin,
		})
	}
}

// Authenticate handles the authentication endpoint for admin login
func (am *AuthMiddleware) Authenticate(c *gin.Context) {
	var requestBody struct {
		IDToken string `json:"idToken" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Verify the ID token
	token, err := am.authClient.VerifyIDToken(context.Background(), requestBody.IDToken)
	if err != nil {
		log.Printf("Token verification failed: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Get user info from token
	userEmail := token.Claims["email"].(string)
	userName, _ := token.Claims["name"].(string)
	userPicture, _ := token.Claims["picture"].(string)

	// Create or update user in database
	user := &models.User{
		Email:   userEmail,
		Name:    userName,
		Picture: userPicture,
	}

	if err := am.contentService.CreateOrUpdateUser(c.Request.Context(), user); err != nil {
		log.Printf("Failed to create/update user: %v", err)
	}

	// Check if user is admin
	isAdmin, err := am.contentService.IsUserAdmin(c.Request.Context(), userEmail)
	if err != nil {
		log.Printf("Failed to check admin status for %s: %v", userEmail, err)
		isAdmin = false
	}

	// Generate session or set cookie here if needed
	// For now, we'll just return the admin status

	c.JSON(http.StatusOK, gin.H{
		"authenticated": true,
		"email":         userEmail,
		"name":          userName,
		"picture":       userPicture,
		"isAdmin":       isAdmin,
	})
}
