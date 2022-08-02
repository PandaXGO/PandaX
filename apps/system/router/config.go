package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/system/api"
	"pandax/apps/system/services"
	"pandax/base/ginx"
)

func InitConfigRouter(router *gin.RouterGroup) {
	s := &api.ConfigApi{
		ConfigApp: services.SysSysConfigModelDao,
	}
	config := router.Group("config")

	config.GET("list", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取配置分页列表").Handle(s.GetConfigList)
	})

	config.GET("configKey", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取配置列表通过ConfigKey").Handle(s.GetConfigListByKey)
	})

	config.GET(":configId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取配置信息").Handle(s.GetConfig)
	})

	config.POST("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("添加配置信息").Handle(s.InsertConfig)
	})

	config.PUT("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改配置信息").Handle(s.UpdateConfig)
	})

	config.DELETE(":configId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("删除配置信息").Handle(s.DeleteConfig)
	})
}
