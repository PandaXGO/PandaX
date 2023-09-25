package httpserver

import (
	"context"
	"github.com/emicklei/go-restful/v3"
	"net/http"
	"pandax/pkg/global"
)

const DefaultPort = ":9002"

type HttpServer struct {
	Addr      string
	srv       *http.Server
	Container *restful.Container
}

func NewHttpServer(addr string) *HttpServer {
	if addr == "" {
		addr = DefaultPort
	}
	c := restful.NewContainer()
	c.EnableContentEncoding(true)
	return &HttpServer{
		Addr:      addr,
		Container: c,
		srv: &http.Server{
			Addr:    addr,
			Handler: c,
		},
	}
}

func (s *HttpServer) GetServe() *http.Server {
	return s.srv
}

func (s *HttpServer) Type() string {
	return "HTTP"
}

func (s *HttpServer) Start(ctx context.Context) error {
	go func() {
		if global.Conf.Server.Tls.Enable {
			global.Log.Infof("HTTPS Server listen: %s", s.Addr)
			if err := s.srv.ListenAndServeTLS(global.Conf.Server.Tls.CertFile, global.Conf.Server.Tls.KeyFile); err != nil {
				global.Log.Errorf("error http serve: %s", err)
			}
		} else {
			global.Log.Infof("HTTP Server listen: %s", s.Addr)
			if err := s.srv.ListenAndServe(); err != nil {
				global.Log.Errorf("error http serve: %s", err)
			}
		}
	}()
	return nil
}

func (s *HttpServer) Stop(ctx context.Context) error {
	s.srv.Shutdown(ctx)
	return nil
}
