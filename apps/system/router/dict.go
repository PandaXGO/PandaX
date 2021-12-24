package router

import (
	"github.com/gin-gonic/gin"
	"pandax/apps/system/api"
	"pandax/apps/system/services"
	"pandax/base/ctx"
)

func InitDictRouter(router *gin.RouterGroup) {
	s := &api.DictApi{
		DictType: services.SysDictTypeModelDao,
		DictData: services.SysDictDataModelDao,
	}
	dict := router.Group("dict")

	typeList := ctx.NewLogInfo("获取字典类型分页列表")
	dict.GET("type/list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(typeList).Handle(s.GetDictTypeList)
	})

	typeLog := ctx.NewLogInfo("获取字典类型信息")
	dict.GET("type/:dictId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(typeLog).Handle(s.GetDictType)
	})

	insertTypeLog := ctx.NewLogInfo("添加字典类型信息")
	dict.POST("type", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(insertTypeLog).Handle(s.InsertDictType)
	})

	updateTypeLog := ctx.NewLogInfo("修改字典类型信息")
	dict.PUT("type", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateTypeLog).Handle(s.UpdateDictType)
	})

	deleteTypeLog := ctx.NewLogInfo("删除字典类型信息")
	dict.DELETE("type/:dictId", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteTypeLog).Handle(s.DeleteDictType)
	})

	exportTypeLog := ctx.NewLogInfo("导出字典类型信息")
	dict.GET("type/export", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(exportTypeLog).Handle(s.ExportDictType)
	})

	dataListLog := ctx.NewLogInfo("获取字典数据分页列表")
	dict.GET("data/list", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(dataListLog).Handle(s.GetDictDataList)
	})

	dataListByDictTypeLog := ctx.NewLogInfo("获取字典数据列表通过字典类型")
	dict.GET("data/type", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(dataListByDictTypeLog).Handle(s.GetDictDataListByDictType)
	})

	dataLog := ctx.NewLogInfo("获取字典数据信息")
	dict.GET("data/:dictCode", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(dataLog).Handle(s.GetDictData)
	})

	insertDataLog := ctx.NewLogInfo("添加字典数据信息")
	dict.POST("data", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(insertDataLog).Handle(s.InsertDictData)
	})

	updateDataLog := ctx.NewLogInfo("修改字典数据信息")
	dict.PUT("data", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(updateDataLog).Handle(s.UpdateDictData)
	})

	deleteDataLog := ctx.NewLogInfo("删除字典数据信息")
	dict.DELETE("data/:dictCode", func(c *gin.Context) {
		ctx.NewReqCtxWithGin(c).WithLog(deleteDataLog).Handle(s.DeleteDictData)
	})

}
