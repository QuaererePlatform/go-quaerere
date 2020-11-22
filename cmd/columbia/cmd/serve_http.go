package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/QuaererePlatform/go-quaerere/internal/config"
	"github.com/QuaererePlatform/go-quaerere/internal/protocol/http/server"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start",
	Run:   serve,
}

func init() {
	serveCmd.Flags().StringP(
		"bind", "b", "0.0.0.0", "bind address for web server")
	serveCmd.Flags().IntP("port", "p", 1323, "port for web server")
	serveCmd.Flags().Bool("disable-tls", false, "disable TLS for web server")

	_ = viper.BindPFlag("bind", serveCmd.Flags().Lookup("bind"))
	_ = viper.BindPFlag("port", serveCmd.Flags().Lookup("port"))
	_ = viper.BindPFlag("tls_disabled", serveCmd.Flags().Lookup("disable-tls"))


	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	c := new(config.AppConfig)

	if err := viper.Unmarshal(c); err != nil {
		log.Fatal(err)
	}

	s, err := server.New(c)
	if err != nil {
		log.Fatal(err)
	}

	go func() { log.Fatal(s.Start()) }()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	log.Fatal(s.Shutdown(ctx))
}
