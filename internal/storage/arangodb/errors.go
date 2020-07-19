package arangodb

import (
	"fmt"
)

type (
	UnknownAuthMethodError struct {}
	UnknownCollectionError struct {
		coll string
	}
)

func (e UnknownAuthMethodError) Error() string {
	return "unknown auth method"
}

func (e UnknownCollectionError) Error() string {
	return fmt.Sprintf("unknown collection: %q", e.coll)
}