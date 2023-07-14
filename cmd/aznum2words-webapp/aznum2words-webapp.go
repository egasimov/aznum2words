package main

import (
	"context"
	"fmt"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/converterapi"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/api/healthapi"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/config"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/handler"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/logger"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/router"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/utils"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/utils/listener"
)

func main() {
	logger := logger.Logger()
	defer logger.Sync()
	defer logger.Info("an application stopped")

	// Create an instance of our handler which satisfies the generated interface
	var converterApi = new(converterapi.Api)
	var healthApi = new(healthapi.Api)

	// Create an app server for handlers
	routerApp := router.NewMainServer()
	handlerAppApi := handler.NewHandler(converterApi, healthApi)
	appApiListener := listener.NewApiListener("app-server", routerApp, handlerAppApi, config.GetConfig().GetAppNetListenAddr())

	// Create a prometheus server for exposing metrics
	routerPrometheus := router.NewPrometheusServer(routerApp)
	promApiListener := listener.NewApiListener("prometheus-server", routerPrometheus, nil, config.GetConfig().GetMetricNetListenAddr())

	tasks := make([]utils.ITask, 0)
	tasks = utils.RegisterTaskWithCloser(tasks, appApiListener)
	tasks = utils.RegisterTaskWithCloser(tasks, promApiListener)

	if err := utils.RunTasksAsync(tasks, context.Background()); err != nil {
		logger.Fatal(fmt.Sprintf("app stopped with error: %s", err))
	}

}
