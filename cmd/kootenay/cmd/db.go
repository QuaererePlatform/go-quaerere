package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/arangodb/go-driver"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/QuaererePlatform/go-quaerere/internal/storage"
	"github.com/QuaererePlatform/go-quaerere/internal/storage/arangodb"
)

type (
	dbInitConfig struct {
		AppEnv         *string `mapstructure:"app_env"`
		DebugMode      bool    `mapstructure:"debug_mode"`
		StorageBackend string  `mapstructure:"storage_backend"`
	}
)

var dbInitCommand = &cobra.Command{
	Use:   "dbinit",
	Short: "Initialize the database",
	Run:   dbInit,
}

func init() {
	rootCommand.AddCommand(dbInitCommand)
}

func dbInit(cmd *cobra.Command, args []string) {
	c := new(dbInitConfig)

	if err := viper.Unmarshal(c); err != nil {
		log.Fatal().Err(err)
	}
	log.Debug().Fields(map[string]interface{}{"config": fmt.Sprintf("%#v", c)}).Msg("loaded config")

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if c.DebugMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Debug().Msg("Debug enabled")

	if strings.ToLower(*c.AppEnv) == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	var store storage.StorageDriver

	switch c.StorageBackend {
	case "arangodb":
		c := new(arangodb.Config)
		c.Endpoints = []string{
			"http://localhost:8529/",
		}
		c.Database = "quaerere"
		c.Username = "quaerere"
		c.Password = "password"
		c.Auth = true
		c.AuthType = driver.AuthenticationTypeBasic
		store = arangodb.NewArangoDBStorage(*c)
	}

	if store != nil {
		if err := store.SetupClient(); err != nil {
			log.Fatal().Err(err).Msg("error setting up client")
		}
		log.Debug().
			Str("method", "dbInit").
			Str("s", fmt.Sprintf("%#v", store)).Msg("StorageDriver object")
		if err := store.InitDB(); err != nil {
			log.Fatal().Err(err).Msg("error initializing database")
		}
	} else {
		log.Fatal().Msg("Data store is nil")
	}
}
