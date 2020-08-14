package storage

import (
	"context"

	"github.com/arangodb/go-driver"

	"github.com/QuaererePlatform/go-quaerere/internal/storage/arangodb"
)

type (
	Item interface {
		GetData() interface{}
	}

	Items []Item

	Meta interface {
		GetMeta() interface{}
	}

	Driver interface {
		Connect(ctx context.Context) error
		GetCollection(ctx context.Context, name string) (Collection, error)
		InitDB() error
	}

	Collection interface {
		CreateItems(ctx context.Context, items Items) (Items, error)
		ReadItems(ctx context.Context, keys []string) (Items, error)
		UpdateItems(ctx context.Context, data map[string]map[string]interface{}) (Items, error)
		DeleteItems(ctx context.Context, keys []string) (Items, error)
		ListItems(ctx context.Context, offset int, limit int) (Items, error)
	}
)

func MakeBackend(backend string) (Driver, error) {

	switch backend {
	case "arangodb":
		c := new(arangodb.Config)
		c.Endpoints = []string{
			"http://localhost:8529/",
		}
		c.Database = "quaerere"
		c.Username = "quaerere"
		c.Password = "password"
		c.Auth = true
		c.AuthType = driver.AuthenticationTypeBasic
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