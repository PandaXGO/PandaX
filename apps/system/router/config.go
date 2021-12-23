package router

import (
	"github.com/gin-gonic/gin"
	api2 "pandax/apps/system/api"
	services "pandax/apps/system/services"
	"pandax/base/ctx"
)

func InitConfigRouter(router *gin.RouterGroup) {
	s := &api2.ConfigApi{
		ConfigApp: services.SysSysConfigModelDao,
	}
	config := router.Group("config")

	configListLog := ctx.NewLogInfo("获取配置分页列表")
	config.GET("list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(configListLog).Handle(s.GetConfigList)
	})

	configListByKeyLog := ctx.NewLogInfo("获取配置列表通过ConfigKey")
	config.GET("configKey", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(configListByKeyLog).Handle(s.GetConfigListByKey)
	})

	configLog := ctx.NewLogInfo("获取配置信息")
	config.GET(":configId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(configLog).Handle(s.GetConfig)
	})

	insertConfigLog := ctx.NewLogInfo("添加配置信息")
	config.POST("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(insertConfigLog).Handle(s.InsertConfig)
	})

	updateConfigLog := ctx.NewLogInfo("修改配置信息")
	config.PUT("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateConfigLog).Handle(s.UpdateConfig)
	})

	deleteConfigLog := ctx.NewLogInfo("删除配置信息")
	config.DELETE(":configId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteConfigLog).Handle(s.DeleteConfig)
	})
}
