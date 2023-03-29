// ==========================================================================
// 生成日期：2023-03-29 20:01:11 +0800 CST
// 生成路径: apps/flow/router/flow_work_info.go
// 生成人：panda
// ==========================================================================
package router

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"pandax/apps/flow/api"
	"pandax/apps/flow/entity"
	"pandax/apps/flow/services"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitFlowWorkInfoRouter(container *restful.Container) {
	s := &api.FlowWorkInfoApi{
		FlowWorkInfoApp: services.FlowWorkInfoModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/flow/workinfo").Produces(restful.MIME_JSON)
	tags := []string{"workinfo"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取WorkInfo分页列表").Handle(s.GetFlowWorkInfoList)
	}).
		Doc("获取WorkInfo分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取WorkInfo信息").Handle(s.GetFlowWorkInfo)
	}).
		Doc("获取WorkInfo信息").
		Param(ws.PathParameter("id", "Id").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.FlowWorkInfo{}). // on the response
		Returns(200, "OK", entity.FlowWorkInfo{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加WorkInfo信息").Handle(s.InsertFlowWorkInfo)
	}).
		Doc("添加WorkInfo信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.FlowWorkInfo{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改WorkInfo信息").Handle(s.UpdateFlowWorkInfo)
	}).
		Doc("修改WorkInfo信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.FlowWorkInfo{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除WorkInfo信息").Handle(s.DeleteFlowWorkInfo)
	}).
		Doc("删除WorkInfo信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
