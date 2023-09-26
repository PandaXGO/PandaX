package tcpserver

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"pandax/pkg/global"
)

const DefaultPort = ":9003"

type TcpServer struct {
	Addr     string
	listener *net.TCPListener
}

func NewTcpServer(addr string) *TcpServer {
	if addr == "" {
		addr = DefaultPort
	}
	return &TcpServer{
		Addr: addr,
	}
}

func (s *TcpServer) GetServe() *net.TCPListener {
	return s.listener
}

func (s *TcpServer) Type() string {
	return "TCP"
}

func (s *TcpServer) Start(ctx context.Context) error {
	addr, _ := net.ResolveTCPAddr("tcp", s.Addr)
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		global.Log.Errorf("error http serve: %s", err)
		return err
	}
	s.listener = listener
	return nil
}

func (s *TcpServer) Stop(ctx context.Context) error {
	s.listener.Close()
	return nil
}

func (s *TcpServer) TlsConfig() (*tls.Config, error) {
	var certificates []tls.Certificate
	cert, err := tls.LoadX509KeyPair(global.Conf.Server.Tls.CertFile, global.Conf.Server.Tls.KeyFile)
	if err != nil {
		return nil, fmt.Errorf("generate x509 key pair failed: %s ", err)
	}
	certificates = append(certificates, cert)

	if len(certificates) == 0 {
		return nil, fmt.Errorf("none valid certs and secret")
	}

	return &tls.Config{Certificates: certificates}, nil
}
