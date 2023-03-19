package converter

import (
	"github.com/egasimov/aznum2words"
	. "github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/converter/models"
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
	var err error
	err = ctx.Bind(&convertRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, UnknownError{Code: constant.ErrRequestBodyInvalid,
			Message: "Request is invalid"})
	}

	var words string
	words, err = aznum2words.SpellNumber(convertRequest.Number)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, UnknownError{Code: constant.ErrInternalServerError,
			Message: "Internal server error occurred"})
	}

	var responseBody = ConvertNumberToWordsResponse{Words: words}
	return ctx.JSON(http.StatusOK, responseBody)
}
