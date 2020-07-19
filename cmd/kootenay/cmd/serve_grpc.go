package cmd

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/QuaererePlatform/go-quaerere/internal/protocol/grpc/server"
)

var serveGRPCCommand = &cobra.Command{
	Use:   "serve-grpc",
	Short: "Start the GRPC service",
	Run:   serve_grpc,
}

func init() {
	rootCommand.AddCommand(serveGRPCCommand)
}

func serve_grpc(cmd *cobra.Command, args []string) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Hello World")

	c := new(server.Config)

	if err := viper.Unmarshal(c); err != nil {
		log.Fatal().Err(err)
	}

	s, err := server.New(c)
	if err != nil {
		log.Fatal().Err(err)
	}

	go func() {
		if err := s.Start(); err != nil {
			log.Fatal().Err(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal().Err(err)
	}

}
