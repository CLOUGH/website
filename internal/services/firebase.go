package services

import (
	"context"
	"log"

	"warrenclough.com/internal/config"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
)

type FirebaseService struct {
	client *firestore.Client
	ctx    context.Context
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

	log.Printf("Firebase initialized for project: %s", cfg.FirebaseProjectID)

	return &FirebaseService{
		client: client,
		ctx:    ctx,
	}, nil
}

func (fs *FirebaseService) GetClient() *firestore.Client {
	return fs.client
}

func (fs *FirebaseService) Close() error {
	return fs.client.Close()
}
