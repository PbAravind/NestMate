package entities

import (
	"time"
)

// Note represents a note entity
type Note struct {
	ID          string
	UserID      string
	Title       string
	Content     string // Markdown format
	Tags        []string
	IsFavorite  bool
	Attachments []Attachment
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// AttachmentType represents the type of attachment
type AttachmentType string

const (
	Link  AttachmentType = "link"
	Image AttachmentType = "image"
)

// Attachment represents a note attachment
type Attachment struct {
	ID       string
	Type     AttachmentType
	URL      string
	Metadata map[string]interface{}
	CreatedAt time.Time
}