package router

import (
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/log/api"
	"pandax/apps/log/entity"
	"pandax/apps/log/services"
)

func InitOperLogRouter(container *restful.Container) {
	// 操作日志
	s := &api.LogOperApi{
		LogOperApp: services.LogOperModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/log/logOper").Produces(restful.MIME_JSON)
	tags := []string{"logOper"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取操作日志列表").Handle(s.GetOperLogList)
	}).
		Doc("获取操作日志列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.LogOper{}).
		Returns(200, "OK", []entity.LogOper{}))

	ws.Route(ws.GET("/{operId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取操作日志信息").Handle(s.GetOperLog)
	}).
		Doc("获取操作日志信息").
		Param(ws.PathParameter("operId", "Id").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.LogOper{}). // on the response
		Returns(200, "OK", entity.LogOper{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.DELETE("/{operId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除操作日志信息").Handle(s.DeleteOperLog)
	}).
		Doc("删除操作日志信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("operId", "多id 1,2,3").DataType("string")))

	ws.Route(ws.DELETE("/all").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("清空操作日志信息").Handle(s.DeleteAll)
	}).
		Doc("清空操作日志信息").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	container.Add(ws)
}
