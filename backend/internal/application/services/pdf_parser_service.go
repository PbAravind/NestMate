package services

import (
	"context"
	"errors"
	"nestmate-backend/internal/domain/entities"
)

// Transaction represents a parsed bank transaction
type Transaction struct {
	Date        string
	Description string
	Amount      string
	Type        string // Credit or Debit
}

// CategorizedTransaction represents a transaction with category
type CategorizedTransaction struct {
	Transaction
	MainCategory entities.MainCategory
	SubCategory  entities.SubCategory
	Confidence   float64 // Confidence score of categorization
}

// ParsedTransactions represents the result of parsing a bank statement
type ParsedTransactions struct {
	Transactions []*Transaction
	BankName     string
	AccountLast4 string
	StatementPeriod string
}

// PDFParserService defines the interface for PDF parsing operations
type PDFParserService interface {
	ParseBankStatement(ctx context.Context, file []byte, bankType string) (*ParsedTransactions, error)
	CategorizeTransactions(ctx context.Context, transactions []*Transaction) ([]*CategorizedTransaction, error)
	ValidateAndSave(ctx context.Context, userID string, transactions []*CategorizedTransaction) error
}

// pdfParserService implements the PDFParserService interface
type pdfParserService struct {
	// Dependencies will be injected here
}

// NewPDFParserService creates a new PDF parser service
func NewPDFParserService() PDFParserService {
	return &pdfParserService{}
}

// ParseBankStatement parses a bank statement PDF
func (s *pdfParserService) ParseBankStatement(ctx context.Context, file []byte, bankType string) (*ParsedTransactions, error) {
	// Implementation will be added in a future task
	return nil, errors.New("not implemented")
}

// CategorizeTransactions categorizes transactions based on rules
func (s *pdfParserService) CategorizeTransactions(ctx context.Context, transactions []*Transaction) ([]*CategorizedTransaction, error) {
	// Implementation will be added in a future task
	return nil, errors.New("not implemented")
}

// ValidateAndSave validates and saves categorized transactions
func (s *pdfParserService) ValidateAndSave(ctx context.Context, userID string, transactions []*CategorizedTransaction) error {
	// Implementation will be added in a future task
	return errors.New("not implemented")
}