package repositories

import (
	"context"
	"time"
)

// TaskStatus represents the status of a task
type TaskStatus int

const (
	Pending TaskStatus = iota
	InProgress
	Done
)

// Priority represents the priority level of a task
type Priority int

const (
	Low Priority = iota
	Medium
	High
)

// TaskRepository defines the interface for task data access
type TaskRepository interface {
	// Create a new task
	Create(ctx context.Context, task *Task) error
	
	// Get a task by ID
	GetByID(ctx context.Context, id string) (*Task, error)
	
	// Get tasks by user ID and filters
	GetByUserID(ctx context.Context, userID string, filters *TaskFilters) ([]*Task, error)
	
	// Update a task
	Update(ctx context.Context, task *Task) error
	
	// Delete a task by ID
	Delete(ctx context.Context, id string) error
	
	// Update task status
	UpdateStatus(ctx context.Context, id string, status TaskStatus) error
}

// Task represents the repository task model
type Task struct {
	ID          string
	UserID      string
	Title       string
	Description string
	DueDate     *time.Time
	Priority    Priority
	Status      TaskStatus
	Labels      []string
	IsRecurring bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// TaskFilters represents filters for task queries
type TaskFilters struct {
	Status    *TaskStatus
	Priority  *Priority
	Labels    []string
	StartDate *time.Time
	EndDate   *time.Time
}