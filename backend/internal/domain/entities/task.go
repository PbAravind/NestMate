package entities

import (
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

// Task represents a task entity
type Task struct {
	ID          string
	UserID      string
	Title       string
	Description string
	DueDate     *time.Time
	Priority    Priority
	Status      TaskStatus
	Labels      []string
	Reminders   []Reminder
	IsRecurring bool
	RecurrenceRule *RecurrenceRule
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Reminder represents a task reminder
type Reminder struct {
	ID        string
	TaskID    string
	Time      time.Time
	Triggered bool
	CreatedAt time.Time
}

// RecurrenceRule represents a task recurrence pattern
type RecurrenceRule struct {
	Frequency  RecurrenceFrequency
	Interval   int
	DaysOfWeek []time.Weekday
	MonthDay   int
	EndDate    *time.Time
	Count      int
}

// RecurrenceFrequency represents how often a task recurs
type RecurrenceFrequency string

const (
	Daily   RecurrenceFrequency = "daily"
	Weekly  RecurrenceFrequency = "weekly"
	Monthly RecurrenceFrequency = "monthly"
	Yearly  RecurrenceFrequency = "yearly"
)