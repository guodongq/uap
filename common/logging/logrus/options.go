package logrus

import (
	"io"
	"os"

	"github.com/guodongq/uap/tools/sys/env"
	"github.com/sirupsen/logrus"
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

type LogrusLoggerOptions struct {
	Level     logrus.Level
	Formatter logrus.Formatter
	Output    io.Writer
}

func getDefaultLogrusLoggerOptions() LogrusLoggerOptions {
	options := LogrusLoggerOptions{
		Level:     logrus.InfoLevel,
		Formatter: &logrus.JSONFormatter{},
		Output:    os.Stderr,
	}

	var envConfig struct {
		level     string
		formatter string
		output    string
	}
	env.SetFromEnvVal(&envConfig.level, loggerLevelEnvKey)
	env.SetFromEnvVal(&envConfig.formatter, loggerFormatterEnvKey)
	env.SetFromEnvVal(&envConfig.output, loggerOutputEnvKey)

	for _, optionFunc := range ParseLogrusLoggerAdapterOptions(envConfig.level, envConfig.formatter, envConfig.output) {
		optionFunc(&options)
	}

	return options
}

func ParseLogrusLoggerAdapterOptions(
	logLevel string,
	logFormatter string,
	logOutput string,
) []func(*LogrusLoggerOptions) {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.InfoLevel
	}

	var formatter logrus.Formatter
	switch logFormatter {
	case "text":
		formatter = &logrus.TextFormatter{}
	case "text_clr":
		formatter = &logrus.TextFormatter{ForceColors: true, FullTimestamp: true}
	case "json":
		fallthrough
	default:
		formatter = &logrus.JSONFormatter{}
	}

	var output io.Writer
	switch logOutput {
	case "stdout":
		output = os.Stdout
	case "stderr":
		fallthrough
	default:
		output = os.Stderr
	}

	return []func(options *LogrusLoggerOptions){
		WithLogrusLoggerOptionsLevel(level),
		WithLogrusLoggerOptionsFormatter(formatter),
		WithLogrusLoggerOptionsOutput(output),
	}
}

func WithLogrusLoggerOptionsLevel(level logrus.Level) func(*LogrusLoggerOptions) {
	return func(options *LogrusLoggerOptions) {
		options.Level = level
	}
}

func WithLogrusLoggerOptionsFormatter(formatter logrus.Formatter) func(*LogrusLoggerOptions) {
	return func(options *LogrusLoggerOptions) {
		options.Formatter = formatter
	}
}

func WithLogrusLoggerOptionsTextFormatter() func(*LogrusLoggerOptions) {
	return func(options *LogrusLoggerOptions) {
		options.Formatter = &logrus.TextFormatter{}
	}
}

func WithLogrusLoggerOptionsTextClrFormatter() func(*LogrusLoggerOptions) {
	return func(options *LogrusLoggerOptions) {
		options.Formatter = &logrus.TextFormatter{ForceColors: true, FullTimestamp: true}
	}
}

func WithLogrusLoggerOptionsOutput(output io.Writer) func(*LogrusLoggerOptions) {
	return func(options *LogrusLoggerOptions) {
		options.Output = output
	}
}

func WithLogrusLoggerOptionsStdoutOutput() func(*LogrusLoggerOptions) {
	return func(options *LogrusLoggerOptions) {
		options.Output = os.Stdout
	}
}
