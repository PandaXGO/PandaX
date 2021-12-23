package router

import (
	"github.com/gin-gonic/gin"
	api2 "pandax/apps/system/api"
	services2 "pandax/apps/system/services"
	"pandax/base/ctx"
)

func InitMenuRouter(router *gin.RouterGroup) {
	s := &api2.MenuApi{
		MenuApp:     services2.SysMenuModelDao,
		RoleApp:     services2.SysRoleModelDao,
		RoleMenuApp: services2.SysRoleMenuModelDao,
		DeptApp:     services2.SysDeptModelDao,
	}
	menu := router.Group("menu")

	getSetRole := ctx.NewLogInfo("获取菜单树")
	menu.GET("menuTreeSelect", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(getSetRole).WithNeedToken(false).WithNeedCasbin(false).Handle(s.GetMenuTreeSelect)
	})

	menuRole := ctx.NewLogInfo("获取角色菜单")
	menu.GET("menuRole", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(menuRole).Handle(s.GetMenuRole)
	})

	roleMenuTreSelect := ctx.NewLogInfo("获取角色菜单树")
	menu.GET("roleMenuTreeSelect/:roleId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(roleMenuTreSelect).Handle(s.GetMenuTreeRoleSelect)
	})

	menuPaths := ctx.NewLogInfo("获取角色菜单路径列表")
	menu.GET("menuPaths", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(menuPaths).Handle(s.GetMenuPaths)
	})

	menuList := ctx.NewLogInfo("获取菜单列表")
	menu.GET("list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(menuList).Handle(s.GetMenuList)
	})

	menuLog := ctx.NewLogInfo("获取菜单信息")
	menu.GET(":menuId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(menuLog).Handle(s.GetMenu)
	})

	insertMenuLog := ctx.NewLogInfo("添加菜单信息")
	menu.POST("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(insertMenuLog).Handle(s.InsertMenu)
	})

	updateMenuLog := ctx.NewLogInfo("修改菜单信息")
	menu.PUT("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateMenuLog).Handle(s.UpdateMenu)
	})

	deleteMenuLog := ctx.NewLogInfo("删除菜单信息")
	menu.DELETE(":menuId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteMenuLog).Handle(s.DeleteMenu)
	})
}
