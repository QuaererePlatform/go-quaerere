package storage

import (
	"github.com/QuaererePlatform/go-quaerere/internal/common"
	"github.com/QuaererePlatform/go-quaerere/internal/common/accounting"
)

type (
	WebSite struct {
		InLanguage       string
		SourceAccounting []accounting.SourceAccounting
		URL              common.StringURL
	}
)

func (ws *WebSite) GetData() interface{} {
	return ws
}
