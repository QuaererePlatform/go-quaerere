package arangodb

import (
	"context"
	"log"
	"sync"

	adb "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type (
	ArangoDBStorage struct {
		client      adb.Client
		config      Config
		db          adb.Database
		collections []string
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

var conn adb.Connection
var once sync.Once

func NewArangoDBStorage(config Config) *ArangoDBStorage {
	s := new(ArangoDBStorage)
	s.config = config
	s.collections = []string{
		WEB_PAGE_COLLECTION,
		WEB_SITE_COLLECTION,
	}
	return s
}

func (s ArangoDBStorage) Init() error {

	for _, c := range s.collections {
		ctx := context.Background()
		if err := s.createCollection(ctx, c, nil); err != nil {
			return err
		}
	}
	return nil
}

func (s ArangoDBStorage) connect(ctx context.Context) (adb.Database, error) {
	err := *new(error)
	once.Do(func() {
		conn, err = http.NewConnection(http.ConnectionConfig{
			Endpoints: s.config.Endpoints,
		})
	})
	if err != nil {
		// Handle error
		log.Printf("Conn err: %+v", err)
		return nil, err
	}
	log.Printf("connect() conn: %+v", conn)

	cc := adb.ClientConfig{
		Connection: conn,
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
	log.Printf("connect() client: %+v", s.client)

	db, err := s.client.Database(ctx, s.config.Database)
	if err != nil {
		// Handle error
		log.Printf("DB err: %+v", err)
		return nil, err
	}
	log.Printf("connect() db: %+v", s.db)

	return db, nil
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
	db, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("getCollection() after connect s: %+v", s)

	coll, err := db.Collection(ctx, name)
	if err != nil {
		return nil, err
	}

	return coll, nil
}

func (d DocumentMeta) GetMeta() (interface{}, error) {
	return d, nil
}
