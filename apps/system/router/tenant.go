package router

/**
 * @Description
 * @Author 熊猫
 * @Date 2022/7/14 17:52
 **/

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/system/api"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
)

func InitSysTenantRouter(container *restful.Container) {
	s := &api.SysTenantsApi{
		SysTenantsApp: services.SysTenantModelDao,
	}
	ws := new(restful.WebService)
	ws.Path("/system/tenant").Produces(restful.MIME_JSON)
	tags := []string{"tenant"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取SysTenant分页列表").Handle(s.GetSysTenantsList)
	}).
		Doc("获取SysTenant分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/lists").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取SysTenant列表").Handle(s.GetSysTenantsAll)
	}).
		Doc("获取SysTenant列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.SysTenants{}).
		Returns(200, "OK", []entity.SysTenants{}))

	ws.Route(ws.GET("/{tenantId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取SysTenant信息").Handle(s.GetSysTenants)
	}).
		Doc("获取SysTenant信息").
		Param(ws.PathParameter("tenantId", "租户Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.SysTenants{}).
		Returns(200, "OK", entity.SysTenants{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加SysTenant信息").Handle(s.InsertSysTenants)
	}).
		Doc("添加SysTenant信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysTenants{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改SysTenant信息").Handle(s.UpdateSysTenants)
	}).
		Doc("修改SysTenant信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysTenants{}))

	ws.Route(ws.DELETE("/{tenantId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除SysTenant信息").Handle(s.DeleteSysTenants)
	}).
		Doc("删除SysTenant信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("tenantId", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
