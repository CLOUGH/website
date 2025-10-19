package config

import (
	"os"
)

type Config struct {
	FirebaseProjectID string
	Port              string
	Environment       string
}

func Load() *Config {
	return &Config{
		FirebaseProjectID: getEnv("FIREBASE_PROJECT_ID", "warrenclough-com"),
		Port:              getEnv("PORT", "8080"),
		Environment:       getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
