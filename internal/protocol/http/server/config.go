package server

import "github.com/QuaererePlatform/go-quaerere/internal/validator"

type (
	Config struct {
		AppEnv         *string `mapstructure:"app_env"`
		Bind           string  `mapstructure:"bind"`
		DebugMode      bool    `mapstructure:"debug_mode"`
		Port           *int    `mapstructure:"port"`
		StorageBackend string  `mapstructure:"storage_backend"`
	}
)

func (c Config) IsValid(errors validator.Error) {
	if c.AppEnv == nil {
		errors.Add("app environment not configured")
	}
	if c.Port == nil {
		errors.Add("port not configured")
	}
}
