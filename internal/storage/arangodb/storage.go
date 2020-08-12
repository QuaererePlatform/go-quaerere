package arangodb

import (
	"context"
	"fmt"
	"reflect"
	"sync"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/QuaererePlatform/go-quaerere/internal/storage"
)

type (
	Storage struct {
		auth     driver.Authentication
		config   Config
		client   driver.Client
		db       driver.Database
		cList    []string
		cTypeMap map[string]reflect.Type
		cOptions map[string]*driver.CreateCollectionOptions
	}

	Config struct {
		Endpoints []string
		Database  string
		Username  string
		Password  string
		AuthType  driver.AuthenticationType
		Auth      bool
	}
)

const WEB_PAGE_COLLECTION = "WebPages"
const WEB_SITE_COLLECTION = "WebSites"

var logger zerolog.Logger
var loggerOnce sync.Once

func init() {
	loggerOnce.Do(func() {
		logger = log.With().Str("component", "storage").Str("storage_driver", "arangodb").Logger()
	})
}

func NewArangoDBStorage(config Config) (*Storage, error) {
	ll := logger.With().Str("method", "NewArangoDBStorage").Logger()
	store := Storage{
		config: config,
		cList: []string{
			WEB_PAGE_COLLECTION,
			WEB_PAGE_COLLECTION,
		},
		cTypeMap: map[string]reflect.Type{
			WEB_PAGE_COLLECTION: reflect.TypeOf(WebPage{}),
			WEB_SITE_COLLECTION: reflect.TypeOf(WebSite{}),
		},
		cOptions: map[string]*driver.CreateCollectionOptions{
			WEB_PAGE_COLLECTION: nil,
			WEB_SITE_COLLECTION: nil,
		},
	}
	if config.Auth == true {
		ll.Debug().
			Str("auth_type", fmt.Sprintf("%#v", config.AuthType)).
			Msg("using auth")
		switch config.AuthType {
		case driver.AuthenticationTypeBasic:
			store.auth = driver.BasicAuthentication(config.Username, config.Password)
		case driver.AuthenticationTypeJWT:
			store.auth = driver.JWTAuthentication(config.Username, config.Password)
		default:
			err := new(UnknownAuthMethodError)
			return nil, err
		}
	}

	return &store, nil
}

func (s *Storage) GetCollection(ctx context.Context, name string) (storage.CollectionStorage, error) {
	cType, ok := s.cTypeMap[name]
	if !ok {
		err := new(UnknownCollectionError)
		err.coll = name
		return nil, err
	}
	coll, err := s.db.Collection(ctx, name)
	if err != nil {
		return nil, err
	}
	c := Collection{
		db:    s.db,
		name:  name,
		cType: cType,
		c:     coll,
	}

	return &c, nil
}

func (s *Storage) Connect(ctx context.Context) error {
	ll := logger.With().Str("method", "Connect").Logger()
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: s.config.Endpoints,
	})
	if err != nil {
		return err
	}
	cc := driver.ClientConfig{
		Connection: conn,
	}
	if s.config.Auth == true {
		cc.Authentication = s.auth
	}
	s.client, err = driver.NewClient(cc)
	if err != nil {
		return err
	}
	e, err := s.client.DatabaseExists(ctx, s.config.Database)
	if err != nil {
		return err
	}
	if e != true {
		err := new(DatabaseDoesNotExistError)
		err.db = s.config.Database
		return err
	}
	s.db, err = s.client.Database(ctx, s.config.Database)
	ll.Debug().Str("s.db", fmt.Sprintf("%#v", s.db))
	return err
}

func (s *Storage) InitDB() error {
	ll := logger.With().Str("method", "InitDB").Logger()
	ll.Debug().Str("store", fmt.Sprintf("%#v", s)).Msg("store inside InitDB")
	for c := range s.cTypeMap {
		ctx := context.TODO()
		defer ctx.Done()
		e, err := s.db.CollectionExists(ctx, c)
		ll.Info().Str("collection", c).Bool("exists", e).Msg("collection exists")
		if err != nil {
			return err
		}
		if e != true {
			opts, ok := s.cOptions[c]
			if !ok {
				ll.Fatal().Str("collection", c).Msg("Missing create options entry for collection")
			}
			ll.Info().Str("collection", c).Msg("creating collection")
			_, err = s.db.CreateCollection(ctx, c, opts)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
