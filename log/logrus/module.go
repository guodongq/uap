package logrus

import (
	logger "github.com/guodongq/uap/log"
	"go.uber.org/fx"
)

func Module(optionFuncs ...func(options *LogrusLoggerOptions)) fx.Option {
	return fx.Module(
		"logging.logrus",
		fx.Provide(func() logger.Logger {
			return New(optionFuncs...)
		}),
		fx.Invoke(func(logger logger.Logger) error {
			if logrusLogger, ok := logger.(*LogrusLoggerAdapter); ok {
				return logrusLogger.Init()
			}
			return nil
		}),
	)
}
