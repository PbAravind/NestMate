package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"nestmate-backend/internal/application/services"
)

// AuthMiddleware handles authentication for protected routes
type AuthMiddleware struct {
	authService *services.AuthService
}

// NewAuthMiddleware creates a new authentication middleware
func NewAuthMiddleware(authService *services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

// RequireAuth middleware that requires authentication
func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
				"code":  "MISSING_AUTH_HEADER",
			})
			c.Abort()
			return
		}
		
		// Check if it's a Bearer token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header format. Expected: Bearer <token>",
				"code":  "INVALID_AUTH_FORMAT",
			})
			c.Abort()
			return
		}
		
		idToken := tokenParts[1]
		
		// Validate the token
		ctx := context.Background()
		user, err := m.authService.ValidateToken(ctx, idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
				"code":  "INVALID_TOKEN",
			})
			c.Abort()
			return
		}
		
		// Set user in context for use in handlers
		c.Set("user", user)
		c.Set("user_id", user.ID)
		
		c.Next()
	}
}

// OptionalAuth middleware that optionally authenticates (doesn't fail if no auth)
func (m *AuthMiddleware) OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}
		
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.Next()
			return
		}
		
		idToken := tokenParts[1]
		
		ctx := context.Background()
		user, err := m.authService.ValidateToken(ctx, idToken)
		if err == nil {
			c.Set("user", user)
			c.Set("user_id", user.ID)
		}
		
		c.Next()
	}
}

// GetUserFromContext extracts user from gin context
func GetUserFromContext(c *gin.Context) (string, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return "", false
	}
	
	uid, ok := userID.(string)
	return uid, ok
}

// RequireUserID middleware that ensures user_id is available in context
func RequireUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := GetUserFromContext(c)
		if !exists || userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User authentication required",
				"code":  "USER_AUTH_REQUIRED",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}