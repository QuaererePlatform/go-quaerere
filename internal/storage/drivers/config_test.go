package drivers_test

import (
	"testing"

	"github.com/QuaererePlatform/go-quaerere/internal/storage/drivers"
	"github.com/QuaererePlatform/go-quaerere/internal/storage/drivers/arangodb"
	"github.com/QuaererePlatform/go-quaerere/internal/validator"
)

func TestConfig_IsValid(t *testing.T) {
	c := new(drivers.Config)

	// no valid configs given, validation fails
	errs := validator.Validate(c)
	if errs == nil {
		t.Errorf("Expected errs but got none")
	}

	// valid config given (arangodb.Config, the only one so far...)
	c = new(drivers.Config)
	c.ArangoDB = new(arangodb.Config)
	errs = validator.Validate(c)
	if errs != nil {
		t.Errorf("Unexpected error(s) occurred: %s", errs.Error())
	}
}
