// ==========================================================================
// 生成日期：2023-04-10 02:51:27 +0000 UTC
// 生成路径: apps/visual/router/visual_data_set_group.go
// 生成人：panda
// ==========================================================================
package router

import (
	"github.com/XM-GO/PandaKit/restfulx"
	"pandax/apps/visual/api"
	"pandax/apps/visual/entity"
	"pandax/apps/visual/services"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitVisualScreenGroupRouter(container *restful.Container) {
	s := &api.VisualScreenGroupApi{
		VisualScreenGroupApp: services.VisualScreenGroupModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/visual/screen/group").Produces(restful.MIME_JSON)
	tags := []string{"datasetgroup"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取ScreenGroup列表").Handle(s.GetScreenGroupList)
	}).
		Doc("获取ScreenGroup列表").
		Param(ws.QueryParameter("name", "名称").Required(false).DataType("string")).
		Param(ws.QueryParameter("status", "状态").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", []entity.VisualScreenGroup{}))

	ws.Route(ws.GET("/list/all").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取ScreenGroup所有列表").Handle(s.GetScreenGroupAllList)
	}).
		Doc("获取ScreenGroup分页列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", []entity.VisualScreenGroup{}))

	ws.Route(ws.GET("/list/tree").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取ScreenGroup树").Handle(s.GetScreenGroupTree)
	}).
		Doc("获取ScreenGroup树").
		Param(ws.QueryParameter("name", "名称").Required(false).DataType("string")).
		Param(ws.QueryParameter("status", "状态").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", []entity.ScreenGroupLabel{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取ScreenGroup信息").Handle(s.GetVisualScreenGroup)
	}).
		Doc("获取ScreenGroup信息").
		Param(ws.PathParameter("", "Id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.VisualScreenGroup{}). // on the response
		Returns(200, "OK", entity.VisualScreenGroup{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加ScreenGroup信息").Handle(s.InsertVisualScreenGroup)
	}).
		Doc("添加ScreenGroup信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualScreenGroup{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改ScreenGroup信息").Handle(s.UpdateVisualScreenGroup)
	}).
		Doc("修改ScreenGroup信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualScreenGroup{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除ScreenGroup信息").Handle(s.DeleteVisualScreenGroup)
	}).
		Doc("删除ScreenGroup信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
