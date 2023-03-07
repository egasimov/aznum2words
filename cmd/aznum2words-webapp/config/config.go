package config

import "github.com/alexflint/go-arg"

// Opts ...
var Opts struct {
	Profile string `short:"p" long:"profile" default:"default" description:"Application run profile"`
}

// Config holds a configuration variables.
type Config struct {
	Port string `arg:"env:PORT,required"`
}

// Conf ...
var Conf Config

// LoadConfig ...
func LoadConfig() {
	arg.MustParse(&Conf)
}
