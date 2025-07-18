package services

import (
	"nestmate-backend/internal/domain/entities"
)

type NotesService interface {
	CreateNote(note *entities.Note) error
	UpdateNote(id string, note *entities.Note) error
	DeleteNote(id string) error
	SearchNotes(userID string, query string) ([]*entities.Note, error)
	GetNotesByTags(userID string, tags []string) ([]*entities.Note, error)
	ExportNote(id string, format string) ([]byte, error)
	AddAttachment(noteID string, attachment *entities.Attachment) error
}