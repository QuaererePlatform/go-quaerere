package storage

import (
	"context"
)

type (
	StorageItem interface {
		GetData() interface{}
	}

	StorageItems []StorageItem

	StorageMeta interface {
		GetMeta() interface{}
	}

	StorageDriver interface {
		Connect(ctx context.Context) error
		GetCollection(ctx context.Context, name string) (CollectionStorage, error)
		InitDB() error
	}

	CollectionStorage interface {
		CreateItems(ctx context.Context, items StorageItems) (StorageItems, error)
		ReadItems(ctx context.Context, keys []string) (StorageItems, error)
		UpdateItems(ctx context.Context, data map[string]map[string]interface{}) (StorageItems, error)
		DeleteItems(ctx context.Context, keys []string) (StorageItems, error)
		ListItems(ctx context.Context, offset int, limit int) (StorageItems, error)
	}
)
