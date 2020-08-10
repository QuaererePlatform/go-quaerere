package storage

import (
	"github.com/QuaererePlatform/go-quaerere/internal/common"
)

type (
	WebPage struct {
		SourceAccounting []SourceAccounting
		Text             string
		URL              common.StringURL
		WebSiteKey       string
	}
)

func (wp *WebPage) GetData() interface{} {
	return wp
}
