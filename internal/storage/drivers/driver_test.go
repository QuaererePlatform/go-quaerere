package drivers_test

import (
	"reflect"
	"testing"

	"github.com/QuaererePlatform/go-quaerere/internal/storage/drivers"
	"github.com/QuaererePlatform/go-quaerere/internal/storage/drivers/arangodb"
)

func TestNewDriver(t *testing.T) {
	c := drivers.Config{}

	// invalid config given to NewDriver
	d, err := drivers.NewDriver(&c)
	if d != nil {
		t.Errorf("unexpected driver given: %+v", d)
	}
	if err == nil {
		t.Errorf("expected error missing")
	}

	// arangodb.Config given
	c.ArangoDB = &arangodb.Config{}
	d, err = drivers.NewDriver(&c)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	e := reflect.TypeOf(&arangodb.Storage{})
	g := reflect.TypeOf(d)

	if g != e {
		t.Errorf("Unexpected Driver returned: got %s, expected %s", g, e)
	}
}
