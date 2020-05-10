package storage

import (
	"log"

	"github.com/arangodb/go-driver"

	"github.com/QuaererePlatform/go-quaerere/internal/data_types/web_pages"
	"github.com/QuaererePlatform/go-quaerere/internal/data_types/web_sites"
	"github.com/QuaererePlatform/go-quaerere/internal/storage/arangodb"
)

type (
	Storage struct {
		driver StorageDriver
	}

	StorageDriver interface {
		CreateWebPage(wp *web_pages.WebPage) (*driver.DocumentMeta, error)
		ReadWebPage(key string) (*web_pages.WebPage, error)
		UpdateWebPage(key string, data map[string]interface{}) (*driver.DocumentMeta, error)
		DeleteWebPage(key string) (*driver.DocumentMeta, error)
		CreateWebSite(wp *web_sites.WebSite) (*driver.DocumentMeta, error)
		ReadWebSite(key string) (*web_sites.WebSite, error)
		UpdateWebSite(key string, data map[string]interface{}) (*driver.DocumentMeta, error)
		DeleteWebSite(key string) (*driver.DocumentMeta, error)
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

func (s Storage) CreateWebPage(wp *web_pages.WebPage) error {
	meta, err := s.driver.CreateWebPage(wp)
	if err != nil {
		return err
	}
	log.Printf("%+v", meta)
	return nil
}
