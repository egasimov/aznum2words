package router

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

var loggerConfig = echomiddleware.LoggerConfig{
	Format: `{"time":"${time_rfc3339_nano}","x-correlation-id":"${header:x-correlation-id}","remote_ip":"${remote_ip}",` +
		`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
		`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"}` + "\n",
	CustomTimeFormat: "2006-01-02 15:04:05.00000",
}

func NewMainServer() *echo.Echo {
	echoMainServer := echo.New()
	echoMainServer.HideBanner = true

	echoMainServer.Use(echomiddleware.CORS())
	echoMainServer.Use(echomiddleware.RequestIDWithConfig(
		echomiddleware.RequestIDConfig{
			TargetHeader: echo.HeaderXCorrelationID,
		}))
	// Log all requests
	echoMainServer.Use(echomiddleware.LoggerWithConfig(loggerConfig))

	return echoMainServer
}

func NewPrometheusServer(echoMainServer *echo.Echo) *echo.Echo {
	// Create Prometheus server and Middleware
	echoPrometheus := echo.New()
	echoPrometheus.HideBanner = true

	echoPrometheus.Use(echomiddleware.CORS())
	echoPrometheus.Use(echomiddleware.RequestIDWithConfig(
		echomiddleware.RequestIDConfig{
			TargetHeader: echo.HeaderXCorrelationID,
		}))
	// Log all requests
	echoPrometheus.Use(echomiddleware.LoggerWithConfig(loggerConfig))

	prom := prometheus.NewPrometheus("echo", nil)

	// Scrape metrics from Main Server
	echoMainServer.Use(prom.HandlerFunc)

	// Setup metrics endpoint at another server
	prom.SetMetricsPath(echoPrometheus)

	return echoPrometheus
}
