package converterapi

import (
	"github.com/egasimov/aznum2words"
	. "github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/models"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/constant"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Api struct {
}

func (a *Api) ConvertWordsToNumber(ctx echo.Context, params ConvertWordsToNumberParams) error {
	return ctx.JSON(http.StatusNotImplemented, nil)
}

func (a *Api) ConvertNumberToWord(ctx echo.Context, params ConvertNumberToWordParams) error {
	var convertRequest ConvertNumberToWordsRequest

	if err := ctx.Bind(&convertRequest); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			UnknownError{
				Code:    constant.ErrRequestBodyInvalid,
				Message: err.Error()},
		)
	}

	result, err := aznum2words.SpellNumber(convertRequest.Number)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			UnknownError{
				Code:    constant.ErrRequestBodyInvalid,
				Message: err.Error(),
			},
		)
	}

	var responseBody = ConvertNumberToWordsResponse{
		Words: result,
	}
	return ctx.JSON(http.StatusOK, responseBody)
}
