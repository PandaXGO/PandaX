package router

import (
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/system/api"
	"pandax/apps/system/api/vo"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
)

func InitMenuRouter(container *restful.Container) {
	s := &api.MenuApi{
		MenuApp:     services.SysMenuModelDao,
		RoleApp:     services.SysRoleModelDao,
		RoleMenuApp: services.SysRoleMenuModelDao,
		DeptApp:     services.SysDeptModelDao,
	}
	ws := new(restful.WebService)
	ws.Path("/system/menu").Produces(restful.MIME_JSON)
	tags := []string{"menu"}

	ws.Route(ws.GET("/menuTreeSelect").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取菜单树").WithNeedToken(false).WithNeedCasbin(false).Handle(s.GetMenuTreeSelect)
	}).
		Doc("获取菜单树").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.MenuLable{}).
		Returns(200, "OK", []entity.MenuLable{}))

	ws.Route(ws.GET("/menuRole").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取角色菜单").Handle(s.GetMenuRole)
	}).
		Doc("获取角色菜单").
		Param(ws.QueryParameter("roleKey", "roleKey").Required(true).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]vo.RouterVo{}).
		Returns(200, "OK", []vo.RouterVo{}))

	ws.Route(ws.GET("/roleMenuTreeSelect/{roleId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取角色菜单树").Handle(s.GetMenuTreeRoleSelect)
	}).
		Doc("获取角色菜单树").
		Param(ws.PathParameter("roleId", "Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(vo.MenuTreeVo{}).
		Returns(200, "OK", vo.MenuTreeVo{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/menuPaths").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取角色菜单路径列表").Handle(s.GetMenuPaths)
	}).
		Doc("获取角色菜单").
		Param(ws.QueryParameter("roleKey", "roleKey").Required(true).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.MenuPath{}).
		Returns(200, "OK", []entity.MenuPath{}))

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取菜单列表").Handle(s.GetMenuList)
	}).
		Doc("获取菜单列表").
		Param(ws.QueryParameter("menuName", "menuName").DataType("string")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.SysMenu{}).
		Returns(200, "OK", []entity.SysMenu{}))

	ws.Route(ws.GET("/{menuId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取菜单信息").Handle(s.GetMenu)
	}).
		Doc("获取菜单信息").
		Param(ws.PathParameter("menuId", "Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.SysMenu{}). // on the response
		Returns(200, "OK", entity.SysMenu{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加菜单信息").Handle(s.InsertMenu)
	}).
		Doc("添加菜单信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysMenu{})) // from the request

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改菜单信息").Handle(s.UpdateMenu)
	}).
		Doc("修改菜单信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysMenu{})) // from the request

	ws.Route(ws.DELETE("/{menuId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除菜单信息").Handle(s.DeleteMenu)
	}).
		Doc("删除SysTenant信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("menuId", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
