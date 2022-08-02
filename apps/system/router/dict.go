package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/system/api"
	"pandax/apps/system/services"
	"pandax/base/ginx"
)

func InitDictRouter(router *gin.RouterGroup) {
	s := &api.DictApi{
		DictType: services.SysDictTypeModelDao,
		DictData: services.SysDictDataModelDao,
	}
	dict := router.Group("dict")

	dict.GET("type/list", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取字典类型分页列表").Handle(s.GetDictTypeList)
	})

	dict.GET("type/:dictId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取字典类型信息").Handle(s.GetDictType)
	})

	dict.POST("type", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("添加字典类型信息").Handle(s.InsertDictType)
	})

	dict.PUT("type", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改字典类型信息").Handle(s.UpdateDictType)
	})

	dict.DELETE("type/:dictId", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("删除字典类型信息").Handle(s.DeleteDictType)
	})

	dict.GET("type/export", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("导出字典类型信息").Handle(s.ExportDictType)
	})

	dict.GET("data/list", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取字典数据分页列表").Handle(s.GetDictDataList)
	})

	dict.GET("data/type", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取字典数据列表通过字典类型").Handle(s.GetDictDataListByDictType)
	})

	dict.GET("data/:dictCode", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("获取字典数据信息").Handle(s.GetDictData)
	})

	dict.POST("data", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("添加字典数据信息").Handle(s.InsertDictData)
	})

	dict.PUT("data", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("修改字典数据信息").Handle(s.UpdateDictData)
	})

	dict.DELETE("data/:dictCode", func(c *gin.Context) {
		ginx.NewReqCtx(c).WithLog("删除字典数据信息").Handle(s.DeleteDictData)
	})

}
