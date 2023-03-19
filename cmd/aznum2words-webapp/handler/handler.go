package handler

import (
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/converter"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/health"
	"github.com/labstack/echo/v4"
)

// Handler ...
type Handler struct {
	converterApi *converter.Api
	healthApi    *health.Api
}

// NewHandler ...

func NewHandler(c *converter.Api, h *health.Api) *Handler {
	return &Handler{
		converterApi: c,
		healthApi:    h,
	}
}

// Register ...
func (h *Handler) Register(e *echo.Echo) {
	converter.RegisterHandlers(e, h.converterApi)
	health.RegisterHandlers(e, h.healthApi)
}
