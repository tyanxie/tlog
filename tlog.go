// Package tlog zap日志库的包装
package tlog

import (
	"context"

	"go.uber.org/zap"
)

var (
	defaultLogger Logger
)

type logkey struct{}

func Debug(msg string) {
	defaultLogger.Debug(msg)
}

func Info(msg string) {
	defaultLogger.Info(msg)
}

func Warn(msg string) {
	defaultLogger.Warn(msg)
}

func Error(msg string) {
	defaultLogger.Error(msg)
}

func Fatal(msg string) {
	defaultLogger.Fatal(msg)
}

func Debugf(template string, args ...interface{}) {
	defaultLogger.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	defaultLogger.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	defaultLogger.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	defaultLogger.Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	defaultLogger.Fatalf(template, args...)
}

func WithField(key string, value interface{}) Logger {
	return &zapLoggerWrapper{
		defaultLogger.(*zapLoggerWrapper).Logger.
			WithOptions(zap.AddCallerSkip(-1)).
			With(zap.Any(key, value)),
	}
}

func WithFields(fields Fields) Logger {
	if fields == nil || len(fields) == 0 {
		return defaultLogger
	}
	fs := make([]zap.Field, len(fields))
	i := 0
	for key, value := range fields {
		fs[i] = zap.Any(key, value)
		i++
	}
	return &zapLoggerWrapper{
		defaultLogger.(*zapLoggerWrapper).Logger.
			WithOptions(zap.AddCallerSkip(-1)).
			With(fs...),
	}
}

func WithError(err error) Logger {
	return &zapLoggerWrapper{
		defaultLogger.(*zapLoggerWrapper).Logger.
			WithOptions(zap.AddCallerSkip(-1)).
			With(zap.Error(err)),
	}
}

func Named(name string) Logger {
	if name == "" {
		return defaultLogger
	}
	return &zapLoggerWrapper{
		defaultLogger.(*zapLoggerWrapper).Logger.
			WithOptions(zap.AddCallerSkip(-1)).
			Named(name),
	}
}

func WithContext(ctx context.Context) (context.Context, Logger) {
	if ctx == nil {
		ctx = context.TODO()
	}

	var logger Logger = &zapLoggerWrapper{
		defaultLogger.(*zapLoggerWrapper).Logger.WithOptions(zap.AddCallerSkip(-1)),
	}
	ctx = context.WithValue(ctx, logkey{}, logger)
	return ctx, logger
}

func FromContext(ctx context.Context) (context.Context, Logger) {
	if ctx == nil {
		return WithContext(nil)
	}
	logger, ok := ctx.Value(logkey{}).(Logger)
	if ok {
		return ctx, logger
	}
	return WithContext(ctx)
}
