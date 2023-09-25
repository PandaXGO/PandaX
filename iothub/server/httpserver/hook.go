package httpserver

import (
	"context"
	"github.com/emicklei/go-restful/v3"
	"io"
	"log"
	"net"
	"net/http"
	"pandax/iothub/hook_message_work"
	"pandax/pkg/global"
	"strings"
)

type HookHttpService struct {
	HookService *hook_message_work.HookService
}

func InitHttpHook(addr string, hs *hook_message_work.HookService) {
	server := NewHttpServer(addr)
	service := NewHookHttpService(hs)
	container := server.Container
	ws := new(restful.WebService)
	ws.Path("/api/v1").Produces(restful.MIME_JSON)
	ws.Route(ws.POST("/{token}/{pathType}").Filter(basicAuthenticate).To(service.hook))
	container.Add(ws)

	server.srv.ConnState = func(conn net.Conn, state http.ConnState) {
		switch state {
		case http.StateNew:
			log.Println("New connection", conn.RemoteAddr())
		case http.StateActive:
			log.Println("Connection active", conn.RemoteAddr())
		case http.StateIdle:
			log.Println("Connection idle", conn.RemoteAddr())
		case http.StateHijacked, http.StateClosed:
			log.Println("Connection closed", conn.RemoteAddr())
		}
	}
	err := server.Start(context.TODO())
	if err != nil {
		global.Log.Error("IOTHUB HTTP服务启动错误", err)
	} else {
		global.Log.Infof("IOTHUB HOOK Start SUCCESS,Grpc Server listen: %s", addr)
	}
}

func NewHookHttpService(hs *hook_message_work.HookService) *HookHttpService {
	hhs := &HookHttpService{
		HookService: hs,
	}
	return hhs
}

// 获取token进行认证
func basicAuthenticate(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	path := req.Request.URL.Path
	log.Println(path)
	split := strings.Split(path, "/")
	log.Println(split)
	chain.ProcessFilter(req, resp)
}

func (hhs *HookHttpService) hook(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "42")
}
