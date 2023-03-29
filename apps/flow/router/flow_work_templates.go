// ==========================================================================
// 生成日期：2023-03-29 19:46:55 +0800 CST
// 生成路径: apps/flow/router/flow_work_templates.go
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

func InitFlowWorkTemplatesRouter(container *restful.Container) {
	s := &api.FlowWorkTemplatesApi{
		FlowWorkTemplatesApp: services.FlowWorkTemplatesModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/flow/worktemplates").Produces(restful.MIME_JSON)
	tags := []string{"worktemplates"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取WorkTemplates分页列表").Handle(s.GetFlowWorkTemplatesList)
	}).
		Doc("获取WorkTemplates分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取WorkTemplates信息").Handle(s.GetFlowWorkTemplates)
	}).
		Doc("获取WorkTemplates信息").
		Param(ws.PathParameter("id", "Id").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.FlowWorkTemplates{}). // on the response
		Returns(200, "OK", entity.FlowWorkTemplates{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加WorkTemplates信息").Handle(s.InsertFlowWorkTemplates)
	}).
		Doc("添加WorkTemplates信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.FlowWorkTemplates{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改WorkTemplates信息").Handle(s.UpdateFlowWorkTemplates)
	}).
		Doc("修改WorkTemplates信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.FlowWorkTemplates{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除WorkTemplates信息").Handle(s.DeleteFlowWorkTemplates)
	}).
		Doc("删除WorkTemplates信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
