package emqxserver

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"pandax/pkg/global"
)

const DefaultPort = ":9001"

type Server struct {
	Addr string
	srv  *grpc.Server
}

func NewServer(addr string) *Server {
	if addr == "" {
		addr = DefaultPort
	}
	return &Server{
		Addr: addr,
		srv:  grpc.NewServer(),
	}
}

func (s *Server) GetServe() *grpc.Server {
	return s.srv
}

func (s *Server) Type() string {
	return "GRPC"
}

func (s *Server) Start(ctx context.Context) error {
	l, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return fmt.Errorf("error listen addr: %w", err)
	}
	go func() {
		if err := s.srv.Serve(l); err != nil {
			global.Log.Error(fmt.Sprintf("error http serve: %s", err))
		}
	}()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.srv.Stop()
	return nil
}
