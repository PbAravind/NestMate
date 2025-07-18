package repositories

import (
	"time"
	"nestmate-backend/internal/domain/entities"
)

type TaskRepository interface {
	Create(task *entities.Task) error
	GetByID(id string) (*entities.Task, error)
	GetByUserID(userID string) ([]*entities.Task, error)
	GetByPeriod(userID string, start, end time.Time) ([]*entities.Task, error)
	Update(task *entities.Task) error
	Delete(id string) error
}