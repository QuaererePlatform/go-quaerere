package accounting

import (
	"time"
)

type (
	SourceAccounting struct {
		DataOrigin       string    `json:"data_origin"`
		DatetimeAcquired time.Time `json:"datetime_acquired"`
	}
)
