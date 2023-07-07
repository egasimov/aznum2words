package router

import (
	"fmt"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/logger"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"strings"
)

func NewMainServer() *echo.Echo {
	echoMainServer := echo.New()
	echoMainServer.HideBanner = true
	echoMainServer.HidePort = true

	echoMainServer.Use(echomiddleware.CORS())
	echoMainServer.Use(echomiddleware.RequestIDWithConfig(
		echomiddleware.RequestIDConfig{
			TargetHeader: echo.HeaderXCorrelationID,
		}))
	// Log all requests
	echoMainServer.Use(echomiddleware.LoggerWithConfig(getLoggerConfig()))

	return echoMainServer
}

func NewPrometheusServer(echoMainServer *echo.Echo) *echo.Echo {
	// Create Prometheus server and Middleware
	echoPrometheus := echo.New()
	echoPrometheus.HideBanner = true
	echoPrometheus.HidePort = true
	echoPrometheus.Use(echomiddleware.CORS())
	echoPrometheus.Use(echomiddleware.RequestIDWithConfig(
		echomiddleware.RequestIDConfig{
			TargetHeader: echo.HeaderXCorrelationID,
		}))
	// Log all requests
	echoPrometheus.Use(echomiddleware.LoggerWithConfig(getLoggerConfig()))

	prom := prometheus.NewPrometheus("echo", nil)

	// Scrape metrics from Main Server
	echoMainServer.Use(prom.HandlerFunc)

	// Setup metrics endpoint at another server
	prom.SetMetricsPath(echoPrometheus)

	return echoPrometheus
}

func getLoggerConfig() echomiddleware.LoggerConfig {
	return echomiddleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","x-correlation-id":"${header:x-correlation-id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			addDefaultLogEntries() +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}
}

func addDefaultLogEntries() string {
	var logEntries []string
	for k, v := range logger.GetDefaultLogEntries() {
		logEntries = append(logEntries, addKeyValPair(k, v, true))
	}

	return strings.Join(logEntries, "")
}

func addKeyValPair(keyName string, valName string, appendComma bool) string {
	var result string
	if appendComma {
		result = fmt.Sprintf(`"%s":"%s",`, keyName, valName)
	} else {
		result = fmt.Sprintf(`"%s":"%s"`, keyName, valName)
	}
	return result
}
