package log

import "context"

var (
	// G shortcut to get logger from the context
	G = GetLogger
	// L shortcut to get the default logger.
	// It is a function so that it always returns the latest default logger
	// even after SetDefaultLogger is called.
	L = DefaultLogger
)

type loggerKey struct{}

// WithLogger returns a new context with the provided logger.
func WithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// GetLogger retrieves the current logger from the context.
// If no logger is available, the default logger is returned.
func GetLogger(ctx context.Context) Logger {
	if ctx == nil {
		return L()
	}

	logger := ctx.Value(loggerKey{})
	if logger == nil {
		return L()
	}

	return logger.(Logger)
}
