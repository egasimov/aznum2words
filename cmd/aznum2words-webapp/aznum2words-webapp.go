package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/converterapi"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/healthapi"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/config"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/handler"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/router"
)

func main() {
	config.LoadConfig()

	// Create an instance of our handler which satisfies the generated interface
	var converterApi = new(converterapi.Api)
	var healthApi = new(healthapi.Api)

	r := router.New()
	h := handler.NewHandler(converterApi, healthApi)
	h.Register(r)

	netListenAddr := spew.Sprintf(
		"%s:%s",
		config.GetConfig().Host,
		config.GetConfig().Port,
	)

	// And we serve HTTP until the world ends.
	r.Logger.Fatal(r.Start(netListenAddr))
}
