package server

import (
	"context"
	"fmt"

	api "github.com/ehazlett/blackbird/api/v1"
	ptypes "github.com/gogo/protobuf/types"
)

func (s *Server) RemoveServer(ctx context.Context, req *api.RemoveServerRequest) (*ptypes.Empty, error) {
	return empty, fmt.Errorf("not implemented")
}
