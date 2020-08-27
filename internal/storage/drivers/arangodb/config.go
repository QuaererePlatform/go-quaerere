package arangodb

import (
	"github.com/arangodb/go-driver"
	"github.com/spf13/viper"
)

const (
	ENV_FLAG_AUTH_TYPE = "datastore.arangodb.auth_type"
	ENV_FLAG_DATABASE = "datastore.arangodb.database"
	ENV_FLAG_ENDPOINTS = "datastore.arangodb.endpoints"
	ENV_FLAG_PASSWORD = "datastore.arangodb.password"
	ENV_FLAG_USERNAME = "datastore.arangodb.username"

	FLAG_AUTH_TYPE_DEFAULT = "basic"
	FLAG_DATABASE_DEFAULT = "quaerere"
	FLAG_ENDPOINTS_DEFAULT = "http://localhost:8529/"
	FLAG_PASSWORD_DEFAULT = "password"
	FLAG_USERNAME_DEFAULT = "quaerere"
)

type Config struct {
	ConfigAuthType string   `mapstructure:"datastore.arangodb.auth_type"`
	Database       string   `mapstructure:"datastore.arangodb.database"`
	Endpoints      []string `mapstructure:"datastore.arangodb.endpoints"`
	Password       string   `mapstructure:"datastore.arangodb.password"`
	Username       string   `mapstructure:"datastore.arangodb.username"`
	authType       driver.AuthenticationType
	auth           bool
}

func init() {
	viper.SetDefault(ENV_FLAG_AUTH_TYPE, FLAG_AUTH_TYPE_DEFAULT)
	viper.SetDefault(ENV_FLAG_DATABASE, FLAG_DATABASE_DEFAULT)
	viper.SetDefault(ENV_FLAG_ENDPOINTS, FLAG_ENDPOINTS_DEFAULT)
	viper.SetDefault(ENV_FLAG_PASSWORD, FLAG_PASSWORD_DEFAULT)
	viper.SetDefault(ENV_FLAG_USERNAME, FLAG_USERNAME_DEFAULT)
}
