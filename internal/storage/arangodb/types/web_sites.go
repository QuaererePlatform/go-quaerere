package types

import (
	"github.com/QuaererePlatform/go-kootenay/internal/data_types/web_sites"
)

type (
	ADBWebSite struct {
		web_sites.WebSite
		ArangoDBInfo
	}
)
