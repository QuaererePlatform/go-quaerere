package arangodb

import (
	"github.com/arangodb/go-driver"
)

type Config struct {
	Endpoints []string `mapstructure:"datastore.arangodb.endpoints"`
	Database  string `mapstructure:"datastore.arangodb.database"`
	Username  string `mapstructure:"datastore.arangodb.username"`
	Password  string `mapstructure:"datastore.arangodb.password"`
	AuthType  driver.AuthenticationType
	Auth      bool
	ConfigAuthType string `mapstructure:"datastore.arangodb.auth_type"`
}

func init() {

}
