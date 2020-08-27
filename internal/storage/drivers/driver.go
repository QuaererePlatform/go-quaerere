package drivers

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/QuaererePlatform/go-quaerere/internal/storage"
	"github.com/QuaererePlatform/go-quaerere/internal/storage/drivers/arangodb"
)

type Driver interface {
	Connect(ctx context.Context) error
	GetCollection(ctx context.Context, name string) (storage.Collection, error)
	InitDB() error
}

func NewDriver(backend string) (Driver, error) {

	switch backend {
	case "arangodb":
		c := new(arangodb.Config)
		if err := viper.Unmarshal(c); err != nil {
			log.Fatal().Err(err)
		}
		log.Debug().Str("config", fmt.Sprintf("%#v", c)).Msg("ADB config")
		store, err := arangodb.NewArangoDBStorage(*c)
		if err != nil {
			return nil, err
		}
		return store, nil
	default:
		err := new(UnknownStorageBackend)
		err.backend = backend
		return nil, err
	}
}