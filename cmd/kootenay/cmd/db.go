package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/QuaererePlatform/go-quaerere/internal/config"
	"github.com/QuaererePlatform/go-quaerere/internal/storage/drivers"
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
	c := new(config.DBInitConfig)

	if err := viper.Unmarshal(c); err != nil {
		log.Fatal().Err(err).Msg("viper unmarshal error")
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if c.DebugMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Debug().Msg("Debug enabled")
	log.Debug().Str("config", fmt.Sprintf("%#v", c)).Msg("loaded config")

	if strings.ToLower(*c.AppEnv) == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	store, err := drivers.NewDriver(&c.Datastore)
	if err != nil {
		log.Fatal().Err(err).Msg("error setting up datastore")
	}

	log.Debug().Str("store", fmt.Sprintf("%#v", store)).Msg("newly created store")
	if store != nil {
		ctx := context.Background()
		ctxWithCancel, cancelFunction := context.WithCancel(ctx)

		defer func() {
			log.Debug().Msg("Main Defer: canceling context")
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
