package server

import (
	"context"
	"fmt"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/QuaererePlatform/go-quaerere/internal/protocol/http/server/handlers/kootenay"
	"github.com/QuaererePlatform/go-quaerere/internal/storage"
	"github.com/QuaererePlatform/go-quaerere/internal/storage/arangodb"
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
		config     *Config
		httpServer *echo.Echo
		storage    storage.StorageDriver
	}
)

func New(c *Config) (Server, error) {
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
		config:     c,
		httpServer: e,
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
	switch s.config.StorageBackend {
	case "arangodb":
		c := new(arangodb.Config)
		c.Endpoints = []string{
			"http://arangodb:8529/",
		}
		c.Database = "quaerere"
		c.Username = "quaerere"
		c.Password = "password"
		c.Auth = true
		c.AuthType = driver.AuthenticationTypeBasic
		s.storage = arangodb.NewArangoDBStorage(*c)
	}

	if err := s.storage.InitDB(); err != nil {
		return err
	}
	return nil
}
