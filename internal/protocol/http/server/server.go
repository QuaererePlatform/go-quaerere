package server

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	"github.com/QuaererePlatform/go-quaerere/internal/config"
	"github.com/QuaererePlatform/go-quaerere/internal/protocol/http/server/handlers/kootenay"
	"github.com/QuaererePlatform/go-quaerere/internal/storage/drivers"
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
		config        *config.HTTPConfig
		httpServer    *echo.Echo
		storage       drivers.Driver
		storageConfig *drivers.Config
	}
)

func New(c *config.AppConfig) (Server, error) {
	e := echo.New()
	e.Debug = c.DebugMode
	e.HideBanner = true
	e.HidePort = true
	e.Server.ReadTimeout = readTimeout
	e.Server.WriteTimeout = writeTimeout
	e.Validator = &customValidator{}

	if err := e.Validator.Validate(c); err != nil {
		return nil, err
	}

	s := &server{
		config:        c.Serve.HTTP,
		httpServer:    e,
		storageConfig: c.Datastore,
	}

	if err := s.setupStorage(); err != nil {
		return nil, err
	}
	s.setupRoutes()
	s.setupMiddleware()

	return s, nil
}

func (s *server) Start() error {
	return s.httpServer.Start(fmt.Sprintf("%s:%d", s.config.Bind, *s.config.Port))
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *server) setupMiddleware() {
	s.httpServer.Use(middleware.Gzip())
	s.httpServer.Use(middleware.Logger())
	s.httpServer.Use(middleware.Recover())
	s.httpServer.Use(middleware.Secure())
}

func (s *server) setupRoutes() {
	wp := new(kootenay.WebPageHandler)
	/*ws := &kootenay.WebSiteHandler{
		*h,
	}*/

	// s.httpServer.GET("/", h.Home).Name = "home"

	// s.httpServer.GET("/api/v0/web-page/:id", wp.Get).Name = "web-page-get"
	s.httpServer.POST("/api/v0/web-page/", wp.Post(s.storage)).Name = "web-page-post"

	// s.httpServer.GET("/api/v0/web-site/:id", ws.Get).Name = "web-site-get"
}

func (s *server) setupStorage() error {
	var err error
	s.storage, err = drivers.NewDriver(s.storageConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("error setting up datastore")
	}

	if err := s.storage.InitDB(); err != nil {
		return err
	}
	return nil
}
