package middleware

import (
	"fmt"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func ContextFillerMiddleware(zlog *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCtx echo.Context) error {
			zlog.Debug(
				fmt.Sprintf(
					"Creating new context with ZLogger for incoming request",
				),
			)

			req := echoCtx.Request()
			newCtx := logger.WithCtx(echoCtx, req.Context(), logger.Logger())
			echoCtx.SetRequest(req.Clone(newCtx))

			return next(echoCtx)
		}
	}
}
