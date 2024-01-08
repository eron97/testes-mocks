package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	logConfiguration := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "mensagem",
			LevelKey:     "nivel",
			TimeKey:      "horario",
			NameKey:      "nome",
			CallerKey:    "quemChamou",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, _ := logConfiguration.Build()
	logger.Info("test")

	/*

			logger, _ := zap.NewProduction()
			defer logger.Sync()
			logger.Error("Failed to fetch URL",
				zap.String("url", "test"),
				zap.Int("attempt", 3),
				zap.Duration("backoff", time.Second),
			)

			OUTPUT

			{"level":"error","ts":1704712295.3010082,"caller":"logger/logger.go:13",
			"msg":"Failed to fetch URL","url":"test","attempt":3,"backoff":1,
			"stacktrace":"main.main\n\tC:/Users/erons/Golang/api-mvc/src/config/logger/logger.go:13\nruntime.main\n\tC:/Program Files/Go/src/runtime/proc.go:267"}

		/*

		/*
			logger, _ := zap.NewProduction()
			defer logger.Sync()
			logger.Info("Failed to fetch URL",
				zap.String("url", "test"),
				zap.Int("attempt", 3),
				zap.Duration("backoff", time.Second),
			)

			OUTPUT

				{"level":"info","ts":1704712024.488738,"caller":"logger/logger.go:13",
				"msg":"Failed to fetch URL","url":"test","attempt":3,"backoff":1}

	*/
}
