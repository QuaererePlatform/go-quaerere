package web_pages

import (
	"net/url"

	"github.com/QuaererePlatform/go-quaerere/internal/data_types/accounting"
)

type (
	WebPage struct {
		SourceAccounting accounting.SourceAccounting `json:"source_accounting"`
		Text             string                      `json:"text"`
		URL              url.URL                     `json:"url"`
		WebSiteKey       string                      `json:"web_site_key"`
	}
)
