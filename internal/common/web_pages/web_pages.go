package web_pages

import (
	"github.com/QuaererePlatform/go-quaerere/internal/common"
	"github.com/QuaererePlatform/go-quaerere/internal/common/accounting"
)

type (
	WebPage struct {
		SourceAccounting []accounting.SourceAccounting `json:"source_accounting"`
		Text             string                        `json:"text"`
		URL              common.StringURL              `json:"url,string"`
		WebSiteKey       string                        `json:"web_site_key"`
	}
)

func (wp *WebPage) GetData() interface{} {
	return wp
}
