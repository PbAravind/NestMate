package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

// MainCategory represents the main expense category
type MainCategory string

const (
	ChennaiHouse  MainCategory = "Chennai House"
	BangaloreHouse MainCategory = "Bangalore House"
	Self          MainCategory = "Self"
	Savings       MainCategory = "Savings"
)

// SubCategory represents the expense sub-category
type SubCategory string

const (
	Food         SubCategory = "Food"
	Entertainment SubCategory = "Entertainment"
	Education    SubCategory = "Education"
	Travel       SubCategory = "Travel"
	Misc         SubCategory = "Misc"
)

// Expense represents an expense entity
type Expense struct {
	ID          string
	UserID      string
	Amount      decimal.Decimal
	Description string
	Date        time.Time
	MainCategory MainCategory
	SubCategory  SubCategory
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Income represents an income entity
type Income struct {
	ID        string
	UserID    string
	Amount    decimal.Decimal
	Month     time.Time
	Source    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// MonthlyBreakdown represents a monthly expense breakdown
type MonthlyBreakdown struct {
	Month          time.Time
	TotalIncome    decimal.Decimal
	TotalExpenses  decimal.Decimal
	Savings        decimal.Decimal
	CategoryBreakdown map[MainCategory]decimal.Decimal
	SubCategoryBreakdown map[SubCategory]decimal.Decimal
}