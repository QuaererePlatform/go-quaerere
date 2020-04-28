package arangodb

import (
	"github.com/QuaererePlatform/go-quaerere/internal/data_types/web_pages"
)

type (
	ADBWebPage struct {
		web_pages.WebPage
		ArangoDBInfo
	}
)
