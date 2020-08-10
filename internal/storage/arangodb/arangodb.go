package arangodb

import (
	"time"

	"github.com/QuaererePlatform/go-quaerere/internal/common"
)

type (
	SourceAccounting struct {
		DataOrigin       string    `json:"data_origin"`
		DatetimeAcquired time.Time `json:"datetime_acquired"`
	}

	WebPage struct {
		SourceAccounting []SourceAccounting `json:"source_accounting"`
		Text             string             `json:"text"`
		URL              common.StringURL   `json:"url,string"`
		WebSiteKey       string             `json:"web_site_key"`
	}

	WebSite struct {
		InLanguage       string             `json:"in_language"`
		SourceAccounting []SourceAccounting `json:"source_accounting"`
		URL              common.StringURL   `json:"url,string"`
	}
)
