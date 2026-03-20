package ctxutil

import "context"

// ContextKey is a typed key for context values.
type ContextKey string

// SetToContext stores a typed value in the context under the given key.
func SetToContext[T any](ctx context.Context, key ContextKey, value T) context.Context {
	return context.WithValue(ctx, key, value)
}

// GetFromContext retrieves a typed value from the context by key.
func GetFromContext[T any](ctx context.Context, key ContextKey) (T, bool) {
	value := ctx.Value(key)
	if value == nil {
		var zero T
		return zero, false
	}
	typedValue, ok := value.(T)
	return typedValue, ok
}
