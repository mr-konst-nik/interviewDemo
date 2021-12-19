package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

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
		return err
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
