package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"warrenclough.com/internal/config"
	"warrenclough.com/internal/models"
	"warrenclough.com/internal/services"
)

// AdminHandler handles admin operations
type AdminHandler struct {
	contentService *services.ContentService
	config         *config.Config
}

// NewAdminHandler creates a new admin handler
func NewAdminHandler(contentService *services.ContentService, cfg *config.Config) *AdminHandler {
	return &AdminHandler{
		contentService: contentService,
		config:         cfg,
	}
}

// Dashboard renders the legacy admin dashboard
func (ah *AdminHandler) Dashboard(c *gin.Context) {
	sections, err := ah.contentService.GetPageSections(c.Request.Context(), "home")
	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin/error.html", gin.H{
			"error": "Failed to load content sections",
		})
		return
	}

	c.HTML(http.StatusOK, "admin/dashboard.html", gin.H{
		"title":       "Admin Dashboard",
		"user_name":   c.GetString("user_name"),
		"user_email":  c.GetString("user_email"),
		"sections":    sections,
		"projects":    []interface{}{}, // Empty for now - will be implemented later
		"experiences": []interface{}{}, // Empty for now - will be implemented later
		"skills":      []interface{}{}, // Empty for now - will be implemented later
	})
}

// NewDashboard renders the new admin dashboard
func (ah *AdminHandler) NewDashboard(c *gin.Context) {
	user, exists := c.Get("user")

	data := gin.H{
		"title":          "Admin Dashboard",
		"FirebaseConfig": ah.config.FirebaseConfig,
	}

	if exists {
		data["user"] = user
	}

	c.HTML(http.StatusOK, "admin/dashboard-new.html", data)
}

// Generic Page Management
func (ah *AdminHandler) GetAllPages(c *gin.Context) {
	pages, err := ah.contentService.GetAllPages(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get pages"})
		return
	}
	c.JSON(http.StatusOK, pages)
}

func (ah *AdminHandler) CreatePage(c *gin.Context) {
	// Custom form handling for page creation
	page := models.Page{
		Slug:        c.PostForm("slug"),
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		UpdatedBy:   c.PostForm("updated_by"),
	}

	// Handle visible checkbox
	page.Visible = c.PostForm("visible") == "true"

	// Handle metadata
	page.Metadata = models.PageMeta{
		Keywords:    c.PostForm("keywords"),
		Description: c.PostForm("meta_description"),
		OGImage:     c.PostForm("og_image"),
		Canonical:   c.PostForm("canonical"),
	}

	err := ah.contentService.CreatePage(c.Request.Context(), &page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create page"})
		return
	}

	c.JSON(http.StatusCreated, page)
}

func (ah *AdminHandler) UpdatePage(c *gin.Context) {
	pageID := c.Param("id")

	// Custom form handling for page update
	page := models.Page{
		ID:          pageID,
		Slug:        c.PostForm("slug"),
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		UpdatedBy:   c.PostForm("updated_by"),
	}

	// Handle visible checkbox
	page.Visible = c.PostForm("visible") == "true"

	// Handle metadata
	page.Metadata = models.PageMeta{
		Keywords:    c.PostForm("keywords"),
		Description: c.PostForm("meta_description"),
		OGImage:     c.PostForm("og_image"),
		Canonical:   c.PostForm("canonical"),
	}

	page.ID = pageID
	err := ah.contentService.UpdatePage(c.Request.Context(), &page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update page"})
		return
	}

	c.JSON(http.StatusOK, page)
}

func (ah *AdminHandler) DeletePage(c *gin.Context) {
	pageID := c.Param("id")

	err := ah.contentService.DeletePage(c.Request.Context(), pageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete page"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Page deleted successfully"})
}

// Generic Section Management
func (ah *AdminHandler) GetPageSections(c *gin.Context) {
	pageSlug := c.Param("pageSlug")

	sections, err := ah.contentService.GetPageSections(c.Request.Context(), pageSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get sections"})
		return
	}

	c.JSON(http.StatusOK, sections)
}

