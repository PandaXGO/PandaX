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
 * @Date 2022/1/13 15:21
 **/

func InitResOssRouter(router *gin.RouterGroup) {
	s := &api.ResOssesApi{
		ResOssesApp: services.ResOssesModelDao,
	}
	routerGroup := router.Group("oss")

	ResOssesListLog := ctx.NewLogInfo("获取ResOsses分页列表")
	routerGroup.GET("list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(ResOssesListLog).Handle(s.GetResOssesList)
	})

	ResOssesLog := ctx.NewLogInfo("获取ResOsses信息")
	routerGroup.GET(":ossId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(ResOssesLog).Handle(s.GetResOsses)
	})

	insertResOssesLog := ctx.NewLogInfo("添加ResOsses信息")
	routerGroup.POST("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(insertResOssesLog).Handle(s.InsertResOsses)
	})

	updateResOssesLog := ctx.NewLogInfo("修改ResOsses信息")
	routerGroup.PUT("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateResOssesLog).Handle(s.UpdateResOsses)
	})

	deleteResOssesLog := ctx.NewLogInfo("删除ResOsses信息")
	routerGroup.DELETE(":ossId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteResOssesLog).Handle(s.DeleteResOsses)
	})
}
