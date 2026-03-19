package fx

import (
	"strings"

	logger "github.com/guodongq/uap/common/logging"
)

// loggerWriter acts as a bridge between io.Writer and uap logger.
type loggerWriter struct {
	logger logger.Logger
}

func (w *loggerWriter) Write(p []byte) (n int, err error) {
	// Trim the trailing newline because the logger implementation (e.g., zap, logrus)
	// usually adds its own newline.
	msg := strings.TrimRight(string(p), "\n")

	// Replace tabs with spaces
	msg = strings.ReplaceAll(msg, "\t", " ")
	// Replace "<=" with "from" to avoid JSON HTML escaping (e.g. \u003c)
	msg = strings.ReplaceAll(msg, "<=", "from")

	if msg != "" {
		w.logger.Info(msg)
	}
	return len(p), nil
}
