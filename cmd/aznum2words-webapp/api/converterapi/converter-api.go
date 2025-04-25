package converterapi

import (
	"net/http"

	"github.com/egasimov/aznum2words"
	. "github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/models"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/constant"
	"github.com/egasimov/aznum2words/datetime"
	"github.com/labstack/echo/v4"
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

func (a *Api) ConvertDateToWords(ctx echo.Context, params ConvertDateToWordsParams) error {
	var convertRequest ConvertDateToWordsRequest

	if err := ctx.Bind(&convertRequest); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			UnknownError{
				Code:    constant.ErrRequestBodyInvalid,
				Message: err.Error()},
		)
	}

	result, err := datetime.ConvertDate(convertRequest.Date)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			UnknownError{
				Code:    constant.ErrRequestBodyInvalid,
				Message: err.Error(),
			},
		)
	}

	var responseBody = Word{
		Words: result,
	}
	return ctx.JSON(http.StatusOK, responseBody)
}

func (a *Api) ConvertTimeToWords(ctx echo.Context, params ConvertTimeToWordsParams) error {
	var convertRequest ConvertTimeToWordsRequest

	if err := ctx.Bind(&convertRequest); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			UnknownError{
				Code:    constant.ErrRequestBodyInvalid,
				Message: err.Error()},
		)
	}

	result, err := datetime.ConvertTime(convertRequest.Time)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			UnknownError{
				Code:    constant.ErrRequestBodyInvalid,
				Message: err.Error(),
			},
		)
	}

	var responseBody = Word{
		Words: result,
	}
	return ctx.JSON(http.StatusOK, responseBody)
}

func (a *Api) ConvertDateTimeToWords(ctx echo.Context, params ConvertDateTimeToWordsParams) error {
	var convertRequest ConvertDateTimeToWordsRequest

	if err := ctx.Bind(&convertRequest); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			UnknownError{
				Code:    constant.ErrRequestBodyInvalid,
				Message: err.Error()},
		)
	}

	result, err := datetime.ConvertDateTime(convertRequest.Datetime)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			UnknownError{
				Code:    constant.ErrRequestBodyInvalid,
				Message: err.Error(),
			},
		)
	}

	var responseBody = Word{
		Words: result,
	}
	return ctx.JSON(http.StatusOK, responseBody)
}
