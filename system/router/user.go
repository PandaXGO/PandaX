package router

import (
	"github.com/gin-gonic/gin"
	"pandax/base/ctx"
	"pandax/system/api"
	"pandax/system/services"
)

func InitUserRouter(router *gin.RouterGroup) {
	s := &api.UserApi{
		RoleApp:     services.SysRoleModelDao,
		MenuApp:     services.SysMenuModelDao,
		RoleMenuApp: services.SysRoleMenuModelDao,
		UserApp:     services.SysUserModelDao,
		LogLogin:    services.LogLoginModelDao,
		DeptApp:     services.SysDeptModelDao,
		PostApp:     services.SysPostModelDao,
	}
	user := router.Group("user")
	// 获取验证码
	user.GET("getCaptcha", s.GenerateCaptcha)

	loginLog := ctx.NewLogInfo("登录")
	user.POST("login", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(loginLog).WithNeedToken(false).WithNeedCasbin(false).Handle(s.Login)
	})

	logoutLog := ctx.NewLogInfo("退出登录")
	user.POST("logout", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(logoutLog).WithNeedToken(false).WithNeedCasbin(false).Handle(s.LogOut)
	})

	sysUserListLog := ctx.NewLogInfo("得到用户分页列表")
	user.GET("list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(sysUserListLog).Handle(s.GetSysUserList)
	})

	avatarLog := ctx.NewLogInfo("修改用户头像")
	user.POST("avatar", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(avatarLog).Handle(s.InsetSysUserAvatar)
	})

	pwdLog := ctx.NewLogInfo("修改用户密码")
	user.PUT("pwd", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(pwdLog).Handle(s.SysUserUpdatePwd)
	})

	userLog := ctx.NewLogInfo("获取用户信息")
	user.GET("getById/:userId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(userLog).Handle(s.GetSysUser)
	})

	getSysUserInitLog := ctx.NewLogInfo("获取初始化角色岗位信息(添加用户初始化)")
	user.GET("getInit", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(getSysUserInitLog).Handle(s.GetSysUserInit)
	})

	getSysUserRoPoLog := ctx.NewLogInfo("获取用户角色岗位信息(添加用户初始化)")
	user.GET("getRoPo", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(getSysUserRoPoLog).Handle(s.GetUserRolePost)
	})

	insertUserLog := ctx.NewLogInfo("添加用户信息")
	user.POST("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(insertUserLog).Handle(s.InsertSysUser)
	})

	updateUserLog := ctx.NewLogInfo("修改用户信息")
	user.PUT("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateUserLog).Handle(s.UpdateSysUser)
	})

	updateUserStuLog := ctx.NewLogInfo("修改用户状态")
	user.PUT("changeStatus", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateUserStuLog).Handle(s.UpdateSysUserStu)
	})

	deleteUserLog := ctx.NewLogInfo("删除用户信息")
	user.DELETE(":userId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteUserLog).Handle(s.DeleteSysUser)
	})

	exportUserLog := ctx.NewLogInfo("导出用户信息")
	user.GET("export", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(exportUserLog).Handle(s.ExportUser)
	})
}
