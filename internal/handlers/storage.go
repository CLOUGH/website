package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"warrenclough.com/internal/services"
)

type StorageHandler struct {
	storageService *services.StorageService
}

func NewStorageHandler(storageService *services.StorageService) *StorageHandler {
	return &StorageHandler{
		storageService: storageService,
	}
}

// GetFile proxies requests to Firebase Storage (with caching)
func (h *StorageHandler) GetFile(c *gin.Context) {
	filename := c.Param("filename")

	// Get public URL
	publicURL := h.storageService.GetPublicURL("static/" + filename)

	// Redirect to public URL (Firebase Storage handles caching)
	c.Redirect(http.StatusMovedPermanently, publicURL)
}

// GetResume specifically handles resume downloads
func (h *StorageHandler) GetResume(c *gin.Context) {
	// Get public URL for resume
	publicURL := h.storageService.GetPublicURL("static/MyResume-20250830-compressed.pdf")

	// Set download headers
	c.Header("Content-Disposition", "attachment; filename=Warren-Clough-Resume.pdf")
	c.Redirect(http.StatusMovedPermanently, publicURL)
}
