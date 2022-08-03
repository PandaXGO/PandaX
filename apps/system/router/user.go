package router

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/system/api"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"

	"github.com/XM-GO/PandaKit/restfulx"
	logServices "pandax/apps/log/services"
)

func InitUserRouter(container *restful.Container) {
	s := &api.UserApi{
		RoleApp:     services.SysRoleModelDao,
		MenuApp:     services.SysMenuModelDao,
		RoleMenuApp: services.SysRoleMenuModelDao,
		UserApp:     services.SysUserModelDao,
		LogLogin:    logServices.LogLoginModelDao,
		DeptApp:     services.SysDeptModelDao,
		PostApp:     services.SysPostModelDao,
	}
	ws := new(restful.WebService)
	ws.Path("/system/user").Produces(restful.MIME_JSON)
	tags := []string{"user"}

	ws.Route(ws.GET("/getCaptcha").To(s.GenerateCaptcha).Doc("获取验证码"))

	ws.Route(ws.POST("/login").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedToken(false).WithNeedCasbin(false).WithLog("登录").Handle(s.Login)
	}).
		Doc("登录").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysUser{})) // from the request

	ws.Route(ws.POST("/logout").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedToken(false).WithNeedCasbin(false).WithLog("退出登录").Handle(s.LogOut)
	}).
		Doc("退出登录").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/auth").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedCasbin(false).WithLog("认证信息").Handle(s.Auth)
	}).
		Doc("认证信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.SysUser{}).
		Returns(200, "OK", entity.SysUser{}))

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("得到用户分页列表").Handle(s.GetSysUserList)
	}).
		Doc("得到用户分页列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.SysUser{}).
		Returns(200, "OK", []entity.SysUser{}))

	ws.Route(ws.GET("/getById/{userId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取用户信息").Handle(s.GetSysUser)
	}).
		Doc("获取用户信息").
		Param(ws.PathParameter("userId", "Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.SysUser{}). // on the response
		Returns(200, "OK", entity.SysUser{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/getInit").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取初始化角色岗位信息(添加用户初始化)").Handle(s.GetSysUserInit)
	}).
		Doc("获取初始化角色岗位信息(添加用户初始化)").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.SysUser{}). // on the response
		Returns(200, "OK", entity.SysUser{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/getRoPo").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取用户角色岗位信息(添加用户初始化)").Handle(s.GetUserRolePost)
	}).
		Doc("获取用户角色岗位信息(添加用户初始化)").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", entity.SysUser{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加用户信息").Handle(s.InsertSysUser)
	}).
		Doc("添加用户信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysUser{})) // from the request

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改用户信息").Handle(s.UpdateSysUser)
	}).
		Doc("修改用户信息").
		Metadata(restfulspec.KeyOpenAPITags, tags)) // from the request

	ws.Route(ws.PUT("/changeStatus").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改用户状态").Handle(s.UpdateSysUserStu)
	}).
		Doc("修改用户状态").
		Metadata(restfulspec.KeyOpenAPITags, tags)) // from the request

	ws.Route(ws.DELETE("/{userId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除用户信息").Handle(s.DeleteSysUser)
	}).
		Doc("删除用户信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("userId", "多id 1,2,3").DataType("string")))

	ws.Route(ws.POST("/avatar").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改用户头像").Handle(s.InsetSysUserAvatar)
	}).
		Doc("修改用户头像").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("pwd").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改用户密码").Handle(s.SysUserUpdatePwd)
	}).
		Doc("修改用户密码").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/export").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("导出用户信息").Handle(s.ExportUser)
	}).
		Doc("导出用户信息").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	container.Add(ws)

}
