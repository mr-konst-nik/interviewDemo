package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log is zap logger
var Log *zap.Logger

// CreateLogger for app
func CreateLogger(logFile string, logLevel string) error {
	var err error
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(defineLogLevel(logLevel)),
		OutputPaths:      []string{logFile, "stderr"},
		ErrorOutputPaths: []string{logFile, "stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	if Log, err = cfg.Build(); err != nil {
		return fmt.Errorf("failed to create logger: %v", err)
	}
	return nil
}

func defineLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "info":
		return zapcore.InfoLevel
	case "debug":
		return zapcore.DebugLevel
	case "error":
		return zapcore.ErrorLevel
	case "warn":
		return zapcore.WarnLevel
	default:
		return zapcore.ErrorLevel
	}
}
