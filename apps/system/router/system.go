package router

import (
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/system/api"
)

func InitSystemRouter(container *restful.Container) {
	s := &api.System{}
	ws := new(restful.WebService)
	ws.Path("/system").Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/").To(s.ConnectWs))
	ws.Route(ws.GET("/server").To(s.ServerInfo))
	container.Add(ws)
}
