package server

import (
	"context"
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/QuaererePlatform/go-quaerere/internal/storage"
	v0 "github.com/QuaererePlatform/go-quaerere/pkg/api/v0"
)

type (
	Server interface {
		Shutdown(context.Context) error
		Start() error
	}

	server struct {
		config     *Config
		grpcServer *grpc.Server
		storage    *storage.Storage
	}
)

func New(c *Config) (Server, error) {

	s := server{
		config:     c,
		grpcServer: grpc.NewServer(),
		storage:    storage.NewStorage(c.StorageBackend),
	}

	return s, nil
}

func (s server) RegisterServices() error {
	v0.RegisterWebPageServiceServer()
	return nil
}

func (s server) Start() error {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.config.Bind, s.config.Port))
	if err != nil {
		return err
	}
	log.Info().Msg("starting gRPC server...")
	return s.grpcServer.Serve(listen)
}

func (s server) Shutdown(ctx context.Context) error {
	log.Info().Msg("stopping gRPC server...")
	s.grpcServer.GracefulStop()
	<-ctx.Done()
	return nil
}
