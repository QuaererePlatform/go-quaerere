package storage

import (
	"time"
)

type (
	SourceAccounting struct {
		DataOrigin       string
		DatetimeAcquired time.Time
	}
)
