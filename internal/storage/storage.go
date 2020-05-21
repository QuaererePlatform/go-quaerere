package storage

import (
	"log"

	"github.com/arangodb/go-driver"

	"github.com/QuaererePlatform/go-quaerere/internal/common/web_pages"
	"github.com/QuaererePlatform/go-quaerere/internal/common/web_sites"
	"github.com/QuaererePlatform/go-quaerere/internal/storage/arangodb"
)

type (
	Storage struct {
		driver    StorageDriver
		itemStore []ItemStore
	}

	StorageItem interface{
		GetData() (interface{}, error)
	}

	StorageMeta interface{
		GetMeta() (interface{}, error)
	}

	ItemStore interface {
		Create(*StorageItem) (*StorageMeta, error)
		Read(string) (*StorageItem, error)
		Update(string, map[string]interface{}) (*StorageMeta, error)
		Delete(string) (*StorageMeta, error)

		Init() error
	}

	StorageDriver interface {
		CreateWebPage(*web_pages.WebPage) (*driver.DocumentMeta, error)
		ReadWebPage(string) (*web_pages.WebPage, error)
		UpdateWebPage(string, map[string]interface{}) (*driver.DocumentMeta, error)
		DeleteWebPage(string) (*driver.DocumentMeta, error)
		CreateWebSite(*web_sites.WebSite) (*driver.DocumentMeta, error)
		ReadWebSite(string) (*web_sites.WebSite, error)
		UpdateWebSite(string, map[string]interface{}) (*driver.DocumentMeta, error)
		DeleteWebSite(string) (*driver.DocumentMeta, error)
	}
)

func NewStorage(backend string) *Storage {

	s := new(Storage)
	switch backend {
	case "arangodb":
		c := new(arangodb.Config)
		c.Endpoints = []string{
			"http://arangodb:8529/",
		}
		c.Database = "quaerere"
		c.Username = "quaerere"
		c.Password = "password"
		c.Auth = true
		c.AuthType = driver.AuthenticationTypeBasic
		s.driver = arangodb.NewArangoDBStorage(*c)
	}

	return s
}

func (s Storage) Init() error {
	return s.driver.Init()
}

func (s Storage) CreateWebPage(wp *web_pages.WebPage) error {
	meta, err := s.driver.CreateWebPage(wp)
	if err != nil {
		return err
	}
	log.Printf("%+v", meta)
	return nil
}
