package sqlite3

import (
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"

	"github.com/QuaererePlatform/go-quaerere/internal/storage"
)

var db SQLStorage
var once sync.Once

type (
	Config struct {
		Path string
	}

	SQLStorage struct {
		conn *sql.DB
	}
)

func NewSQLStorage(config Config) (*SQLStorage, error) {
	var err error
	once.Do(func() {
		db.conn, err = sql.Open("sqlite3", config.Path)
	})

	return &db, nil
}

func (s SQLStorage) Create(i storage.Item) (storage.Meta, error) {
	return nil, nil
}

func (s SQLStorage) Read(key string) (storage.Item, error) {
	return nil, nil
}

func (s SQLStorage) Update(string, map[string]interface{}) (storage.Meta, error) {
	return nil, nil
}

func (s SQLStorage) Delete(string) (storage.Meta, error) {
	return nil, nil
}
