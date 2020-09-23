package arangodb

import (
	"fmt"
)

type (
	DatabaseDoesNotExistError struct {
		db string
	}
	UnknownAuthMethodError struct{}
	UnknownCollectionError struct {
		coll string
	}
)

func (e DatabaseDoesNotExistError) Error() string {
	return fmt.Sprintf("requested database does not exist: %s", e.db)
}

func (e UnknownAuthMethodError) Error() string {
	return "unknown auth method"
}

func (e UnknownCollectionError) Error() string {
	return fmt.Sprintf("unknown collection: %s", e.coll)
}
