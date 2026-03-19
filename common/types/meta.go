package types

import (
	"fmt"
	"time"

	"github.com/guodongq/uap/common/errors"
)

// Metadata represents audit information for domain entities.
// It tracks creation, updates, and deletion timestamps along with the user who performed each action.
// The Revision field supports optimistic locking for concurrent update detection.
type Metadata struct {
	CreatedAt time.Time  `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" bson:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" bson:"deleted_at,omitempty"`

	CreatedBy string  `json:"createdBy" bson:"created_by"`
	UpdatedBy string  `json:"updatedBy" bson:"updated_by"`
	DeletedBy *string `json:"deletedBy,omitempty" bson:"deleted_by,omitempty"`

	Revision int64 `json:"revision" bson:"revision"`
}

// ============================================================================
// Factory Functions
// ============================================================================

// NewMetadata creates a new Metadata instance with the given creator.
func NewMetadata(createdBy string) *Metadata {
	now := time.Now().UTC()
	return &Metadata{
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: createdBy,
		UpdatedBy: createdBy,
		Revision:  1,
	}
}

// DefaultMetadata creates a new Metadata instance with default values (system user).
func DefaultMetadata() *Metadata {
	return NewMetadata("system")
}

// ============================================================================
// Update Operations
// ============================================================================

// MarkUpdated updates the UpdatedAt timestamp, UpdatedBy user, and increments the Revision.
// This should be called whenever the entity is modified.
func (m *Metadata) MarkUpdated(updatedBy string) {
	m.UpdatedAt = time.Now().UTC()
	m.UpdatedBy = updatedBy
	m.Revision++
}

// MarkDeleted marks the entity as deleted with the current timestamp and user.
// This implements soft delete pattern.
func (m *Metadata) MarkDeleted(deletedBy string) {
	now := time.Now().UTC()
	m.DeletedAt = &now
	m.DeletedBy = &deletedBy
	m.UpdatedAt = now
	m.UpdatedBy = deletedBy
	m.Revision++
}

// Restore restores a soft-deleted entity by clearing the deletion metadata.
func (m *Metadata) Restore(restoredBy string) error {
	if !m.IsDeleted() {
		return errors.New("entity is not deleted")
	}
	m.DeletedAt = nil
	m.DeletedBy = nil
	m.MarkUpdated(restoredBy)
	return nil
}

// IncrementRevision increments the revision number for optimistic locking.
func (m *Metadata) IncrementRevision() {
	m.Revision++
}

// ============================================================================
// Query Operations
// ============================================================================

// IsDeleted checks if the entity has been soft-deleted.
func (m *Metadata) IsDeleted() bool {
	return m.DeletedAt != nil
}

// IsActive checks if the entity is active (not deleted).
func (m *Metadata) IsActive() bool {
	return !m.IsDeleted()
}

// Age returns the duration since the entity was created.
func (m *Metadata) Age() time.Duration {
	return time.Since(m.CreatedAt)
}

// TimeSinceLastUpdate returns the duration since the entity was last updated.
func (m *Metadata) TimeSinceLastUpdate() time.Duration {
	return time.Since(m.UpdatedAt)
}

// WasCreatedBy checks if the entity was created by the specified user.
func (m *Metadata) WasCreatedBy(userID string) bool {
	return m.CreatedBy == userID
}

// WasLastUpdatedBy checks if the entity was last updated by the specified user.
func (m *Metadata) WasLastUpdatedBy(userID string) bool {
	return m.UpdatedBy == userID
}

// WasDeletedBy checks if the entity was deleted by the specified user.
func (m *Metadata) WasDeletedBy(userID string) bool {
	return m.DeletedBy != nil && *m.DeletedBy == userID
}

// ============================================================================
// Validation
// ============================================================================

// Validate performs basic validation on the metadata.
func (m *Metadata) Validate() error {
	if m.CreatedAt.IsZero() {
		return errors.New("created_at cannot be zero")
	}
	if m.UpdatedAt.IsZero() {
		return errors.New("updated_at cannot be zero")
	}
	if m.UpdatedAt.Before(m.CreatedAt) {
		return errors.New("updated_at cannot be before created_at")
	}
	if m.CreatedBy == "" {
		return errors.New("created_by cannot be empty")
	}
	if m.UpdatedBy == "" {
		return errors.New("updated_by cannot be empty")
	}
	if m.Revision < 1 {
		return errors.New("revision must be at least 1")
	}
	if m.IsDeleted() && m.DeletedAt.Before(m.CreatedAt) {
		return errors.New("deleted_at cannot be before created_at")
	}
	return nil
}

// HasConflict checks if there's a revision conflict for optimistic locking.
// Returns true if the expected revision doesn't match the current revision.
func (m *Metadata) HasConflict(expectedRevision int64) bool {
	return m.Revision != expectedRevision
}

// ============================================================================
// Comparison
// ============================================================================

// IsNewerThan checks if this metadata is newer than another based on UpdatedAt.
func (m *Metadata) IsNewerThan(other *Metadata) bool {
	if other == nil {
		return true
	}
	return m.UpdatedAt.After(other.UpdatedAt)
}

// IsOlderThan checks if this metadata is older than another based on UpdatedAt.
func (m *Metadata) IsOlderThan(other *Metadata) bool {
	if other == nil {
		return false
	}
	return m.UpdatedAt.Before(other.UpdatedAt)
}

// IsSameRevision checks if this metadata has the same revision as another.
func (m *Metadata) IsSameRevision(other *Metadata) bool {
	if other == nil {
		return false
	}
	return m.Revision == other.Revision
}

// ============================================================================
// Cloning
// ============================================================================

// Clone creates a deep copy of the metadata.
func (m *Metadata) Clone() *Metadata {
	clone := &Metadata{
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		CreatedBy: m.CreatedBy,
		UpdatedBy: m.UpdatedBy,
		Revision:  m.Revision,
	}

	if m.DeletedAt != nil {
		deletedAt := *m.DeletedAt
		clone.DeletedAt = &deletedAt
	}

	if m.DeletedBy != nil {
		deletedBy := *m.DeletedBy
		clone.DeletedBy = &deletedBy
	}

	return clone
}

// ============================================================================
// Helper Functions
// ============================================================================

// GetTimestamps returns all timestamps as a map for easy logging or debugging.
func (m *Metadata) GetTimestamps() map[string]interface{} {
	timestamps := map[string]interface{}{
		"created_at": m.CreatedAt,
		"updated_at": m.UpdatedAt,
	}

	if m.DeletedAt != nil {
		timestamps["deleted_at"] = *m.DeletedAt
	}

	return timestamps
}

// GetUsers returns all user IDs as a map for easy logging or debugging.
func (m *Metadata) GetUsers() map[string]string {
	users := map[string]string{
		"created_by": m.CreatedBy,
		"updated_by": m.UpdatedBy,
	}

	if m.DeletedBy != nil {
		users["deleted_by"] = *m.DeletedBy
	}

	return users
}

// String returns a human-readable string representation of the metadata.
func (m *Metadata) String() string {
	status := "active"
	if m.IsDeleted() {
		status = "deleted"
	}
	return fmt.Sprintf("Metadata{status=%s, revision=%d, created_by=%s, updated_by=%s}",
		status, m.Revision, m.CreatedBy, m.UpdatedBy)
}
