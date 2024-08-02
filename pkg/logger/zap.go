package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/marcelofabianov/my-cash/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
	config.LogConfig
}

func NewLogger(cfg config.LogConfig) (*Logger, error) {
	level := defineLevel(cfg.Level)
	encoderConfig := defineEncoderConfig(cfg.Format)

	core, err := defineOutputConfig(cfg, encoderConfig, level)
	if err != nil {
		return nil, err
	}

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	return &Logger{Logger: logger}, nil
}

func (l *Logger) Close() {
	if l.LogConfig.Output == "file" {
		if l == nil {
			return
		}
		if err := l.Sync(); err != nil {
			fmt.Printf("error syncing logger: %v\n", err)
		}
	}
}

func defineLevel(level string) zap.AtomicLevel {
	switch level {
	case "debug":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "panic":
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	default:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	}
}

func defineEncoderConfig(configFormat string) zapcore.EncoderConfig {
	switch configFormat {
	case "json":
		return zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			TimeKey:     "timestamp",
			EncodeTime:  zapcore.ISO8601TimeEncoder,
		}
	default:
		return zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			TimeKey:      "timestamp",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		}
	}
}

func defineOutputConfig(
	cfg config.LogConfig,
	encoderConfig zapcore.EncoderConfig,
	level zap.AtomicLevel,
) (zapcore.Core, error) {
	var core zapcore.Core
	switch cfg.Output {
	case "stdout":
		encoder := zapcore.NewJSONEncoder(encoderConfig)
		writer := zapcore.AddSync(os.Stdout)
		core = zapcore.NewCore(encoder, writer, level)
	case "file":
		file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		encoder := zapcore.NewJSONEncoder(encoderConfig)
		writer := zapcore.AddSync(file)
		core = zapcore.NewCore(encoder, writer, level)
	default:
		return nil, fmt.Errorf("invalid log output: %s", cfg.Output)
	}
	return core, nil
}

func (l *Logger) FieldString(key string, value string) zap.Field {
	return zap.String(key, value)
}

func (l *Logger) FieldInt(key string, value int) zap.Field {
	return zap.Int(key, value)
}

func (l *Logger) FieldDuration(key string, value time.Duration) zap.Field {
	return zap.Duration(key, value)
}

func (l *Logger) FieldError(err error) zap.Field {
	return zap.Error(err)
}

func (l *Logger) Field(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func ErrorField(err error) zap.Field {
	return zap.Error(err)
}
