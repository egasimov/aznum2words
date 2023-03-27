package main

import (
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

	// Create a app server for handlers
	routerApp := router.NewMainServer()
	handlerApi := handler.NewHandler(converterApi, healthApi)
	handlerApi.Register(routerApp)

	// Create a prometheus server for exposing metrics
	routerPrometheus := router.NewPrometheusServer(routerApp)

	go func() {
		routerPrometheus.Logger.Fatal(routerPrometheus.Start(
			config.GetConfig().GetMetricNetListenAddr()),
		)
	}()

	// And we serve HTTP until the world ends.
	routerApp.Logger.Fatal(routerApp.Start(
		config.GetConfig().GetAppNetListenAddr(),
	))
}
