package repositories

import (
	"time"
	"nestmate-backend/internal/domain/entities"
)

type ExpenseRepository interface {
	Create(expense *entities.Expense) error
	GetByID(id string) (*entities.Expense, error)
	GetByUserID(userID string) ([]*entities.Expense, error)
	GetByPeriod(userID string, start, end time.Time) ([]*entities.Expense, error)
	Update(expense *entities.Expense) error
	Delete(id string) error
}

type IncomeRepository interface {
	Create(income *entities.Income) error
	GetByUserID(userID string) ([]*entities.Income, error)
	GetByMonth(userID string, month time.Time) (*entities.Income, error)
	Update(income *entities.Income) error
	Delete(id string) error
}