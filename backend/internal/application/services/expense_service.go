package services

import (
	"context"
	"errors"
	"nestmate-backend/internal/domain/entities"
	"time"

	"github.com/shopspring/decimal"
)

// ExpenseService defines the interface for expense operations
type ExpenseService interface {
	AddExpense(ctx context.Context, expense *entities.Expense) error
	UpdateExpense(ctx context.Context, id string, expense *entities.Expense) error
	DeleteExpense(ctx context.Context, id string) error
	GetExpensesByPeriod(ctx context.Context, userID string, start, end time.Time) ([]*entities.Expense, error)
	GetMonthlyBreakdown(ctx context.Context, userID string, month time.Time) (*entities.MonthlyBreakdown, error)
	CalculateSavings(ctx context.Context, userID string, month time.Time) (decimal.Decimal, error)
	ExportData(ctx context.Context, userID string, format string) ([]byte, error)
}

// expenseService implements the ExpenseService interface
type expenseService struct {
	// Dependencies will be injected here
}

// NewExpenseService creates a new expense service
func NewExpenseService() ExpenseService {
	return &expenseService{}
}

// AddExpense adds a new expense
func (s *expenseService) AddExpense(ctx context.Context, expense *entities.Expense) error {
	// Implementation will be added in a future task
	return errors.New("not implemented")
}

// UpdateExpense updates an existing expense
func (s *expenseService) UpdateExpense(ctx context.Context, id string, expense *entities.Expense) error {
	// Implementation will be added in a future task
	return errors.New("not implemented")
}

// DeleteExpense deletes an expense
func (s *expenseService) DeleteExpense(ctx context.Context, id string) error {
	// Implementation will be added in a future task
	return errors.New("not implemented")
}

// GetExpensesByPeriod gets expenses for a specific period
func (s *expenseService) GetExpensesByPeriod(ctx context.Context, userID string, start, end time.Time) ([]*entities.Expense, error) {
	// Implementation will be added in a future task
	return nil, errors.New("not implemented")
}

// GetMonthlyBreakdown gets a monthly breakdown of expenses
func (s *expenseService) GetMonthlyBreakdown(ctx context.Context, userID string, month time.Time) (*entities.MonthlyBreakdown, error) {
	// Implementation will be added in a future task
	return nil, errors.New("not implemented")
}

// CalculateSavings calculates savings for a month
func (s *expenseService) CalculateSavings(ctx context.Context, userID string, month time.Time) (decimal.Decimal, error) {
	// Implementation will be added in a future task
	return decimal.Zero, errors.New("not implemented")
}

// ExportData exports expense data in the specified format
func (s *expenseService) ExportData(ctx context.Context, userID string, format string) ([]byte, error) {
	// Implementation will be added in a future task
	return nil, errors.New("not implemented")
}