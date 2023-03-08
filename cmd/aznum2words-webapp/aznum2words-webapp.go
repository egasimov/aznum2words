package main

import (
	"context"
	"fmt"
	"github.com/alexliesenfeld/health"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/config"
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"os"
)

var LoggerConfig = echomiddleware.LoggerConfig{
	Format: `{"time":"${time_rfc3339_nano}","x-correlation-id":"${header:x-correlation-id}","remote_ip":"${remote_ip}",` +
		`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
		`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"}` + "\n",
	CustomTimeFormat: "2006-01-02 15:04:05.00000",
}

func main() {
	var err error
	_, err = flags.Parse(&config.Conf)
	if err != nil {
		panic(err)
	}

	initEnvVars()
	config.LoadConfig()

	// Create an instance of our handler which satisfies the generated interface
	var azNum2WordsApi = new(api.AzNum2WordsApi)

	// This is how you set up a basic Echo router
	e := echo.New()
	//Add x-correlation-id header to response
	e.Use(echomiddleware.RequestIDWithConfig(
		echomiddleware.RequestIDConfig{
			TargetHeader: echo.HeaderXCorrelationID,
		}))
	// Log all requests
	e.Use(echomiddleware.LoggerWithConfig(LoggerConfig))

	// We now register our api above as the handler for the interface
	api.RegisterHandlers(e, azNum2WordsApi)
	registerHealthCheckProbes(e)

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(config.Conf.Host + ":" + config.Conf.Port))
}

func initEnvVars() {
	var env = os.Getenv("ENVIRONMENT")
	if len(env) == 0 {
		env = "default"
	}
	profileFileName := "cmd/aznum2words-webapp/profile/" + env + ".env"
	if godotenv.Load(profileFileName) != nil {
		log.Fatal("Error in loading environment variables from: ", profileFileName)
	} else {
		log.Info("Environment variables overloaded from: ", profileFileName)
	}
}

func registerHealthCheckProbes(e *echo.Echo) {
	checker := health.NewChecker(
		// Set a status listener that will be invoked when the health status changes.
		health.WithStatusListener(func(ctx context.Context, state health.CheckerState) {
			log.Info(fmt.Sprintf("health status changed to %s", state.Status))
		}),
	)
	e.GET("/health", echo.WrapHandler(health.NewHandler(checker)))
}
