package router

import (
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/develop/api"
	"pandax/apps/develop/services"
)

func InitGenRouter(container *restful.Container) {

	// 登录日志
	s := &api.GenApi{
		GenTableApp: services.DevGenTableModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/develop/code/gen").Produces(restful.MIME_JSON)
	tags := []string{"codegen"}

	ws.Route(ws.GET("/preview/{tableId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取生成代码视图").Handle(s.Preview)
	}).
		Doc("获取生成代码视图").
		Param(ws.PathParameter("tableId", "Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", map[string]any{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/code/{tableId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("生成代码").Handle(s.GenCode)
	}).
		Doc("生成代码").
		Param(ws.PathParameter("tableId", "Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", map[string]any{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/configure/{tableId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("生成配置").Handle(s.GenConfigure)
	}).
		Doc("生成配置").
		Param(ws.PathParameter("tableId", "Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", map[string]any{}).
		Returns(404, "Not Found", nil))

	container.Add(ws)
}
