package storage

import (
	"github.com/arangodb/go-driver"

	"github.com/QuaererePlatform/go-quaerere/internal/common/web_pages"
	"github.com/QuaererePlatform/go-quaerere/internal/common/web_sites"
)

type (
	StorageItem interface {
		GetData() interface{}
	}

	StorageMeta interface {
		GetMeta() interface{}
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

		Create(StorageItem) (StorageMeta, error)

		Init() error
	}
)
