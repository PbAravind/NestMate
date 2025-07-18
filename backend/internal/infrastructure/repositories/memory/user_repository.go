package memory

import (
	"context"
	"fmt"
	"sync"

	"nestmate-backend/internal/domain/entities"
	"nestmate-backend/internal/domain/repositories"
)

// InMemoryUserRepository implements UserRepository using in-memory storage
// This is a temporary implementation for development/testing
type InMemoryUserRepository struct {
	users map[string]*entities.User
	mutex sync.RWMutex
}

// NewInMemoryUserRepository creates a new in-memory user repository
func NewInMemoryUserRepository() repositories.UserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*entities.User),
	}
}

// Create creates a new user
func (r *InMemoryUserRepository) Create(ctx context.Context, user *entities.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	// Check if user already exists
	if _, exists := r.users[user.ID]; exists {
		return fmt.Errorf("user with ID %s already exists", user.ID)
	}
	
	// Check if email already exists
	for _, existingUser := range r.users {
		if existingUser.Email == user.Email {
			return fmt.Errorf("user with email %s already exists", user.Email)
		}
	}
	
	r.users[user.ID] = user
	return nil
}

// GetByID retrieves a user by ID
func (r *InMemoryUserRepository) GetByID(ctx context.Context, id string) (*entities.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}
	
	return user, nil
}

// GetByEmail retrieves a user by email
func (r *InMemoryUserRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	
	return nil, fmt.Errorf("user with email %s not found", email)
}

// Update updates an existing user
func (r *InMemoryUserRepository) Update(ctx context.Context, user *entities.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	if _, exists := r.users[user.ID]; !exists {
		return fmt.Errorf("user with ID %s not found", user.ID)
	}
	
	r.users[user.ID] = user
	return nil
}

// Delete deletes a user by ID
func (r *InMemoryUserRepository) Delete(ctx context.Context, id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	
	if _, exists := r.users[id]; !exists {
		return fmt.Errorf("user with ID %s not found", id)
	}
	
	delete(r.users, id)
	return nil
}