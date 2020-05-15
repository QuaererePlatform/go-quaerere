package web_sites

import (
	"github.com/QuaererePlatform/go-quaerere/internal/common/accounting"
	"github.com/QuaererePlatform/go-quaerere/internal/common"
)

type (
	WebSite struct {
		InLanguage       string                      `json:"in_language"`
		SourceAccounting accounting.SourceAccounting `json:"source_accounting"`
		URL              common.StringURL            `json:"url,string"`
	}
)
