package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/job/api"
	"pandax/apps/job/services"
	"pandax/base/ctx"
)

func InitJobRouter(router *gin.RouterGroup) {
	// 登录日志
	jobApi := &api.JobApi{
		JobApp: services.JobModelDao,
	}
	job := router.Group("")

	jobListLog := ctx.NewLogInfo("获取Job列表")
	job.GET("list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(jobListLog).Handle(jobApi.GetJobList)
	})

	getJobLog := ctx.NewLogInfo("获取Job信息")
	job.GET(":jobId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(getJobLog).Handle(jobApi.GetJob)
	})

	insertJobLog := ctx.NewLogInfo("添加Job信息")
	job.POST("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(insertJobLog).Handle(jobApi.CreateJob)
	})

	updateJobLog := ctx.NewLogInfo("修改Job信息")
	job.PUT("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateJobLog).Handle(jobApi.UpdateJob)
	})

	deleteJobLog := ctx.NewLogInfo("删除Job信息")
	job.DELETE(":jobId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteJobLog).Handle(jobApi.DeleteJob)
	})

	stopJobLog := ctx.NewLogInfo("停止一个job")
	job.GET("/stop/:jobId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(stopJobLog).Handle(jobApi.StopJobForService)
	})

	starteJobLog := ctx.NewLogInfo("开启一个job")
	job.GET("/start/:jobId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(starteJobLog).Handle(jobApi.StartJobForService)
	})
}
