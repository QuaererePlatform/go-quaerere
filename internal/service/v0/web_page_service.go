package v0

import (
	"context"
	"errors"

	"github.com/QuaererePlatform/go-quaerere/internal/storage"
	"github.com/QuaererePlatform/go-quaerere/pkg/api/v0"
)

const apiVersion = "v0"

type webPageServiceServer struct {
	storage *storage.Storage
}

/*func (s *webPageServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}*/

func NewWebPageServiceServer(s *storage.StorageDriver) v0.WebPageServiceServer {

}

func (s *webPageServiceServer) Create(ctx context.Context, req *v0.CreateWebPageRequest) (*v0.CreateWebPageResponse, error) {


	e := errors.New("unimplemented")
	return nil, e
}
