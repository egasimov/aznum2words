package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Api struct {
}

func (a *Api) GetHealth(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}
