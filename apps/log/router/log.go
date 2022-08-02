package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/log/api"
	"pandax/apps/log/services"
	"pandax/base/ginx"
)

func InitLogRouter(router *gin.RouterGroup) {
	// 登录日志
	login := &api.LogLoginApi{
		LogLoginApp: services.LogLoginModelDao,
	}
	logLogin := router.Group("logLogin")

	logLogin.GET("list", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取登录日志列表").Handle(login.GetLoginLogList)
	})

	logLogin.GET(":infoId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取登录日志信息").Handle(login.GetLoginLog)
	})

	logLogin.PUT("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改登录日志信息").Handle(login.UpdateLoginLog)
	})

	logLogin.DELETE("all", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("清空登录日志信息").Handle(login.DeleteAll)
	})

	logLogin.DELETE(":infoId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("删除登录日志信息").Handle(login.DeleteLoginLog)
	})

	// 操作日志
	oper := &api.LogOperApi{
		LogOperApp: services.LogOperModelDao,
	}
	logOper := router.Group("logOper")

	logOper.GET("list", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取操作日志列表").Handle(oper.GetOperLogList)
	})

	logOper.GET(":operId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取操作日志信息").Handle(oper.GetOperLog)
	})

	logOper.DELETE("all", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("清空操作日志信息").Handle(oper.DeleteAll)
	})

	logOper.DELETE(":operId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("删除操作日志信息").Handle(oper.DeleteOperLog)
	})

	// Job日志
	job := &api.LogJobApi{
		LogJobApp: services.LogJobModelDao,
	}
	logJob := router.Group("logJob")

	logJob.GET("list", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取操作日志列表").Handle(job.GetJobLogList)
	})

	logJob.DELETE("all", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("清空操作日志信息").Handle(job.DeleteAll)
	})

	logJob.DELETE(":logId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("删除操作日志信息").Handle(job.DeleteJobLog)
	})
}
