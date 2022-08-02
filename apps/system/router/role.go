package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/system/api"
	"pandax/apps/system/services"
	"pandax/base/ginx"
)

func InitRoleRouter(router *gin.RouterGroup) {
	s := &api.RoleApi{
		RoleApp:     services.SysRoleModelDao,
		RoleMenuApp: services.SysRoleMenuModelDao,
		RoleDeptApp: services.SysRoleDeptModelDao,
		UserApp:     services.SysUserModelDao,
	}
	role := router.Group("role")

	role.GET("list", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取角色分页列表").Handle(s.GetRoleList)
	})

	role.GET(":roleId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取角色信息").Handle(s.GetRole)
	})

	role.POST("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("添加角色信息").Handle(s.InsertRole)
	})

	role.PUT("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改角色信息").Handle(s.UpdateRole)
	})

	role.PUT("changeStatus", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改角色状态").Handle(s.UpdateRoleStatus)
	})

	role.PUT("dataScope", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改角色部门权限").Handle(s.UpdateRoleDataScope)
	})

	role.DELETE(":roleId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("删除角色信息").Handle(s.DeleteRole)
	})

	role.GET("export", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("导出角色信息").Handle(s.ExportRole)
	})
}
