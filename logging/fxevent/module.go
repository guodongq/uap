package fxevent

import (
	logger "github.com/guodongq/uap/logging"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func Module() fx.Option {
	return fx.WithLogger(func(logger logger.Logger) fxevent.Logger {
		return New(logger)
	})
}
