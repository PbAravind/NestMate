package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"nestmate-backend/internal/domain/entities"
)

func TestAuthEndpoints(t *testing.T) {
	// This test verifies that the auth endpoints are properly configured
	// Note: This will fail without proper Firebase credentials, but tests the routing
	
	server := NewServer()
	
	// Test health endpoint (should work)
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	server.router.ServeHTTP(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Health endpoint failed: expected %d, got %d", http.StatusOK, w.Code)
	}
	
	// Test register endpoint structure (will fail auth but should reach handler)
	registerReq := entities.AuthRequest{
		Email:    "test@example.com",
		Password: "password123",
		Name:     "Test User",
	}
	
	jsonData, _ := json.Marshal(registerReq)
	req, _ = http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	server.router.ServeHTTP(w, req)
	
	// Should reach the handler (will fail due to missing Firebase config, but that's expected)
	if w.Code == http.StatusNotFound {
		t.Error("Register endpoint not found - routing issue")
	}
	
	// Test protected endpoint without auth (should return 401)
	req, _ = http.NewRequest("GET", "/api/v1/auth/profile", nil)
	w = httptest.NewRecorder()
	server.router.ServeHTTP(w, req)
	
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Protected endpoint should return 401 without auth: expected %d, got %d", http.StatusUnauthorized, w.Code)
	}
	
	// Test protected endpoint with invalid auth (should return 401)
	req, _ = http.NewRequest("GET", "/api/v1/auth/profile", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")
	w = httptest.NewRecorder()
	server.router.ServeHTTP(w, req)
	
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Protected endpoint should return 401 with invalid token: expected %d, got %d", http.StatusUnauthorized, w.Code)
	}
}

func TestAuthMiddleware(t *testing.T) {
	server := NewServer()
	
	// Test that protected routes require authentication
	protectedRoutes := []string{
		"/api/v1/expenses",
		"/api/v1/tasks",
		"/api/v1/notes",
	}
	
	for _, route := range protectedRoutes {
		req, _ := http.NewRequest("GET", route, nil)
		w := httptest.NewRecorder()
		server.router.ServeHTTP(w, req)
		
		if w.Code != http.StatusUnauthorized {
			t.Errorf("Route %s should require authentication: expected %d, got %d", route, http.StatusUnauthorized, w.Code)
		}
	}
}