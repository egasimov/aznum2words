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
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	handlerApi := handler.NewHandler(converterApi, healthApi)
	handlerApi.Register(routerApp)

	// Create a prometheus server for exposing metrics
	routerPrometheus := router.NewPrometheusServer(routerApp)

	go func() {
		logger.
			Info(fmt.Sprintf("prometheus server started at %s", config.GetConfig().GetAppNetListenAddr()))

		if err := routerPrometheus.Start(config.GetConfig().GetMetricNetListenAddr()); err != nil && err != http.ErrServerClosed {
			routerPrometheus.Logger.Fatal("shutting down the prometheus server")
		}
	}()

	go func() {
		logger.
			Info(fmt.Sprintf("app server started at %s", config.GetConfig().GetAppNetListenAddr()))

		if err := routerApp.Start(config.GetConfig().GetAppNetListenAddr()); err != nil && err != http.ErrServerClosed {
			routerApp.Logger.Fatal("shutting down the app server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := routerApp.Shutdown(ctx); err != nil {
		logger.
			Error(fmt.Sprintf("an error occurred when shutdown app server, err: %s", err))
	} else {
		logger.
			Info("gracefully shut downed app server")
	}

	if err := routerPrometheus.Shutdown(ctx); err != nil {
		logger.
			Error(fmt.Sprintf("an error occurred when shutdown prometheus server, err: %s", err))
	} else {
		logger.
			Info("gracefully shut downed prometheus server")
	}
}
