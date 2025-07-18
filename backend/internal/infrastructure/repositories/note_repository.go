package repositories

import (
	"nestmate-backend/internal/domain/entities"
)

type NoteRepository interface {
	Create(note *entities.Note) error
	GetByID(id string) (*entities.Note, error)
	GetByUserID(userID string) ([]*entities.Note, error)
	SearchByContent(userID string, query string) ([]*entities.Note, error)
	GetByTags(userID string, tags []string) ([]*entities.Note, error)
	Update(note *entities.Note) error
	Delete(id string) error
}