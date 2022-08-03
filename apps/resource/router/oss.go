package router

import (
	"github.com/emicklei/go-restful/v3"
)

/**
 * @Description 添加qq群467890197 交流学习
 * @Author 熊猫
 * @Date 2022/1/13 15:21
 **/

func InitResOssRouter(container *restful.Container) {
	/*s := &api.ResOssesApi{
		ResOssesApp: services.ResOssesModelDao,
	}
	routerGroup := router.Group("oss")

	routerGroup.GET("list", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("获取ResOsses分页列表").Handle(s.GetResOssesList)
	})

	routerGroup.GET(":ossId", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("获取ResOsses信息").Handle(s.GetResOsses)
	})

	routerGroup.POST("", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("添加ResOsses信息").Handle(s.InsertResOsses)
	})

	routerGroup.PUT("", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("修改ResOsses信息").Handle(s.UpdateResOsses)
	})

	routerGroup.DELETE(":ossId", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("删除ResOsses信息").Handle(s.DeleteResOsses)
	})

	routerGroup.POST("uploadFile", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("测试文件上传").Handle(s.UplaodResOsses)
	})

	routerGroup.PUT("/changeStatus", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("修改状态").Handle(s.UpdateOssStatus)
	})*/
}
