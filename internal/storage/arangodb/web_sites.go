package arangodb

import (
	"github.com/QuaererePlatform/go-quaerere/internal/data_types/web_sites"
)

type (
	ADBWebSite struct {
		web_sites.WebSite
		ArangoDBInfo
	}
)
