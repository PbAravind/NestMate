package services

import (
	"context"
	"fmt"
	"time"

	"nestmate-backend/internal/domain/entities"
	"nestmate-backend/internal/domain/repositories"
	"nestmate-backend/internal/infrastructure/auth"
)

// AuthService handles authentication business logic
type AuthService struct {
	firebaseAuth   *auth.FirebaseAuthService
	userRepository repositories.UserRepository
}

// NewAuthService creates a new authentication service
func NewAuthService(firebaseAuth *auth.FirebaseAuthService, userRepo repositories.UserRepository) *AuthService {
	return &AuthService{
		firebaseAuth:   firebaseAuth,
		userRepository: userRepo,
	}
}

// Register creates a new user account
func (s *AuthService) Register(ctx context.Context, req *entities.AuthRequest) (*entities.AuthResponse, error) {
	if s.firebaseAuth == nil {
		return nil, fmt.Errorf("Firebase Auth is not configured")
	}
	
	// Check if user already exists
	existingUser, err := s.userRepository.GetByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", req.Email)
	}
	
	// Create user in Firebase Auth
	user, err := s.firebaseAuth.CreateUser(ctx, req.Email, req.Password, req.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to create user in Firebase: %w", err)
	}
	
	// Save user to our database
	err = s.userRepository.Create(ctx, user)
	if err != nil {
		// If saving to our DB fails, we should clean up the Firebase user
		_ = s.firebaseAuth.DeleteUser(ctx, user.ID)
		return nil, fmt.Errorf("failed to save user to database: %w", err)
	}
	
	// Create custom token for immediate login
	customToken, err := s.firebaseAuth.CreateCustomToken(ctx, user.ID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create custom token: %w", err)
	}
	
	token := &entities.AuthToken{
		AccessToken: customToken,
		ExpiresAt:   time.Now().Add(time.Hour * 24), // 24 hours
		TokenType:   "Bearer",
	}
	
	return &entities.AuthResponse{
		User:  user,
		Token: token,
	}, nil
}

// Login authenticates a user (Firebase handles the actual authentication on client side)
func (s *AuthService) Login(ctx context.Context, idToken string) (*entities.AuthResponse, error) {
	if s.firebaseAuth == nil {
		return nil, fmt.Errorf("Firebase Auth is not configured")
	}
	
	// Verify the ID token from Firebase client SDK
	user, err := s.firebaseAuth.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("invalid ID token: %w", err)
	}
	
	// Get or create user in our database
	dbUser, err := s.userRepository.GetByID(ctx, user.ID)
	if err != nil {
		// User doesn't exist in our DB, create them
		dbUser = user
		err = s.userRepository.Create(ctx, dbUser)
		if err != nil {
			return nil, fmt.Errorf("failed to create user in database: %w", err)
		}
	}
	
	// Create custom token for API access
	customToken, err := s.firebaseAuth.CreateCustomToken(ctx, user.ID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create custom token: %w", err)
	}
	
	token := &entities.AuthToken{
		AccessToken: customToken,
		ExpiresAt:   time.Now().Add(time.Hour * 24), // 24 hours
		TokenType:   "Bearer",
	}
	
	return &entities.AuthResponse{
		User:  dbUser,
		Token: token,
	}, nil
}

// Logout revokes user's refresh tokens
func (s *AuthService) Logout(ctx context.Context, uid string) error {
	if s.firebaseAuth == nil {
		return fmt.Errorf("Firebase Auth is not configured")
	}
	
	err := s.firebaseAuth.RevokeRefreshTokens(ctx, uid)
	if err != nil {
		return fmt.Errorf("failed to logout user: %w", err)
	}
	return nil
}

// ValidateToken validates an ID token and returns the user
func (s *AuthService) ValidateToken(ctx context.Context, idToken string) (*entities.User, error) {
	if s.firebaseAuth == nil {
		return nil, fmt.Errorf("Firebase Auth is not configured")
	}
	
	user, err := s.firebaseAuth.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	
	// Get user from our database to ensure they still exist
	dbUser, err := s.userRepository.GetByID(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("user not found in database: %w", err)
	}
	
	return dbUser, nil
}

// RefreshToken creates a new custom token (Firebase client SDK handles refresh tokens)
func (s *AuthService) RefreshToken(ctx context.Context, uid string) (*entities.AuthToken, error) {
	if s.firebaseAuth == nil {
		return nil, fmt.Errorf("Firebase Auth is not configured")
	}
	
	// Verify user exists
	_, err := s.userRepository.GetByID(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	
	// Create new custom token
	customToken, err := s.firebaseAuth.CreateCustomToken(ctx, uid, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create custom token: %w", err)
	}
	
	token := &entities.AuthToken{
		AccessToken: customToken,
		ExpiresAt:   time.Now().Add(time.Hour * 24), // 24 hours
		TokenType:   "Bearer",
	}
	
	return token, nil
}

// GetUserProfile gets user profile information
func (s *AuthService) GetUserProfile(ctx context.Context, uid string) (*entities.User, error) {
	user, err := s.userRepository.GetByID(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return user, nil
}

// UpdateUserProfile updates user profile information
func (s *AuthService) UpdateUserProfile(ctx context.Context, uid string, updates map[string]interface{}) (*entities.User, error) {
	if s.firebaseAuth == nil {
		return nil, fmt.Errorf("Firebase Auth is not configured")
	}
	
	// Update in Firebase Auth
	user, err := s.firebaseAuth.UpdateUser(ctx, uid, updates)
	if err != nil {
		return nil, fmt.Errorf("failed to update user in Firebase: %w", err)
	}
	
	// Update in our database
	user.UpdatedAt = time.Now()
	err = s.userRepository.Update(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user in database: %w", err)
	}
	
	return user, nil
}