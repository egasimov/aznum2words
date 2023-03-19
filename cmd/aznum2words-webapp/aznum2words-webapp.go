package main

import (
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/converter"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/health"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/config"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/constant"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/handler"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/router"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
)

func main() {
	initEnvVars()
	config.LoadConfig()

	// Create an instance of our handler which satisfies the generated interface
	var converterApi = new(converter.Api)
	var healthApi = new(health.Api)

	r := router.New()
	h := handler.NewHandler(converterApi, healthApi)
	h.Register(r)

	// And we serve HTTP until the world ends.
	r.Logger.Fatal(r.Start(config.Conf.Host + ":" + config.Conf.Port))
}

func initEnvVars() {
	var env = os.Getenv(constant.DeployEnvKey)
	if len(env) == 0 {
		env = "default"
	}
	profileFileName := "cmd/aznum2words-webapp/profile/" + env + ".env"
	if err := godotenv.Load(profileFileName); err != nil {
		log.Fatalf("Error in loading environment variables from '%s' with error description '%s'",
			profileFileName, err)
	} else {
		log.Infof("Environment variables overloaded from: %s", profileFileName)
	}
}
