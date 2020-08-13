package arangodb

import (
	"context"
	"reflect"

	"github.com/arangodb/go-driver"

	"github.com/QuaererePlatform/go-quaerere/internal/storage"
)

type Collection struct {
	c       driver.Collection
	db      driver.Database
	name    string
	cType   reflect.Type
}

func (c *Collection) CreateItems(ctx context.Context, items storage.Items) (storage.Items, error) {
	/*	db, err := c.store.connect(ctx)
		if err != nil {
			return nil, err
		}

		coll, err := db.Collection(ctx, c.name)
		if err != nil {
			return nil, err
		}

		adbMetas, errs, err := coll.CreateDocuments(ctx, items.GetData())
		if err != nil {
			return nil, err
		}

		var meta DocumentMeta

		err = copier.Copy(&meta, adbMetas)
		if err != nil {
			return nil, err
		}*/
	return nil, nil
}

func (c *Collection) ReadItems(ctx context.Context, keys []string) (storage.Items, error) {
	/*	db, err := c.store.connect(ctx)
		if err != nil {
			return nil, err
		}

		coll, err := db.Collection(ctx, c.name)
		if err != nil {
			return nil, err
		}
		item, err := makeStorageItem(c.name)
		if err != nil {
			return nil, err
		}
		adbMeta, errs, err := coll.ReadDocuments(ctx, keys, item)
		log.Printf("Meta: %+v", adbMeta)
		if err != nil {
			return nil, err
		}*/

	return nil, nil
}

func (c *Collection) UpdateItems(ctx context.Context, data map[string]map[string]interface{}) (storage.Items, error) {
	/*db, err := c.store.connect(ctx)
	if err != nil {
		return nil, err
	}

	coll, err := db.Collection(ctx, c.name)
	if err != nil {
		return nil, err
	}

	adbMeta, errs, err := coll.UpdateDocuments(ctx, keys, data)
	if err != nil {
		return nil, err
	}

	var meta DocumentMeta

	err = copier.Copy(&meta, adbMeta)
	if err != nil {
		return nil, err
	}*/
	return nil, nil
}

func (c *Collection) DeleteItems(ctx context.Context, keys []string) (storage.Items, error) {
	/*db, err := c.store.connect(ctx)
	if err != nil {
		return nil, err
	}

	coll, err := db.Collection(ctx, c.name)
	if err != nil {
		return nil, err
	}
	adbMeta, errs, err := coll.RemoveDocuments(ctx, keys)
	if err != nil {
		return nil, err
	}

	var meta DocumentMeta

	err = copier.Copy(&meta, adbMeta)
	if err != nil {
		return nil, err
	}*/
	return nil, nil
}

func (c *Collection) ListItems(ctx context.Context, offset int, limit int) (storage.Items, error) {
	/*db, err := c.store.connect(ctx)
	if err != nil {
		return nil, err
	}
	// TODO: Protobuf is only returning keys, need to reconcile returning data here
	q := fmt.Sprintf("FOR i IN %s SORT i._key LIMIT %d, %d RETURN i", c.name, offset, limit)
	cur, err := db.Query(ctx, q, nil)
	if err != nil {
		return nil, err
	}
	zlog.Debug().Fields(map[string]interface{}{"query": q, "num_results": cur.Count()})
	items := make(storage.StorageItems, 0)
	for {
		item, err := makeStorageItem(c.name)
		if err != nil {
			return nil, err
		}
		// TODO: Integrate DB metadata into Item
		_, err = cur.ReadDocument(ctx, &item)
		zlog.Debug().Fields(map[string]interface{}{"cursor_stats": cur.Statistics()})
		items = append(items, item)
		if adb.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return items, err
		}
	}*/
	return nil, nil
}
