package handler

import (
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/converterapi"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/healthapi"
	"github.com/labstack/echo/v4"
)

// Handler ...
type Handler struct {
	converterApi *converterapi.Api
	healthApi    *healthapi.Api
}

// NewHandler ...

func NewHandler(c *converterapi.Api, h *healthapi.Api) *Handler {
	return &Handler{
		converterApi: c,
		healthApi:    h,
	}
}

// Register ...
func (h *Handler) Register(e *echo.Echo) {
	converterapi.RegisterHandlers(e, h.converterApi)
	healthapi.RegisterHandlers(e, h.healthApi)
}
