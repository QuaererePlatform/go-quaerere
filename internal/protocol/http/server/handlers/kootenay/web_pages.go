package kootenay

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/QuaererePlatform/go-quaerere/internal/common"
	"github.com/QuaererePlatform/go-quaerere/internal/protocol/http/server/handlers"
	"github.com/QuaererePlatform/go-quaerere/internal/storage"
)

type (
	SourceAccounting struct {
		DataOrigin       string    `json:"data_origin"`
		DatetimeAcquired time.Time `json:"datetime_acquired"`
	}

	WebPageHandler struct {
		handlers.Handler
	}

	WebPageData struct {
		Text string           `json:"text"`
		URL  common.StringURL `json:"url,string"`
	}

	WebPageMeta struct {
		SourceAccounting []SourceAccounting `json:"sourceAccounting"`
	}

	WebPageRequest struct {
		SourceAccounting []SourceAccounting `json:"sourceAccounting"`
		Text             string             `json:"text"`
		URL              common.StringURL   `json:"url,string"`
		WebSiteKey       string             `json:"webSiteKey"`
	}

	WebPageResponse struct {
		Data             WebPageData        `json:"data"`
		Meta             WebPageMeta        `json:"meta"`
		SourceAccounting []SourceAccounting `json:"sourceAccounting"`
		Text             string             `json:"text"`
		URL              common.StringURL   `json:"url,string"`
		WebSiteKey       string             `json:"webSiteKey"`
	}
)

func (w *WebPageHandler) Post(s storage.StorageDriver) echo.HandlerFunc {
	return func(c echo.Context) error {
		/*		log.Printf("WebPageHandler.Post() s: %+v", s)
				wp := new(WebPageRequest)
				if err := c.Bind(wp); err != nil {
					return c.String(http.StatusInternalServerError, err.Error())
				}
				meta, err := s.Create(wp)
				if err != nil {
					return c.String(http.StatusInternalServerError, err.Error())
				}
				log.Printf("WebPageHandler.Post() meta: %#v", meta)
				return c.JSON(http.StatusCreated, meta.GetMeta())*/
		return nil
	}
}

func (w *WebPageHandler) Delete(s storage.StorageDriver) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (w *WebPageHandler) List(s storage.StorageDriver) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (w *WebPageHandler) Get(s storage.StorageDriver) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (w *WebPageHandler) Put(s storage.StorageDriver) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
