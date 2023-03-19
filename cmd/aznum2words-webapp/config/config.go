package config

import (
	"github.com/alexflint/go-arg"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/constant"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config holds a configuration variables.
type Config struct {
	Port string `arg:"-p, env:PORT" default:"8080"`
	Host string `arg:"-h, env:HOST" default:"0.0.0.0"`
}

var conf *Config = &Config{}

// LoadConfig ...
func LoadConfig() {
	// Local development purpose
	if env := os.Getenv(constant.DEPLOY_ENV_KEY); env == "" {
		errLoad := godotenv.Load(".env.local")

		if errLoad != nil {
			log.Fatalln(errLoad)
		}
	}

	//Processes command line arguments, environment variables
	// and exits upon failure
	arg.MustParse(conf)
}

func GetConfig() *Config {
	if conf == nil {
		log.Fatalln("Config not loaded...")
	}

	return conf
}
