package router

/**
 * @Description
 * @Author 熊猫
 * @Date 2022/7/14 17:52
 **/

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/system/api"
	"pandax/apps/system/services"
	"pandax/base/ginx"
)

func InitSysTenantRouter(router *gin.RouterGroup) {
	s := &api.SysTenantsApi{
		SysTenantsApp: services.SysTenantModelDao,
	}
	routerGroup := router.Group("tenant")

	routerGroup.GET("list", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取SysTenant分页列表").Handle(s.GetSysTenantsList)
	})

	routerGroup.GET("lists", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取SysTenant列表").Handle(s.GetSysTenantsAll)
	})

	routerGroup.GET(":tenantId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取SysTenant信息").Handle(s.GetSysTenants)
	})

	routerGroup.POST("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("添加SysTenant信息").Handle(s.InsertSysTenants)
	})

	routerGroup.PUT("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改SysTenant信息").Handle(s.UpdateSysTenants)
	})

	routerGroup.DELETE(":tenantId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("删除SysTenant信息").Handle(s.DeleteSysTenants)
	})
}
