package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/system/api"
	"pandax/apps/system/services"

	logServices "pandax/apps/log/services"
	"pandax/base/ginx"
)

func InitUserRouter(router *gin.RouterGroup) {
	s := &api.UserApi{
		RoleApp:     services.SysRoleModelDao,
		MenuApp:     services.SysMenuModelDao,
		RoleMenuApp: services.SysRoleMenuModelDao,
		UserApp:     services.SysUserModelDao,
		LogLogin:    logServices.LogLoginModelDao,
		DeptApp:     services.SysDeptModelDao,
		PostApp:     services.SysPostModelDao,
	}
	user := router.Group("user")
	// 获取验证码
	user.GET("getCaptcha", s.GenerateCaptcha)

	user.POST("login", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("登录").WithNeedToken(false).WithNeedCasbin(false).Handle(s.Login)
	})
	user.GET("auth", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("认证信息").WithNeedCasbin(false).Handle(s.Auth)
	})
	user.POST("logout", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("退出登录").WithNeedToken(false).WithNeedCasbin(false).Handle(s.LogOut)
	})

	user.GET("list", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("得到用户分页列表").Handle(s.GetSysUserList)
	})

	user.POST("avatar", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改用户头像").Handle(s.InsetSysUserAvatar)
	})

	user.PUT("pwd", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改用户密码").Handle(s.SysUserUpdatePwd)
	})

	user.GET("getById/:userId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取用户信息").Handle(s.GetSysUser)
	})

	user.GET("getInit", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取初始化角色岗位信息(添加用户初始化)").Handle(s.GetSysUserInit)
	})

	user.GET("getRoPo", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取用户角色岗位信息(添加用户初始化)").Handle(s.GetUserRolePost)
	})

	user.POST("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("添加用户信息").Handle(s.InsertSysUser)
	})

	user.PUT("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改用户信息").Handle(s.UpdateSysUser)
	})

	user.PUT("changeStatus", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改用户状态").Handle(s.UpdateSysUserStu)
	})

	user.DELETE(":userId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("删除用户信息").Handle(s.DeleteSysUser)
	})

	user.GET("export", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("导出用户信息").Handle(s.ExportUser)
	})
}
