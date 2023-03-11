package handler

import (
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api"
	"github.com/labstack/echo/v4"
)

// Handler ...
type Handler struct {
	aznum2wordsApi *api.AzNum2WordsApi
}

// NewHandler ...

func NewHandler(
	a *api.AzNum2WordsApi) *Handler {
	return &Handler{
		aznum2wordsApi: a,
	}
}

// Register ...
func (h *Handler) Register(e *echo.Echo) {
	api.RegisterHandlers(e, h.aznum2wordsApi)

	e.GET("/health", h.Health)
	e.GET("/readiness", h.Health)
}
