package web_sites

import (
	"net/url"

	"github.com/QuaererePlatform/go-quaerere/internal/data_types/accounting"
)

type (
	WebSite struct {
		InLanguage       string                      `json:"in_language"`
		SourceAccounting accounting.SourceAccounting `json:"source_accounting"`
		URL              url.URL                     `json:"url"`
	}
)
