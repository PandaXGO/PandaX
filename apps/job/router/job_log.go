package router

import (
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"pandax/apps/job/api"
	"pandax/apps/job/services"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitJobLogRouter(container *restful.Container) {
	// Job日志
	s := &api.JobLogApi{
		JobLogApp: services.JobLogModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/job/log").Produces(restful.MIME_JSON)
	tags := []string{"任务日志"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取操作日志列表").Handle(s.GetJobLogList)
	}).
		Doc("获取操作日志列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("name", "name").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除操作日志信息").Handle(s.DeleteJobLog)
	}).
		Doc("删除操作日志信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	ws.Route(ws.DELETE("/all").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("清空操作日志信息").Handle(s.DeleteAll)
	}).
		Doc("清空操作日志信息").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	container.Add(ws)
}
