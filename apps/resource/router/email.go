package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/resource/api"
	"pandax/apps/resource/services"
	"pandax/base/ctx"
)

/**
 * @Description 添加qq群467890197 交流学习
 * @Author 熊猫
 * @Date 2022/1/14 15:24
 **/

func InitResEmailsRouter(router *gin.RouterGroup) {
	s := &api.ResEmailsApi{
		ResEmailsApp: services.ResEmailsModelDao,
	}
	routerGroup := router.Group("email")

	ResEmailsListLog := ctx.NewLogInfo("获取ResEmails分页列表")
	routerGroup.GET("list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(ResEmailsListLog).Handle(s.GetResEmailsList)
	})

	ResEmailsLog := ctx.NewLogInfo("获取ResEmails信息")
	routerGroup.GET(":mailId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(ResEmailsLog).Handle(s.GetResEmails)
	})

	insertResEmailsLog := ctx.NewLogInfo("添加ResEmails信息")
	routerGroup.POST("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(insertResEmailsLog).Handle(s.InsertResEmails)
	})

	updateResEmailsLog := ctx.NewLogInfo("修改ResEmails信息")
	routerGroup.PUT("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateResEmailsLog).Handle(s.UpdateResEmails)
	})

	deleteResEmailsLog := ctx.NewLogInfo("删除ResEmails信息")
	routerGroup.DELETE(":mailId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteResEmailsLog).Handle(s.DeleteResEmails)
	})

	updateStatusEmailLog := ctx.NewLogInfo("修改状态")
	routerGroup.PUT("changeStatus", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateStatusEmailLog).Handle(s.UpdateMailStatus)
	})

	debugMailEmailLog := ctx.NewLogInfo("修改状态")
	routerGroup.POST("debugMail", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(debugMailEmailLog).Handle(s.DebugMail)
	})
}
