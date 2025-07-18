package http

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"nestmate-backend/internal/application/services"
	"nestmate-backend/internal/domain/entities"
	"nestmate-backend/internal/infrastructure/auth"
	"nestmate-backend/internal/infrastructure/config"
	"nestmate-backend/internal/infrastructure/repositories/memory"
	"nestmate-backend/internal/interfaces/http/middleware"
)

type Server struct {
	router      *gin.Engine
	config      *config.Config
	authService *services.AuthService
	authMiddleware *middleware.AuthMiddleware
}

func NewServer() *Server {
	cfg := config.Load()
	router := gin.Default()
	
	// Initialize Firebase Auth service
	firebaseAuth, err := auth.NewFirebaseAuthService(&cfg.Firebase)
	if err != nil {
		log.Printf("Warning: Failed to initialize Firebase Auth: %v", err)
		log.Println("Firebase Auth will not be available. Make sure to set FIREBASE_PROJECT_ID, FIREBASE_PRIVATE_KEY, and FIREBASE_CLIENT_EMAIL environment variables.")
	}
	
	// Initialize repositories
	userRepo := memory.NewInMemoryUserRepository()
	
	// Initialize services
	authService := services.NewAuthService(firebaseAuth, userRepo)
	
	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(authService)
	
	server := &Server{
		router:         router,
		config:         cfg,
		authService:    authService,
		authMiddleware: authMiddleware,
	}
	
	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	// Health check
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	
	// API routes
	api := s.router.Group("/api/v1")
	{
		// Auth routes (public)
		auth := api.Group("/auth")
		{
			auth.POST("/register", s.handleRegister)
			auth.POST("/login", s.handleLogin)
			auth.POST("/logout", s.authMiddleware.RequireAuth(), s.handleLogout)
			auth.POST("/refresh", s.authMiddleware.RequireAuth(), s.handleRefreshToken)
			auth.GET("/profile", s.authMiddleware.RequireAuth(), s.handleGetProfile)
			auth.PUT("/profile", s.authMiddleware.RequireAuth(), s.handleUpdateProfile)
		}
		
		// Protected routes - require authentication
		protected := api.Group("/")
		protected.Use(s.authMiddleware.RequireAuth())
		{
			// Expense routes
			expenses := protected.Group("/expenses")
			{
				expenses.POST("", s.handleCreateExpense)
				expenses.GET("", s.handleGetExpenses)
				expenses.GET("/:id", s.handleGetExpense)
				expenses.PUT("/:id", s.handleUpdateExpense)
				expenses.DELETE("/:id", s.handleDeleteExpense)
				expenses.GET("/breakdown", s.handleGetMonthlyBreakdown)
			}
			
			// Task routes
			tasks := protected.Group("/tasks")
			{
				tasks.POST("", s.handleCreateTask)
				tasks.GET("", s.handleGetTasks)
				tasks.GET("/:id", s.handleGetTask)
				tasks.PUT("/:id", s.handleUpdateTask)
				tasks.DELETE("/:id", s.handleDeleteTask)
				tasks.PATCH("/:id/status", s.handleUpdateTaskStatus)
			}
			
			// Notes routes
			notes := protected.Group("/notes")
			{
				notes.POST("", s.handleCreateNote)
				notes.GET("", s.handleGetNotes)
				notes.GET("/:id", s.handleGetNote)
				notes.PUT("/:id", s.handleUpdateNote)
				notes.DELETE("/:id", s.handleDeleteNote)
				notes.GET("/search", s.handleSearchNotes)
			}
		}
	}
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

// Authentication handlers
func (s *Server) handleRegister(c *gin.Context) {
	var req entities.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
			"code":  "INVALID_REQUEST",
			"details": err.Error(),
		})
		return
	}
	
	ctx := context.Background()
	response, err := s.authService.Register(ctx, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Registration failed",
			"code":  "REGISTRATION_FAILED",
			"details": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusCreated, response)
}

func (s *Server) handleLogin(c *gin.Context) {
	var req struct {
		IDToken string `json:"id_token" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
			"code":  "INVALID_REQUEST",
			"details": err.Error(),
		})
		return
	}
	
	ctx := context.Background()
	response, err := s.authService.Login(ctx, req.IDToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Login failed",
			"code":  "LOGIN_FAILED",
			"details": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, response)
}

func (s *Server) handleLogout(c *gin.Context) {
	userID, exists := middleware.GetUserFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
			"code":  "USER_NOT_AUTHENTICATED",
		})
		return
	}
	
	ctx := context.Background()
	err := s.authService.Logout(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Logout failed",
			"code":  "LOGOUT_FAILED",
			"details": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}

func (s *Server) handleRefreshToken(c *gin.Context) {
	userID, exists := middleware.GetUserFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
			"code":  "USER_NOT_AUTHENTICATED",
		})
		return
	}
	
	ctx := context.Background()
	token, err := s.authService.RefreshToken(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Token refresh failed",
			"code":  "TOKEN_REFRESH_FAILED",
			"details": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (s *Server) handleGetProfile(c *gin.Context) {
	userID, exists := middleware.GetUserFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
			"code":  "USER_NOT_AUTHENTICATED",
		})
		return
	}
	
	ctx := context.Background()
	user, err := s.authService.GetUserProfile(ctx, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User profile not found",
			"code":  "USER_NOT_FOUND",
			"details": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (s *Server) handleUpdateProfile(c *gin.Context) {
	userID, exists := middleware.GetUserFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
			"code":  "USER_NOT_AUTHENTICATED",
		})
		return
	}
	
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
			"code":  "INVALID_REQUEST",
			"details": err.Error(),
		})
		return
	}
	
	ctx := context.Background()
	user, err := s.authService.UpdateUserProfile(ctx, userID, updates)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Profile update failed",
			"code":  "PROFILE_UPDATE_FAILED",
			"details": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// Placeholder handlers for other modules - will be implemented in later tasks
func (s *Server) handleCreateExpense(c *gin.Context)    { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleGetExpenses(c *gin.Context)      { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleGetExpense(c *gin.Context)       { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleUpdateExpense(c *gin.Context)    { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleDeleteExpense(c *gin.Context)    { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleGetMonthlyBreakdown(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleCreateTask(c *gin.Context)       { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleGetTasks(c *gin.Context)         { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleGetTask(c *gin.Context)          { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleUpdateTask(c *gin.Context)       { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleDeleteTask(c *gin.Context)       { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleUpdateTaskStatus(c *gin.Context) { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleCreateNote(c *gin.Context)       { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleGetNotes(c *gin.Context)         { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleGetNote(c *gin.Context)          { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleUpdateNote(c *gin.Context)       { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleDeleteNote(c *gin.Context)       { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleSearchNotes(c *gin.Context)      { c.JSON(501, gin.H{"error": "not implemented"}) }