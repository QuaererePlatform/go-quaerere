package server

import "github.com/QuaererePlatform/go-quaerere/internal/validator"

type customValidator struct {}

// Validate wraps the Validate command so that it meets Echo's validator
func (cv *customValidator) Validate(i interface{}) error {
	return validator.Validate(i.(validator.Validatable))
}