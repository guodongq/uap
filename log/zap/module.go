package zap

import (
	logger "github.com/guodongq/uap/log"
	"go.uber.org/fx"
)

func Module(optionFuncs ...func(*ZapLoggerAdapterOptions)) fx.Option {
	return fx.Module(
		"logging.zap",
		fx.Provide(
			func() logger.Logger {
				return New(optionFuncs...)
			},
		),
		fx.Invoke(func(logger logger.Logger) error {
			if zapLogger, ok := logger.(*ZapLoggerAdapter); ok {
				return zapLogger.Init()
			}
			return nil
		}),
	)
}
