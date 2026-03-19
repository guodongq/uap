package zap

import (
	"context"

	"github.com/guodongq/uap/common/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapToLevel = map[zapcore.Level]logging.Level{
	zapcore.DebugLevel: logging.DebugLevel,
	zapcore.InfoLevel:  logging.InfoLevel,
	zapcore.WarnLevel:  logging.WarnLevel,
	zapcore.ErrorLevel: logging.ErrorLevel,
	zapcore.FatalLevel: logging.FatalLevel,
}

type ZapLoggerAdapter struct {
	logger *zap.Logger
}

func New(optionFuncs ...func(*ZapLoggerAdapterOptions)) *ZapLoggerAdapter {
	options := getDefaultZapLoggerAdapterOptions()

	for _, optionFunc := range optionFuncs {
		optionFunc(&options)
	}

	core := zapcore.NewCore(
		options.Encoder,
		options.Output,
		options.Level,
	)

	logger := zap.New(core, zap.AddCallerSkip(1))
	return &ZapLoggerAdapter{
		logger: logger,
	}
}

func (p *ZapLoggerAdapter) Init() error {
	defaultLevel := logging.InfoLevel
	if level, exists := zapToLevel[p.logger.Level()]; exists {
		defaultLevel = level
	}
	logging.SetDefaultLevel(defaultLevel)
	logging.SetDefaultLogger(p)
	return nil
}

func (p *ZapLoggerAdapter) WithContext(ctx context.Context) logging.Logger {
	newLogger := p.logger
	if reqID, ok := ctx.Value("request-id").(string); ok && reqID != "" {
		newLogger = newLogger.With(zap.String("request-id", reqID))
	}
	return &ZapLoggerAdapter{
		logger: newLogger,
	}
}

func (p *ZapLoggerAdapter) WithError(err error) logging.Logger {
	newLogger := p.logger.With(zap.Error(err))

	return &ZapLoggerAdapter{
		logger: newLogger,
	}
}

func (p *ZapLoggerAdapter) WithField(key string, value any) logging.Logger {
	newLogger := p.logger.With(zap.Any(key, value))

	return &ZapLoggerAdapter{
		logger: newLogger,
	}
}

func (p *ZapLoggerAdapter) WithFields(fields logging.Fields) logging.Logger {
	zapFields := make([]zap.Field, 0, len(fields))
	for key, value := range fields {
		if key == "" {
			continue
		}
		zapFields = append(zapFields, zap.Any(key, value))
	}

	newLogger := p.logger.With(zapFields...)

	return &ZapLoggerAdapter{
		logger: newLogger,
	}
}

func (p *ZapLoggerAdapter) Logger() *zap.Logger {
	return p.logger
}

func (p *ZapLoggerAdapter) Debug(args ...interface{}) {
	p.logger.Sugar().Debug(args...)
}

func (p *ZapLoggerAdapter) Debugf(format string, args ...interface{}) {
	p.logger.Sugar().Debugf(format, args...)
}

func (p *ZapLoggerAdapter) Info(args ...interface{}) {
	p.logger.Sugar().Info(args...)
}

func (p *ZapLoggerAdapter) Infof(format string, args ...interface{}) {
	p.logger.Sugar().Infof(format, args...)
}

func (p *ZapLoggerAdapter) Warn(args ...interface{}) {
	p.logger.Sugar().Warn(args...)
}

func (p *ZapLoggerAdapter) Warnf(format string, args ...interface{}) {
	p.logger.Sugar().Warnf(format, args...)
}

func (p *ZapLoggerAdapter) Error(args ...interface{}) {
	p.logger.Sugar().Error(args...)
}

func (p *ZapLoggerAdapter) Errorf(format string, args ...interface{}) {
	p.logger.Sugar().Errorf(format, args...)
}

func (p *ZapLoggerAdapter) Fatal(args ...interface{}) {
	p.logger.Sugar().Fatal(args...)
}

func (p *ZapLoggerAdapter) Fatalf(format string, args ...interface{}) {
	p.logger.Sugar().Fatalf(format, args...)
}
