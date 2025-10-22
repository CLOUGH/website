package handlers

import (
	"fmt"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/gin-gonic/gin"
	"warrenclough.com/internal/config"
)

type ContactRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Subject string `json:"subject"`
	Message string `json:"message" binding:"required"`
}

type ContactResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// NewContactHandler creates a contact handler with the given config
func NewContactHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		handleContact(c, cfg)
	}
}

// HandleContact processes contact form submissions and sends email notifications
func HandleContact(c *gin.Context) {
	// For backwards compatibility, create a config from environment
	cfg := config.Load()
	handleContact(c, cfg)
}

// handleContact is the internal handler that uses the config
func handleContact(c *gin.Context, cfg *config.Config) {
	fmt.Printf("[CONTACT] Received %s request to %s from IP: %s\n", c.Request.Method, c.Request.URL.Path, c.ClientIP())

	var req ContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("[CONTACT] JSON binding failed: %v\n", err)
		fmt.Printf("[CONTACT] Request body: %v\n", c.Request.Body)
		c.JSON(http.StatusBadRequest, ContactResponse{
			Success: false,
			Message: "Invalid form data",
		})
		return
	}

	fmt.Printf("[CONTACT] Parsed request - Name: '%s', Email: '%s', Message length: %d\n", req.Name, req.Email, len(req.Message))

	// Spam detection: check for suspicious patterns
	if isSpam(req) {
		fmt.Printf("[CONTACT] Message marked as spam for IP: %s\n", c.ClientIP())
		c.JSON(http.StatusBadRequest, ContactResponse{
			Success: false,
			Message: "Message rejected",
		})
		return
	}

	fmt.Printf("[CONTACT] Spam detection passed\n")

	// Auto-generate subject if not provided
	if req.Subject == "" {
		req.Subject = "New Contact Form Message from " + req.Name
	}

	fmt.Printf("[CONTACT] Attempting to send email with subject: '%s'\n", req.Subject)

	// Send email notification
	if err := sendContactEmail(req, cfg); err != nil {
		fmt.Printf("[CONTACT] Email sending failed: %v\n", err)
		c.JSON(http.StatusInternalServerError, ContactResponse{
			Success: false,
			Message: "Failed to send email",
		})
		return
	}

	fmt.Printf("[CONTACT] Email sent successfully\n")

	c.JSON(http.StatusOK, ContactResponse{
		Success: true,
		Message: "Message sent successfully",
	})
}

func sendContactEmail(req ContactRequest, cfg *config.Config) error {
	// Get SMTP configuration from config
	smtpHost := cfg.SMTPConfig.Host
	smtpPort := cfg.SMTPConfig.Port
	smtpUsername := cfg.SMTPConfig.Username
	smtpPassword := cfg.SMTPConfig.Password
	recipientEmail := cfg.SMTPConfig.Username // Send to yourself
	// If SMTP is not configured, log the message instead
	if smtpHost == "" || smtpPort == "" || smtpUsername == "" || smtpPassword == "" {
		fmt.Printf("\n=== Contact Form Submission (SMTP Not Configured) ===\n")
		fmt.Printf("From: %s <%s>\n", req.Name, req.Email)
		fmt.Printf("Subject: %s\n", req.Subject)
		fmt.Printf("Message:\n%s\n", req.Message)
		fmt.Printf("=====================================================\n\n")
		return nil
	}

	// Prepare email message
	from := smtpUsername
	to := []string{recipientEmail}

	message := []byte(
		"From: " + from + "\r\n" +
			"To: " + recipientEmail + "\r\n" +
			"Subject: [Website Contact] " + req.Subject + "\r\n" +
			"MIME-version: 1.0;\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
			"<html><body>" +
			"<h2>New Contact Form Submission</h2>" +
			"<p><strong>Name:</strong> " + req.Name + "</p>" +
			"<p><strong>Email:</strong> " + req.Email + "</p>" +
			"<p><strong>Subject:</strong> " + req.Subject + "</p>" +
			"<p><strong>Message:</strong></p>" +
			"<p>" + req.Message + "</p>" +
			"<hr>" +
			"<p><em>Reply to: " + req.Email + "</em></p>" +
			"</body></html>",
	)

	// Authentication
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Send email
	addr := smtpHost + ":" + smtpPort
	err := smtp.SendMail(addr, auth, from, to, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func isSpam(req ContactRequest) bool {
	fmt.Printf("[SPAM CHECK] Checking message from '%s' (%s)\n", req.Name, req.Email)
	fmt.Printf("[SPAM CHECK] Message: '%s'\n", req.Message)

	// Check for excessive links in message
	linkCount := strings.Count(strings.ToLower(req.Message), "http://") +
		strings.Count(strings.ToLower(req.Message), "https://")
	if linkCount > 5 { // Increased from 3 to 5
		fmt.Printf("[SPAM CHECK] REJECTED: too many links (%d)\n", linkCount)
		return true
	}
	fmt.Printf("[SPAM CHECK] Link count OK: %d\n", linkCount)

	// Check for suspicious keywords
	spamKeywords := []string{
		"viagra", "cialis", "casino", "lottery", "winner",
		"click here", "buy now", "limited offer", "act now",
		"congratulations", "you've won", "claim your",
	}
	messageLower := strings.ToLower(req.Message)
	for _, keyword := range spamKeywords {
		if strings.Contains(messageLower, keyword) {
			fmt.Printf("[SPAM CHECK] REJECTED: keyword '%s'\n", keyword)
			return true
		}
	}
	fmt.Printf("[SPAM CHECK] No spam keywords found\n")

	// Check if message is too short (less than 5 characters) - reduced from 10
	messageLen := len(strings.TrimSpace(req.Message))
	if messageLen < 5 {
		fmt.Printf("[SPAM CHECK] REJECTED: message too short (%d chars)\n", messageLen)
		return true
	}
	fmt.Printf("[SPAM CHECK] Message length OK: %d chars\n", messageLen)

	// More lenient name validation - allow numbers and more special characters
	for i, char := range req.Name {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') || char == ' ' || char == '-' ||
			char == '\'' || char == '.' || char == '_') {
			fmt.Printf("[SPAM CHECK] REJECTED: invalid character in name at position %d: '%c' (unicode: %d)\n", i, char, char)
			return true
		}
	}
	fmt.Printf("[SPAM CHECK] Name validation passed\n")

	fmt.Printf("[SPAM CHECK] PASSED: Message approved\n")
	return false
}
