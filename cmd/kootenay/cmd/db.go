package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/QuaererePlatform/go-quaerere/internal/storage"
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

	store := storage.NewStorage(c.StorageBackend)

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
}
