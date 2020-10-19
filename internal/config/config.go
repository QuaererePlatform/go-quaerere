package config

import (
	grpc "github.com/QuaererePlatform/go-quaerere/internal/protocol/grpc/server"
	http "github.com/QuaererePlatform/go-quaerere/internal/protocol/http/server"
	"github.com/QuaererePlatform/go-quaerere/internal/storage/drivers"
	"github.com/QuaererePlatform/go-quaerere/internal/validator"
)

type (
	DBInitConfig struct {
		AppEnv         *string        `mapstructure:"app_env"`
		DebugMode      bool           `mapstructure:"debug_mode"`
		StorageBackend string         `mapstructure:"storage_backend"`
		Datastore      drivers.Config `mapstructure:"datastore"`
		Serve          *ServeConfig   `mapstructure:"serve"`
	}

	AppConfig struct {
		AppEnv    *string         `mapstructure:"app_env"`
		DebugMode bool            `mapstructure:"debug_mode"`
		Datastore *drivers.Config `mapstructure:"datastore"`
		Serve     *ServeConfig    `mapstructure:"serve"`
	}

	ServeConfig struct {
		GRPC *grpc.Config `mapstructure:"grpc"`
		HTTP *http.Config `mapstructure:"http"`
	}
)

func (c *AppConfig) IsValid(errors validator.Error) {
	if c.AppEnv == nil {
		errors.Add("app environment not configured")
	}
}
