package server

import (
	"fmt"
	"net"

	"github.com/ehazlett/blackbird"
	api "github.com/ehazlett/blackbird/api/v1"
	"github.com/ehazlett/blackbird/ds"
	"github.com/ehazlett/blackbird/version"
	ptypes "github.com/gogo/protobuf/types"
	"github.com/mholt/caddy"
	_ "github.com/mholt/caddy/caddyhttp"
	"github.com/mholt/caddy/caddytls"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	empty = &ptypes.Empty{}
)

type Server struct {
	config     *blackbird.Config
	grpcServer *grpc.Server
	instance   *caddy.Instance
	// TODO: replace with interface to store in remote datastore
	servers   map[string]*api.Server
	datastore ds.Datastore
}

func NewServer(cfg *blackbird.Config) (*Server, error) {
	grpcServer := grpc.NewServer()
	srv := &Server{
		config:     cfg,
		grpcServer: grpcServer,
		servers:    make(map[string]*api.Server),
	}

	datastore, err := getDatastore(cfg.DatastoreUri)
	if err != nil {
		return nil, err
	}
	srv.datastore = datastore

	logrus.WithFields(logrus.Fields{
		"type": datastore.Name(),
	}).Debug("registered datastore")

	api.RegisterProxyServer(grpcServer, srv)

	caddy.AppName = version.Name
	caddy.AppVersion = version.FullVersion()
	caddy.Quiet = true
	caddytls.Agreed = true
	caddytls.DefaultCAUrl = "https://acme-v02.api.letsencrypt.org/directory"
	caddytls.DefaultEmail = cfg.TLSEmail

	return srv, nil
}

func (s *Server) Run() error {
	// start grpc
	l, err := net.Listen("tcp", s.config.GRPCAddr)
	if err != nil {
		return err
	}

	go s.grpcServer.Serve(l)

	caddy.SetDefaultCaddyfileLoader("default", caddy.LoaderFunc(s.defaultLoader))

	caddyfile, err := caddy.LoadCaddyfile("http")
	if err != nil {
		return err
	}
	s.instance, err = caddy.Start(caddyfile)
	if err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{
		"name":    version.Name,
		"version": version.BuildVersion(),
	}).Info("server started")

	return nil
}

func (s *Server) defaultLoader(serverType string) (caddy.Input, error) {
	return caddy.CaddyfileInput{
		Contents:       []byte(fmt.Sprintf(":%d", s.config.HTTPPort)),
		Filepath:       caddy.DefaultConfigFile,
		ServerTypeName: serverType,
	}, nil
}

func (s *Server) getCaddyConfig() (caddy.Input, error) {
	data, err := s.generateConfig()
	if err != nil {
		return nil, err

	}
	return caddy.CaddyfileInput{
		Contents:       data,
		Filepath:       caddy.DefaultConfigFile,
		ServerTypeName: "http",
	}, nil
}
