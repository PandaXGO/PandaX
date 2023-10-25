package updserver

import (
	"context"
	"net"
	"pandax/pkg/global"
)

const DefaultPort = ":9003"

type UdpServer struct {
	Addr     string
	listener *net.UDPConn
}

func NewUdpServer(addr string) *UdpServer {
	if addr == "" {
		addr = DefaultPort
	}
	return &UdpServer{
		Addr: addr,
	}
}

func (s *UdpServer) GetServe() *net.UDPConn {
	return s.listener
}

func (s *UdpServer) Type() string {
	return "UDP"
}

func (s *UdpServer) Start(ctx context.Context) error {
	addr, _ := net.ResolveUDPAddr("udp", s.Addr)
	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		global.Log.Errorf("error http serve: %s", err)
		return err
	}
	s.listener = listener
	return nil
}

func (s *UdpServer) Stop(ctx context.Context) error {
	s.listener.Close()
	return nil
}
