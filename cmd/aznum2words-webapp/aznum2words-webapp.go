package main

import (
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/config"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/handler"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/router"
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
)

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

	r := router.New()
	h := handler.NewHandler(azNum2WordsApi)
	h.Register(r)

	// And we serve HTTP until the world ends.
	r.Logger.Fatal(r.Start(config.Conf.Host + ":" + config.Conf.Port))
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
