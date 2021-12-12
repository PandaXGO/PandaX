package router

import (
	"github.com/gin-gonic/gin"
	"pandax/base/ctx"
	"pandax/system/api"
	"pandax/system/services"
)

func InitSettingRouter(router *gin.RouterGroup) {
	s := &api.SettingApi{
		SettingApp: services.SysSettingsModelDao,
	}
	setting := router.Group("setting")

	getSetLog := ctx.NewLogInfo("获取配置信息")
	setting.GET("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(getSetLog).WithNeedToken(false).Handle(s.GetSettingsInfo)
	})

	setSetLog := ctx.NewLogInfo("设置配置信息")
	setting.POST("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(setSetLog).Handle(s.SetSettingsInfo)
	})
}
