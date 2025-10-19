package models

import (
	"time"
)

// BlogPost represents a blog post
type BlogPost struct {
	ID          string    `json:"id" firestore:"id"`
	Title       string    `json:"title" firestore:"title"`
	Slug        string    `json:"slug" firestore:"slug"`
	Content     string    `json:"content" firestore:"content"`
	Excerpt     string    `json:"excerpt" firestore:"excerpt"`
	Author      string    `json:"author" firestore:"author"`
	PublishedAt time.Time `json:"published_at" firestore:"published_at"`
	UpdatedAt   time.Time `json:"updated_at" firestore:"updated_at"`
	Tags        []string  `json:"tags" firestore:"tags"`
	Featured    bool      `json:"featured" firestore:"featured"`
	ImageURL    string    `json:"image_url" firestore:"image_url"`
}

// PortfolioProject represents a portfolio project
type PortfolioProject struct {
	ID          string    `json:"id" firestore:"id"`
	Title       string    `json:"title" firestore:"title"`
	Description string    `json:"description" firestore:"description"`
	LongDescription string `json:"long_description" firestore:"long_description"`
	ImageURL    string    `json:"image_url" firestore:"image_url"`
	LiveURL     string    `json:"live_url" firestore:"live_url"`
	GitHubURL   string    `json:"github_url" firestore:"github_url"`
	Technologies []string `json:"technologies" firestore:"technologies"`
	Featured    bool      `json:"featured" firestore:"featured"`
	CreatedAt   time.Time `json:"created_at" firestore:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" firestore:"updated_at"`
}

// Resume represents resume information
type Resume struct {
	ID          string    `json:"id" firestore:"id"`
	Name        string    `json:"name" firestore:"name"`
	Title       string    `json:"title" firestore:"title"`
	Email       string    `json:"email" firestore:"email"`
	Phone       string    `json:"phone" firestore:"phone"`
	Location    string    `json:"location" firestore:"location"`
	Summary     string    `json:"summary" firestore:"summary"`
	Experience  []Experience `json:"experience" firestore:"experience"`
	Education   []Education  `json:"education" firestore:"education"`
	Skills      []Skill      `json:"skills" firestore:"skills"`
	UpdatedAt   time.Time `json:"updated_at" firestore:"updated_at"`
}

type Experience struct {
	Company     string    `json:"company" firestore:"company"`
	Position    string    `json:"position" firestore:"position"`
	Location    string    `json:"location" firestore:"location"`
	StartDate   time.Time `json:"start_date" firestore:"start_date"`
	EndDate     *time.Time `json:"end_date" firestore:"end_date"`
	Description string    `json:"description" firestore:"description"`
	Technologies []string `json:"technologies" firestore:"technologies"`
}

type Education struct {
	Institution string    `json:"institution" firestore:"institution"`
	Degree      string    `json:"degree" firestore:"degree"`
	Field       string    `json:"field" firestore:"field"`
	StartDate   time.Time `json:"start_date" firestore:"start_date"`
	EndDate     time.Time `json:"end_date" firestore:"end_date"`
	GPA         string    `json:"gpa" firestore:"gpa"`
}

type Skill struct {
	Name        string `json:"name" firestore:"name"`
	Category    string `json:"category" firestore:"category"`
	Proficiency int    `json:"proficiency" firestore:"proficiency"` // 1-5 scale
}
