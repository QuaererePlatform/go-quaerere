package drivers

import (
	"github.com/QuaererePlatform/go-quaerere/internal/storage/drivers/arangodb"
	"github.com/QuaererePlatform/go-quaerere/internal/validator"
)

type Config struct {
	ArangoDB *arangodb.Config `mapstructure:"arangodb"`
}

func (c *Config) IsValid(errors validator.Error) {
	if c.ArangoDB == nil {
		errors.Add("At least 1 storage backend is required")
	}
}
