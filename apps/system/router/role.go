package router

import (
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/system/api"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
)

func InitRoleRouter(container *restful.Container) {
	s := &api.RoleApi{
		RoleApp:     services.SysRoleModelDao,
		RoleMenuApp: services.SysRoleMenuModelDao,
		RoleDeptApp: services.SysRoleDeptModelDao,
		UserApp:     services.SysUserModelDao,
	}
	ws := new(restful.WebService)
	ws.Path("/system/role").Produces(restful.MIME_JSON)
	tags := []string{"role"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取角色分页列表").Handle(s.GetRoleList)
	}).
		Doc("获取角色分页列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.SysRole{}).
		Returns(200, "OK", []entity.SysRole{}))

	ws.Route(ws.GET("/{roleId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取角色信息").Handle(s.GetRole)
	}).
		Doc("获取角色信息").
		Param(ws.PathParameter("roleId", "Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.SysRole{}). // on the response
		Returns(200, "OK", entity.SysRole{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加角色信息").Handle(s.InsertRole)
	}).
		Doc("添加角色信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysRole{})) // from the request

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改角色信息").Handle(s.UpdateRole)
	}).
		Doc("修改角色信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysRole{}))

	ws.Route(ws.PUT("/changeStatus").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改角色状态").Handle(s.UpdateRoleStatus)
	}).
		Doc("修改角色状态").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysRole{}))

	ws.Route(ws.PUT("/dataScope").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改角色部门权限").Handle(s.UpdateRoleDataScope)
	}).
		Doc("修改角色部门权限").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysRole{}))

	ws.Route(ws.DELETE("/{roleId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除角色信息").Handle(s.DeleteRole)
	}).
		Doc("删除角色信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("roleId", "多id 1,2,3").DataType("string")))

	ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("导出角色信息").Handle(s.ExportRole)
	}).
		Doc("导出角色信息").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	container.Add(ws)
}
