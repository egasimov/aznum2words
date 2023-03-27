package config

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/constant"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config holds a configuration variables.
type Config struct {
	AppHost    string `arg:"-h, env:APP_HOST" default:"0.0.0.0"`
	AppPort    string `arg:"-p, env:APP_PORT" default:"8080"`
	MetricPort string `arg:"-m, env:METRIC_PORT" default:"9090"`
}

func (this *Config) GetAppNetListenAddr() string {
	return fmt.Sprintf(
		"%s:%s",
		this.AppHost,
		this.AppPort,
	)
}

func (this *Config) GetMetricNetListenAddr() string {
	return fmt.Sprintf(
		"%s:%s",
		this.AppHost,
		this.MetricPort,
	)
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
