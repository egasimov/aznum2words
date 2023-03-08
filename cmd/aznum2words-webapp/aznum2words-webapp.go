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
	"time"
)

var LoggerConfig = echomiddleware.LoggerConfig{
	Format: `{"time":"${time_rfc3339_nano}","x-correlation-id":"${header:x-correlation-id}","remote_ip":"${remote_ip}",` +
		`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
		`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"}` + "\n",
	CustomTimeFormat: "2006-01-02 15:04:05.00000",
}

func main() {
	var err error
	_, err = flags.Parse(&config.Opts)
	if err != nil {
		panic(err)
	}

	initEnvVars()
	config.LoadConfig()

	// Create an instance of our handler which satisfies the generated interface
	var azNum2WordsApi = new(api.AzNum2WordsApi)

	// This is how you set up a basic Echo router
	e := echo.New()
	// Log all requests
	e.Use(echomiddleware.LoggerWithConfig(LoggerConfig))

	// We now register our api above as the handler for the interface
	api.RegisterHandlers(e, azNum2WordsApi)
	registerHealthCheckProbes(e)

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start("0.0.0.0:" + config.Conf.Port))
}

func initEnvVars() {
	if godotenv.Load("cmd/aznum2words-webapp/profile/default.env") != nil {
		log.Fatal("Error in loading environment variables from: profiles/default.env")
	} else {
		log.Info("Environment variables loaded from: profiles/default.env")
	}

	if config.Opts.Profile != "default" {
		profileFileName := "profile/" + config.Opts.Profile + ".env"
		if godotenv.Overload(profileFileName) != nil {
			log.Fatal("Error in loading environment variables from: ", profileFileName)
		} else {
			log.Info("Environment variables overloaded from: ", profileFileName)
		}
	}
}

func registerHealthCheckProbes(e *echo.Echo) {
	checker := health.NewChecker(

		// Set the time-to-live for our cache to 1 second (default).
		health.WithCacheDuration(1*time.Second),

		// Configure a global timeout that will be applied to all checks.
		health.WithTimeout(10*time.Second),

		// The following check will be executed periodically every 15 seconds
		// started with an initial delay of 3 seconds. The check function will NOT
		// be executed for each HTTP request.
		health.WithPeriodicCheck(15*time.Second, 3*time.Second, health.Check{
			Name: "aznum2words",
			// The check function checks the health of a component. If an error is
			// returned, the component is considered unavailable (or "down").
			// The context contains a deadline according to the configured timeouts.
			Check: func(ctx context.Context) error {
				return nil
			},
		}),

		// Set a status listener that will be invoked when the health status changes.
		// More powerful hooks are also available (see docs).
		health.WithStatusListener(func(ctx context.Context, state health.CheckerState) {
			log.Info(fmt.Sprintf("health status changed to %s", state.Status))
		}),
	)

	e.GET("/health", echo.WrapHandler(health.NewHandler(checker)))
}
