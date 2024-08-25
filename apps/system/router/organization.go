package router

import (
	"github.com/PandaXGO/PandaKit/restfulx"
	"pandax/apps/system/api"
	"pandax/apps/system/api/vo"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitOrganizationRouter(container *restful.Container) {
	s := &api.OrganizationApi{
		OrganizationApp: services.SysOrganizationModelDao,
		RoleApp:         services.SysRoleModelDao,
		UserApp:         services.SysUserModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/system/organization").Produces(restful.MIME_JSON)
	tags := []string{"system", "组织"}

	ws.Route(ws.GET("/roleOrganizationTreeSelect/{roleId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取角色组织树").Handle(s.GetOrganizationTreeRoleSelect)
	}).
		Doc("获取角色组织树").
		Param(ws.PathParameter("roleId", "角色Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(vo.OrganizationTreeVo{}).
		Returns(200, "OK", vo.OrganizationTreeVo{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/organizationTree").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取所有组织树").Handle(s.GetOrganizationTree)
	}).
		Doc("获取所有组织树").
		Param(ws.QueryParameter("organizationName", "organizationName").DataType("string")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("organizationId", "organizationId").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.SysOrganization{}).
		Returns(200, "OK", []entity.SysOrganization{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取组织列表").Handle(s.GetOrganizationList)
	}).
		Doc("获取组织列表").
		Param(ws.QueryParameter("organizationName", "organizationName").DataType("string")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("organizationId", "organizationId").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.SysOrganization{}).
		Returns(200, "OK", []entity.SysOrganization{}))

	ws.Route(ws.GET("/{organizationId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取组织信息").Handle(s.GetOrganization)
	}).
		Doc("获取组织信息").
		Param(ws.PathParameter("organizationId", "组织Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.SysOrganization{}). // on the response
		Returns(200, "OK", entity.SysOrganization{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加组织信息").Handle(s.InsertOrganization)
	}).
		Doc("添加组织信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysOrganization{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改组织信息").Handle(s.UpdateOrganization)
	}).
		Doc("修改组织信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysOrganization{}))

	ws.Route(ws.DELETE("/{organizationId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除组织信息").Handle(s.DeleteOrganization)
	}).
		Doc("删除组织信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("organizationId", "多id 1,2,3").DataType("int")))

	container.Add(ws)

}
