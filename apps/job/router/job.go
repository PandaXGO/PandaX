package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/job/api"
	"pandax/apps/job/services"
	"pandax/base/ginx"
)

func InitJobRouter(router *gin.RouterGroup) {
	// 登录日志
	jobApi := &api.JobApi{
		JobApp: services.JobModelDao,
	}
	job := router.Group("")

	job.GET("list", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取Job列表").Handle(jobApi.GetJobList)
	})

	job.GET(":jobId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取Job信息").Handle(jobApi.GetJob)
	})

	job.POST("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("添加Job信息").Handle(jobApi.CreateJob)
	})

	job.PUT("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改Job信息").Handle(jobApi.UpdateJob)
	})

	job.DELETE(":jobId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("删除Job信息").Handle(jobApi.DeleteJob)
	})

	job.GET("/stop/:jobId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("停止一个job").Handle(jobApi.StopJobForService)
	})

	job.GET("/start/:jobId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("开启一个job").Handle(jobApi.StartJobForService)
	})

	job.PUT("/changeStatus", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改状态").Handle(jobApi.UpdateStatusJob)
	})
}
