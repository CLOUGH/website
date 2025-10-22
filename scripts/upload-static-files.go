package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"warrenclough.com/internal/config"
	"warrenclough.com/internal/services"

	"github.com/joho/godotenv"
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

	// Directory containing static files
	staticDir := "./web/static"

	// Get list of files to upload
	files, err := os.ReadDir(staticDir)
	if err != nil {
		log.Fatalf("Failed to read static directory: %v", err)
	}

	log.Println("Starting file upload to Firebase Storage...")
	uploadedFiles := make(map[string]string) // local path -> public URL

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Skip manifest.json as it should stay local
		if file.Name() == "manifest.json" {
			log.Printf("Skipping manifest.json (keeping local)")
			continue
		}

		localPath := filepath.Join(staticDir, file.Name())
		storagePath := "static/" + file.Name()

		// Open file
		f, err := os.Open(localPath)
		if err != nil {
			log.Printf("Error opening %s: %v", localPath, err)
			continue
		}

		// Upload file
		publicURL, err := storageService.UploadFile(storagePath, f, true)
		f.Close()

		if err != nil {
			log.Printf("Error uploading %s: %v", file.Name(), err)
			continue
		}

		uploadedFiles[file.Name()] = publicURL
		log.Printf("✓ Uploaded: %s -> %s", file.Name(), publicURL)
	}

	// Print summary
	fmt.Println("\n=== Upload Summary ===")
	fmt.Printf("Total files uploaded: %d\n\n", len(uploadedFiles))

	fmt.Println("File URLs:")
	for filename, url := range uploadedFiles {
		fmt.Printf("  %s:\n    %s\n", filename, url)
	}

	fmt.Println("\n=== Next Steps ===")
	fmt.Println("1. Update your templates to reference these URLs")
	fmt.Println("2. Store URLs in Firestore documents (e.g., resume_url field)")
	fmt.Println("3. Test that all images load correctly")
	fmt.Println("4. Optional: Delete local files from web/static after verification")
}
