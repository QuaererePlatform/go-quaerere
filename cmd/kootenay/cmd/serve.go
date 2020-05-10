package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/QuaererePlatform/go-quaerere/internal/server"
)

const (
	FLAG_BIND = "bind"
	FLAG_BIND_SHORT = "b"
	FLAG_BIND_DEFAULT = "0.0.0.0"
	FLAG_BIND_DESCR = "bind address for web server"

	FLAG_PORT = "port"
	FLAG_PORT_SHORT = "p"
	FLAG_PORT_DEFAULT = 1323
	FLAG_PORT_DESCR = "port for web server"

	FLAG_TLS_DISABLE = "disable-tls"
	FLAG_TLS_DISABLE_DEFAULT = false
	FLAG_TLS_DISABLE_DESCR = "disable TLS for web server"

	ENV_FLAG_BIND = FLAG_BIND
	ENV_FLAG_PORT = FLAG_PORT
	ENV_FLAG_TLS_DISABLE = "tls_disabled"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the http service",
	Run:   serve,
}

func init() {
	serveCmd.Flags().StringP(FLAG_BIND, FLAG_BIND_SHORT, FLAG_BIND_DEFAULT, FLAG_BIND_DESCR)
	serveCmd.Flags().IntP(FLAG_PORT, FLAG_PORT_SHORT, FLAG_PORT_DEFAULT, FLAG_PORT_DESCR)
	serveCmd.Flags().Bool(FLAG_TLS_DISABLE, FLAG_TLS_DISABLE_DEFAULT, FLAG_TLS_DISABLE_DESCR)

	_ = viper.BindPFlag(ENV_FLAG_BIND, serveCmd.Flags().Lookup(FLAG_BIND))
	_ = viper.BindPFlag(ENV_FLAG_PORT, serveCmd.Flags().Lookup(FLAG_PORT))
	_ = viper.BindPFlag(ENV_FLAG_TLS_DISABLE, serveCmd.Flags().Lookup(FLAG_TLS_DISABLE))

	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	c := new(server.Config)

	/*for _, i := range []interface{}{
		c,
	} {
		if err := viper.Unmarshal(i); err != nil {
			log.Fatal(err)
		}
	}*/

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
