package server

import (
	"github.com/ehazlett/blackbird"
	api "github.com/ehazlett/blackbird/api/v1"
	ptypes "github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
)

var (
	empty = &ptypes.Empty{}
)

type Server struct {
	config     *blackbird.Config
	grpcServer *grpc.Server
}

func NewServer(cfg *blackbird.Config) (*Server, error) {
	grpcServer := grpc.NewServer()
	srv := &Server{
		config:     cfg,
		grpcServer: grpcServer,
	}

	api.RegisterProxyServer(grpcServer, srv)

	return srv, nil
}
