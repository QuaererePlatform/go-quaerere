package web_sites

import (
	"github.com/QuaererePlatform/go-quaerere/internal/common"
	"github.com/QuaererePlatform/go-quaerere/internal/common/accounting"
)

type (
	WebSite struct {
		InLanguage       string                        `json:"in_language"`
		SourceAccounting []accounting.SourceAccounting `json:"source_accounting"`
		URL              common.StringURL              `json:"url,string"`
	}
)

func (ws *WebSite) GetData() interface{} {
	return ws
}