func (ah *AdminHandler) CreateSection(c *gin.Context) {
	pageSlug := c.Param("pageSlug")

	// Custom form handling for section creation
	section := models.Section{
		PageSlug:  pageSlug,
		Slug:      c.PostForm("slug"),
		Type:      c.PostForm("type"),
		Title:     c.PostForm("title"),
		Subtitle:  c.PostForm("subtitle"),
		Content:   c.PostForm("content"),
		Template:  c.PostForm("template"),
		CSS:       c.PostForm("css"),
		JS:        c.PostForm("js"),
		UpdatedBy: c.PostForm("updated_by"),
	}

	// Handle order field
	if orderStr := c.PostForm("order"); orderStr != "" {
		if order, err := strconv.Atoi(orderStr); err == nil {
			section.Order = order
		} else {
			section.Order = 1 // Default value
		}
	} else {
		section.Order = 1 // Default value
	}

	// Handle visible checkbox
	visibleValues := c.Request.PostForm["visible"]
	section.Visible = len(visibleValues) > 0 && visibleValues[0] == "true"

	// Handle JSON data field
	if dataStr := c.PostForm("data"); dataStr != "" {
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON in data field: " + err.Error()})
			return
		}
		section.Data = data
	}

	section.PageSlug = pageSlug
	err := ah.contentService.CreateSection(c.Request.Context(), &section)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create section"})
		return
	}

	c.JSON(http.StatusCreated, section)
}

func (ah *AdminHandler) UpdateSection(c *gin.Context) {
	sectionID := c.Param("id")

	// Custom form handling for section update
	section := models.Section{
		ID:        sectionID,
		PageSlug:  c.PostForm("page_slug"),
		Slug:      c.PostForm("slug"),
		Type:      c.PostForm("type"),
		Title:     c.PostForm("title"),
		Subtitle:  c.PostForm("subtitle"),
		Content:   c.PostForm("content"),
		Template:  c.PostForm("template"),
		CSS:       c.PostForm("css"),
		JS:        c.PostForm("js"),
		UpdatedBy: c.PostForm("updated_by"),
	}

	// Handle order field
	if orderStr := c.PostForm("order"); orderStr != "" {
		if order, err := strconv.Atoi(orderStr); err == nil {
			section.Order = order
		}
	}

	// Handle visible checkbox - if checkbox is checked, value="true" is sent, if unchecked, no value is sent
	visibleValues := c.Request.PostForm["visible"]
	section.Visible = len(visibleValues) > 0 && visibleValues[0] == "true"

	// Handle JSON data field
	if dataStr := c.PostForm("data"); dataStr != "" {
		var data map[string]interface{}
		if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON in data field: " + err.Error()})
			return
		}
		section.Data = data
	}

	section.ID = sectionID
	err := ah.contentService.UpdateSection(c.Request.Context(), &section)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update section"})
		return
	}

	c.JSON(http.StatusOK, section)
}

func (ah *AdminHandler) DeleteSection(c *gin.Context) {
	sectionID := c.Param("id")

	err := ah.contentService.DeleteSection(c.Request.Context(), sectionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete section"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Section deleted successfully"})
}

func (ah *AdminHandler) UpdateSectionOrder(c *gin.Context) {
	sectionID := c.Param("id")
	targetID := c.PostForm("target_id")

	err := ah.contentService.UpdateSectionOrder(c.Request.Context(), sectionID, targetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update section order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Section order updated"})
}

// Form rendering methods
func (ah *AdminHandler) RenderSectionForm(c *gin.Context) {
	sectionID := c.Query("id")

	data := gin.H{
		"title": "Create Section",
	}

	if sectionID != "" {
		section, err := ah.contentService.GetSection(c.Request.Context(), sectionID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Section not found"})
			return
		}

		data["title"] = "Edit Section"
		data["section"] = section
	}

	c.HTML(http.StatusOK, "admin/section-form.html", data)
}

func (ah *AdminHandler) RenderSectionEditForm(c *gin.Context) {
	sectionID := c.Param("id")

	section, err := ah.contentService.GetSection(c.Request.Context(), sectionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Section not found"})
		return
	}

	data := gin.H{
		"title":   "Edit Section",
		"section": section,
	}

	c.HTML(http.StatusOK, "admin/section-form.html", data)
}

// GetPageSectionsHTML returns the sections list as HTML for HTMX
func (ah *AdminHandler) GetPageSectionsHTML(c *gin.Context) {
	pageSlug := c.Param("pageSlug")

	sections, err := ah.contentService.GetPageSections(c.Request.Context(), pageSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get sections"})
		return
	}

	data := gin.H{
		"sections": sections,
	}

	c.HTML(http.StatusOK, "admin/sections-list.html", data)
}

// ToggleSectionVisibility toggles a section's visibility
func (ah *AdminHandler) ToggleSectionVisibility(c *gin.Context) {
	sectionID := c.Param("id")

	section, err := ah.contentService.GetSection(c.Request.Context(), sectionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Section not found"})
		return
	}

	section.Visible = !section.Visible

	err = ah.contentService.UpdateSection(c.Request.Context(), section)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update section"})
		return
	}

	// Return updated sections list for the page
	sections, err := ah.contentService.GetPageSections(c.Request.Context(), "home")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get sections"})
		return
	}

	data := gin.H{
		"sections": sections,
	}

	c.HTML(http.StatusOK, "admin/sections-list.html", data)
}
