package logger

import (
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/config"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/constant"
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	encoderConfig := zap.NewProductionConfig()
	encoderConfig.Level.SetLevel(zapcore.InfoLevel)
	if config.GetConfig().DeployEnv == constant.LOCAL_ENVIRONMENT {
		encoderConfig.Level.SetLevel(zapcore.DebugLevel)
	}
	encoderConfig.EncoderConfig.TimeKey = "timestamp"
	encoderConfig.EncoderConfig.MessageKey = "message"
	encoderConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339Nano)

	zapLogger, err := encoderConfig.Build()
	if err != nil {
		log.Fatalf("fail to build log. err: %s", err)
	}

	for k, v := range GetDefaultLogEntries() {
		zapLogger = zapLogger.
			With(zap.Any(k, v))
	}

	logger = zapLogger
}

func Logger() *zap.Logger {
	return logger
}

func GetDefaultLogEntries() map[string]string {
	return map[string]string{
		"app": config.GetConfig().AppName,
		"env": config.GetConfig().DeployEnv,
	}
}
