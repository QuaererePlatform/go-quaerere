package server

import "github.com/QuaererePlatform/go-quaerere/internal/validator"

type (
	Config struct {
		Bind           string  `mapstructure:"bind"`
		Port           *int    `mapstructure:"port"`
		TLSDisabled    bool    `mapstructure:"tls_disabled"`
	}
)

func (c *Config) IsValid(errors validator.Error) {
	if c.Port == nil {
		errors.Add("port not configured")
	}
}
