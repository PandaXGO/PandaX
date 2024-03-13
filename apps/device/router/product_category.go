package router

import (
	"pandax/apps/device/api"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	"pandax/kit/restfulx"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitProductCategoryRouter(container *restful.Container) {
	s := &api.ProductCategoryApi{
		ProductCategoryApp: services.ProductCategoryModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/device/product/category").Produces(restful.MIME_JSON)
	tags := []string{"产品分类"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取ProductCategory列表").Handle(s.GetProductCategoryList)
	}).
		Doc("获取ProductCategory列表").
		Param(ws.QueryParameter("name", "名称").Required(false).DataType("string")).
		Param(ws.QueryParameter("status", "状态").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", []entity.ProductCategory{}))

	ws.Route(ws.GET("/list/all").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取ProductCategory所有列表").Handle(s.GetProductCategoryAllList)
	}).
		Doc("获取ProductCategory分页列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", []entity.ProductCategory{}))

	ws.Route(ws.GET("/list/tree").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取ProductCategory树").Handle(s.GetProductCategoryTree)
	}).
		Doc("获取ProductCategory树").
		Param(ws.QueryParameter("name", "名称").Required(false).DataType("string")).
		Param(ws.QueryParameter("status", "状态").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", []entity.ProductCategory{}))

	ws.Route(ws.GET("/list/tree/label").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取ProductCategory树").Handle(s.GetProductCategoryTreeLabel)
	}).
		Doc("获取ProductCategory树").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", []entity.ProductCategoryLabel{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取ProductCategory信息").Handle(s.GetProductCategory)
	}).
		Doc("获取ProductCategory信息").
		Param(ws.PathParameter("", "Id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.ProductCategory{}). // on the response
		Returns(200, "OK", entity.ProductCategory{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加ProductCategory信息").Handle(s.InsertProductCategory)
	}).
		Doc("添加ProductCategory信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.ProductCategory{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改ProductCategory信息").Handle(s.UpdateProductCategory)
	}).
		Doc("修改ProductCategory信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.ProductCategory{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除ProductCategory信息").Handle(s.DeleteProductCategory)
	}).
		Doc("删除ProductCategory信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
