package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// NullableValue is a generic struct representing a nullable value that tracks whether it's been explicitly set.
// It distinguishes between "unset" and "set to nil" states.
type NullableValue[T any] struct {
	value *T
	isSet bool
}

// NewNullableValue creates a new NullableValue instance from a pointer.
// It's marked as "set" if the input pointer is non-nil.
func NewNullableValue[T any](val *T) *NullableValue[T] {
	return &NullableValue[T]{
		value: val,
		isSet: val != nil,
	}
}

// FromValue creates a new NullableValue directly from a concrete value (auto-wraps in a pointer).
func FromValue[T any](val T) *NullableValue[T] {
	return NewNullableValue(&val)
}

// Get returns the underlying pointer value. Returns nil if no value is set.
func (v *NullableValue[T]) Get() *T {
	return v.value
}

// Set updates the value and marks it as "set".
// Accepts nil to explicitly set a null value while maintaining "set" status.
func (v *NullableValue[T]) Set(val *T) {
	v.value = val
	v.isSet = true
}

// IsSet returns true if the value has been explicitly set (including being set to nil).
func (v *NullableValue[T]) IsSet() bool {
	return v.isSet
}

// Unset resets the value to nil and marks it as "unset".
func (v *NullableValue[T]) Unset() {
	v.value = nil
	v.isSet = false
}

// MarshalJSON implements the json.Marshaler interface.
// Serializes to the underlying value's JSON representation if set; otherwise returns "null".
func (v *NullableValue[T]) MarshalJSON() ([]byte, error) {
	if v.IsSet() {
		return json.Marshal(v.value)
	}
	return []byte("null"), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// Marks the value as "set" after deserialization, even if the result is null.
func (v *NullableValue[T]) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// Value returns the concrete value and a validity flag.
// The flag is true only if the underlying value is non-nil.
func (v *NullableValue[T]) Value() (T, bool) {
	if v.value == nil {
		var zero T
		return zero, false
	}
	return *v.value, true
}

// ------------------------------
// Specific nullable type aliases
// ------------------------------

// NullableBool is a type alias for NullableValue[bool]
type NullableBool = NullableValue[bool]

// NullableInt is a type alias for NullableValue[int]
type NullableInt = NullableValue[int]

// NullableInt32 is a type alias for NullableValue[int32]
type NullableInt32 = NullableValue[int32]

// NullableInt64 is a type alias for NullableValue[int64]
type NullableInt64 = NullableValue[int64]

// NullableFloat32 is a type alias for NullableValue[float32]
type NullableFloat32 = NullableValue[float32]

// NullableFloat64 is a type alias for NullableValue[float64]
type NullableFloat64 = NullableValue[float64]

// NullableString is a type alias for NullableValue[string]
type NullableString = NullableValue[string]

// NullableTime is a type alias for NullableValue[time.Time]
type NullableTime = NullableValue[time.Time]

// ------------------------------
// Helper functions (optional)
// ------------------------------

func NewNullableBool(val *bool) *NullableBool {
	return NewNullableValue(val)
}

func NewNullableInt(val *int) *NullableInt {
	return NewNullableValue(val)
}

func NewNullableInt32(val *int32) *NullableInt32 {
	return NewNullableValue(val)
}

func NewNullableInt64(val *int64) *NullableInt64 {
	return NewNullableValue(val)
}

func NewNullableFloat32(val *float32) *NullableFloat32 {
	return NewNullableValue(val)
}

func NewNullableFloat64(val *float64) *NullableFloat64 {
	return NewNullableValue(val)
}

func NewNullableString(val *string) *NullableString {
	return NewNullableValue(val)
}

func NewNullableTime(val *time.Time) *NullableTime {
	return NewNullableValue(val)
}

// ------------------------------
// Other utility functions
// ------------------------------

// IsNil checks if an input is nil
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	case reflect.Array:
		return reflect.ValueOf(i).IsZero()
	default:
		return false
	}
}

type MappedNullable interface {
	ToMap() (map[string]interface{}, error)
}

// NewStrictDecoder creates a JSON decoder that disallows unknown fields
func NewStrictDecoder(data []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	return dec
}

// ReportError wraps fmt.Errorf for consistent error reporting
func ReportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
