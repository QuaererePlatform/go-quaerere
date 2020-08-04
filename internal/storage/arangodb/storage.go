package arangodb

import (
	"context"
	"fmt"
	"log"
	"sync"

	adb "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"

	"github.com/QuaererePlatform/go-quaerere/internal/common/web_pages"
	"github.com/QuaererePlatform/go-quaerere/internal/common/web_sites"
	"github.com/QuaererePlatform/go-quaerere/internal/storage"
)

type (
	ArangoDBStorage struct {
		conn         adb.Connection
		client       adb.Client
		clientConfig *adb.ClientConfig
		config       Config
		db           adb.Database
		collMap      map[string]string
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
var store ArangoDBStorage
var storeOnce sync.Once

func init() {
	loggerOnce.Do(func() {
		logger = zlog.With().Str("component", "storage").Str("storage_driver", "arangodb").Logger()
	})
}

func makeStorageItem(collName string) (storage.StorageItem, error) {
	switch collName {
	case WEB_PAGE_COLLECTION:
		return new(web_pages.WebPage), nil
	case WEB_SITE_COLLECTION:
		return new(web_sites.WebSite), nil
	default:
		e := new(UnknownCollectionError)
		e.coll = collName
		return nil, e
	}
}

func NewArangoDBStorage(config Config) *ArangoDBStorage {
	storeOnce.Do(func() {
		store = ArangoDBStorage{
			config: config,
			collMap: map[string]string{
				"*web_pages.WebPage": WEB_PAGE_COLLECTION,
				"*web_sites.WebSite": WEB_SITE_COLLECTION,
			},
		}
	})

	return &store
}

func (s ArangoDBStorage) connect(ctx context.Context) (adb.Database, error) {
	var err error
	ll := logger.With().Str("method", "connect").Logger()

	ll.Debug().
		Strs("endpoints", s.config.Endpoints).
		Msg("creating connection")
	s.conn, err = http.NewConnection(http.ConnectionConfig{
		Endpoints: s.config.Endpoints,
	})
	if err != nil {
		return nil, err
	}
	ll.Debug().
		Str("connection", fmt.Sprintf("%#v", s.conn)).
		Msg("connection info")

	cc := adb.ClientConfig{
		Connection: s.conn,
	}
	if s.config.Auth == true {
		ll.Debug().
			Str("auth_type", fmt.Sprintf("%#v", s.config.AuthType)).
			Msg("using auth")
		switch s.config.AuthType {
		case adb.AuthenticationTypeBasic:
			cc.Authentication = adb.BasicAuthentication(s.config.Username, s.config.Password)
		case adb.AuthenticationTypeJWT:
			cc.Authentication = adb.JWTAuthentication(s.config.Username, s.config.Password)
		default:
			err := new(UnknownAuthMethodError)
			return nil, err
		}
	}
	ll.Debug().
		Str("clientConfig", fmt.Sprintf("%#v", s.clientConfig)).
		Msg("config info")

	s.client, err = adb.NewClient(cc)
	if err != nil {
		return nil, err
	}
	ll.Debug().
		Str("client", fmt.Sprintf("%#v", s.client)).
		Msg("client info")

	e, err := s.client.DatabaseExists(ctx, s.config.Database)
	if err != nil {
		return nil, err
	}
	if e != true {
		err := new(DatabaseDoesNotExistError)
		err.db = s.config.Database
		return nil, err
	}
	ll.Info().
		Str("database_exists", fmt.Sprintf("%#v", e)).
		Msg("connected to server")
	s.db, err = s.client.Database(ctx, s.config.Database)
	if err != nil {
		return nil, err
	}
	ll.Info().
		Str("database", fmt.Sprintf("%#v", s.db)).
		Msg("connected to database")

	return s.db, nil
}

func (s ArangoDBStorage) createCollection(name string, options *adb.CreateCollectionOptions) error {
	ll := logger.With().Str("method", "createCollection").Logger()
	cc, err := s.getClientConfig()
	if err != nil {
		return err
	}
	ll.Info().Str("cc", fmt.Sprintf("%#v", cc)).Msg("")
	c, err := s.getClient()
	if err != nil {
		return err
	}
	ll.Info().Str("c", fmt.Sprintf("%#v", c)).Msg("")
	ctx := context.TODO()
	defer ctx.Done()
	db, err := s.getDatabase(ctx)
	if err != nil {
		return err
	}

	e, err := db.CollectionExists(ctx, name)
	if err != nil {
		return err
	}
	var c_err error
	if e != true {
		_, c_err = db.CreateCollection(ctx, name, options)
	}

	return c_err
}

func (s ArangoDBStorage) getCollection(ctx context.Context, name string) (adb.Collection, error) {
	ll := logger.With().Str("method", "getCollection").Logger()

	ll.Debug().
		Str("database", fmt.Sprintf("%#v", s.db)).
		Msg("before getDatabase")
	var err error
	if s.db == nil {
		s.db, err = s.getDatabase(ctx)
		if err != nil {
			return nil, err
		}
	}

	ll.Debug().
		Str("database", fmt.Sprintf("%#v", s.db)).
		Msg("after getDatabase")

	coll, err := s.db.Collection(ctx, name)
	if err != nil {
		return nil, err
	}
	ll.Debug().
		Str("collection", fmt.Sprintf("%#v", coll)).
		Msg("retrieved collection")

	return coll, nil
}

func (s ArangoDBStorage) getDatabase(ctx context.Context) (adb.Database, error) {
	ll := logger.With().Str("method", "getDatabase").Logger()

	ll.Debug().
		Str("s", fmt.Sprintf("%#v", s)).
		Msg("ArangoDBStorage object")
	ll.Debug().
		Str("s.db", fmt.Sprintf("%#v", s.db)).
		Msg("stored database")
	ll.Debug().
		Str("s.client", fmt.Sprintf("%#v", s.client)).
		Msg("stored client")
	if s.db == nil {
		e, err := s.client.DatabaseExists(ctx, s.config.Database)
		if err != nil {
			return nil, err
		}
		if e != true {
			err := new(DatabaseDoesNotExistError)
			err.db = s.config.Database
			return nil, err
		}
		db, err := s.client.Database(ctx, s.config.Database)
		ll.Debug().
			Str("database", fmt.Sprintf("%#v", db)).
			Msg("retrieved database")
		if err != nil {
			return nil, err
		}
		s.db = db
	}
	ll.Debug().
		Str("s.db", fmt.Sprintf("%#v", s.db)).
		Msg("stored database")

	return s.db, nil
}

func (s ArangoDBStorage) getClient() (adb.Client, error) {
	ll := logger.With().Str("method", "getClient").Logger()
	if s.client == nil {
		cc, err := s.getClientConfig()
		if err != nil {
			return nil, err
		}
		client, err := adb.NewClient(*cc)
		if err != nil {
			return nil, err
		}
		ll.Debug().
			Str("client", fmt.Sprintf("%#v", client)).
			Msg("new client")
		s.client = client
	}
	ll.Debug().
		Str("s.client", fmt.Sprintf("%#v", s.client)).
		Msg("stored client")

	return s.client, nil
}

func (s ArangoDBStorage) getClientConfig() (*adb.ClientConfig, error) {
	ll := logger.With().Str("method", "getClientConfig").Logger()
	if s.clientConfig == nil {
		cc := adb.ClientConfig{
			Connection: s.conn,
		}
		if s.config.Auth == true {
			ll.Debug().
				Str("auth_type", fmt.Sprintf("%#v", s.config.AuthType)).
				Msg("using auth")
			switch s.config.AuthType {
			case adb.AuthenticationTypeBasic:
				cc.Authentication = adb.BasicAuthentication(s.config.Username, s.config.Password)
			case adb.AuthenticationTypeJWT:
				cc.Authentication = adb.JWTAuthentication(s.config.Username, s.config.Password)
			default:
				err := new(UnknownAuthMethodError)
				return nil, err
			}
		}
		ll.Debug().
			Str("cc", fmt.Sprintf("%#v", cc)).
			Msg("new client config")

		s.clientConfig = &cc
	}
	ll.Debug().
		Str("clientConfig", fmt.Sprintf("%#v", s.clientConfig)).
		Msg("stored client config")

	return s.clientConfig, nil
}

func (s ArangoDBStorage) InitDB() error {
	ll := logger.With().Str("method", "InitDB").Logger()
	for _, c := range s.collMap {
		ll.Info().Str("collection", c).Msg("calling createCollection")
		if err := s.createCollection(c, nil); err != nil {
			return err
		}
	}
	return nil
}

func (s ArangoDBStorage) Create(item storage.StorageItem) (storage.StorageMeta, error) {
	log.Printf("s.Create() T(i): %T", item)
	log.Printf("arangodb.CreateWebPage() before getCollection s: %+v", s)
	ctx := context.TODO()
	itemType := fmt.Sprintf("%T", item)
	collName, ok := s.collMap[itemType]
	if !ok {
		err := new(UnknownCollectionError)
		err.coll = itemType
		return nil, err
	}
	coll, err := s.getCollection(ctx, collName)
	log.Printf("arangodb.CreateWebPage() coll: %+v", coll)
	log.Printf("arangodb.CreateWebPage() after getCollections s: %+v", s)
	if err != nil {
		return nil, err
	}

	adbMeta, err := coll.CreateDocument(ctx, item.GetData())
	if err != nil {
		return nil, err
	}

	var meta DocumentMeta

	err = copier.Copy(&meta, adbMeta)
	if err != nil {
		return nil, err
	}

	return &meta, nil

}

func (s ArangoDBStorage) Read(key string, itemType string) (storage.StorageItem, storage.StorageMeta, error) {
	ctx := context.Background()
	collName, ok := s.collMap[itemType]
	if !ok {
		err := new(UnknownCollectionError)
		err.coll = itemType
		return nil, nil, err
	}
	coll, err := s.getCollection(ctx, collName)
	if err != nil {
		return nil, nil, err
	}
	item, err := makeStorageItem(collName)
	if err != nil {
		return nil, nil, err
	}
	adbMeta, err := coll.ReadDocument(ctx, key, item)
	log.Printf("Meta: %+v", adbMeta)
	if err != nil {
		return nil, nil, err
	}

	var meta DocumentMeta

	err = copier.Copy(&meta, adbMeta)
	if err != nil {
		return nil, nil, err
	}

	return item, meta, nil
}

func (s ArangoDBStorage) Update(key string, data map[string]interface{}, itemType string) (storage.StorageMeta, error) {
	ctx := context.Background()
	collName, ok := s.collMap[itemType]
	if !ok {
		err := new(UnknownCollectionError)
		err.coll = itemType
		return nil, err
	}
	coll, _ := s.getCollection(ctx, collName)
	adbMeta, err := coll.UpdateDocument(ctx, key, data)
	if err != nil {
		return nil, err
	}

	var meta DocumentMeta

	err = copier.Copy(&meta, adbMeta)
	if err != nil {
		return nil, err
	}

	return &meta, nil
}

func (s ArangoDBStorage) Delete(key string, itemType string) (storage.StorageMeta, error) {
	ctx := context.Background()
	collName, ok := s.collMap[itemType]
	if !ok {
		err := new(UnknownCollectionError)
		err.coll = itemType
		return nil, err
	}
	coll, _ := s.getCollection(ctx, collName)
	adbMeta, err := coll.RemoveDocument(ctx, key)
	if err != nil {
		return nil, err
	}

	var meta DocumentMeta

	err = copier.Copy(&meta, adbMeta)
	if err != nil {
		return nil, err
	}

	return &meta, nil
}

func (s ArangoDBStorage) List(itemType string, offset int, limit int) (storage.StorageItems, error) {
	ctx := context.Background()
	collName, ok := s.collMap[itemType]
	if !ok {
		err := new(UnknownCollectionError)
		err.coll = itemType
		return nil, err
	}
	// TODO: Protobuf is only returning keys, need to reconcile returning data here
	q := fmt.Sprintf("FOR i IN %s SORT i._key LIMIT %i, %i RETURN i", collName, offset, limit)
	cur, err := s.db.Query(ctx, q, nil)
	if err != nil {
		return nil, err
	}
	zlog.Debug().Fields(map[string]interface{}{"query": q, "num_results": cur.Count()})
	items := make(storage.StorageItems, 0)
	for {
		item, err := makeStorageItem(collName)
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
	}
	return items, nil
}

func (d DocumentMeta) GetMeta() interface{} {
	return d
}
