package storage

import (
	"time"

	"github.com/QuaererePlatform/go-quaerere/internal/common"
)

type (
	SourceAccounting struct {
		DataOrigin       string
		DatetimeAcquired time.Time
	}
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
