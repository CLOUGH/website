package models

import "time"

// Page represents a web page (home, blog, etc.)
type Page struct {
	ID          string    `json:"id" firestore:"id"`
	Slug        string    `json:"slug" firestore:"slug"` // "home", "blog", "about", etc.
	Title       string    `json:"title" firestore:"title"`
	Description string    `json:"description" firestore:"description"`
	Metadata    PageMeta  `json:"metadata" firestore:"metadata"`
	Visible     bool      `json:"visible" firestore:"visible"`
	CreatedAt   time.Time `json:"created_at" firestore:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" firestore:"updated_at"`
	UpdatedBy   string    `json:"updated_by" firestore:"updated_by"`
}

// PageMeta represents page metadata for SEO
type PageMeta struct {
	Keywords    string `json:"keywords" firestore:"keywords"`
	Description string `json:"description" firestore:"description"`
	OGImage     string `json:"og_image,omitempty" firestore:"og_image,omitempty"`
	Canonical   string `json:"canonical,omitempty" firestore:"canonical,omitempty"`
}

// Section represents a content section within a page
type Section struct {
	ID        string                 `json:"id" firestore:"id"`
	PageSlug  string                 `json:"page_slug" firestore:"page_slug"` // Reference to parent page
	Slug      string                 `json:"slug" firestore:"slug"`           // Unique identifier within page
	Type      string                 `json:"type" firestore:"type"`           // "hero", "about", "experience", "projects", "skills", "contact"
	Title     string                 `json:"title" firestore:"title"`
	Subtitle  string                 `json:"subtitle,omitempty" firestore:"subtitle,omitempty"`
	Content   string                 `json:"content" firestore:"content"`
	Data      map[string]interface{} `json:"data,omitempty" firestore:"data,omitempty"` // Flexible JSON data for section-specific content
	Order     int                    `json:"order" firestore:"order"`
	Visible   bool                   `json:"visible" firestore:"visible"`
	Template  string                 `json:"template,omitempty" firestore:"template,omitempty"` // Optional custom template name
	CSS       string                 `json:"css,omitempty" firestore:"css,omitempty"`           // Optional custom CSS
	JS        string                 `json:"js,omitempty" firestore:"js,omitempty"`             // Optional custom JavaScript
	CreatedAt time.Time              `json:"created_at" firestore:"created_at"`
	UpdatedAt time.Time              `json:"updated_at" firestore:"updated_at"`
	UpdatedBy string                 `json:"updated_by" firestore:"updated_by"`
}

// SectionData represents common section data structures
type SectionData struct {
	// Hero section data
	HeroImage   string   `json:"hero_image,omitempty"`
	HeroButtons []Button `json:"hero_buttons,omitempty"`

	// Experience section data
	Experiences []WorkExperience `json:"experiences,omitempty"`

	// Projects section data
	Projects []Project `json:"projects,omitempty"`

	// Skills section data
	SkillCategories map[string][]TechSkill `json:"skill_categories,omitempty"`

	// Contact section data
	ContactForm ContactFormConfig `json:"contact_form,omitempty"`
	SocialLinks []SocialLink      `json:"social_links,omitempty"`

	// Generic list data
	Items []ContentItem `json:"items,omitempty"`
}

// Button represents a button element
type Button struct {
	Text   string `json:"text"`
	URL    string `json:"url"`
	Style  string `json:"style"`            // "primary", "secondary", "outline"
	Target string `json:"target,omitempty"` // "_blank", etc.
}

// ContactFormConfig represents contact form configuration
type ContactFormConfig struct {
	Enabled bool     `json:"enabled"`
	Fields  []string `json:"fields"` // "name", "email", "subject", "message"
	Action  string   `json:"action"` // Form submission endpoint
}

// SocialLink represents a social media link
type SocialLink struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Icon string `json:"icon"` // Icon class or SVG
}

