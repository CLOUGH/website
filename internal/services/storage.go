package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime"
	"path/filepath"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type StorageService struct {
	client     *storage.Client
	bucketName string
	ctx        context.Context
}

func NewStorageService(firebaseService *FirebaseService) *StorageService {
	return &StorageService{
		client:     firebaseService.GetStorageClient(),
		bucketName: firebaseService.GetStorageBucket(),
		ctx:        firebaseService.ctx,
	}
}

// UploadFile uploads a file to Firebase Storage
func (s *StorageService) UploadFile(filePath string, fileData io.Reader, makePublic bool) (string, error) {
	bucket := s.client.Bucket(s.bucketName)
	obj := bucket.Object(filePath)

	// Create writer
	writer := obj.NewWriter(s.ctx)

	// Set content type based on file extension
	contentType := mime.TypeByExtension(filepath.Ext(filePath))
	if contentType != "" {
		writer.ContentType = contentType
	}

	// Set cache control for better performance
	writer.CacheControl = "public, max-age=31536000" // 1 year

	// Copy file data
	if _, err := io.Copy(writer, fileData); err != nil {
		writer.Close()
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %w", err)
	}

	// Make public if requested
	if makePublic {
		if err := obj.ACL().Set(s.ctx, storage.AllUsers, storage.RoleReader); err != nil {
			log.Printf("Warning: Failed to make file public: %v", err)
		}
	}

	// Return public URL
	publicURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", s.bucketName, filePath)
	return publicURL, nil
}

// GetPublicURL returns the public URL for a file
func (s *StorageService) GetPublicURL(filePath string) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", s.bucketName, filePath)
}

// GetSignedURL creates a temporary signed URL for private files
func (s *StorageService) GetSignedURL(filePath string, expiration time.Duration) (string, error) {
	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: time.Now().Add(expiration),
	}

	url, err := storage.SignedURL(s.bucketName, filePath, opts)
	if err != nil {
		return "", fmt.Errorf("failed to create signed URL: %w", err)
	}

	return url, nil
}

// DeleteFile deletes a file from storage
func (s *StorageService) DeleteFile(filePath string) error {
	bucket := s.client.Bucket(s.bucketName)
	obj := bucket.Object(filePath)

	if err := obj.Delete(s.ctx); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

// ListFiles lists all files in a directory
func (s *StorageService) ListFiles(prefix string) ([]string, error) {
	bucket := s.client.Bucket(s.bucketName)
	query := &storage.Query{Prefix: prefix}

	var files []string
	it := bucket.Objects(s.ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to list files: %w", err)
		}
		files = append(files, attrs.Name)
	}

	return files, nil
}

// FileExists checks if a file exists in storage
func (s *StorageService) FileExists(filePath string) (bool, error) {
	bucket := s.client.Bucket(s.bucketName)
	obj := bucket.Object(filePath)

	_, err := obj.Attrs(s.ctx)
	if err == storage.ErrObjectNotExist {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
