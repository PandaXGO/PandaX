package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/develop/api"
	"pandax/apps/develop/services"
	"pandax/base/ctx"
)

func InitGenTableRouter(router *gin.RouterGroup) {
	// 登录日志
	genApi := &api.GenTableApi{
		GenTableApp: services.DevGenTableModelDao,
	}
	gen := router.Group("table")

	genDbListLog := ctx.NewLogInfo("获取数据库列表")
	gen.GET("/db/list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(genDbListLog).Handle(genApi.GetDBTableList)
	})

	genListLog := ctx.NewLogInfo("获取表列表")
	gen.GET("list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(genListLog).Handle(genApi.GetTablePage)
	})

	genInfoNameLog := ctx.NewLogInfo("获取表信息By tableName")
	gen.GET("/info/tableName", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(genInfoNameLog).Handle(genApi.GetTableInfoByName)
	})

	genInfoLog := ctx.NewLogInfo("获取表信息")
	gen.GET("info/:tableId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(genInfoLog).Handle(genApi.GetTableInfo)
	})

	genTreeLog := ctx.NewLogInfo("获取表树")
	gen.GET("tableTree", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(genTreeLog).Handle(genApi.GetTableTree)
	})

	genInsterLog := ctx.NewLogInfo("新增表")
	gen.POST("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(genInsterLog).Handle(genApi.Insert)
	})

	genUpdateLog := ctx.NewLogInfo("修改表")
	gen.PUT("", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(genUpdateLog).Handle(genApi.Update)
	})

	genDeleteLog := ctx.NewLogInfo("删除表")
	gen.DELETE(":tableId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(genDeleteLog).Handle(genApi.Delete)
	})
}
