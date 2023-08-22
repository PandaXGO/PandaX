package router

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/resource/api"
	"pandax/apps/resource/entity"
	"pandax/apps/resource/services"
)

/**
 * @Description 添加qq群467890197 交流学习
 * @Author 熊猫
 * @Date 2022/1/13 15:21
 **/

func InitResOssRouter(container *restful.Container) {
	s := &api.ResOssesApi{
		ResOssesApp: services.ResOssesModelDao,
	}
	ws := new(restful.WebService)
	ws.Path("/resource/oss").Produces(restful.MIME_JSON)
	tags := []string{"oss"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取ResOsses分页列表").Handle(s.GetResOssesList)
	}).
		Doc("获取ResOsses分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("category", "category").DataType("string")).
		Param(ws.QueryParameter("ossCode", "ossCode").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{ossId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取ResOsses信息").Handle(s.GetResOsses)
	}).
		Doc("获取ResOsses信息").
		Param(ws.PathParameter("ossId", "Id").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.ResOss{}). // on the response
		Returns(200, "OK", entity.ResOss{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加ResOsses信息").Handle(s.InsertResOsses)
	}).
		Doc("添加ResOsses信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.ResOss{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改ResOsses信息").Handle(s.UpdateResOsses)
	}).
		Doc("修改ResOsses信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.ResOss{})) // from the request

	ws.Route(ws.DELETE("/{ossId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除ResOsses信息").Handle(s.DeleteResOsses)
	}).
		Doc("删除ResOsses信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("ossId", "多id 1,2,3").DataType("string")))

	ws.Route(ws.POST("/uploadFile").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("测试文件上传").Handle(s.UplaodResOsses)
	}).
		Doc("测试文件上传").
		Param(ws.QueryParameter("ossCode", "ossCode").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/changeStatus").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改状态").Handle(s.UpdateOssStatus)
	}).
		Doc("修改状态").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	container.Add(ws)
}
