package config

import "github.com/alexflint/go-arg"

// Config holds a configuration variables.
type Config struct {
	Port string `arg:"env:PORT,required"`
	Host string `arg:"env:HOST, required"`
}

// Conf ...
var Conf Config

// LoadConfig ...
func LoadConfig() {
	arg.MustParse(&Conf)
}
