package config

import (
	"os"
	"strings"
)

type Config struct {
	FirebaseProjectID string
	Port              string
	Environment       string
	FirebaseConfig    FirebaseWebConfig
	SMTPConfig        SMTPConfig
	AdminEmails       []string
}

type FirebaseWebConfig struct {
	APIKey            string
	AuthDomain        string
	ProjectID         string
	StorageBucket     string
	MessagingSenderID string
	AppID             string
}

type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

func Load() *Config {
	projectID := getEnv("FIREBASE_PROJECT_ID", "my-website-a3970")

	return &Config{
		FirebaseProjectID: projectID,
		Port:              getEnv("PORT", "8080"),
		Environment:       getEnv("ENVIRONMENT", "development"),
		FirebaseConfig: FirebaseWebConfig{
			APIKey:            getEnv("FIREBASE_API_KEY", ""),
			AuthDomain:        getEnv("FIREBASE_AUTH_DOMAIN", projectID+".firebaseapp.com"),
			ProjectID:         projectID,
			StorageBucket:     getEnv("FIREBASE_STORAGE_BUCKET", projectID+".appspot.com"),
			MessagingSenderID: getEnv("FIREBASE_MESSAGING_SENDER_ID", ""),
			AppID:             getEnv("FIREBASE_APP_ID", ""),
		},
		SMTPConfig: SMTPConfig{
			Host:     getEnv("SMTP_HOST", "smtp.gmail.com"),
			Port:     getEnv("SMTP_PORT", "587"),
			Username: getEnv("SMTP_USERNAME", ""),
			Password: getEnv("SMTP_PASSWORD", ""),
		},
		AdminEmails: getAdminEmails(),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getAdminEmails() []string {
	adminEmailsStr := getEnv("ADMIN_EMAILS", "clough.warren@gmail.com")
	emails := strings.Split(adminEmailsStr, ",")

	// Trim whitespace from each email
	for i, email := range emails {
		emails[i] = strings.TrimSpace(email)
	}

	return emails
}
