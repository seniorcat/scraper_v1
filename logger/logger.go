package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// По умолчанию установлен no-op-логер, который не выводит никаких сообщений.
var Log *zap.Logger = zap.NewNop()

// Initialize инициализирует синглтон логера с необходимым уровнем логирования.
func Initialize(level string) error {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}
	cfg := zap.NewProductionConfig()

	cfg.Level = lvl
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05 MST")
	cfg.DisableCaller = true

	zl, err := cfg.Build()

	if err != nil {
		return err
	}

	defer zl.Sync()
	Log = zl
	return nil
}
