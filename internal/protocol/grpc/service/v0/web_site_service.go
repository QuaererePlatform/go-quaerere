package v0

import (
	"context"
	"errors"

	"github.com/QuaererePlatform/go-quaerere/internal/storage/drivers"
	"github.com/QuaererePlatform/go-quaerere/pkg/api/v0"
)

type webSiteServiceServer struct {
	v0.WebSiteServiceServer
	storage *drivers.Driver
}

func (s *webSiteServiceServer) CreateWebSite(ctx context.Context, req *v0.CreateWebSiteRequest) (*v0.CreateWebSiteResponse, error) {
	e := errors.New("unimplemented")
	return nil, e
}

func (s *webSiteServiceServer) ReadWebSite(ctx context.Context, request *v0.ReadWebSiteRequest) (*v0.ReadWebSiteResponse, error) {
	panic("implement me")
}

func (s *webSiteServiceServer) UpdateWebSite(ctx context.Context, request *v0.UpdateWebSiteRequest) (*v0.UpdateWebSiteResponse, error) {
	panic("implement me")
}

func (s *webSiteServiceServer) DeleteWebSite(ctx context.Context, request *v0.DeleteWebSiteRequest) (*v0.DeleteWebSiteResponse, error) {
	panic("implement me")
}

func (s *webSiteServiceServer) ListWebSites(ctx context.Context, request *v0.ListWebSiteRequest) (*v0.ListWebSiteResults, error) {
	panic("implement me")
}

func NewWebSiteServiceServer(s *drivers.Driver) v0.WebSiteServiceServer {
	return &webSiteServiceServer{storage: s}
}
