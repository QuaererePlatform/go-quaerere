package types

import (
	"github.com/QuaererePlatform/go-kootenay/internal/data_types/web_pages"
)

type (
	ADBWebPage struct {
		web_pages.WebPage
		ArangoDBInfo
	}
)
