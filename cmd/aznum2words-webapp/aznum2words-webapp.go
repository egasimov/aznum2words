package main

import (
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/config"
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
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

	// We now register our petStore above as the handler for the interface
	api.RegisterHandlers(e, azNum2WordsApi)

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
