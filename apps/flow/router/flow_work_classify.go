// ==========================================================================
// 生成日期：2022-08-24 22:02:33 +0800 CST
// 生成路径: apps/flow/router/flow_work_classify.go
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

func InitFlowWorkClassifyRouter(container *restful.Container) {
	s := &api.FlowWorkClassifyApi{
		FlowWorkClassifyApp: services.FlowWorkClassifyModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/flow/workclassify").Produces(restful.MIME_JSON)
	tags := []string{"workclassify"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Classify分页列表").Handle(s.GetFlowWorkClassifyList)
	}).
		Doc("获取Classify分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Classify信息").Handle(s.GetFlowWorkClassify)
	}).
		Doc("获取Classify信息").
		Param(ws.PathParameter("id", "Id").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.FlowWorkClassify{}). // on the response
		Returns(200, "OK", entity.FlowWorkClassify{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加Classify信息").Handle(s.InsertFlowWorkClassify)
	}).
		Doc("添加Classify信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.FlowWorkClassify{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改Classify信息").Handle(s.UpdateFlowWorkClassify)
	}).
		Doc("修改Classify信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.FlowWorkClassify{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除Classify信息").Handle(s.DeleteFlowWorkClassify)
	}).
		Doc("删除Classify信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
