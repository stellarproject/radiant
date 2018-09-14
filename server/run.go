package server

import "net"

func (s *Server) Run() error {
	// start grpc
	l, err := net.Listen("tcp", s.config.GRPCAddr)
	if err != nil {
		return err
	}

	go s.grpcServer.Serve(l)

	// TODO: start caddy
	return nil
}
