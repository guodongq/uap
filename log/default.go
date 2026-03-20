package log

import (
	"context"
	"sync"
)

var (
	defaultLogger Logger = &NoOpLogger{}
	defaultLevel         = InfoLevel
	mu            sync.RWMutex
)

func DefaultLogger() Logger {
	mu.RLock()
	defer mu.RUnlock()
	return defaultLogger
}

// SetDefaultLogger sets the default logger.
// This is not concurrency safe, which means it should only be called during init.
func SetDefaultLogger(l Logger) {
	if l == nil {
		panic("logger must not be nil")
	}
	mu.Lock()
	defer mu.Unlock()
	defaultLogger = l
}

// SetDefaultLevel sets the Level of logs below which logs will not be Output.
// The default log Level is LevelTrace.
func SetDefaultLevel(level Level) {
	if level < DebugLevel || level > FatalLevel {
		panic("invalid log defaultLevel")
	}
	mu.Lock()
	defer mu.Unlock()
	defaultLevel = level
}

var _ Logger = (*NoOpLogger)(nil)

type NoOpLogger struct{}

func (l *NoOpLogger) WithFields(fields Fields) Logger {
	return l
}

func (l *NoOpLogger) WithField(key string, value any) Logger {
	return l
}

func (l *NoOpLogger) WithError(err error) Logger {
	return l
}

func (l *NoOpLogger) WithContext(ctx context.Context) Logger {
	return l
}
func (l *NoOpLogger) Debug(_ ...interface{}) {
}
func (l *NoOpLogger) Debugf(_ string, _ ...interface{}) {}

func (l *NoOpLogger) Info(_ ...interface{}) {
}
func (l *NoOpLogger) Infof(_ string, _ ...interface{}) {
}

func (l *NoOpLogger) Warn(_ ...interface{}) {
}

func (l *NoOpLogger) Warnf(_ string, _ ...interface{}) {
}

func (l *NoOpLogger) Error(_ ...interface{}) {
}
func (l *NoOpLogger) Errorf(_ string, _ ...interface{}) {
}
func (l *NoOpLogger) Fatal(_ ...interface{}) {
}
func (l *NoOpLogger) Fatalf(_ string, _ ...interface{}) {
}
