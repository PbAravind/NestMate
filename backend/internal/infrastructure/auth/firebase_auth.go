package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"nestmate-backend/internal/domain/entities"
	"nestmate-backend/internal/infrastructure/config"
)

// FirebaseAuthService implements authentication using Firebase Auth
type FirebaseAuthService struct {
	client *auth.Client
	config *config.FirebaseConfig
}

// NewFirebaseAuthService creates a new Firebase Auth service
func NewFirebaseAuthService(cfg *config.FirebaseConfig) (*FirebaseAuthService, error) {
	ctx := context.Background()
	
	// Create Firebase credentials from config
	credentials := map[string]interface{}{
		"type":                        "service_account",
		"project_id":                  cfg.ProjectID,
		"private_key":                 cfg.PrivateKey,
		"client_email":                cfg.ClientEmail,
		"auth_uri":                    "https://accounts.google.com/o/oauth2/auth",
		"token_uri":                   "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
	}
	
	credentialsJSON, err := json.Marshal(credentials)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Firebase credentials: %w", err)
	}
	
	// Initialize Firebase app
	opt := option.WithCredentialsJSON(credentialsJSON)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %w", err)
	}
	
	// Get Auth client
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get Firebase Auth client: %w", err)
	}
	
	return &FirebaseAuthService{
		client: client,
		config: cfg,
	}, nil
}

// CreateUser creates a new user in Firebase Auth
func (f *FirebaseAuthService) CreateUser(ctx context.Context, email, password, name string) (*entities.User, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password).
		DisplayName(name).
		EmailVerified(false)
	
	userRecord, err := f.client.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	
	user := &entities.User{
		ID:        userRecord.UID,
		Email:     userRecord.Email,
		Name:      userRecord.DisplayName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	return user, nil
}

// VerifyIDToken verifies a Firebase ID token and returns the user
func (f *FirebaseAuthService) VerifyIDToken(ctx context.Context, idToken string) (*entities.User, error) {
	token, err := f.client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("failed to verify ID token: %w", err)
	}
	
	userRecord, err := f.client.GetUser(ctx, token.UID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	
	user := &entities.User{
		ID:    userRecord.UID,
		Email: userRecord.Email,
		Name:  userRecord.DisplayName,
	}
	
	return user, nil
}

// GetUser retrieves a user by UID
func (f *FirebaseAuthService) GetUser(ctx context.Context, uid string) (*entities.User, error) {
	userRecord, err := f.client.GetUser(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	
	user := &entities.User{
		ID:    userRecord.UID,
		Email: userRecord.Email,
		Name:  userRecord.DisplayName,
	}
	
	return user, nil
}

// GetUserByEmail retrieves a user by email
func (f *FirebaseAuthService) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	userRecord, err := f.client.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	
	user := &entities.User{
		ID:    userRecord.UID,
		Email: userRecord.Email,
		Name:  userRecord.DisplayName,
	}
	
	return user, nil
}

// UpdateUser updates a user's information
func (f *FirebaseAuthService) UpdateUser(ctx context.Context, uid string, updates map[string]interface{}) (*entities.User, error) {
	params := &auth.UserToUpdate{}
	
	if email, ok := updates["email"].(string); ok {
		params = params.Email(email)
	}
	if name, ok := updates["name"].(string); ok {
		params = params.DisplayName(name)
	}
	if password, ok := updates["password"].(string); ok {
		params = params.Password(password)
	}
	
	userRecord, err := f.client.UpdateUser(ctx, uid, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	
	user := &entities.User{
		ID:        userRecord.UID,
		Email:     userRecord.Email,
		Name:      userRecord.DisplayName,
		UpdatedAt: time.Now(),
	}
	
	return user, nil
}

// DeleteUser deletes a user from Firebase Auth
func (f *FirebaseAuthService) DeleteUser(ctx context.Context, uid string) error {
	err := f.client.DeleteUser(ctx, uid)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// CreateCustomToken creates a custom token for a user
func (f *FirebaseAuthService) CreateCustomToken(ctx context.Context, uid string, claims map[string]interface{}) (string, error) {
	token, err := f.client.CustomToken(ctx, uid)
	if err != nil {
		return "", fmt.Errorf("failed to create custom token: %w", err)
	}
	return token, nil
}

// RevokeRefreshTokens revokes all refresh tokens for a user
func (f *FirebaseAuthService) RevokeRefreshTokens(ctx context.Context, uid string) error {
	err := f.client.RevokeRefreshTokens(ctx, uid)
	if err != nil {
		return fmt.Errorf("failed to revoke refresh tokens: %w", err)
	}
	return nil
}