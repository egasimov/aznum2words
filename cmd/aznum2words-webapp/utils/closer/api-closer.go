package closer

import (
	"context"
	"fmt"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/logger"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/utils/listener"
	"net/http"
)

type ApiCloser struct {
	target listener.IApiListener
}

func (s *ApiCloser) GetName() string {
	return "ApiCloser::" + s.target.GetName()
}

func NewApiCloser(t listener.IApiListener) *ApiCloser {
	return &ApiCloser{
		target: t,
	}
}

func (s *ApiCloser) Run(ctx context.Context) error {
	logger := logger.Logger()
	defer logger.Sync()

	<-ctx.Done()

	logger.Info(fmt.Sprintf("ApiCloser::[%s] Cancel event is received from patent context,"+
		" going to gracefully shut down", s.GetName()))

	if err := s.target.GetRouter().Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		logger.
			Error(fmt.Sprintf("ApiCloser::[%s] an error occurred when shutdown, err: %s", s.GetName(), err))
	} else {
		logger.
			Info(fmt.Sprintf("ApiCloser::[%s] gracefully shut downed", s.GetName()))
	}

	return nil
}
