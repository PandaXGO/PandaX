package router

import (
	"github.com/emicklei/go-restful/v3"
)

func InitGenRouter(container *restful.Container) {
	// 登录日志
	/*	genApi := &api.GenApi{
			GenTableApp: services.DevGenTableModelDao,
		}
		gen := router.Group("gen")

		gen.GET("preview/:tableId", func(c *gin.Context) {
			restfulx.NewReqCtx(c).WithLog("获取生成代码视图").Handle(genApi.Preview)
		})

		gen.GET("code/:tableId", func(c *gin.Context) {
			restfulx.NewReqCtx(c).WithLog("生成代码").Handle(genApi.GenCode)
		})

		gen.GET("configure/:tableId", func(c *gin.Context) {
			restfulx.NewReqCtx(c).WithLog("生成配置").Handle(genApi.GenConfigure)
		})*/
}
