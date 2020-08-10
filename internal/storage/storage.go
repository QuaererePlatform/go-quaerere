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
		Create(StorageItem) (StorageMeta, error)
		Read(string, string) (StorageItem, StorageMeta, error)
		Update(string, map[string]interface{}, string) (StorageMeta, error)
		Delete(string, string) (StorageMeta, error)
		List(itemType string, offset int, limit int) (StorageItems, error)

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
