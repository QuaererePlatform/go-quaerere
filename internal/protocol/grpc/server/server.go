package server

import (
	"context"
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	v0service "github.com/QuaererePlatform/go-quaerere/internal/service/v0"
	"github.com/QuaererePlatform/go-quaerere/internal/storage/drivers"
	v0api "github.com/QuaererePlatform/go-quaerere/pkg/api/v0"
)

type (
	Server interface {
		Shutdown(context.Context) error
		Start() error
	}

	server struct {
		config     *Config
		grpcServer *grpc.Server
		storage    drivers.Driver
	}
)

func New(c *Config) (Server, error) {

	s := server{
		config:     c,
		grpcServer: grpc.NewServer(),
	}

	if err := s.setupStorage(); err != nil {
		return nil, err
	}

	return &s, nil
}

func (s *server) RegisterServices() error {
	wp := v0service.NewWebPageServiceServer(&s.storage)
	v0api.RegisterWebPageServiceServer(s.grpcServer, wp)
	return nil
}

func (s *server) Start() error {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.config.Bind, s.config.Port))
	if err != nil {
		return err
	}
	log.Info().Msg("starting gRPC server...")
	return s.grpcServer.Serve(listen)
}

func (s *server) Shutdown(ctx context.Context) error {
	log.Info().Msg("stopping gRPC server...")
	s.grpcServer.GracefulStop()
	<-ctx.Done()
	return nil
}

func (s *server) setupStorage() error {
	var err error
	s.storage, err = drivers.NewDriver(s.config.StorageBackend)
	if err != nil {
		log.Fatal().Err(err).Msg("error setting up datastore")
	}

	if err := s.storage.InitDB(); err != nil {
		return err
	}
	return nil
}
