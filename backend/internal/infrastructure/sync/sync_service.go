package sync

import (
	"context"
	"errors"
)

// ChangeRecord represents a local change to be synced
type ChangeRecord struct {
	ID        string
	Type      string // "expense", "task", "note"
	Operation string // "create", "update", "delete"
	Data      map[string]interface{}
	Timestamp int64
}

// SyncConflict represents a sync conflict
type SyncConflict struct {
	ID          string
	Type        string
	LocalData   map[string]interface{}
	RemoteData  map[string]interface{}
	LocalTime   int64
	RemoteTime  int64
}

// SyncStatus represents the current sync status
type SyncStatus struct {
	LastSyncTime   int64
	PendingChanges int
	SyncInProgress bool
	LastError      string
}

// SyncService defines the interface for data synchronization
type SyncService interface {
	SyncUserData(ctx context.Context, userID string) error
	QueueLocalChanges(ctx context.Context, changes []*ChangeRecord) error
	ResolveConflicts(ctx context.Context, conflicts []*SyncConflict) error
	GetSyncStatus(ctx context.Context, userID string) (*SyncStatus, error)
}

// syncService implements the SyncService interface
type syncService struct {
	// Dependencies will be injected here
}

// NewSyncService creates a new sync service
func NewSyncService() SyncService {
	return &syncService{}
}

// SyncUserData syncs user data between local and remote
func (s *syncService) SyncUserData(ctx context.Context, userID string) error {
	// Implementation will be added in a future task
	return errors.New("not implemented")
}

// QueueLocalChanges queues local changes for sync
func (s *syncService) QueueLocalChanges(ctx context.Context, changes []*ChangeRecord) error {
	// Implementation will be added in a future task
	return errors.New("not implemented")
}

// ResolveConflicts resolves sync conflicts
func (s *syncService) ResolveConflicts(ctx context.Context, conflicts []*SyncConflict) error {
	// Implementation will be added in a future task
	return errors.New("not implemented")
}

// GetSyncStatus gets the current sync status
func (s *syncService) GetSyncStatus(ctx context.Context, userID string) (*SyncStatus, error) {
	// Implementation will be added in a future task
	return nil, errors.New("not implemented")
}