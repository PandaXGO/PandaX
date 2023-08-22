package router

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/system/api"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
)

func InitConfigRouter(container *restful.Container) {
	s := &api.ConfigApi{
		ConfigApp: services.SysSysConfigModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/system/config").Produces(restful.MIME_JSON)
	tags := []string{"config"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取配置分页列表").Handle(s.GetConfigList)
	}).
		Doc("获取配置分页列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("configName", "configName").DataType("string")).
		Param(ws.QueryParameter("configKey", "configKey").DataType("string")).
		Param(ws.QueryParameter("configType", "configType").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/configKey").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取配置列表通过ConfigKey").Handle(s.GetConfigListByKey)
	}).
		Doc("获取配置列表通过ConfigKey").
		Param(ws.QueryParameter("configKey", "configKey").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.SysConfig{}).
		Returns(200, "OK", []entity.SysConfig{}))

	ws.Route(ws.GET("/{configId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取配置信息").Handle(s.GetConfig)
	}).
		Doc("获取配置信息").
		Param(ws.PathParameter("configId", "configId").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.SysConfig{}).
		Returns(200, "OK", entity.SysConfig{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加配置信息").Handle(s.InsertConfig)
	}).
		Doc("添加配置信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysConfig{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改配置信息").Handle(s.UpdateConfig)
	}).
		Doc("修改配置信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysConfig{})) // from the request

	ws.Route(ws.DELETE("/{configId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除配置信息").Handle(s.DeleteConfig)
	}).
		Doc("删除配置信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("configId", "多id 1,2,3").DataType("string")))

	container.Add(ws)

}
