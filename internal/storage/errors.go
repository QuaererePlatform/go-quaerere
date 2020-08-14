package storage

import (
	"fmt"
)

type (
	UnknownStorageBackend struct {
		backend string
	}
)

func (e *UnknownStorageBackend) Error() string {
	return fmt.Sprintf("unknown storage backend: %s", e.backend)
}
