package main

import (
"context"
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
	
	// Get page
	page, err := contentService.GetPage(ctx, "home")
	if err != nil {
		log.Fatalf("Failed to get page: %v", err)
	}
	
	fmt.Println("📄 PAGE")
	fmt.Printf("   Title: %s\n", page.Title)
	fmt.Printf("   Slug: %s\n", page.Slug)
	fmt.Printf("   Description: %s\n", page.Description)
	fmt.Println()
	
	// Get all sections
	sections, err := contentService.GetPageSections(ctx, "home")
	if err != nil {
		log.Fatalf("Failed to get sections: %v", err)
	}
	
	fmt.Printf("📦 SECTIONS (%d total)\n\n", len(sections))
	for _, section := range sections {
		fmt.Printf("%d. %s\n", section.Order, section.Title)
		fmt.Printf("   Type: %s\n", section.Type)
		fmt.Printf("   Slug: %s\n", section.Slug)
		fmt.Printf("   Visible: %v\n", section.Visible)
		
		// Show data summary
		if section.Data != nil {
			fmt.Printf("   Data Keys: ")
			i := 0
			for key := range section.Data {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Print(key)
				i++
			}
			fmt.Println()
		}
		fmt.Println()
	}
	
	fmt.Println("✅ All content verified in database!")
}
