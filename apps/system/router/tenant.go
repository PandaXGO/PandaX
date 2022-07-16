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
	"pandax/base/ctx"
)

func InitSysTenantRouter(router *gin.RouterGroup) {
	s := &api.SysTenantsApi{
		SysTenantsApp: services.SysTenantModelDao,
	}
	routerGroup := router.Group("tenant")

	SysTenantListLog := ctx.NewLogInfo("获取SysTenant分页列表")
	routerGroup.GET("list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(SysTenantListLog).Handle(s.GetSysTenantsList)
	})

	SysTenantLog := ctx.NewLogInfo("获取SysTenant信息")
	routerGroup.GET(":tenantId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(SysTenantLog).Handle(s.GetSysTenants)
	})

	insertSysTenantLog := ctx.NewLogInfo("添加SysTenant信息")
	routerGroup.POST("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(insertSysTenantLog).Handle(s.InsertSysTenants)
	})

	updateSysTenantLog := ctx.NewLogInfo("修改SysTenant信息")
	routerGroup.PUT("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateSysTenantLog).Handle(s.UpdateSysTenants)
	})

	deleteSysTenantLog := ctx.NewLogInfo("删除SysTenant信息")
	routerGroup.DELETE(":tenantId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteSysTenantLog).Handle(s.DeleteSysTenants)
	})
}
