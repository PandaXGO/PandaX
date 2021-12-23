package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/log/api"
	"pandax/apps/log/services"
	"pandax/base/ctx"
)

func InitLogRouter(router *gin.RouterGroup) {
	// 登录日志
	login := &api.LogLoginApi{
		LogLoginApp: services.LogLoginModelDao,
	}
	logLogin := router.Group("logLogin")

	logLoginListLog := ctx.NewLogInfo("获取登录日志列表")
	logLogin.GET("list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(logLoginListLog).Handle(login.GetLoginLogList)
	})

	getLogLoginLog := ctx.NewLogInfo("获取登录日志信息")
	logLogin.GET(":infoId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(getLogLoginLog).Handle(login.GetLoginLog)
	})

	updateLogLoginLog := ctx.NewLogInfo("修改登录日志信息")
	logLogin.PUT("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateLogLoginLog).Handle(login.UpdateLoginLog)
	})

	deleteLogLoginAllLog := ctx.NewLogInfo("删除登录日志信息")
	logLogin.DELETE("all", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteLogLoginAllLog).Handle(login.DeleteAll)
	})

	deleteLogLoginLog := ctx.NewLogInfo("删除登录日志信息")
	logLogin.DELETE(":infoId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteLogLoginLog).Handle(login.DeleteLoginLog)
	})

	// 操作日志
	oper := &api.LogOperApi{
		LogOperApp: services.LogOperModelDao,
	}
	logOper := router.Group("logOper")

	logOperListLog := ctx.NewLogInfo("获取操作日志列表")
	logOper.GET("list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(logOperListLog).Handle(oper.GetOperLogList)
	})

	getLogOperLog := ctx.NewLogInfo("获取操作日志信息")
	logOper.GET(":operId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(getLogOperLog).Handle(oper.GetOperLog)
	})

	deleteLogOperAllLog := ctx.NewLogInfo("删除操作日志信息")
	logOper.DELETE("all", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteLogOperAllLog).Handle(oper.DeleteAll)
	})

	deleteLogOperLog := ctx.NewLogInfo("删除操作日志信息")
	logOper.DELETE(":operId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteLogOperLog).Handle(oper.DeleteOperLog)
	})
}
