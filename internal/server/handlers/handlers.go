package handlers

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type (
	Handler struct {

	}
)

func (h Handler) WrapError(err error, s string, c echo.Context) error {
	if s != "" {
		return errors.New(s + ": " + err.Error())
	}
	return err
}
