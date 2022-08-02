package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/system/api"
	"pandax/apps/system/services"
	"pandax/base/ginx"
)

func InitMenuRouter(router *gin.RouterGroup) {
	s := &api.MenuApi{
		MenuApp:     services.SysMenuModelDao,
		RoleApp:     services.SysRoleModelDao,
		RoleMenuApp: services.SysRoleMenuModelDao,
		DeptApp:     services.SysDeptModelDao,
	}
	menu := router.Group("menu")

	menu.GET("menuTreeSelect", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取菜单树").WithNeedToken(false).WithNeedCasbin(false).Handle(s.GetMenuTreeSelect)
	})

	menu.GET("menuRole", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取角色菜单").Handle(s.GetMenuRole)
	})

	menu.GET("roleMenuTreeSelect/:roleId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取角色菜单树").Handle(s.GetMenuTreeRoleSelect)
	})

	menu.GET("menuPaths", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取角色菜单路径列表").Handle(s.GetMenuPaths)
	})

	menu.GET("list", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取菜单列表").Handle(s.GetMenuList)
	})

	menu.GET(":menuId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取菜单信息").Handle(s.GetMenu)
	})

	menu.POST("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("添加菜单信息").Handle(s.InsertMenu)
	})

	menu.PUT("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改菜单信息").Handle(s.UpdateMenu)
	})

	menu.DELETE(":menuId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("删除菜单信息").Handle(s.DeleteMenu)
	})
}
