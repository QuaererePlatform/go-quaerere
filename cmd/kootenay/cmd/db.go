package cmd

import (
	"log"

	"github.com/arangodb/go-driver"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/QuaererePlatform/go-quaerere/internal/storage"
	"github.com/QuaererePlatform/go-quaerere/internal/storage/arangodb"
)

type (
	dbInitConfig struct {
		StorageBackend string `mapstructure:"storage_backend"`
	}
)

var dbInitCmd = &cobra.Command{
	Use:   "dbinit",
	Short: "Start the http service",
	Run:   dbInit,
}

func init() {
	rootCmd.AddCommand(dbInitCmd)
}

func dbInit(cmd *cobra.Command, args []string) {
	c := new(dbInitConfig)

	if err := viper.Unmarshal(c); err != nil {
		log.Fatal(err)
	}

	var store storage.StorageDriver

	switch c.StorageBackend {
	case "arangodb":
		c := new(arangodb.Config)
		c.Endpoints = []string{
			"http://arangodb:8529/",
		}
		c.Database = "quaerere"
		c.Username = "quaerere"
		c.Password = "password"
		c.Auth = true
		c.AuthType = driver.AuthenticationTypeBasic
		store = arangodb.NewArangoDBStorage(*c)
	}

	if store == nil {
		log.Fatal()
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
}
