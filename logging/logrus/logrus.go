package logrus

import (
	"context"

	logger "github.com/guodongq/uap/logging"
	"github.com/sirupsen/logrus"
)

var logrusToLevel = map[logrus.Level]logger.Level{
	logrus.DebugLevel: logger.DebugLevel,
	logrus.InfoLevel:  logger.InfoLevel,
	logrus.WarnLevel:  logger.WarnLevel,
	logrus.ErrorLevel: logger.ErrorLevel,
	logrus.FatalLevel: logger.FatalLevel,
}

type LogrusLoggerAdapter struct {
	entry *logrus.Entry
}

func New(optionFuncs ...func(*LogrusLoggerOptions)) *LogrusLoggerAdapter {
	options := getDefaultLogrusLoggerOptions()

	for _, optionFunc := range optionFuncs {
		optionFunc(&options)
	}

	log := logrus.New()
	log.Out = options.Output
	log.Formatter = options.Formatter
	log.Level = options.Level

	return &LogrusLoggerAdapter{
		entry: logrus.NewEntry(log),
	}
}

func (l *LogrusLoggerAdapter) Init() error {
	defaultLevel := logger.InfoLevel
	if level, exists := logrusToLevel[l.entry.Logger.Level]; exists {
		defaultLevel = level
	}
	logger.SetDefaultLevel(defaultLevel)
	logger.SetDefaultLogger(l)
	return nil
}

func (l *LogrusLoggerAdapter) WithContext(ctx context.Context) logger.Logger {
	newEntry := l.entry.WithContext(ctx)
	return &LogrusLoggerAdapter{
		entry: newEntry,
	}
}

func (l *LogrusLoggerAdapter) WithError(err error) logger.Logger {
	newEntry := l.entry.WithError(err)
	return &LogrusLoggerAdapter{
		entry: newEntry,
	}
}

func (l *LogrusLoggerAdapter) WithField(key string, value any) logger.Logger {
	newEntry := l.entry.WithField(key, value)
	return &LogrusLoggerAdapter{
		entry: newEntry,
	}
}

func (l *LogrusLoggerAdapter) WithFields(fields logger.Fields) logger.Logger {
	newEntry := l.entry.WithFields(map[string]interface{}(fields))
	return &LogrusLoggerAdapter{
		entry: newEntry,
	}
}

func (l *LogrusLoggerAdapter) Logger() *logrus.Entry {
	return l.entry
}

func (l *LogrusLoggerAdapter) Debug(v ...any) {
	l.entry.Debug(v...)
}
func (l *LogrusLoggerAdapter) Debugf(format string, v ...any) {
	l.entry.Debugf(format, v...)
}

func (l *LogrusLoggerAdapter) Info(v ...any) {
	l.entry.Info(v...)
}

func (l *LogrusLoggerAdapter) Infof(format string, v ...any) {
	l.entry.Infof(format, v...)
}

func (l *LogrusLoggerAdapter) Warn(v ...any) {
	l.entry.Warn(v...)
}

func (l *LogrusLoggerAdapter) Warnf(format string, v ...any) {
	l.entry.Warnf(format, v...)
}

func (l *LogrusLoggerAdapter) Error(v ...any) {
	l.entry.Error(v...)
}

func (l *LogrusLoggerAdapter) Errorf(format string, v ...any) {
	l.entry.Errorf(format, v...)
}

func (l *LogrusLoggerAdapter) Fatal(v ...any) {
	l.entry.Panic(v...)
}

func (l *LogrusLoggerAdapter) Fatalf(format string, v ...any) {
	l.entry.Panicf(format, v...)
}
