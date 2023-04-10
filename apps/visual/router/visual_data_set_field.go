// ==========================================================================
// 生成日期：2023-04-10 03:05:21 +0800 CST
// 生成路径: apps/visual/router/visual_data_set_field.go
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

func InitVisualDataSetFieldRouter(container *restful.Container) {
	s := &api.VisualDataSetFieldApi{
		VisualDataSetFieldApp: services.VisualDataSetFieldModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/visual/dataset/field").Produces(restful.MIME_JSON)
	tags := []string{"datasetfield"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取DataSetField分页列表").Handle(s.GetVisualDataSetFieldList)
	}).
		Doc("获取DataSetField分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{fieldId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取DataSetField信息").Handle(s.GetVisualDataSetField)
	}).
		Doc("获取DataSetField信息").
		Param(ws.PathParameter("fieldId", "Id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.VisualDataSetField{}). // on the response
		Returns(200, "OK", entity.VisualDataSetField{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加DataSetField信息").Handle(s.InsertVisualDataSetField)
	}).
		Doc("添加DataSetField信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualDataSetField{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改DataSetField信息").Handle(s.UpdateVisualDataSetField)
	}).
		Doc("修改DataSetField信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualDataSetField{}))

	ws.Route(ws.DELETE("/{fieldId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除DataSetField信息").Handle(s.DeleteVisualDataSetField)
	}).
		Doc("删除DataSetField信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("fieldId", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