// ContentItem represents a generic content item
type ContentItem struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Image       string                 `json:"image,omitempty"`
	URL         string                 `json:"url,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
	Order       int                    `json:"order"`
	Visible     bool                   `json:"visible"`
}

// Legacy models for backward compatibility and data migration

// Project represents a project item (for migration)
type Project struct {
	ID               string    `json:"id" firestore:"id"`
	Title            string    `json:"title" firestore:"title"`
	ShortDescription string    `json:"short_description" firestore:"short_description"`
	Description      string    `json:"description" firestore:"description"`
	FullDescription  string    `json:"full_description" firestore:"full_description"`
	Technologies     []string  `json:"technologies" firestore:"technologies"`
	Screenshots      []string  `json:"screenshots" firestore:"screenshots"`
	GitHubURL        string    `json:"github_url,omitempty" firestore:"github_url,omitempty"`
	LiveURL          string    `json:"live_url,omitempty" firestore:"live_url,omitempty"`
	VideoURL         string    `json:"video_url,omitempty" firestore:"video_url,omitempty"`
	Featured         bool      `json:"featured" firestore:"featured"`
	Order            int       `json:"order" firestore:"order"`
	Visible          bool      `json:"visible" firestore:"visible"`
	CreatedAt        time.Time `json:"created_at" firestore:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" firestore:"updated_at"`
}

// WorkExperience represents a work experience item (for migration)
type WorkExperience struct {
	ID           string    `json:"id" firestore:"id"`
	Company      string    `json:"company" firestore:"company"`
	Position     string    `json:"position" firestore:"position"`
	Location     string    `json:"location" firestore:"location"`
	StartDate    string    `json:"start_date" firestore:"start_date"`
	EndDate      string    `json:"end_date,omitempty" firestore:"end_date,omitempty"`
	Description  []string  `json:"description" firestore:"description"`
	Technologies []string  `json:"technologies" firestore:"technologies"`
	Order        int       `json:"order" firestore:"order"`
	Visible      bool      `json:"visible" firestore:"visible"`
	CreatedAt    time.Time `json:"created_at" firestore:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" firestore:"updated_at"`
}

// TechSkill represents a skill/technology (for migration)
type TechSkill struct {
	ID        string    `json:"id" firestore:"id"`
	Name      string    `json:"name" firestore:"name"`
	Category  string    `json:"category" firestore:"category"` // "languages", "cloud", "databases", etc.
	Order     int       `json:"order" firestore:"order"`
	Visible   bool      `json:"visible" firestore:"visible"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
	UpdatedAt time.Time `json:"updated_at" firestore:"updated_at"`
}

// ContentSection represents a content section on the website (legacy, for migration)
type ContentSection struct {
	ID        string                 `json:"id" firestore:"id"`
	Type      string                 `json:"type" firestore:"type"` // "hero", "about", "experience", "project", "skill", "contact"
	Title     string                 `json:"title" firestore:"title"`
	Subtitle  string                 `json:"subtitle,omitempty" firestore:"subtitle,omitempty"`
	Content   string                 `json:"content" firestore:"content"`
	Data      map[string]interface{} `json:"data,omitempty" firestore:"data,omitempty"` // Additional structured data
	Order     int                    `json:"order" firestore:"order"`
	Visible   bool                   `json:"visible" firestore:"visible"`
	CreatedAt time.Time              `json:"created_at" firestore:"created_at"`
	UpdatedAt time.Time              `json:"updated_at" firestore:"updated_at"`
	UpdatedBy string                 `json:"updated_by" firestore:"updated_by"`
}

// User represents an authenticated user
type User struct {
	ID        string    `json:"id" firestore:"id"`
	Email     string    `json:"email" firestore:"email"`
	Name      string    `json:"name" firestore:"name"`
	Picture   string    `json:"picture" firestore:"picture"`
	IsAdmin   bool      `json:"is_admin" firestore:"is_admin"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
	LastLogin time.Time `json:"last_login" firestore:"last_login"`
}
