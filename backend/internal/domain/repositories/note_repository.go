package repositories

import (
	"context"
	"time"
)

// NoteRepository defines the interface for note data access
type NoteRepository interface {
	// Create a new note
	Create(ctx context.Context, note *Note) error
	
	// Get a note by ID
	GetByID(ctx context.Context, id string) (*Note, error)
	
	// Get notes by user ID
	GetByUserID(ctx context.Context, userID string) ([]*Note, error)
	
	// Get notes by tags
	GetByTags(ctx context.Context, userID string, tags []string) ([]*Note, error)
	
	// Search notes by content
	Search(ctx context.Context, userID string, query string) ([]*Note, error)
	
	// Update a note
	Update(ctx context.Context, note *Note) error
	
	// Delete a note by ID
	Delete(ctx context.Context, id string) error
}

// Note represents the repository note model
type Note struct {
	ID          string
	UserID      string
	Title       string
	Content     string
	Tags        []string
	IsFavorite  bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Attachment represents a note attachment
type Attachment struct {
	ID       string
	NoteID   string
	Type     string
	URL      string
	Metadata map[string]interface{}
	CreatedAt time.Time
}