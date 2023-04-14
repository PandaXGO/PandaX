// ==========================================================================
// 生成日期：2023-04-10 11:26:49 +0800 CST
// 生成路径: apps/visual/router/visual_screen.go
// 生成人：panda
// ==========================================================================
package router

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"pandax/apps/visual/api"
	"pandax/apps/visual/entity"
	"pandax/apps/visual/services"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitVisualScreenRouter(container *restful.Container) {
	s := &api.VisualScreenApi{
		VisualScreenApp: services.VisualScreenModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/visual/screen").Produces(restful.MIME_JSON)
	tags := []string{"screen"}

	ws.Route(ws.GET("/twin").To(s.ScreenTwin)).Doc("Websocket孪生")

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Screen分页列表").Handle(s.GetVisualScreenList)
	}).
		Doc("获取Screen分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{screenId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Screen信息").Handle(s.GetVisualScreen)
	}).
		Doc("获取Screen信息").
		Param(ws.PathParameter("screenId", "Id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.VisualScreen{}). // on the response
		Returns(200, "OK", entity.VisualScreen{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加Screen信息").Handle(s.InsertVisualScreen)
	}).
		Doc("添加Screen信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualScreen{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改Screen信息").Handle(s.UpdateVisualScreen)
	}).
		Doc("修改Screen信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualScreen{}))

	ws.Route(ws.DELETE("/{screenId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除Screen信息").Handle(s.DeleteVisualScreen)
	}).
		Doc("删除Screen信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("screenId", "多id 1,2,3").DataType("string")))

	ws.Route(ws.PUT("/changeStatus").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改状态").Handle(s.UpdateScreenStatus)
	}).
		Doc("修改状态").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualScreen{}))

	container.Add(ws)
}
