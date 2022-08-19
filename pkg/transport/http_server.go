package transport

import (
	"context"
	"github.com/XM-GO/PandaKit/logger"
	"github.com/emicklei/go-restful/v3"
	"net/http"
	"pandax/pkg/global"
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
		if err := s.srv.ListenAndServe(); err != nil {
			global.Log.Errorf("error http serve: %s", err)
		}
	}()
	return nil
}

func (s *HttpServer) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

type httpLog struct{}

func (t *httpLog) Print(v ...any) {
	logger.Log.Debug(v...)
}

func (t *httpLog) Printf(format string, v ...any) {
	logger.Log.Debugf(format, v...)
}
