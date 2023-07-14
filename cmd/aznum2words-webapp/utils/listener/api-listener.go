package listener

import (
	"context"
	"fmt"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/handler"
	"github.com/egasimov/aznum2words/cmd/aznum2words-webapp/logger"
	"github.com/labstack/echo/v4"
)

type IApiListener interface {
	Run(context.Context) error
	GetName() string
	GetRouter() *echo.Echo
	GetAddress() string
}

type ApiListener struct {
	name    string
	router  *echo.Echo
	handler *handler.Handler
	address string
}

func NewApiListener(name string, echoRouter *echo.Echo, handlerApi *handler.Handler, address string) *ApiListener {
	aLisnr := &ApiListener{
		name:    name,
		router:  echoRouter,
		handler: handlerApi,
		address: address,
	}

	if handlerApi != nil {
		handlerApi.Register(echoRouter)
	}

	return aLisnr
}

func (s *ApiListener) Run(ctx context.Context) error {
	logger := logger.Logger()
	defer logger.Sync()

	logger.Info(fmt.Sprintf("ApiListener::[%s] started at %s", s.GetName(), s.address))

	if err := s.router.Start(s.address); err != nil {
		s.router.Logger.Errorf(fmt.Sprintf("ApiListener::[%s] got error: %s", s.GetName(), err))
		logger.Error(fmt.Sprintf("ApiListener::[%s] got error: %s", s.GetName(), err))

		return err
	}

	return nil
}

func (s *ApiListener) GetName() string {
	return s.name
}

func (s *ApiListener) GetAddress() string {
	return s.address
}

func (s *ApiListener) GetRouter() *echo.Echo {
	return s.router
}
