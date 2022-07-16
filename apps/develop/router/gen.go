package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/develop/api"
	"pandax/apps/develop/services"
	"pandax/base/ctx"
)

func InitGenRouter(router *gin.RouterGroup) {
	// 登录日志
	genApi := &api.GenApi{
		GenTableApp: services.DevGenTableModelDao,
	}
	gen := router.Group("gen")

	genViewLog := ctx.NewLogInfo("获取生成代码视图")
	gen.GET("preview/:tableId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(genViewLog).Handle(genApi.Preview)
	})

	genCodeLog := ctx.NewLogInfo("生成代码")
	gen.GET("code/:tableId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(genCodeLog).Handle(genApi.GenCode)
	})

	genConfigureLog := ctx.NewLogInfo("生成配置")
	gen.GET("configure/:tableId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(genConfigureLog).Handle(genApi.GenCode)
	})
}
