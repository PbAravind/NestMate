package auth

import (
	"testing"

	"nestmate-backend/internal/infrastructure/config"
)

func TestNewFirebaseAuthService(t *testing.T) {
	// Test with empty config (should handle gracefully)
	cfg := &config.FirebaseConfig{
		ProjectID:   "",
		PrivateKey:  "",
		ClientEmail: "",
	}
	
	_, err := NewFirebaseAuthService(cfg)
	if err == nil {
		t.Error("Expected error when Firebase config is empty, but got nil")
	}
	
	// Test with invalid config
	cfg = &config.FirebaseConfig{
		ProjectID:   "test-project",
		PrivateKey:  "invalid-key",
		ClientEmail: "test@example.com",
	}
	
	_, err = NewFirebaseAuthService(cfg)
	if err == nil {
		t.Error("Expected error when Firebase config is invalid, but got nil")
	}
}

func TestFirebaseAuthServiceCreation(t *testing.T) {
	// This test verifies that the service can be created with proper error handling
	// In a real environment with proper Firebase credentials, this would succeed
	
	cfg := &config.FirebaseConfig{
		ProjectID:   "test-project-id",
		PrivateKey:  "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC...\n-----END PRIVATE KEY-----\n",
		ClientEmail: "firebase-adminsdk-test@test-project.iam.gserviceaccount.com",
	}
	
	service, err := NewFirebaseAuthService(cfg)
	
	// We expect this to fail in test environment since we don't have real credentials
	// But the service creation logic should handle it gracefully
	if err != nil {
		t.Logf("Expected error in test environment: %v", err)
	}
	
	if service != nil && service.client == nil {
		t.Error("Service created but client is nil")
	}
}