package storage

import (
	"github.com/QuaererePlatform/go-quaerere/internal/common"
)

type (
	WebSite struct {
		InLanguage       string
		SourceAccounting []SourceAccounting
		URL              common.StringURL
	}
)

func (ws *WebSite) GetData() interface{} {
	return ws
}
