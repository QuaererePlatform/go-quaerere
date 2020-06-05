package server

import (
	"context"
	"fmt"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/QuaererePlatform/go-quaerere/internal/server/handlers/kootenay"
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
		echo    *echo.Echo
		config  *Config
		storage storage.StorageDriver
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
		echo:    e,
		config:  c,
	}

	if err := s.setupStorage(); err != nil {
		return nil, err
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
	wp := new(kootenay.WebPageHandler)
	/*ws := &kootenay.WebSiteHandler{
		*h,
	}*/

	//s.echo.GET("/", h.Home).Name = "home"

	//s.echo.GET("/api/v1/web-page/:id", wp.Get).Name = "web-page-get"
	s.echo.POST("/api/v1/web-page/", wp.Post(s.storage)).Name = "web-page-post"

	//s.echo.GET("/api/v1/web-site/:id", ws.Get).Name = "web-site-get"
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

	if err := s.storage.Init(); err != nil {
		return err
	}
	return nil
}
