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

// Debug 打印Debug等级的日志
func Debug(msg string) {
	defaultLogger.Debug(msg)
}

// Info 打印Info等级的日志
func Info(msg string) {
	defaultLogger.Info(msg)
}

// Warn 打印Warn等级的日志
func Warn(msg string) {
	defaultLogger.Warn(msg)
}

// Error 打印Error等级的日志
func Error(msg string) {
	defaultLogger.Error(msg)
}

// Fatal 打印Fatal等级的日志
func Fatal(msg string) {
	defaultLogger.Fatal(msg)
}

// Debugf 使用fmt.Sprintf打印Debug等级的日志
func Debugf(template string, args ...interface{}) {
	defaultLogger.Debugf(template, args...)
}

// Infof 使用fmt.Sprintf打印Info等级的日志
func Infof(template string, args ...interface{}) {
	defaultLogger.Infof(template, args...)
}

// Warnf 使用fmt.Sprintf打印Warn等级的日志
func Warnf(template string, args ...interface{}) {
	defaultLogger.Warnf(template, args...)
}

// Errorf 使用fmt.Sprintf打印Error等级的日志
func Errorf(template string, args ...interface{}) {
	defaultLogger.Errorf(template, args...)
}

// Fatalf 使用fmt.Sprintf打印Fatal等级的日志
func Fatalf(template string, args ...interface{}) {
	defaultLogger.Fatalf(template, args...)
}

// WithField 向日志增加一个自定义字段
func WithField(key string, value interface{}) Logger {
	return &zapLoggerWrapper{
		defaultLogger.(*zapLoggerWrapper).Logger.
			WithOptions(zap.AddCallerSkip(-1)).
			With(zap.Any(key, value)),
	}
}

// WithFields 向日志增加多个自定义字段
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

// WithError 向日志增加error错误类型字段
func WithError(err error) Logger {
	return &zapLoggerWrapper{
		defaultLogger.(*zapLoggerWrapper).Logger.
			WithOptions(zap.AddCallerSkip(-1)).
			With(zap.Error(err)),
	}
}

// Named 向日志器增加标题
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

// WithContext 向ctx中新增日志器，如果logger为nil则按照默认选项创建新的logger并返回
func WithContext(ctx context.Context, logger Logger) (context.Context, Logger) {
	if ctx == nil {
		ctx = context.TODO()
	}
	if logger == nil {
		logger = &zapLoggerWrapper{
			defaultLogger.(*zapLoggerWrapper).Logger.WithOptions(zap.AddCallerSkip(-1)),
		}
	}
	ctx = context.WithValue(ctx, logkey{}, logger)
	return ctx, logger
}

// FromContext 从ctx中获取日志器，如果没有则创建新的日志器并返回新的context
func FromContext(ctx context.Context) (context.Context, Logger) {
	if ctx == nil {
		return WithContext(context.TODO(), nil)
	}
	logger, ok := ctx.Value(logkey{}).(Logger)
	if ok {
		return ctx, logger
	}
	return WithContext(ctx, nil)
}
