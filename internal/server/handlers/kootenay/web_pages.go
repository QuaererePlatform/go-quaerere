package kootenay

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/QuaererePlatform/go-quaerere/internal/common/web_pages"
	"github.com/QuaererePlatform/go-quaerere/internal/server/handlers"
	"github.com/QuaererePlatform/go-quaerere/internal/storage"
)

type (
	WebPageHandler struct {
		handlers.Handler
	}
)

/*func (w WebPageHandler) Post(c echo.Context) error {
	wp := new(web_pages.WebPage)
	if err := c.Bind(wp); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if err := w.Storage.CreateWebPage(wp); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}*/

func (w WebPageHandler) Post(s *storage.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Printf("WebPageHandler.Post() s: %+v", s)
		wp := new(web_pages.WebPage)
		if err := c.Bind(wp); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		if err := s.CreateWebPage(wp); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.NoContent(http.StatusCreated)
	}
}

func (w WebPageHandler) Delete(c echo.Context) error {
	return nil
}

func (w WebPageHandler) List(c echo.Context) error {
	return nil
}

func (w WebPageHandler) Get(c echo.Context) error {
	return nil
}

func (w WebPageHandler) Put(c echo.Context) error {
	return nil
}
