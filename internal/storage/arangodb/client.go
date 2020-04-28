package arangodb

import (
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type (
	Client struct {
		client driver.Client
		config Config
		db     driver.Database
	}

	Config struct {
		Endpoints []string
		Database  string
		Username  string
		Password  string
		AuthType  *driver.AuthenticationType
	}
)

func NewClient(config Config) *Client {
	cl := new(Client)
	cl.config = config
	return cl
}

func (c Client) Connect() {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: c.config.Endpoints,
	})
	if err != nil {
		// Handle error
	}
	cc := driver.ClientConfig{
		Connection: conn,
	}
	if c.config.AuthType != nil {
		switch *c.config.AuthType {
		case driver.AuthenticationTypeBasic:
			cc.Authentication = driver.BasicAuthentication(c.config.Username, c.config.Password)
		case driver.AuthenticationTypeJWT:
			cc.Authentication = driver.JWTAuthentication(c.config.Username, c.config.Password)
		default:
			// Handle error
		}
	}
	c.client, err = driver.NewClient(cc)
	if err != nil {
		// Handle error
	}
	c.db, err = c.client.Database(nil, c.config.Database)
	if err != nil {
		// Handle error
	}
}

func (c Client) GetCollection(name string) driver.Collection {
	coll, err := c.db.Collection(nil, name)
	if err != nil {
		// Handle error
		return nil
	}
	return coll
}
