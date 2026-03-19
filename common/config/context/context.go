package uapcontext

import "context"

type contextKey string

func setToContext[T any](ctx context.Context, key contextKey, value T) context.Context {
	return context.WithValue(ctx, key, value)
}

func getFromContext[T any](ctx context.Context, key contextKey) (T, bool) {
	value := ctx.Value(key)
	if value == nil {
		var zero T
		return zero, false
	}
	typedValue, ok := value.(T)
	return typedValue, ok
}
