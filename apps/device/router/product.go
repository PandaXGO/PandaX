package router

import (
	"pandax/apps/device/api"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	ruleService "pandax/apps/rule/services"
	"pandax/kit/model"
	"pandax/kit/restfulx"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitProductRouter(container *restful.Container) {
	s := &api.ProductApi{
		ProductApp:  services.ProductModelDao,
		DeviceApp:   services.DeviceModelDao,
		TemplateApp: services.ProductTemplateModelDao,
		OtaAPP:      services.ProductOtaModelDao,
		RuleApp:     ruleService.RuleChainModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/device/product").Produces(restful.MIME_JSON)
	tags := []string{"产品管理"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Product分页列表").Handle(s.GetProductList)
	}).
		Doc("获取Product分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("status", "状态").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "名称").Required(false).DataType("string")).
		Param(ws.QueryParameter("productCategoryId", "产品分类Id").Required(false).DataType("string")).
		Param(ws.QueryParameter("protocolName", "协议名").Required(false).DataType("string")).
		Param(ws.QueryParameter("deviceType", "设备类型").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/list/all").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Product分页列表").Handle(s.GetProductListAll)
	}).
		Doc("获取Product分页列表").
		Param(ws.QueryParameter("status", "状态").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "名称").Required(false).DataType("string")).
		Param(ws.QueryParameter("productCategoryId", "产品分类Id").Required(false).DataType("string")).
		Param(ws.QueryParameter("protocolName", "协议名").Required(false).DataType("string")).
		Param(ws.QueryParameter("deviceType", "设备类型").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.Product{}).
		Returns(200, "OK", entity.Product{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Product信息").Handle(s.GetProduct)
	}).
		Doc("获取Product信息").
		Param(ws.PathParameter("id", "Id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.Product{}). // on the response
		Returns(200, "OK", entity.Product{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/{id}/tsl").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Product的TSL信息").Handle(s.GetProductTsl)
	}).
		Doc("获取Product的TSL信息").
		Param(ws.PathParameter("id", "Id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", map[string]interface{}{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加Product信息").Handle(s.InsertProduct)
	}).
		Doc("添加Product信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.Product{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改Product信息").Handle(s.UpdateProduct)
	}).
		Doc("修改Product信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.Product{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除Product信息").Handle(s.DeleteProduct)
	}).
		Doc("删除Product信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
