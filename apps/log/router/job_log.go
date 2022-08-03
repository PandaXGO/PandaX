package router

import (
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/log/api"
	"pandax/apps/log/entity"
	"pandax/apps/log/services"
)

func InitJobLogRouter(container *restful.Container) {
	// Job日志
	s := &api.LogJobApi{
		LogJobApp: services.LogJobModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/log/logJob").Produces(restful.MIME_JSON)
	tags := []string{"logJob"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取操作日志列表").Handle(s.GetJobLogList)
	}).
		Doc("获取操作日志列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.LogJob{}).
		Returns(200, "OK", []entity.LogJob{}))

	ws.Route(ws.DELETE("/{logId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除操作日志信息").Handle(s.DeleteJobLog)
	}).
		Doc("删除操作日志信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("logId", "多id 1,2,3").DataType("string")))

	ws.Route(ws.DELETE("/all").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("清空操作日志信息").Handle(s.DeleteAll)
	}).
		Doc("清空操作日志信息").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	container.Add(ws)
}
