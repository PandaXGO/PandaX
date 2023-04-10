// ==========================================================================
// 生成日期：2023-04-10 11:31:34 +0800 CST
// 生成路径: apps/visual/router/visual_data_set_table_task_log.go
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

func InitVisualDataSetTableTaskLogRouter(container *restful.Container) {
	s := &api.VisualDataSetTableTaskLogApi{
		VisualDataSetTableTaskLogApp: services.VisualDataSetTableTaskLogModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/visual/dataset/tabletask/log").Produces(restful.MIME_JSON)
	tags := []string{"datasettabletasklog"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取DataSetTableTaskLog分页列表").Handle(s.GetVisualDataSetTableTaskLogList)
	}).
		Doc("获取DataSetTableTaskLog分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取DataSetTableTaskLog信息").Handle(s.GetVisualDataSetTableTaskLog)
	}).
		Doc("获取DataSetTableTaskLog信息").
		Param(ws.PathParameter("id", "Id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.VisualDataSetTableTaskLog{}). // on the response
		Returns(200, "OK", entity.VisualDataSetTableTaskLog{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加DataSetTableTaskLog信息").Handle(s.InsertVisualDataSetTableTaskLog)
	}).
		Doc("添加DataSetTableTaskLog信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualDataSetTableTaskLog{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改DataSetTableTaskLog信息").Handle(s.UpdateVisualDataSetTableTaskLog)
	}).
		Doc("修改DataSetTableTaskLog信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualDataSetTableTaskLog{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除DataSetTableTaskLog信息").Handle(s.DeleteVisualDataSetTableTaskLog)
	}).
		Doc("删除DataSetTableTaskLog信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
