package logger

import (
	"context"
	"fmt"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/config"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/constant"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

var _logger *zap.Logger

func init() {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level.SetLevel(zapcore.DebugLevel)
	if config.GetConfig().DeployEnv == constant.LOCAL_ENVIRONMENT {
		zapConfig.Level.SetLevel(zapcore.DebugLevel)
	}

	// create encoder config
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.MessageKey = "message"
	encoderConfig.EncodeTime = func(time time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(time.Format("2006-01-02T15:04:05.999Z"))
	}

	zapConfig.EncoderConfig = encoderConfig

	zapLogger, err := zapConfig.Build()
	if err != nil {
		log.Fatalf("fail to build log. err: %s", err)
	}

	var defaultFields []zap.Field
	defaultLogEntries, ok := getDefaultLogEntries()

	if !ok {
		log.Fatalf("fail to build log. err: %s", err)
	}

	for k, v := range defaultLogEntries {
		defaultFields = append(defaultFields, zap.String(k, v))
	}

	_logger = zapLogger.
		With(defaultFields...)
}

func Logger() *zap.Logger {
	return _logger
}

func getDefaultLogEntries() (map[string]string, bool) {
	buildInfo, ok := debug.ReadBuildInfo()

	if !ok {
		return nil, false
	}

	return map[string]string{
		constant.LogAppName:   config.GetConfig().AppName,
		constant.LogEnvName:   config.GetConfig().DeployEnv,
		constant.LogGoVersion: buildInfo.GoVersion,
		constant.LogProcessId: fmt.Sprintf("%d", os.Getpid()),
		constant.LogGoArch:    runtime.GOARCH,
	}, true
}

func FromCtx(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(constant.CtxZLogger).(*zap.Logger); ok {
		l.Warn(
			fmt.Sprintf("ZLogger found in the context, returning it to caller"),
		)
		return l
	}
	_logger.Warn(
		fmt.Sprintf("ZLogger could NOT found in the context, returning default ZLogger"),
	)
	return _logger
}

func WithCtx(echoCtx echo.Context, ctx context.Context, zlogger *zap.Logger) context.Context {
	if foundZLogger, ok := ctx.Value(constant.CtxZLogger).(*zap.Logger); ok {
		if foundZLogger == zlogger {
			return ctx
		}
	}

	newZLogger := buildZLoggerFromEchoCtx(echoCtx, zlogger)
	newCtx := context.WithValue(ctx, constant.CtxZLogger, newZLogger)
	newCtx = context.WithValue(
		newCtx, constant.CtxCorrelationId, echoCtx.Request().Header.Get(echo.HeaderXCorrelationID),
	)
	return newCtx
}

func buildZLoggerFromEchoCtx(echoCtx echo.Context, zlogger *zap.Logger) *zap.Logger {
	fields := []zap.Field{
		zap.String(constant.LogCorrelationId, echoCtx.Request().Header.Get(constant.HeaderXCorrelationId)),
		zap.String(constant.LogUserAgent, echoCtx.Request().Header.Get(constant.HeaderUserAgent)),
		// Add more key-value pairs as needed
	}
	return zlogger.With(fields...)
}
