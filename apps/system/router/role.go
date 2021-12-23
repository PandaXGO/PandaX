package router

import (
	"github.com/gin-gonic/gin"
	api2 "pandax/apps/system/api"
	services2 "pandax/apps/system/services"
	"pandax/base/ctx"
)

func InitRoleRouter(router *gin.RouterGroup) {
	s := &api2.RoleApi{
		RoleApp:     services2.SysRoleModelDao,
		RoleMenuApp: services2.SysRoleMenuModelDao,
		RoleDeptApp: services2.SysRoleDeptModelDao,
		UserApp:     services2.SysUserModelDao,
	}
	role := router.Group("role")

	roleListLog := ctx.NewLogInfo("获取角色分页列表")
	role.GET("list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(roleListLog).Handle(s.GetRoleList)
	})

	roleLog := ctx.NewLogInfo("获取角色信息")
	role.GET(":roleId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(roleLog).Handle(s.GetRole)
	})

	insertRoleLog := ctx.NewLogInfo("添加角色信息")
	role.POST("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(insertRoleLog).Handle(s.InsertRole)
	})

	updateRoleLog := ctx.NewLogInfo("修改角色信息")
	role.PUT("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateRoleLog).Handle(s.UpdateRole)
	})
	updateStaRoleLog := ctx.NewLogInfo("修改角色状态")
	role.PUT("changeStatus", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateStaRoleLog).Handle(s.UpdateRoleStatus)
	})

	updateDaSRoleLog := ctx.NewLogInfo("修改角色部门权限")
	role.PUT("dataScope", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateDaSRoleLog).Handle(s.UpdateRoleDataScope)
	})

	deleteRoleLog := ctx.NewLogInfo("删除角色信息")
	role.DELETE(":roleId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteRoleLog).Handle(s.DeleteRole)
	})

	exportRoleLog := ctx.NewLogInfo("导出角色信息")
	role.GET("export", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(exportRoleLog).Handle(s.ExportRole)
	})
}
