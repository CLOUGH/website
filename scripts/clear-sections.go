package main

import (
	"context"
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

	log.Printf("Found %d sections to delete", len(sections))

	// Delete all sections
	for _, section := range sections {
		err := contentService.DeleteSection(ctx, section.ID)
		if err != nil {
			log.Printf("Failed to delete section %s: %v", section.ID, err)
		} else {
			log.Printf("Deleted section: %s (%s)", section.Title, section.ID)
		}
	}

	log.Println("✅ All sections deleted!")
}
