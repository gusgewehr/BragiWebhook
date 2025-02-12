package infrastructure

import (
	"go.uber.org/zap"
	"log"
	"os"
)

func NewLogger() *zap.Logger {

	logLevel := zap.NewAtomicLevelAt(zap.InfoLevel)

	if os.Getenv("APP_ENV") == "development" {
		logLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	config := zap.Config{
		Level:            logLevel,
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := config.Build() //zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	defer logger.Sync() // flushes buffer, if any

	return logger
}
