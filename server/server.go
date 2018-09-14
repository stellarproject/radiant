package server

import (
	"context"
	"fmt"

	api "github.com/ehazlett/blackbird/api/v1"
)

func (s *Server) Servers(ctx context.Context, req *api.ServersRequest) (*api.ServersResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
