package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Warren Clough - Software Developer & Designer",
		"meta": gin.H{
			"description": "Personal website of Warren Clough - Software Developer, Designer, and Technology Enthusiast. Portfolio, blog, and professional information.",
			"keywords":    "Warren Clough, software developer, designer, portfolio, blog, technology",
		},
	})
}

func (h *HomeHandler) Contact(c *gin.Context) {
	c.HTML(http.StatusOK, "contact.html", gin.H{
		"title": "Contact - Warren Clough",
		"meta": gin.H{
			"description": "Get in touch with Warren Clough for collaboration, opportunities, or just to say hello.",
			"keywords":    "contact, Warren Clough, collaboration, opportunities",
		},
	})
}

func (h *HomeHandler) ContactForm(c *gin.Context) {
	var form struct {
		Name    string `json:"name" binding:"required"`
		Email   string `json:"email" binding:"required,email"`
		Subject string `json:"subject" binding:"required"`
		Message string `json:"message" binding:"required"`
	}

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}

	// Here you would typically save to Firebase or send an email
	// For now, we'll just return a success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Thank you for your message! I'll get back to you soon.",
	})
}

func (h *HomeHandler) NewsletterSubscribe(c *gin.Context) {
	var form struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}

	// Here you would typically save to Firebase
	// For now, we'll just return a success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Thank you for subscribing to my newsletter!",
	})
}
