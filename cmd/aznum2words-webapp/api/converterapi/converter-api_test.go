package converterapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestConvertNumberToWord(t *testing.T) {
	e := echo.New()
	api := &Api{}

	testCases := []struct {
		name           string
		requestBody    ConvertNumberToWords
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Valid number",
			requestBody: ConvertNumberToWords{
				Number: "123",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"words":"bir yüz iyirmi üç"}`,
		},
		{
			name: "Invalid number",
			requestBody: ConvertNumberToWords{
				Number: "invalid",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"code":"request_body_invalid","message":"an invalid argument provided"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/conversion/to-word", bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := api.ConvertNumberToWord(c, ConvertNumberToWordParams{})
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.JSONEq(t, tc.expectedBody, rec.Body.String())
		})
	}
}

func TestConvertDateToWords(t *testing.T) {
	e := echo.New()
	api := &Api{}

	testCases := []struct {
		name           string
		requestBody    ConvertDateToWords
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Valid date",
			requestBody: ConvertDateToWords{
				Date: "1990-01-01",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"words":"bir yanvar min doqquz yüz doxsan"}`,
		},
		{
			name: "Invalid date",
			requestBody: ConvertDateToWords{
				Date: "invalid-date",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"code":"request_body_invalid","message":"invalid date format: parsing time \"invalid-date\" as \"2006-01-02\": cannot parse \"invalid-date\" as \"2006\""}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/conversion/date-to-words", bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := api.ConvertDateToWords(c, ConvertDateToWordsParams{})
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.JSONEq(t, tc.expectedBody, rec.Body.String())
		})
	}
}

func TestConvertTimeToWords(t *testing.T) {
	e := echo.New()
	api := &Api{}

	testCases := []struct {
		name           string
		requestBody    ConvertTimeToWords
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Valid time",
			requestBody: ConvertTimeToWords{
				Time: "14:30",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"words":"on dörd otuz"}`,
		},
		{
			name: "Invalid time",
			requestBody: ConvertTimeToWords{
				Time: "invalid-time",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"code":"request_body_invalid","message":"invalid time format"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/conversion/time-to-words", bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := api.ConvertTimeToWords(c, ConvertTimeToWordsParams{})
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.JSONEq(t, tc.expectedBody, rec.Body.String())
		})
	}
}

func TestConvertDateTimeToWords(t *testing.T) {
	e := echo.New()
	api := &Api{}

	testCases := []struct {
		name           string
		requestBody    ConvertDateTimeToWords
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Valid datetime",
			requestBody: ConvertDateTimeToWords{
				Datetime: "1990-01-01 09:15",
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"words":"bir yanvar min doqquz yüz doxsan, doqquz on beş"}`,
		},
		{
			name: "Invalid datetime",
			requestBody: ConvertDateTimeToWords{
				Datetime: "invalid-datetime",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"code":"request_body_invalid","message":"invalid datetime format"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/conversion/datetime-to-words", bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := api.ConvertDateTimeToWords(c, ConvertDateTimeToWordsParams{})
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.JSONEq(t, tc.expectedBody, rec.Body.String())
		})
	}
}
