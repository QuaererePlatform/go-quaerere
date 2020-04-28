package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "kootenay",
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

	rootCmd.PersistentFlags().BoolP("debug", "d", false, "enable debug mode")
}

func initConfig() {
	envAppEnv := "app_env"

	viper.SetDefault(envAppEnv, "development")
	viper.SetDefault("debug_mode", false)

	viper.SetEnvPrefix("columbia")

	/*	for k, v := range map[string]string{
			"debug": "debug_mode",
		} {
			if err := viper.BindPFlag(v, serveCmd.PersistentFlags().Lookup(k)); err != nil {
				log.Fatal(err)
			}
		}

		for _, e := range []string{
			"app_env",
			"debug_mode",
		} {
			if err := viper.BindEnv(e); err != nil {
				log.Fatal(err)
			}
		}*/
}
