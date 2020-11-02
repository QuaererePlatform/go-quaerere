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

	ENV_FLAG_APP_ENV = "app_env"
	ENV_FLAG_DEBUG   = "debug_mode"
)

var rootCommand = &cobra.Command{
	Use:   "kootenay",
	Short: "The kootenay microservice, part of The QuaererePlatform",
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCommand.PersistentFlags().BoolP(FLAG_DEBUG, FLAG_DEBUG_SHORT, FLAG_DEBUG_DEFAULT, FLAG_DEBUG_DESCR)
}

func initConfig() {
	viper.SetEnvPrefix(ENV_PREFIX)

	viper.SetDefault(ENV_FLAG_APP_ENV, FLAG_APP_ENV_DEFAULT)
	viper.SetDefault(ENV_FLAG_DEBUG, FLAG_DEBUG_DEFAULT)

	_ = viper.BindPFlag(ENV_FLAG_DEBUG, rootCommand.Flags().Lookup(FLAG_DEBUG))

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/kootenay/")
	viper.AddConfigPath("$HOME/.kootenay/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
