package config

import (
	"github.com/QuaererePlatform/go-quaerere/internal/storage/drivers"
	"github.com/QuaererePlatform/go-quaerere/internal/validator"
)

type (
	AppConfig struct {
		AppEnv    *string         `mapstructure:"app_env"`
		DebugMode bool            `mapstructure:"debug_mode"`
		Datastore *drivers.Config `mapstructure:"datastore"`
		Serve     *ServeConfig    `mapstructure:"serve"`
	}

	GRPCConfig struct {
		Bind string `mapstructure:"bind"`
		Port *int   `mapstructure:"port"`
	}

	HTTPConfig struct {
		Bind        string `mapstructure:"bind"`
		Port        *int   `mapstructure:"port"`
		TLSDisabled bool   `mapstructure:"tls_disabled"`
	}

	ServeConfig struct {
		GRPC *GRPCConfig `mapstructure:"grpc"`
		HTTP *HTTPConfig `mapstructure:"http"`
	}
)

func (c *AppConfig) IsValid(errors validator.Error) {
	if c.AppEnv == nil {
		errors.Add("app environment not configured")
	}
}

func (c *GRPCConfig) IsValid(errors validator.Error) {
	if c.Port == nil {
		errors.Add("port not configured")
	}
}

func (c *HTTPConfig) IsValid(errors validator.Error) {
	if c.Port == nil {
		errors.Add("port not configured")
	}
}
