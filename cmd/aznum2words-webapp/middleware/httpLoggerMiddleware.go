package middleware

import (
	"fmt"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/constant"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"time"
)

func HttpLoggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCtx echo.Context) error {
			start := time.Now()

			req := echoCtx.Request()
			res := echoCtx.Response()

			zlog := logger.FromCtx(req.Context())

			zlog.Debug(
				fmt.Sprintf(
					"[Request]  | Uri: %s %s | RawQuery: %s",
					req.Method, req.RequestURI,
					req.URL.RawQuery,
				),
			)

			if err := next(echoCtx); err != nil {
				zlog.Error(
					fmt.Sprintf("An error occurred: %v", err),
					zap.Error(err),
				)
				echoCtx.Error(err)
			}

			latency := time.Since(start)

			zlog.Debug(
				fmt.Sprintf(
					"[Response] | Uri: %s %s | Elapsed time: %s | Result: %d",
					req.Method, req.RequestURI,
					latency.String(),
					res.Status,
				),
			)

			zlog.Info(
				fmt.Sprintf("[Request] Uri: %s %s | [Response] Status: %d Elapsed time: %s",
					req.Method, req.RequestURI,
					res.Status,
					latency.String(),
				),
				zap.String(constant.LogRequestMethod, req.Method),
				zap.String(constant.LogRequestUri, req.RequestURI),
				zap.String(constant.LogUserAgent, req.UserAgent()),
				zap.String(constant.LogClientRealIp, req.Header.Get(echo.HeaderXRealIP)),
				zap.Int(constant.LogResponseStatus, res.Status),
				zap.String(constant.LogLatency, latency.String()),
			)

			return nil
		}
	}
}
