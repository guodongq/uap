package model

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Ider is a generic interface for types that have an ID.
type Ider[T comparable] interface {
	ID() T
}

// BaseIder is a generic implementation of the Ider interface.
type BaseIder[T comparable] struct {
	id T
}

// NewIder creates a new BaseIder with the given value.
func NewIder[T comparable](id T) *BaseIder[T] {
	return &BaseIder[T]{
		id: id,
	}
}

// ID returns the stored identifier.
func (b *BaseIder[T]) ID() T {
	return b.id
}

// StringIder is a specific implementation of Ider for string type.
// It is an alias for BaseIder[string].
type StringIder = BaseIder[string]

// NewStringIder creates a new StringIder instance.
func NewStringIder(value string) *StringIder {
	return NewIder(value)
}

// IntIder is a specific implementation of Ider for int type.
// It is an alias for BaseIder[int].
type IntIder = BaseIder[int]

// NewIntIder creates a new IntIder instance.
func NewIntIder(value int) *IntIder {
	return NewIder(value)
}

type ObjectIDIder = BaseIder[bson.ObjectID]

func NewObjectIDIder(value bson.ObjectID) *ObjectIDIder {
	return NewIder(value)
}

type UUIDIder = BaseIder[uuid.UUID]

func NewUUIDIder(value uuid.UUID) *UUIDIder {
	return NewIder(value)
}
