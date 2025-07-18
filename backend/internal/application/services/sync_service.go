package services

import (
	"time"
)

type ChangeRecord struct {
	ID        string                 `json:"id"`
	Type      string                 `json:"type"`      // expense, task, note
	Operation string                 `json:"operation"` // create, update, delete
	Data      map[string]interface{} `json:"data"`
	Timestamp time.Time              `json:"timestamp"`
	UserID    string                 `json:"user_id"`
}

type SyncConflict struct {
	ID         string      `json:"id"`
	LocalData  interface{} `json:"local_data"`
	RemoteData interface{} `json:"remote_data"`
	Type       string      `json:"type"`
}

type SyncStatus struct {
	LastSync    time.Time `json:"last_sync"`
	PendingSync int       `json:"pending_sync"`
	InProgress  bool      `json:"in_progress"`
	Errors      []string  `json:"errors,omitempty"`
}

type SyncService interface {
	SyncUserData(userID string) error
	QueueLocalChanges(changes []*ChangeRecord) error
	ResolveConflicts(conflicts []*SyncConflict) error
	GetSyncStatus(userID string) (*SyncStatus, error)
}