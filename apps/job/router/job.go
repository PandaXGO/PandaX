package router

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/job/api"
	"pandax/apps/job/entity"
	"pandax/apps/job/services"
)

func InitJobRouter(container *restful.Container) {
	// 登录日志
	s := &api.JobApi{
		JobApp: services.JobModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/job").Produces(restful.MIME_JSON)
	tags := []string{"job"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Job列表").Handle(s.GetJobList)
	}).
		Doc("获取Job列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("jobName", "jobName").DataType("string")).
		Param(ws.QueryParameter("jobGroup", "jobGroup").DataType("string")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{jobId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Job列表").Handle(s.GetJob)
	}).
		Doc("获取Job列表").
		Param(ws.PathParameter("jobId", "Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.SysJob{}). // on the response
		Returns(200, "OK", entity.SysJob{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加Job信息").Handle(s.CreateJob)
	}).
		Doc("添加Job信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysJob{})) // from the request

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改Job信息").Handle(s.UpdateJob)
	}).
		Doc("修改Job信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysJob{})) // from the request

	ws.Route(ws.DELETE("/{jobId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除Job信息").Handle(s.DeleteJob)
	}).
		Doc("删除Job信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("jobId", "多id 1,2,3").DataType("string")))

	ws.Route(ws.GET("/stop/{jobId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("停止一个job").Handle(s.StopJobForService)
	}).
		Doc("停止一个job").
		Param(ws.PathParameter("jobId", "Id").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/start/{jobId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("开启一个job").Handle(s.StartJobForService)
	}).
		Doc("开启一个job").
		Param(ws.PathParameter("jobId", "Id").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/changeStatus").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改状态").Handle(s.UpdateStatusJob)
	}).
		Doc("修改状态").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	container.Add(ws)
}
