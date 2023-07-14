package handler

import (
	"fmt"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/converterapi"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/healthapi"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/config"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/constant"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/logger"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"io/ioutil"
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

	// Add swagger UI endpoints
	if config.GetConfig().DeployEnv != constant.PRODUCTION_ENVIRONMENT {
		registerSwaggerUI(e)
	}
}

func registerSwaggerUI(e *echo.Echo) {
	logger := logger.Logger()
	defer logger.Sync()

	swaggerUiPath := "/swagger/*"

	// Serve the OpenAPI documentation
	e.GET("/openapi.yaml", func(c echo.Context) error {
		// Read the OpenAPI specification YAML file
		filePath := "api/open-api-spec.yaml"
		openapiBytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			return err
		}

		// Set the response headers
		c.Response().Header().Set(echo.HeaderContentType, "application/x-yaml")
		c.Response().Write(openapiBytes)
		return nil
	})

	// Enable Swagger UI
	e.GET("/swagger/*", echoSwagger.EchoWrapHandler(
		func(c *echoSwagger.Config) {
			c.URL = "/openapi.yaml"
		}),
	)

	logger.Info(fmt.Sprintf("Swagger-UI is available at %s%s", config.GetConfig().GetAppNetListenAddr(), swaggerUiPath))
}
