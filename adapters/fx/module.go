package uapfx

import (
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
