package router

import (
	"fmt"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/logger"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/middleware"
	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

var _requestIDConfig = echomiddleware.RequestIDConfig{
	TargetHeader: echo.HeaderXCorrelationID,
	Generator: func() string {
		return uuid.New().String()
	},
	RequestIDHandler: func(echoCtx echo.Context, newRequestId string) {
		zlog := logger.FromCtx(echoCtx.Request().Context())

		req := echoCtx.Request()
		if val := req.Header.Get(echo.HeaderXCorrelationID); val == "" {
			zlog.Info(
				fmt.Sprintf(
					"Correlation ID not found in the request, going to set ID %s to request",
					newRequestId,
				),
			)
			echoCtx.Request().Header.
				Set(echo.HeaderXCorrelationID, newRequestId)
		}
	},
}

func NewMainServer() *echo.Echo {
	echoMainServer := echo.New()
	echoMainServer.HideBanner = true
	echoMainServer.HidePort = true

	echoMainServer.Use(echomiddleware.CORS())
	echoMainServer.Use(echomiddleware.RequestIDWithConfig(_requestIDConfig))

	// Log all requests
	echoMainServer.Use(middleware.ContextFillerMiddleware(logger.Logger()))
	echoMainServer.Use(middleware.HttpLoggerMiddleware())

	return echoMainServer
}

func NewPrometheusServer(echoMainServer *echo.Echo) *echo.Echo {
	// Create Prometheus server and Middleware
	echoPrometheus := echo.New()
	echoPrometheus.HideBanner = true
	echoPrometheus.HidePort = true

	echoPrometheus.Use(echomiddleware.CORS())
	echoPrometheus.Use(echomiddleware.RequestIDWithConfig(_requestIDConfig))

	// Log all requests
	echoPrometheus.Use(middleware.ContextFillerMiddleware(logger.Logger()))
	echoPrometheus.Use(middleware.HttpLoggerMiddleware())

	prom := prometheus.NewPrometheus("echo", nil)

	// Scrape metrics from Main Server
	echoMainServer.Use(prom.HandlerFunc)

	// Setup metrics endpoint at another server
	prom.SetMetricsPath(echoPrometheus)

	return echoPrometheus
}
