package services

import (
	"context"
	"errors"
	"nestmate-backend/internal/domain/entities"
)

// NotesService defines the interface for note operations
type NotesService interface {
	CreateNote(ctx context.Context, note *entities.Note) error
	UpdateNote(ctx context.Context, id string, note *entities.Note) error
	DeleteNote(ctx context.Context, id string) error
	SearchNotes(ctx context.Context, userID string, query string) ([]*entities.Note, error)
	GetNotesByTags(ctx context.Context, userID string, tags []string) ([]*entities.Note, error)
	ExportNote(ctx context.Context, id string, format string) ([]byte, error)
	AddAttachment(ctx context.Context, noteID string, attachment *entities.Attachment) error
}

// noteService implements the NotesService interface
type noteService struct {
	// Dependencies will be injected here
}

// NewNoteService creates a new note service
func NewNoteService() NotesService {
	return &noteService{}
}

// CreateNote creates a new note
func (s *noteService) CreateNote(ctx context.Context, note *entities.Note) error {
	// Implementation will be added in a future task
	return errors.New("not implemented")
}

// UpdateNote updates an existing note
func (s *noteService) UpdateNote(ctx context.Context, id string, note *entities.Note) error {
	// Implementation will be added in a future task
	return errors.New("not implemented")
}

// DeleteNote deletes a note
func (s *noteService) DeleteNote(ctx context.Context, id string) error {
	// Implementation will be added in a future task
	return errors.New("not implemented")
}

// SearchNotes searches notes by content
func (s *noteService) SearchNotes(ctx context.Context, userID string, query string) ([]*entities.Note, error) {
	// Implementation will be added in a future task
	return nil, errors.New("not implemented")
}

// GetNotesByTags gets notes by tags
func (s *noteService) GetNotesByTags(ctx context.Context, userID string, tags []string) ([]*entities.Note, error) {
	// Implementation will be added in a future task
	return nil, errors.New("not implemented")
}

// ExportNote exports a note in the specified format
func (s *noteService) ExportNote(ctx context.Context, id string, format string) ([]byte, error) {
	// Implementation will be added in a future task
	return nil, errors.New("not implemented")
}

// AddAttachment adds an attachment to a note
func (s *noteService) AddAttachment(ctx context.Context, noteID string, attachment *entities.Attachment) error {
	// Implementation will be added in a future task
	return errors.New("not implemented")
}