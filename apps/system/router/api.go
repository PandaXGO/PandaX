package router

import (
	"github.com/XM-GO/PandaKit/casbin"
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/system/api"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
)

func InitApiRouter(container *restful.Container) {
	s := &api.SystemApiApi{
		ApiApp: services.SysApiModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/system/api").Produces(restful.MIME_JSON)
	tags := []string{"api"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取api分页列表").Handle(s.GetApiList)
	}).
		Doc("获取api分页列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("path", "路径").DataType("string")).
		Param(ws.QueryParameter("description", "描述").DataType("string")).
		Param(ws.QueryParameter("method", "方法").DataType("string")).
		Param(ws.QueryParameter("apiGroup", "API组").DataType("string")).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/all").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取所有api").Handle(s.GetAllApis)
	}).
		Doc("获取所有api").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.SysApi{}).
		Returns(200, "OK", []entity.SysApi{}))

	ws.Route(ws.GET("/getPolicyPathByRoleId").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取角色拥有的api权限").Handle(s.GetPolicyPathByRoleId)
	}).
		Doc("获取角色拥有的api权限").
		Param(ws.QueryParameter("roleKey", "校色key").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]casbin.CasbinRule{}).
		Returns(200, "OK", []casbin.CasbinRule{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取api信息").Handle(s.GetApiById)
	}).
		Doc("获取api信息").
		Param(ws.PathParameter("id", "Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.SysApi{}). // on the response
		Returns(200, "OK", entity.SysApi{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加api信息").Handle(s.CreateApi)
	}).
		Doc("添加api信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysApi{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改api信息").Handle(s.UpdateApi)
	}).
		Doc("修改api信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysApi{})) // from the request

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除api信息").Handle(s.DeleteApi)
	}).
		Doc("删除api信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "id").DataType("int")))

	container.Add(ws)
}
