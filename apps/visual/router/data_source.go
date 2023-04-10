// ==========================================================================
// 生成日期：2023-04-10 02:51:27 +0000 UTC
// 生成路径: apps/visual/router/visual_data_source.go
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

func InitVisualDataSourceRouter(container *restful.Container) {
	s := &api.VisualDataSourceApi{
		VisualDataSourceApp: services.VisualDataSourceModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/visual/datasource").Produces(restful.MIME_JSON)
	tags := []string{"datasource"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取DataSource分页列表").Handle(s.GetVisualDataSourceList)
	}).
		Doc("获取DataSource分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{sourceId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取DataSource信息").Handle(s.GetVisualDataSource)
	}).
		Doc("获取DataSource信息").
		Param(ws.PathParameter("sourceId", "Id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.VisualDataSource{}). // on the response
		Returns(200, "OK", entity.VisualDataSource{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加DataSource信息").Handle(s.InsertVisualDataSource)
	}).
		Doc("添加DataSource信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualDataSource{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改DataSource信息").Handle(s.UpdateVisualDataSource)
	}).
		Doc("修改DataSource信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualDataSource{}))

	ws.Route(ws.DELETE("/{sourceId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除DataSource信息").Handle(s.DeleteVisualDataSource)
	}).
		Doc("删除DataSource信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("sourceId", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
