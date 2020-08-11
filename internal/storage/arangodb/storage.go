package arangodb

import (
	"context"
	"fmt"
	"sync"

	adb "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"

	"github.com/QuaererePlatform/go-quaerere/internal/storage"
)

type (
	Collection struct {
		store *Storage
		name string
	}

	Storage struct {
		auth     adb.Authentication
		config       Config
		collMap      map[string]string
		db adb.Database
	}

	Config struct {
		Endpoints []string
		Database  string
		Username  string
		Password  string
		AuthType  adb.AuthenticationType
		Auth      bool
	}

	DocumentMeta struct {
		adb.DocumentMeta
	}
)

const WEB_PAGE_COLLECTION = "WebPages"
const WEB_SITE_COLLECTION = "WebSites"

var logger zerolog.Logger
var loggerOnce sync.Once

func init() {
	loggerOnce.Do(func() {
		logger = zlog.With().Str("component", "storage").Str("storage_driver", "arangodb").Logger()
	})
}

func makeStorageItem(collName string) (storage.StorageItem, error) {
	switch collName {
	case WEB_PAGE_COLLECTION:
		return new(storage.WebPage), nil
	case WEB_SITE_COLLECTION:
		return new(storage.WebSite), nil
	default:
		e := new(UnknownCollectionError)
		e.coll = collName
		return nil, e
	}
}

func NewArangoDBStorage(config Config) (*Storage, error) {
	ll := logger.With().Str("method", "NewArangoDBStorage").Logger()
	store := Storage{
			config: config,
			collMap: map[string]string{
				"*web_pages.WebPage": WEB_PAGE_COLLECTION,
				"*web_sites.WebSite": WEB_SITE_COLLECTION,
			},
	}
	if config.Auth == true {
		ll.Debug().
			Str("auth_type", fmt.Sprintf("%#v", config.AuthType)).
			Msg("using auth")
		switch config.AuthType {
		case adb.AuthenticationTypeBasic:
			store.auth = adb.BasicAuthentication(config.Username, config.Password)
		case adb.AuthenticationTypeJWT:
			store.auth = adb.JWTAuthentication(config.Username, config.Password)
		default:
			err := new(UnknownAuthMethodError)
			return nil, err
		}
	}

	return &store, nil
}

func (s *Storage) connect(ctx context.Context) (adb.Database, error) {
	ll := logger.With().Str("method", "createCollection").Logger()
	conn, err := s.getConnection()
	if err != nil {
		return nil, err
	}
	ll.Info().Str("cc", fmt.Sprintf("%#v", conn)).Msg("")

	cc, err := s.getClientConfig(conn)
	if err != nil {
		return nil, err
	}
	ll.Info().Str("cc", fmt.Sprintf("%#v", cc)).Msg("")
	c, err := s.getClient(cc)
	if err != nil {
		return nil, err
	}
	ll.Info().Str("c", fmt.Sprintf("%#v", c)).Msg("")
	db, err := s.getDatabase(ctx, c)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (s *Storage) createCollection(name string, options *adb.CreateCollectionOptions) error {
	ll := logger.With().Str("method", "createCollection").Logger()
	ctx := context.TODO()
	defer ctx.Done()
	//db, err := s.connect(ctx)
	//if err != nil {
	//	return err
	//}
	ll.Debug().Str("store", fmt.Sprintf("%#v", s)).Msg("store inside createCollection")
	e, err := s.db.CollectionExists(ctx, name)
	if err != nil {
		return err
	}
	var c_err error
	if e != true {
		_, c_err = s.db.CreateCollection(ctx, name, options)
	}

	return c_err
}

func (s *Storage) getCollection(ctx context.Context, client adb.Client, name string) (adb.Collection, error) {
	ll := logger.With().Str("method", "getCollection").Logger()

	var err error
	db, err := s.getDatabase(ctx, client)
	if err != nil {
			return nil, err
		}


	ll.Debug().
		Str("database", fmt.Sprintf("%#v", db)).
		Msg("after getDatabase")

	coll, err := db.Collection(ctx, name)
	if err != nil {
		return nil, err
	}
	ll.Debug().
		Str("collection", fmt.Sprintf("%#v", coll)).
		Msg("retrieved collection")

	return coll, nil
}

func (s *Storage) getDatabase(ctx context.Context, c adb.Client) (adb.Database, error) {
	ll := logger.With().Str("method", "getDatabase").Logger()

		e, err := c.DatabaseExists(ctx, s.config.Database)
		if err != nil {
			return nil, err
		}
		if e != true {
			err := new(DatabaseDoesNotExistError)
			err.db = s.config.Database
			return nil, err
		}
		db, err := c.Database(ctx, s.config.Database)
		ll.Debug().
			Str("database", fmt.Sprintf("%#v", db)).
			Msg("retrieved database")
		if err != nil {
			return nil, err
		}

	return db, nil
}

func (s *Storage) getClient(cc *adb.ClientConfig) (adb.Client, error) {
	ll := logger.With().Str("method", "getClient").Logger()

	client, err := adb.NewClient(*cc)
		if err != nil {
			return nil, err
		}
		ll.Debug().
			Str("client", fmt.Sprintf("%#v", client)).
			Msg("new client")

	return client, nil
}

func (s *Storage) getClientConfig(conn adb.Connection) (*adb.ClientConfig, error) {
	ll := logger.With().Str("method", "getClientConfig").Logger()

		cc := adb.ClientConfig{
			Connection: conn,
		}
		if s.config.Auth == true {
			cc.Authentication = s.auth
		}
		ll.Debug().
			Str("cc", fmt.Sprintf("%#v", cc)).
			Msg("new client config")

	return &cc, nil
}

func (s *Storage) getConnection() (adb.Connection, error) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: s.config.Endpoints,
	})
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (s *Storage) NewCollection(name string) (Collection) {
	c := Collection{
		store: s,
		name: name,
	}

	return c
}

func (s *Storage) Connect(ctx context.Context) error {
	ll := logger.With().Str("method", "Connect").Logger()
	var err error
	s.db, err = s.connect(ctx)
	ll.Debug().Str("s.db", fmt.Sprintf("%#v", s.db))
	return err
}

func (s *Storage) InitDB() error {
	ll := logger.With().Str("method", "InitDB").Logger()
	ll.Debug().Str("store", fmt.Sprintf("%#v", s)).Msg("store inside InitDB")
	for _, c := range s.collMap {
		ll.Info().Str("collection", c).Msg("calling createCollection")
		if err := s.createCollection(c, nil); err != nil {
			return err
		}
	}
	return nil
}

func (c *Collection) CreateItems(ctx context.Context, items storage.StorageItems) (storage.StorageItems, error) {
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

func (c *Collection) ReadItems(ctx context.Context, keys []string) (storage.StorageItems, error) {
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

func (c *Collection) UpdateItems(ctx context.Context, data map[string]map[string]interface{}) (storage.StorageItems, error) {
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

func (c *Collection) DeleteItems(ctx context.Context, keys []string) (storage.StorageItems, error) {
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

func (c *Collection) ListItems(ctx context.Context, offset int, limit int) (storage.StorageItems, error) {
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
		// TODO: Integrate DB metadata into StorageItem
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

func (d *DocumentMeta) GetMeta() interface{} {
	return d
}
