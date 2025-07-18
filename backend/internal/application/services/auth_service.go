package services

import (
	"context"
	"errors"
	"nestmate-backend/internal/domain/entities"
)

// AuthService defines the interface for authentication operations
type AuthService interface {
	Register(ctx context.Context, email, password string) (*entities.User, error)
	Login(ctx context.Context, email, password string) (*entities.AuthToken, error)
	Logout(ctx context.Context, token string) error
	ValidateToken(ctx context.Context, token string) (*entities.User, error)
	RefreshToken(ctx context.Context, refreshToken string) (*entities.AuthToken, error)
}

// authService implements the AuthService interface
type authService struct {
	// Dependencies will be injected here
}

// NewAuthService creates a new auth service
func NewAuthService() AuthService {
	return &authService{}
}

// Register creates a new user account
func (s *authService) Register(ctx context.Context, email, password string) (*entities.User, error) {
	// Implementation will be added in a future task
	return nil, errors.New("not implemented")
}

// Login authenticates a user and returns a token
func (s *authService) Login(ctx context.Context, email, password string) (*entities.AuthToken, error) {
	// Implementation will be added in a future task
	return nil, errors.New("not implemented")
}

// Logout invalidates a user's token
func (s *authService) Logout(ctx context.Context, token string) error {
	// Implementation will be added in a future task
	return errors.New("not implemented")
}

// ValidateToken validates a token and returns the associated user
func (s *authService) ValidateToken(ctx context.Context, token string) (*entities.User, error) {
	// Implementation will be added in a future task
	return nil, errors.New("not implemented")
}

// RefreshToken refreshes an expired token
func (s *authService) RefreshToken(ctx context.Context, refreshToken string) (*entities.AuthToken, error) {
	// Implementation will be added in a future task
	return nil, errors.New("not implemented")
}