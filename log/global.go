package log

import (
	"context"
	"runtime/debug"
)

func WithFields(fields Fields) Logger {
	return DefaultLogger().WithFields(fields)
}

func WithField(key string, value any) Logger {
	return DefaultLogger().WithField(key, value)
}

func WithError(err error) Logger {
	return DefaultLogger().WithError(err)
}

func WithContext(ctx context.Context) Logger {
	return DefaultLogger().WithContext(ctx)
}

// Debug calls the default logger's Debug method.
func Debug(v ...interface{}) {
	if shouldLog(DebugLevel) {
		DefaultLogger().Debug(v...)
	}
}

// Debugf calls the default logger's Debugf method.
func Debugf(format string, v ...interface{}) {
	if shouldLog(DebugLevel) {
		DefaultLogger().Debugf(format, v...)
	}
}

// Info calls the default logger's Info method.
func Info(v ...interface{}) {
	if shouldLog(InfoLevel) {
		DefaultLogger().Info(v...)
	}
}

// Infof calls the default logger's Infof method.
func Infof(format string, v ...interface{}) {
	if shouldLog(InfoLevel) {
		DefaultLogger().Infof(format, v...)
	}
}

// Warn calls the default logger's Warn method.
func Warn(v ...interface{}) {
	if shouldLog(WarnLevel) {
		DefaultLogger().Warn(v...)
	}
}

// Warnf calls the default logger's Warnf method.
func Warnf(format string, v ...interface{}) {
	if shouldLog(WarnLevel) {
		DefaultLogger().Warnf(format, v...)
	}
}

// Error calls the default logger's Error method.
func Error(v ...interface{}) {
	if shouldLog(ErrorLevel) {
		DefaultLogger().Error(v...)
	}
}

// Errorf calls the default logger's Errorf method.
func Errorf(format string, v ...interface{}) {
	if shouldLog(ErrorLevel) {
		DefaultLogger().Errorf(format, v...)
	}
}

// Fatal calls the default logger's Panic method
func Fatal(v ...interface{}) {
	if shouldLog(FatalLevel) {
		DefaultLogger().Fatal(v...)
		debug.PrintStack()
	}
}

// Fatalf calls the default logger's Fatalf method
func Fatalf(format string, v ...interface{}) {
	if shouldLog(FatalLevel) {
		DefaultLogger().Fatalf(format, v...)
	}
}

// shouldLog checks if the given defaultLevel should be logged based on the global defaultLevel.
func shouldLog(lvl Level) bool {
	mu.RLock()
	defer mu.RUnlock()
	return lvl >= defaultLevel
}
