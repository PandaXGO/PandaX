package router

import (
	"github.com/gin-gonic/gin"
	api2 "pandax/apps/system/api"
	services "pandax/apps/system/services"
	"pandax/base/ctx"
)

func InitApiRouter(router *gin.RouterGroup) {
	s := &api2.SystemApiApi{
		ApiApp: services.SysSysApiModelDao,
	}
	api := router.Group("api")

	apiListLog := ctx.NewLogInfo("获取api分页列表")
	api.GET("list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(apiListLog).Handle(s.GetApiList)
	})

	apiListAllLog := ctx.NewLogInfo("获取所有api")
	api.GET("all", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(apiListAllLog).Handle(s.GetAllApis)
	})

	apiListByRoleLog := ctx.NewLogInfo("获取角色拥有的api权限")
	api.GET("getPolicyPathByRoleId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(apiListByRoleLog).Handle(s.GetPolicyPathByRoleId)
	})

	apiLog := ctx.NewLogInfo("获取api信息")
	api.GET(":id", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(apiLog).Handle(s.GetApiById)
	})

	insertApiLog := ctx.NewLogInfo("添加api信息")
	api.POST("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(insertApiLog).Handle(s.CreateApi)
	})

	updateApiLog := ctx.NewLogInfo("修改api信息")
	api.PUT("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateApiLog).Handle(s.UpdateApi)
	})

	deleteApiLog := ctx.NewLogInfo("删除api信息")
	api.DELETE(":id", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteApiLog).Handle(s.DeleteApi)
	})
}
