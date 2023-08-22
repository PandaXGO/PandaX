package router

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"pandax/apps/device/api"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitProductTemplateRouter(container *restful.Container) {
	s := &api.ProductTemplateApi{
		ProductTemplateApp: services.ProductTemplateModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/device/template").Produces(restful.MIME_JSON)
	tags := []string{"template"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Template分页列表").Handle(s.GetProductTemplateList)
	}).
		Doc("获取Template分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pid", "产品ID").Required(false).DataType("string")).
		Param(ws.QueryParameter("classify", "分类").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "名称").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/list/all").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Template列表").Handle(s.GetProductTemplateListAll)
	}).
		Doc("获取Template列表").
		Param(ws.QueryParameter("pid", "产品ID").Required(false).DataType("string")).
		Param(ws.QueryParameter("classify", "分类").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "名称").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", []entity.ProductTemplate{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Template信息").Handle(s.GetProductTemplate)
	}).
		Doc("获取Template信息").
		Param(ws.PathParameter("id", "Id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.ProductTemplate{}). // on the response
		Returns(200, "OK", entity.ProductTemplate{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加Template信息").Handle(s.InsertProductTemplate)
	}).
		Doc("添加Template信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.ProductTemplate{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改Template信息").Handle(s.UpdateProductTemplate)
	}).
		Doc("修改Template信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.ProductTemplate{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除Template信息").Handle(s.DeleteProductTemplate)
	}).
		Doc("删除Template信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
