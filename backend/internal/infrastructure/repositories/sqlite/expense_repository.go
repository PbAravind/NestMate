package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"nestmate-backend/internal/domain/repositories"
	"time"
)

// ExpenseRepository implements the repositories.ExpenseRepository interface
type ExpenseRepository struct {
	db *sql.DB
}

// NewExpenseRepository creates a new SQLite expense repository
func NewExpenseRepository(db *sql.DB) repositories.ExpenseRepository {
	return &ExpenseRepository{
		db: db,
	}
}

// Create creates a new expense
func (r *ExpenseRepository) Create(ctx context.Context, expense *repositories.Expense) error {
	// Implementation will be added in a future task
	return errors.New("not implemented")
}

// GetByID gets an expense by ID
func (r *ExpenseRepository) GetByID(ctx context.Context, id string) (*repositories.Expense, error) {
	// Implementation will be added in a future task
	return nil, errors.New("not implemented")
}

// GetByUserIDAndDateRange gets expenses by user ID and date range
func (r *ExpenseRepository) GetByUserIDAndDateRange(ctx context.Context, userID string, startDate, endDate time.Time) ([]*repositories.Expense, error) {
	// Implementation will be added in a future task
	return nil, errors.New("not implemented")
}

// Update updates an expense
func (r *ExpenseRepository) Update(ctx context.Context, expense *repositories.Expense) error {
	// Implementation will be added in a future task
	return errors.New("not implemented")
}

// Delete deletes an expense by ID
func (r *ExpenseRepository) Delete(ctx context.Context, id string) error {
	// Implementation will be added in a future task
	return errors.New("not implemented")
}