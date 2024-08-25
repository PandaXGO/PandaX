package transport

import (
	"context"
	"net/http"
	"pandax/pkg/global"

	"github.com/emicklei/go-restful/v3"
)

type HttpServer struct {
	Addr string
	srv  *http.Server

	Container *restful.Container
}

func NewHttpServer(addr string) *HttpServer {
	c := restful.NewContainer()
	c.EnableContentEncoding(true)
	restful.TraceLogger(&httpLog{})
	restful.SetLogger(&httpLog{})
	return &HttpServer{
		Addr:      addr,
		Container: c,
		srv: &http.Server{
			Addr:    addr,
			Handler: c,
		},
	}
}

func (s *HttpServer) Type() Type {
	return TypeHTTP
}

func (s *HttpServer) Start(ctx context.Context) error {
	global.Log.Infof("HTTP Server listen: %s", s.Addr)
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
	return s.srv.Shutdown(ctx)
}

type httpLog struct{}

func (t *httpLog) Print(v ...any) {
	global.Log.Debug(v...)
}

func (t *httpLog) Printf(format string, v ...any) {
	global.Log.Debugf(format, v...)
}
