package arangodb

import (
	"context"
	"fmt"
	"log"
	"sync"

	adb "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/jinzhu/copier"
	zlog "github.com/rs/zerolog/log"

	"github.com/QuaererePlatform/go-quaerere/internal/common/web_pages"
	"github.com/QuaererePlatform/go-quaerere/internal/storage"
)

type (
	ArangoDBStorage struct {
		conn      adb.Connection
		client    adb.Client
		config    Config
		db        adb.Database
		collNames []string
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

var store ArangoDBStorage
var once sync.Once

func NewArangoDBStorage(config Config) *ArangoDBStorage {
	once.Do(func() {
		store = ArangoDBStorage{
			config: config,
			collNames: []string{
				WEB_PAGE_COLLECTION,
				WEB_SITE_COLLECTION,
			},
		}
	})

	return &store
}

func (s ArangoDBStorage) Init() error {

	for _, c := range s.collNames {
		ctx := context.Background()
		if err := s.createCollection(ctx, c, nil); err != nil {
			return err
		}
	}
	return nil
}

func (s ArangoDBStorage) connect(ctx context.Context) (adb.Database, error) {
	var err error

	s.conn, err = http.NewConnection(http.ConnectionConfig{
		Endpoints: s.config.Endpoints,
	})

	if err != nil {
		// Handle error
		log.Printf("Conn err: %+v", err)
		return nil, err
	}
	log.Printf("connect() s.conn: %+v", s.conn)
	zlog.Debug()

	cc := adb.ClientConfig{
		Connection: s.conn,
	}
	if s.config.Auth == true {
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
	log.Printf("connect() cc: %+v", cc)

	s.client, err = adb.NewClient(cc)
	if err != nil {
		// Handle error
		log.Printf("Client err: %+v", err)
		return nil, err
	}
	log.Printf("connect() s.client: %#v", s.client)

	s.db, err = s.client.Database(ctx, s.config.Database)
	if err != nil {
		// Handle error
		log.Printf("DB err: %+v", err)
		return nil, err
	}
	log.Printf("connect() s.db: %#v", s.db)

	return s.db, nil
}

func (s ArangoDBStorage) createCollection(ctx context.Context, name string, options *adb.CreateCollectionOptions) error {
	db, err := s.connect(ctx)
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
	log.Printf("getCollection() before connect s: %+v", s)
	var db adb.Database
	if s.db == nil {
		var err error
		db, err = s.connect(ctx)
		if err != nil {
			return nil, err
		}
	}

	log.Printf("getCollection() after connect s: %+v", s)

	coll, err := db.Collection(ctx, name)
	if err != nil {
		return nil, err
	}

	return coll, nil
}

func collectionName(i storage.StorageItem) (string, error) {
	switch i.(type) {
	case *web_pages.WebPage:
		log.Printf("s.Create() i: %#v", i)
		return WEB_PAGE_COLLECTION, nil
	default:
		err := new(UnknownCollectionError)
		err.coll = fmt.Sprintf("%T", i)
		return "", err
	}
}

func (s ArangoDBStorage) Create(i storage.StorageItem) (storage.StorageMeta, error) {
	log.Printf("s.Create() T(i): %T", i)
	log.Printf("arangodb.CreateWebPage() before getCollection s: %+v", s)
	ctx := context.TODO()
	log.Printf("arangodb.CreateWebPage() ctx fresh: %+v", ctx)
	collName, err := collectionName(i)
	if err != nil {
		return nil, err
	}
	coll, err := s.getCollection(ctx, collName)
	log.Printf("arangodb.CreateWebPage() ctx after getCollections: %+v", ctx)
	log.Printf("arangodb.CreateWebPage() coll: %+v", coll)
	log.Printf("arangodb.CreateWebPage() after getCollections s: %+v", s)
	if err != nil {
		return nil, err
	}

	adbMeta, err := coll.CreateDocument(ctx, i.GetData())
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

func (s ArangoDBStorage) Read(key string) (storage.StorageItem, error) {
	return nil, nil
}

func (s ArangoDBStorage) Update(string, map[string]interface{}) (storage.StorageMeta, error) {
	return nil, nil
}

func (s ArangoDBStorage) Delete(string) (storage.StorageMeta, error) {
	return nil, nil
}

func (d DocumentMeta) GetMeta() interface{} {
	return d
}
