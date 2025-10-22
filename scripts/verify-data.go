package main

import (
"context"
"encoding/json"
"fmt"
"log"

"warrenclough.com/internal/config"
"warrenclough.com/internal/services"
)

func main() {
	ctx := context.Background()
	cfg := config.Load()
	
	firebaseService, err := services.NewFirebaseService(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}
	defer firebaseService.Close()
	
	contentService := services.NewContentService(firebaseService)
	
	// Get all sections
	sections, err := contentService.GetPageSections(ctx, "home")
	if err != nil {
		log.Fatalf("Failed to get sections: %v", err)
	}
	
	fmt.Printf("Found %d sections:\n\n", len(sections))
	for _, section := range sections {
		fmt.Printf("📦 %s (%s)\n", section.Title, section.Type)
		fmt.Printf("   Order: %d, Visible: %v\n", section.Order, section.Visible)
		if section.Data != nil {
			data, _ := json.MarshalIndent(section.Data, "   ", "  ")
			fmt.Printf("   Data keys: %s\n", string(data)[:min(200, len(data))])
		}
		fmt.Println()
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
