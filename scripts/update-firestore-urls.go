package main

import (
	"context"
	"log"

	"warrenclough.com/internal/config"
	"warrenclough.com/internal/services"

	"cloud.google.com/go/firestore"
	"github.com/joho/godotenv"
	"google.golang.org/api/iterator"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize configuration
	cfg := config.Load()

	// Initialize Firebase
	ctx := context.Background()
	firebaseService, err := services.NewFirebaseService(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}
	defer firebaseService.Close()

	// Initialize storage service
	storageService := services.NewStorageService(firebaseService)
	firestoreClient := firebaseService.GetFirestore()

	log.Println("Updating Firestore documents with Firebase Storage URLs...")

	// Base URL for Firebase Storage
	baseURL := "https://storage.googleapis.com/my-website-a3970.appspot.com/static/"

	// Update projects with new screenshot URLs
	projectsRef := firestoreClient.Collection("sections")
	query := projectsRef.Where("type", "==", "project")

	iter := query.Documents(ctx)
	defer iter.Stop()

	updatedCount := 0
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error iterating projects: %v", err)
			continue
		}

		data := doc.Data()
		projectTitle := data["title"].(string)
		screenshots, ok := data["screenshots"].([]interface{})

		if !ok || len(screenshots) == 0 {
			continue
		}

		// Update screenshot URLs
		newScreenshots := make([]string, len(screenshots))
		for i, screenshot := range screenshots {
			oldURL := screenshot.(string)
			// Extract filename from old path
			filename := oldURL
			if len(oldURL) > 8 && oldURL[:8] == "/static/" {
				filename = oldURL[8:]
			}
			newScreenshots[i] = baseURL + filename
		}

		// Update document
		_, err = doc.Ref.Update(ctx, []firestore.Update{
			{
				Path:  "screenshots",
				Value: newScreenshots,
			},
		})

		if err != nil {
			log.Printf("Error updating project %s: %v", projectTitle, err)
			continue
		}

		log.Printf("✓ Updated project: %s (%d screenshots)", projectTitle, len(newScreenshots))
		updatedCount++
	}

	// Update hero section with background image
	heroQuery := projectsRef.Where("type", "==", "hero")
	heroIter := heroQuery.Documents(ctx)
	defer heroIter.Stop()

	for {
		doc, err := heroIter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error iterating hero sections: %v", err)
			continue
		}

		data := doc.Data()
		if bgImage, ok := data["background_image"].(string); ok && bgImage != "" {
			filename := bgImage
			if len(bgImage) > 8 && bgImage[:8] == "/static/" {
				filename = bgImage[8:]
			}
			newURL := baseURL + filename

			_, err = doc.Ref.Update(ctx, []firestore.Update{
				{
					Path:  "background_image",
					Value: newURL,
				},
			})

			if err != nil {
				log.Printf("Error updating hero background: %v", err)
			} else {
				log.Printf("✓ Updated hero background image")
				updatedCount++
			}
		}
	}

	// Add resume URL to a config document (create if doesn't exist)
	configRef := firestoreClient.Collection("config").Doc("site")
	resumeURL := storageService.GetPublicURL("static/MyResume-20250830-compressed.pdf")

	_, err = configRef.Set(ctx, map[string]interface{}{
		"resume_url":      resumeURL,
		"static_base_url": baseURL,
	})

	if err != nil {
		log.Printf("Error updating config: %v", err)
	} else {
		log.Printf("✓ Updated site config with resume URL")
		updatedCount++
	}

	log.Printf("\n=== Update Complete ===")
	log.Printf("Total documents updated: %d\n", updatedCount)
	log.Println("\nAll static files are now served from Firebase Storage!")
	log.Println("Resume URL:", resumeURL)
}
