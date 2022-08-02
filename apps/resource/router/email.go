package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/resource/api"
	"pandax/apps/resource/services"
	"pandax/base/ginx"
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

	routerGroup.GET("list", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取ResEmails分页列表").Handle(s.GetResEmailsList)
	})

	routerGroup.GET(":mailId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取ResEmails信息").Handle(s.GetResEmails)
	})

	routerGroup.POST("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("添加ResEmails信息").Handle(s.InsertResEmails)
	})

	routerGroup.PUT("", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改ResEmails信息").Handle(s.UpdateResEmails)
	})

	routerGroup.DELETE(":mailId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("删除ResEmails信息").Handle(s.DeleteResEmails)
	})

	routerGroup.PUT("changeStatus", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改状态").Handle(s.UpdateMailStatus)
	})

	routerGroup.POST("debugMail", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改状态").Handle(s.DebugMail)
	})
}
