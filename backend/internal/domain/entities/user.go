package entities

import (
	"time"
)

// User represents a user entity
type User struct {
	ID        string
	Email     string
	Password  string // Hashed password
	CreatedAt time.Time
	UpdatedAt time.Time
}

// AuthToken represents an authentication token
type AuthToken struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}