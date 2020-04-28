package server

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/QuaererePlatform/go-quaerere/internal/server/handlers"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 10 * time.Second
)

type (
	Server interface {
		Shutdown(context.Context) error
		Start() error
	}

	server struct {
		echo   *echo.Echo
		config *Config
	}
)

func New(c *Config) (Server, error) {
	e := echo.New()
	e.Debug = c.DebugMode
	e.Server.ReadTimeout = readTimeout
	e.Server.WriteTimeout = writeTimeout
	e.Validator = &customValidator{}

	if err := e.Validator.Validate(c); err != nil {
		return nil, err
	}

	s := &server{
		echo:   e,
		config: c,
	}

	s.setupRoutes()
	s.setupMiddleware()

	return s, nil
}

func (s *server) Start() error {
	return s.echo.Start(fmt.Sprintf("%s:%d", s.config.Bind, *s.config.Port))
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func (s *server) setupMiddleware() {
	s.echo.Use(middleware.Gzip())
	s.echo.Use(middleware.Logger())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.Secure())
}

func (s *server) setupRoutes() {
	h := &handlers.Handler{}
	s.echo.GET("/", h.Home).Name = "home"
}
