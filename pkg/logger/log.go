package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(cfg *Config) *zap.Logger {
	core := zapcore.NewCore(
		prepEncoder(cfg),
		zapcore.Lock(zapcore.AddSync(os.Stderr)),
		prepLevel(cfg),
	)

	return zap.New(zapcore.NewTee(core), getOptions(cfg)...)
}

func prepEncoder(cfg *Config) zapcore.Encoder {
	var encoderConfig zapcore.EncoderConfig
	if cfg.Development {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if cfg.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	return encoder
}

func prepLevel(cfg *Config) zap.AtomicLevel {
	var level zapcore.Level

	if err := level.Set(cfg.Level); err != nil {
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	return zap.NewAtomicLevelAt(level)
}

func getOptions(cfg *Config) []zap.Option {
	return []zap.Option{
		zap.AddStacktrace(zap.ErrorLevel),
		zap.AddCaller(),
	}
}
