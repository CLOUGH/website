package services

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"warrenclough.com/internal/config"
	"warrenclough.com/internal/models"
)

// ContentService handles content operations
type ContentService struct {
	fs     *FirebaseService
	config *config.Config
}

// NewContentService creates a new content service
func NewContentService(fs *FirebaseService, cfg *config.Config) *ContentService {
	return &ContentService{
		fs:     fs,
		config: cfg,
	}
}

// Page operations
func (cs *ContentService) GetAllPages(ctx context.Context) ([]models.Page, error) {
	var pages []models.Page

	iter := cs.fs.client.Collection("pages").Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var page models.Page
		if err := doc.DataTo(&page); err != nil {
			return nil, err
		}
		page.ID = doc.Ref.ID
		pages = append(pages, page)
	}

	return pages, nil
}

func (cs *ContentService) GetPage(ctx context.Context, pageSlug string) (*models.Page, error) {
	iter := cs.fs.client.Collection("pages").Where("slug", "==", pageSlug).Limit(1).Documents(ctx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err == iterator.Done {
		return nil, fmt.Errorf("page not found")
	}
	if err != nil {
		return nil, err
	}

	var page models.Page
	if err := doc.DataTo(&page); err != nil {
		return nil, err
	}
	page.ID = doc.Ref.ID

	return &page, nil
}

func (cs *ContentService) CreatePage(ctx context.Context, page *models.Page) error {
	page.CreatedAt = time.Now()
	page.UpdatedAt = time.Now()

	docRef, _, err := cs.fs.client.Collection("pages").Add(ctx, page)
	if err != nil {
		return err
	}

	page.ID = docRef.ID
	return nil
}

func (cs *ContentService) UpdatePage(ctx context.Context, page *models.Page) error {
	page.UpdatedAt = time.Now()

	_, err := cs.fs.client.Collection("pages").Doc(page.ID).Set(ctx, page)
	return err
}

func (cs *ContentService) DeletePage(ctx context.Context, pageID string) error {
	_, err := cs.fs.client.Collection("pages").Doc(pageID).Delete(ctx)
	return err
}

// Section operations
func (cs *ContentService) GetPageSections(ctx context.Context, pageSlug string) ([]models.Section, error) {
	var sections []models.Section

	iter := cs.fs.client.Collection("sections").
		Where("page_slug", "==", pageSlug).
		OrderBy("order", firestore.Asc).
		Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var section models.Section
		if err := doc.DataTo(&section); err != nil {
			return nil, err
		}
		section.ID = doc.Ref.ID
		sections = append(sections, section)
	}

	return sections, nil
}

func (cs *ContentService) GetSection(ctx context.Context, sectionID string) (*models.Section, error) {
	doc, err := cs.fs.client.Collection("sections").Doc(sectionID).Get(ctx)
	if err != nil {
		return nil, err
	}

	var section models.Section
	if err := doc.DataTo(&section); err != nil {
		return nil, err
	}
	section.ID = doc.Ref.ID

	return &section, nil
}

func (cs *ContentService) CreateSection(ctx context.Context, section *models.Section) error {
	section.CreatedAt = time.Now()
	section.UpdatedAt = time.Now()

	docRef, _, err := cs.fs.client.Collection("sections").Add(ctx, section)
	if err != nil {
		return err
	}

	section.ID = docRef.ID
	return nil
}

func (cs *ContentService) UpdateSection(ctx context.Context, section *models.Section) error {
	section.UpdatedAt = time.Now()

	_, err := cs.fs.client.Collection("sections").Doc(section.ID).Set(ctx, section)
	return err
}

func (cs *ContentService) DeleteSection(ctx context.Context, sectionID string) error {
	_, err := cs.fs.client.Collection("sections").Doc(sectionID).Delete(ctx)
	return err
}

func (cs *ContentService) UpdateSectionOrder(ctx context.Context, sectionID, targetID string) error {
	// Get the section to move
	sourceDoc, err := cs.fs.client.Collection("sections").Doc(sectionID).Get(ctx)
	if err != nil {
		return err
	}

	var sourceSection models.Section
	if err := sourceDoc.DataTo(&sourceSection); err != nil {
		return err
	}

	// Get the target section
	targetDoc, err := cs.fs.client.Collection("sections").Doc(targetID).Get(ctx)
	if err != nil {
		return err
	}

	var targetSection models.Section
	if err := targetDoc.DataTo(&targetSection); err != nil {
		return err
	}

	// Swap the order values
	sourceSection.Order, targetSection.Order = targetSection.Order, sourceSection.Order
	sourceSection.UpdatedAt = time.Now()
	targetSection.UpdatedAt = time.Now()

	// Update both sections
	batch := cs.fs.client.Batch()
	batch.Set(cs.fs.client.Collection("sections").Doc(sectionID), sourceSection)
	batch.Set(cs.fs.client.Collection("sections").Doc(targetID), targetSection)

	_, err = batch.Commit(ctx)
	return err
}

// User management methods (placeholder implementations)
func (cs *ContentService) CreateOrUpdateUser(ctx context.Context, user *models.User) error {
	// TODO: Implement user creation/update logic
	return nil
}

func (cs *ContentService) IsUserAdmin(ctx context.Context, email string) (bool, error) {
	// Get admin emails from config
	adminEmails := cs.config.AdminEmails

	// Check if user email is in admin list
	for _, adminEmail := range adminEmails {
		if email == adminEmail {
			return true, nil
		}
	}

	return false, nil
}
