package zap

import (
	"os"

	"github.com/guodongq/uap/tools/sys/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loggerLevelEnvKey = []string{
		"LOGGER_LEVEL",
	}

	loggerFormatterEnvKey = []string{
		"LOGGER_FORMATTER",
	}

	loggerOutputEnvKey = []string{
		"LOGGER_OUTPUT",
	}
)

type ZapLoggerAdapterOptions struct {
	Level   zapcore.Level
	Encoder zapcore.Encoder
	Output  zapcore.WriteSyncer
}

func getDefaultZapLoggerAdapterOptions() ZapLoggerAdapterOptions {
	options := ZapLoggerAdapterOptions{
		Level:   zapcore.InfoLevel,
		Encoder: zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
		Output:  zapcore.Lock(zapcore.AddSync(os.Stdout)),
	}

	var envConfig struct {
		level     string
		formatter string
		output    string
	}
	env.SetFromEnvVal(&envConfig.level, loggerLevelEnvKey)
	env.SetFromEnvVal(&envConfig.formatter, loggerFormatterEnvKey)
	env.SetFromEnvVal(&envConfig.output, loggerOutputEnvKey)

	for _, optionFunc := range ParseZapLoggerAdapterOptions(envConfig.level, envConfig.formatter, envConfig.output) {
		optionFunc(&options)
	}

	return options
}

func ParseZapLoggerAdapterOptions(
	logLevel string,
	logFormat string,
	logOutput string,
) []func(*ZapLoggerAdapterOptions) {
	level, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		level = zapcore.InfoLevel // Default to Info level if parsing fails
	}

	var formatter zapcore.Encoder
	switch logFormat {
	case "text":
		formatter = zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	case "json":
		fallthrough
	default:
		formatter = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	}

	var output zapcore.WriteSyncer
	switch logOutput {
	case "stdout":
		output = zapcore.Lock(zapcore.AddSync(os.Stdout))
	case "stderr":
		fallthrough
	default:
		output = zapcore.Lock(zapcore.AddSync(os.Stderr))
	}

	return []func(*ZapLoggerAdapterOptions){
		WithZapLoggerAdapterOptionsLevel(level),
		WithZapLoggerAdapterOptionsEncoder(formatter),
		WithZapLoggerAdapterOptionsOutput(output),
	}
}

func WithZapLoggerAdapterOptionsLevel(level zapcore.Level) func(*ZapLoggerAdapterOptions) {
	return func(options *ZapLoggerAdapterOptions) {
		options.Level = level
	}
}

func WithZapLoggerAdapterOptionsEncoder(encoder zapcore.Encoder) func(*ZapLoggerAdapterOptions) {
	return func(options *ZapLoggerAdapterOptions) {
		options.Encoder = encoder
	}
}

func WithZapLoggerAdapterOptionsOutput(output zapcore.WriteSyncer) func(*ZapLoggerAdapterOptions) {
	return func(options *ZapLoggerAdapterOptions) {
		options.Output = output
	}
}
