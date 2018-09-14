package server

import (
	"context"
	"fmt"

	api "github.com/ehazlett/blackbird/api/v1"
	ptypes "github.com/gogo/protobuf/types"
)

func (s *Server) Reload(ctx context.Context, req *api.ReloadRequest) (*ptypes.Empty, error) {
	return empty, fmt.Errorf("not implemented")
}
