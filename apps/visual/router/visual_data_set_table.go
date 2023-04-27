// ==========================================================================
// 生成日期：2023-04-10 03:05:21 +0800 CST
// 生成路径: apps/visual/router/visual_data_set_table.go
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

func InitVisualDataSetTableRouter(container *restful.Container) {
	s := &api.VisualDataSetTableApi{
		VisualDataSetTableApp: services.VisualDataSetTableModelDao,
		VisualDataSourceApp:   services.VisualDataSourceModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/visual/dataset/table").Produces(restful.MIME_JSON)
	tags := []string{"datasettable"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedCasbin(false).WithLog("获取DataSetTable分页列表").Handle(s.GetVisualDataSetTableList)
	}).
		Doc("获取DataSetTable分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{tableId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedCasbin(false).WithLog("获取DataSetTable信息").Handle(s.GetVisualDataSetTable)
	}).
		Doc("获取DataSetTable信息").
		Param(ws.PathParameter("tableId", "Id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.VisualDataSetTable{}). // on the response
		Returns(200, "OK", entity.VisualDataSetTable{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedCasbin(false).WithLog("添加DataSetTable信息").Handle(s.InsertVisualDataSetTable)
	}).
		Doc("添加DataSetTable信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualDataSetTable{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedCasbin(false).WithLog("修改DataSetTable信息").Handle(s.UpdateVisualDataSetTable)
	}).
		Doc("修改DataSetTable信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualDataSetTable{}))

	ws.Route(ws.DELETE("/{tableId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedCasbin(false).WithLog("删除DataSetTable信息").Handle(s.DeleteVisualDataSetTable)
	}).
		Doc("删除DataSetTable信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("tableId", "多id 1,2,3").DataType("string")))

	ws.Route(ws.POST("/resultPreview").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedCasbin(false).WithLog("运行结果").Handle(s.GetVisualDataSetTableResult)
	}).
		Doc("运行结果").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualDataSetRequest{}).
		Returns(200, "OK", entity.VisualDataSetRes{}))

	ws.Route(ws.POST("/up/excel").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedCasbin(false).WithLog("上传Excel").Handle(s.GetVisualDataSetTableByExcelResult)
	}).
		Doc("上传Excel").
		Param(ws.FormParameter("excelFile", "Excel文件")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", entity.VisualDataSetRes{}))

	container.Add(ws)
}
