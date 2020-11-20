package common_test

import (
	"testing"

	"github.com/QuaererePlatform/go-quaerere/internal/common"
)

func TestStringURL_MarshalJSON(t *testing.T) {
	expected := "\"https://example.com\""
	su := new(common.StringURL)
	su.Scheme = "https"
	su.Host = "example.com"
	m, _ := su.MarshalJSON()
	if string(m) != expected {
		t.Errorf("got %s, expected %s", string(m), expected)
	}
}
