package uapfx

import (
	"strings"

	logger "github.com/guodongq/uap/logging"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

// Module provides the fx logger option using the UAP logger.
// It utilizes the official fxevent.ConsoleLogger for event formatting
// and redirects the output to the provided UAP logger.
func Module() fx.Option {
	return fx.WithLogger(func(logger logger.Logger) fxevent.Logger {
		return &fxevent.ConsoleLogger{
			W: &loggerWriter{logger: logger},
		}
	})
}

// loggerWriter acts as a bridge between io.Writer and uap logger.
type loggerWriter struct {
	logger logger.Logger
}

func (w *loggerWriter) Write(p []byte) (n int, err error) {
	// Trim the trailing newline because the logger implementation (e.g., zap, logrus)
	// usually adds its own newline.
	msg := strings.TrimRight(string(p), "\n")
	if msg != "" {
		w.logger.Info(msg)
	}
	return len(p), nil
}
