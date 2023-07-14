package utils

import (
	"context"
	"fmt"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/logger"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/utils/closer"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/utils/listener"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

type ITask interface {
	Run(ctx context.Context) error
	GetName() string
}

func RegisterTaskWithCloser(taskList []ITask, task ITask) []ITask {
	logger := logger.Logger()
	defer logger.Sync()

	taskList = append(taskList, task)

	switch task.(type) {
	case listener.IApiListener:
		v, ok := task.(listener.IApiListener)
		if !ok {
			logger.Fatal(fmt.Sprintf("type cast failed. from %s to %s",
				reflect.TypeOf(task).String(), ""))
		}
		apiCloser := closer.NewApiCloser(v)

		return append(taskList, apiCloser)
	default:
		logger.Error(fmt.Sprintf("Unable to find appropriate closer."))
	}

	return taskList
}

func RunTasksAsync(tasks []ITask, parentCtx context.Context) error {
	logger := logger.Logger()
	defer logger.Sync()

	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	enableGracefullyShutdown(ctx, cancel)

	g, egCtx := errgroup.WithContext(ctx)
	for _, s := range tasks {
		lcs := s // create local copy of listener pointer https://golang.org/doc/faq#closures_and_goroutines

		g.TryGo(func() error {
			logger.Debug(fmt.Sprintf("task with name: %s starting...", lcs.GetName()))

			if err := lcs.Run(egCtx); err != nil {
				logger.Error(fmt.Sprintf("error: %s, when running task: %s", err, lcs.GetName()))
				return err
			}

			logger.Debug(fmt.Sprintf("task with name: %s succesfully finished", lcs.GetName()))
			return nil
		})
	}

	logger.Debug(fmt.Sprintf("All tasks are running now..."))
	return g.Wait()
}

func enableGracefullyShutdown(ctx context.Context, cancel context.CancelFunc) {
	go func() { // catch signal and invoke graceful termination
		logger := logger.Logger()
		defer logger.Sync()

		// Use a buffered channel to avoid missing signals as recommended for signal.Notify
		stopCh := make(chan os.Signal, 1)

		// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
		signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		defer signal.Stop(stopCh)

		select {
		case <-stopCh:
			logger.Debug("interrupt signal received, going to cancel tasks")
			cancel()

			return
		case <-ctx.Done():

			return
		}
	}()
}
