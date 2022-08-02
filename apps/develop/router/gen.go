package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/develop/api"
	"pandax/apps/develop/services"
	"pandax/base/ginx"
)

func InitGenRouter(router *gin.RouterGroup) {
	// 登录日志
	genApi := &api.GenApi{
		GenTableApp: services.DevGenTableModelDao,
	}
	gen := router.Group("gen")

	gen.GET("preview/:tableId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取生成代码视图").Handle(genApi.Preview)
	})

	gen.GET("code/:tableId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("生成代码").Handle(genApi.GenCode)
	})

	gen.GET("configure/:tableId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("生成配置").Handle(genApi.GenConfigure)
	})
}
