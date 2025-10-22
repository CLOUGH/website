package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"warrenclough.com/internal/services"
)

type HomeHandler struct {
	contentService *services.ContentService
}

func NewHomeHandler(contentService *services.ContentService) *HomeHandler {
	return &HomeHandler{
		contentService: contentService,
	}
}

func (h *HomeHandler) Home(c *gin.Context) {
	ctx := c.Request.Context()

	// Fetch page data
	page, err := h.contentService.GetPage(ctx, "home")
	if err != nil {
		// Fallback to static home page if database isn't set up yet
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "Warren Clough - Software Developer & Designer",
			"meta": gin.H{
				"description": "Personal website of Warren Clough - Software Developer, Designer, and Technology Enthusiast. Portfolio, blog, and professional information.",
				"keywords":    "Warren Clough, software developer, designer, portfolio, blog, technology",
			},
		})
		return
	}

	// Fetch sections
	sections, err := h.contentService.GetPageSections(ctx, "home")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Failed to load page content",
		})
		return
	}

	// Render dynamic home page with components
	c.HTML(http.StatusOK, "home-dynamic", gin.H{
		"Page":     page,
		"Sections": sections,
	})
}

// API Endpoints for generic page and section access

func (h *HomeHandler) GetPage(c *gin.Context) {
	pageSlug := c.Param("pageSlug")

	page, err := h.contentService.GetPage(c.Request.Context(), pageSlug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}

	c.JSON(http.StatusOK, page)
}

func (h *HomeHandler) GetPageSections(c *gin.Context) {
	pageSlug := c.Param("pageSlug")

	sections, err := h.contentService.GetPageSections(c.Request.Context(), pageSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch sections"})
		return
	}

	c.JSON(http.StatusOK, sections)
}

func (h *HomeHandler) GetPageSection(c *gin.Context) {
	sectionID := c.Param("sectionID")

	section, err := h.contentService.GetSection(c.Request.Context(), sectionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Section not found"})
		return
	}

	c.JSON(http.StatusOK, section)
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
