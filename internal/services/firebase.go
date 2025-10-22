package services

import (
	"context"
	"log"

	"warrenclough.com/internal/config"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

type FirebaseService struct {
	client        *firestore.Client
	authClient    *auth.Client
	storageClient *storage.Client
	storageBucket string
	ctx           context.Context
}

func NewFirebaseService(ctx context.Context, cfg *config.Config) (*FirebaseService, error) {
	// Initialize Firebase app
	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: cfg.FirebaseProjectID,
	})
	if err != nil {
		return nil, err
	}

	// Initialize Firestore client
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	// Initialize Auth client
	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	// Initialize Storage client
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	// Storage bucket name (default: projectID.appspot.com)
	storageBucket := cfg.FirebaseProjectID + ".appspot.com"

	log.Printf("Firebase initialized for project: %s", cfg.FirebaseProjectID)

	return &FirebaseService{
		client:        client,
		authClient:    authClient,
		storageClient: storageClient,
		storageBucket: storageBucket,
		ctx:           ctx,
	}, nil
}

func (fs *FirebaseService) GetFirestore() *firestore.Client {
	return fs.client
}

func (fs *FirebaseService) GetAuthClient() *auth.Client {
	return fs.authClient
}

func (fs *FirebaseService) GetStorageClient() *storage.Client {
	return fs.storageClient
}

func (fs *FirebaseService) GetStorageBucket() string {
	return fs.storageBucket
}

func (fs *FirebaseService) Close() error {
	if err := fs.storageClient.Close(); err != nil {
		log.Printf("Error closing storage client: %v", err)
	}
	return fs.client.Close()
}
