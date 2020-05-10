package handlers

import (
	"errors"

	"github.com/labstack/echo/v4"

	"github.com/QuaererePlatform/go-quaerere/internal/storage"
)

type (
	Handler struct {
		Storage *storage.Storage
	}
)

func (h Handler) WrapError(err error, s string, c echo.Context) error {
	if s != "" {
		return errors.New(s + ": " + err.Error())
	}
	return err
}
