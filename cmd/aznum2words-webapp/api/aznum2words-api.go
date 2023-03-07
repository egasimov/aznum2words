package api

import (
	"github.com/egasimov/aznum2words"
	. "github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/models"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/enums"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AzNum2WordsApi struct {
}

func (a *AzNum2WordsApi) ConvertWordsToNumber(ctx echo.Context, params ConvertWordsToNumberParams) error {
	return ctx.JSON(http.StatusNotImplemented, nil)
}

func (a *AzNum2WordsApi) ConvertNumberToWord(ctx echo.Context, params ConvertNumberToWordParams) error {
	var convertRequest ConvertNumberToWordsRequest
	var err error
	err = ctx.Bind(&convertRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, UnknownError{Code: enums.ErrRequestBodyInvalid,
			Message: "Request is invalid"})
	}

	var words string
	words, err = aznum2words.SpellNumber(convertRequest.Number)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, UnknownError{Code: enums.ErrInternalServerError,
			Message: "Internal server error occurred"})
	}

	var responseBody = ConvertNumberToWordsResponse{Words: words}
	return ctx.JSON(http.StatusOK, responseBody)
}
