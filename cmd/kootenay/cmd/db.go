package cmd

import (
	"context"
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
		log.Fatal().Err(err).Msg("viper unmarshal error")
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
		var err error
		store, err = arangodb.NewArangoDBStorage(*c)
		if err != nil {
			log.Fatal().Err(err)
		}
	}

	log.Debug().Str("store", fmt.Sprintf("%#v", store)).Msg("newly created store")
	if store != nil {
		ctx := context.Background()
		ctxWithCancel, cancelFunction := context.WithCancel(ctx)

		defer func() {
			log.Info().Msg("Main Defer: canceling context")
			cancelFunction()
		}()

		if err := store.Connect(ctxWithCancel); err != nil {
			log.Fatal().Err(err).Msg("error connecting to database")
		}
		log.Debug().Str("store", fmt.Sprintf("%#v", store)).Msg("store after Connect()")
		if err := store.InitDB(); err != nil {
			log.Fatal().Err(err).Msg("error initializing database")
		}
	} else {
		log.Fatal().Msg("Data store is nil")
	}
}
