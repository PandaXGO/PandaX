// ==========================================================================
// 生成日期：2023-04-10 11:31:34 +0800 CST
// 生成路径: apps/visual/router/visual_data_set_table_task.go
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

func InitVisualDataSetTableTaskRouter(container *restful.Container) {
	s := &api.VisualDataSetTableTaskApi{
		VisualDataSetTableTaskApp: services.VisualDataSetTableTaskModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/visual/dataset/tabletask").Produces(restful.MIME_JSON)
	tags := []string{"datasettabletask"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取DataSetTableTask分页列表").Handle(s.GetVisualDataSetTableTaskList)
	}).
		Doc("获取DataSetTableTask分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取DataSetTableTask信息").Handle(s.GetVisualDataSetTableTask)
	}).
		Doc("获取DataSetTableTask信息").
		Param(ws.PathParameter("id", "Id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.VisualDataSetTableTask{}). // on the response
		Returns(200, "OK", entity.VisualDataSetTableTask{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加DataSetTableTask信息").Handle(s.InsertVisualDataSetTableTask)
	}).
		Doc("添加DataSetTableTask信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualDataSetTableTask{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改DataSetTableTask信息").Handle(s.UpdateVisualDataSetTableTask)
	}).
		Doc("修改DataSetTableTask信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualDataSetTableTask{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除DataSetTableTask信息").Handle(s.DeleteVisualDataSetTableTask)
	}).
		Doc("删除DataSetTableTask信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
