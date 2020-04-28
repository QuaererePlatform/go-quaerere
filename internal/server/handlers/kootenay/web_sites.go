package kootenay

import (
	"github.com/labstack/echo/v4"

	"github.com/QuaererePlatform/go-quaerere/internal/data_types/web_sites"
	"github.com/QuaererePlatform/go-quaerere/internal/server/handlers"
)

type (
	WebSiteHandler struct {
		handlers.Handler
		dataType struct{}
	}
)

func (w WebSiteHandler) Create(c echo.Context) error {
	ws := new(web_sites.WebSite)
	if err := c.Bind(ws); err != nil {
		return err
	}
	return nil
}

func (w WebSiteHandler) Delete(c echo.Context) error {
	return nil
}

func (w WebSiteHandler) List(c echo.Context) error {
	return nil
}

func (w WebSiteHandler) Read(c echo.Context) error {
	return nil
}

func (w WebSiteHandler) Update(c echo.Context) error {
	return nil
}
