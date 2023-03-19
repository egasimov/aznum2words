package router

import (
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

var loggerConfig = echomiddleware.LoggerConfig{
	Format: `{"time":"${time_rfc3339_nano}","x-correlation-id":"${header:x-correlation-id}","remote_ip":"${remote_ip}",` +
		`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
		`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"}` + "\n",
	CustomTimeFormat: "2006-01-02 15:04:05.00000",
}

func New() *echo.Echo {
	e := echo.New()
	e.Use(echomiddleware.RequestIDWithConfig(
		echomiddleware.RequestIDConfig{
			TargetHeader: echo.HeaderXCorrelationID,
		}))
	// Log all requests
	e.Use(echomiddleware.LoggerWithConfig(loggerConfig))

	return e
}
