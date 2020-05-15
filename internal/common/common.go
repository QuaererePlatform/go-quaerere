package common

import (
	"encoding/json"
	"net/url"
)

type StringURL struct {
	url.URL
}

func (u *StringURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}
