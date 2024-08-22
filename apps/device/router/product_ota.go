package router

import (
	"pandax/apps/device/api"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	"pandax/kit/model"
	"pandax/kit/restfulx"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitProductOtaRouter(container *restful.Container) {
	s := &api.ProductOtaApi{
		ProductOtaApp: services.ProductOtaModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/device/ota").Produces(restful.MIME_JSON)
	tags := []string{"产品OTA"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Ota分页列表").Handle(s.GetProductOtaList)
	}).
		Doc("获取Ota分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pid", "产品Id").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "名称").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Ota信息").Handle(s.GetProductOta)
	}).
		Doc("获取Ota信息").
		Param(ws.PathParameter("id", "Id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.ProductOta{}). // on the response
		Returns(200, "OK", entity.ProductOta{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加Ota信息").Handle(s.InsertProductOta)
	}).
		Doc("添加Ota信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.ProductOta{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改Ota信息").Handle(s.UpdateProductOta)
	}).
		Doc("修改Ota信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.ProductOta{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除Ota信息").Handle(s.DeleteProductOta)
	}).
		Doc("删除Ota信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	ws.Route(ws.DELETE("/{id}/down").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("Ota升级").Handle(s.OtaDown)
	}).
		Doc("Ota升级").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("pid", "产品Id").Required(false).DataType("string")))

	container.Add(ws)
}
