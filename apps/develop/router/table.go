package router

import (
	"github.com/emicklei/go-restful/v3"
)

func InitGenTableRouter(container *restful.Container) {
	// 登录日志
	/*genApi := &api.GenTableApi{
		GenTableApp: services.DevGenTableModelDao,
	}
	gen := router.Group("table")

	gen.GET("/db/list", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("获取数据库列表").Handle(genApi.GetDBTableList)
	})

	gen.GET("list", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("获取表列表").Handle(genApi.GetTablePage)
	})

	gen.GET("/info/tableName", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("获取表信息By tableName").Handle(genApi.GetTableInfoByName)
	})

	gen.GET("info/:tableId", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("获取表信息").Handle(genApi.GetTableInfo)
	})

	gen.GET("tableTree", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("获取表树").Handle(genApi.GetTableTree)
	})

	gen.POST("", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("新增表").Handle(genApi.Insert)
	})

	gen.PUT("", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("修改表").Handle(genApi.Update)
	})

	gen.DELETE(":tableId", func(c *gin.Context) {
		restfulx.NewReqCtx(c).WithLog("删除表").Handle(genApi.Delete)
	})*/
}
