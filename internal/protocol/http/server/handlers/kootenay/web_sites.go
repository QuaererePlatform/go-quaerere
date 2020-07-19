package kootenay

import (
	"github.com/labstack/echo/v4"

	"github.com/QuaererePlatform/go-quaerere/internal/common/web_sites"
	"github.com/QuaererePlatform/go-quaerere/internal/protocol/http/server/handlers"
)

type (
	WebSiteHandler struct {
		handlers.Handler
	}
)

func (w WebSiteHandler) Post(c echo.Context) error {
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

func (w WebSiteHandler) Get(c echo.Context) error {
	return nil
}

func (w WebSiteHandler) Update(c echo.Context) error {
	return nil
}
