package http

import (
	"github.com/gin-gonic/gin"
	"nestmate-backend/internal/infrastructure/config"
)

type Server struct {
	router *gin.Engine
	config *config.Config
}

func NewServer() *Server {
	cfg := config.Load()
	router := gin.Default()
	
	server := &Server{
		router: router,
		config: cfg,
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
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", s.handleRegister)
			auth.POST("/login", s.handleLogin)
			auth.POST("/logout", s.handleLogout)
			auth.POST("/refresh", s.handleRefreshToken)
		}
		
		// Protected routes (will add middleware later)
		protected := api.Group("/")
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

// Placeholder handlers - will be implemented in later tasks
func (s *Server) handleRegister(c *gin.Context)         { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleLogin(c *gin.Context)            { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleLogout(c *gin.Context)           { c.JSON(501, gin.H{"error": "not implemented"}) }
func (s *Server) handleRefreshToken(c *gin.Context)     { c.JSON(501, gin.H{"error": "not implemented"}) }
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