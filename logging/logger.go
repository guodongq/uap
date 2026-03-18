package logging

import "context"

type Fields map[string]any

// Logger is a logger interface that provides logging function with levels.
type Logger interface {
	WithFields(fields Fields) Logger
	WithField(key string, value any) Logger
	WithError(err error) Logger
	WithContext(ctx context.Context) Logger

	Debug(v ...any)
	Debugf(format string, v ...any)

	Info(v ...any)
	Infof(format string, v ...any)

	Warn(v ...any)
	Warnf(format string, v ...any)

	Error(v ...any)
	Errorf(format string, v ...any)

	Fatal(v ...any)
	Fatalf(format string, v ...any)
}
