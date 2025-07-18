package repositories

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

// ExpenseRepository defines the interface for expense data access
type ExpenseRepository interface {
	// Create a new expense
	Create(ctx context.Context, expense *Expense) error
	
	// Get an expense by ID
	GetByID(ctx context.Context, id string) (*Expense, error)
	
	// Get expenses by user ID and date range
	GetByUserIDAndDateRange(ctx context.Context, userID string, startDate, endDate time.Time) ([]*Expense, error)
	
	// Update an expense
	Update(ctx context.Context, expense *Expense) error
	
	// Delete an expense by ID
	Delete(ctx context.Context, id string) error
}

// Expense represents the repository expense model
type Expense struct {
	ID          string
	UserID      string
	Amount      decimal.Decimal
	Description string
	Date        time.Time
	MainCategory string
	SubCategory  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}