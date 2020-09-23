package storage

import (
	"context"
)

type (
	Item interface {
		GetData() interface{}
	}

	Items []Item

	Meta interface {
		GetMeta() interface{}
	}

	Collection interface {
		CreateItems(ctx context.Context, items Items) (Items, error)
		ReadItems(ctx context.Context, keys []string) (Items, error)
		UpdateItems(ctx context.Context, data map[string]map[string]interface{}) (Items, error)
		DeleteItems(ctx context.Context, keys []string) (Items, error)
		ListItems(ctx context.Context, offset int, limit int) (Items, error)
	}
)
