package log

import (
	"fmt"
	"strings"
)

// Level defines the priority of a log message.
// When a logger is configured with a Level, any log message with a lower
// log Level (smaller by integer comparison) will not be Output.
type Level int

// The levels of logs.
const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// String returns a string representation of the log defaultLevel.
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func ParseLevel(levelStr string) (level Level, err error) {
	switch strings.ToLower(levelStr) {
	case "debug":
		level = DebugLevel
	case "info":
		level = InfoLevel
	case "warn":
		level = WarnLevel
	case "error":
		level = ErrorLevel
	case "fatal":
		level = FatalLevel
	default:
		err = fmt.Errorf("invalid log level: %s", levelStr)
	}
	return
}
