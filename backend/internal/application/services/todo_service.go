package services

import (
	"time"
	"nestmate-backend/internal/domain/entities"
)

type TaskFilter struct {
	Status   *entities.TaskStatus `json:"status,omitempty"`
	Priority *entities.Priority   `json:"priority,omitempty"`
	Labels   []string             `json:"labels,omitempty"`
	Search   string               `json:"search,omitempty"`
}

type Reminder struct {
	ID       string    `json:"id"`
	TaskID   string    `json:"task_id"`
	DateTime time.Time `json:"date_time"`
	Message  string    `json:"message"`
}

type TodoService interface {
	CreateTask(task *entities.Task) error
	UpdateTask(id string, task *entities.Task) error
	DeleteTask(id string) error
	GetTasksByFilter(userID string, filter *TaskFilter) ([]*entities.Task, error)
	GetTasksForPeriod(userID string, start, end time.Time) ([]*entities.Task, error)
	MarkTaskComplete(id string) error
	SetReminder(taskID string, reminder *Reminder) error
}