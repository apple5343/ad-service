package logger

import (
	"log"
	"os"
	"server/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

func Init(cfg config.LoggerConfig, options ...zap.Option) {
	globalLogger = zap.New(getCore(cfg), options...)
}

func Debug(msg string, fields ...zap.Field) {
	if globalLogger == nil {
		return
	}
	globalLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	if globalLogger == nil {
		return
	}
	globalLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	globalLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	if globalLogger == nil {
		return
	}
	globalLogger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	if globalLogger == nil {
		return
	}
	globalLogger.Fatal(msg, fields...)
}

func WithOptions(opts ...zap.Option) *zap.Logger {
	return globalLogger.WithOptions(opts...)
}

func getCore(cfg config.LoggerConfig) zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)

	loggerCfg := zap.NewDevelopmentEncoderConfig()
	loggerCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	loggerCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(loggerCfg)

	level := getAtomicLevel(cfg.Level())

	return zapcore.NewCore(consoleEncoder, stdout, level)
}

func getAtomicLevel(logLevel string) zap.AtomicLevel {
	var level zapcore.Level
	if err := level.Set(logLevel); err != nil {
		log.Fatalf("failed to set log level: %v", err)
	}

	return zap.NewAtomicLevelAt(level)
}
