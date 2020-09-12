package webserver

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger ...
func NewLogger(logLevel string) *zap.Logger {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:     logLevel,
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.FullCallerEncoder,
		},
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Println(err)
	}
	return logger
}
