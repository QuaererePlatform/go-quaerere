package drivers

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/QuaererePlatform/go-quaerere/internal/storage"
	"github.com/QuaererePlatform/go-quaerere/internal/storage/drivers/arangodb"
)

type (
	Driver interface {
		Connect(ctx context.Context) error
		GetCollection(ctx context.Context, name string) (storage.Collection, error)
		InitDB() error
	}
)

func NewDriver(config *Config) (Driver, error) {
	if config.ArangoDB != nil {
		log.Debug().Str("config", fmt.Sprintf("%#v", config.ArangoDB)).Msg("ADB config")
		store, err := arangodb.NewArangoDBStorage(config.ArangoDB)
		if err != nil {
			return nil, err
		}
		return store, nil
	} else {
		err := new(UnknownStorageBackend)
		err.backend = "nil"
		return nil, err
	}
}
