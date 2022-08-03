package transport

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"pandax/pkg/global"
)

type GrpcServer struct {
	Addr string
	srv  *grpc.Server
}

func NewServer(addr string) *GrpcServer {
	return &GrpcServer{
		Addr: addr,
		srv:  grpc.NewServer(),
	}
}

func (s *GrpcServer) GetServe() *grpc.Server {
	return s.srv
}

func (s *GrpcServer) Type() Type {
	return TypeGRPC
}

func (s *GrpcServer) Start(ctx context.Context) error {
	l, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return fmt.Errorf("error listen addr: %w", err)
	}
	global.Log.Debugf("GRPC Server listen: %s", s.Addr)
	go func() {
		if err := s.srv.Serve(l); err != nil {
			global.Log.Errorf("error http serve: %s", err)
		}
	}()
	return nil
}

func (s *GrpcServer) Stop(ctx context.Context) error {
	s.srv.Stop()
	return nil
}
