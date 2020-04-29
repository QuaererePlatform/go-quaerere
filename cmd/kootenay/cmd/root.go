package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	ENV_PREFIX = "kootenay"

	FLAG_APP_ENV_DEFAULT = "development"
	FLAG_DEBUG           = "debug"
	FLAG_DEBUG_SHORT     = "d"
	FLAG_DEBUG_DEFAULT   = false
	FLAG_DEBUG_DESCR     = "enable debug mode"
	FLAG_STORAGE         = "storage-backend"
	FLAG_STORAGE_SHORT   = "s"
	FLAG_STORAGE_DEFAULT = "arangodb"
	FLAG_STORAGE_DESCR   = "set storage backend"

	ENV_FLAG_APP_ENV = "app_env"
	ENV_FLAG_DEBUG   = "debug_mode"
)

var rootCmd = &cobra.Command{
	Use:   "kootenay",
	Short: "The kootenay microservice, part of The QuaererePlatform",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolP(FLAG_DEBUG, FLAG_DEBUG_SHORT, FLAG_DEBUG_DEFAULT, FLAG_DEBUG_DESCR)
	rootCmd.PersistentFlags().StringP(FLAG_STORAGE, FLAG_STORAGE_SHORT, FLAG_STORAGE_DEFAULT, FLAG_STORAGE_DESCR)
}

func initConfig() {

	viper.SetDefault(ENV_FLAG_APP_ENV, FLAG_APP_ENV_DEFAULT)
	viper.SetDefault(ENV_FLAG_DEBUG, FLAG_DEBUG_DEFAULT)

	viper.SetEnvPrefix(ENV_PREFIX)

/*	for k, v := range map[string]string{
		FLAG_DEBUG: ENV_FLAG_DEBUG,
	} {
		if err := viper.BindPFlag(v, serveCmd.PersistentFlags().Lookup(k)); err != nil {
			log.Fatal(err)
		}
	}

	for _, e := range []string{
		ENV_FLAG_APP_ENV,
		ENV_FLAG_DEBUG,
	} {
		if err := viper.BindEnv(e); err != nil {
			log.Fatal(err)
		}
	}*/
}
