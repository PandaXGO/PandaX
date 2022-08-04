package router

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/system/api"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
)

func InitDictRouter(container *restful.Container) {
	s := &api.DictApi{
		DictType: services.SysDictTypeModelDao,
		DictData: services.SysDictDataModelDao,
	}
	ws := new(restful.WebService)
	ws.Path("/system/dict").Produces(restful.MIME_JSON)
	tags := []string{"dict"}

	ws.Route(ws.GET("/type/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取字典类型分页列表").Handle(s.GetDictTypeList)
	}).
		Doc("获取字典类型分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("dictName", "dictName").DataType("string")).
		Param(ws.QueryParameter("dictType", "dictType").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/type/{dictId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取字典类型信息").Handle(s.GetDictType)
	}).
		Doc("获取字典类型信息").
		Param(ws.PathParameter("dictId", "Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", entity.SysDictType{}))

	ws.Route(ws.POST("/type").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加字典类型信息").Handle(s.InsertDictType)
	}).
		Doc("添加字典类型信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysDictType{}))

	ws.Route(ws.PUT("/type").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改字典类型信息").Handle(s.UpdateDictType)
	}).
		Doc("修改字典类型信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysDictType{}))

	ws.Route(ws.DELETE("/type/{dictId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除字典类型信息").Handle(s.DeleteDictType)
	}).
		Doc("删除字典类型信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("dictId", "多id 1,2,3").DataType("string")))

	ws.Route(ws.GET("/type/export").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("导出字典类型信息").Handle(s.ExportDictType)
	}).
		Doc("导出字典类型信息").
		Param(ws.QueryParameter("filename", "filename").DataType("string")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Param(ws.QueryParameter("dictName", "dictName").DataType("string")).
		Param(ws.QueryParameter("dictType", "dictType").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/data/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取字典数据分页列表").Handle(s.GetDictDataList)
	}).
		Doc("获取字典数据分页列表").
		Param(ws.QueryParameter("dictLabel", "dictLabel").DataType("string")).
		Param(ws.QueryParameter("dictType", "dictType").DataType("string")).
		Param(ws.QueryParameter("status", "status").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", []entity.SysDictData{}))

	ws.Route(ws.GET("/data/type").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取字典数据列表通过字典类型").Handle(s.GetDictDataListByDictType)
	}).
		Doc("获取字典数据列表通过字典类型").
		Param(ws.QueryParameter("dictType", "dictType").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", []entity.SysDictData{}))

	ws.Route(ws.GET("/data/{dictCode}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取字典数据信息").Handle(s.GetDictData)
	}).
		Doc("获取字典数据信息").
		Param(ws.PathParameter("dictCode", "dictCode").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", entity.SysDictData{}))

	ws.Route(ws.POST("/data").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加字典数据信息").Handle(s.InsertDictData)
	}).
		Doc("添加字典数据信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysDictData{}))

	ws.Route(ws.PUT("/data").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改字典数据信息").Handle(s.UpdateDictData)
	}).
		Doc("修改字典数据信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysDictData{}))

	ws.Route(ws.DELETE("data/{dictCode}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除字典数据信息").Handle(s.DeleteDictData)
	}).
		Doc("删除字典数据信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("dictCode", "多id 1,2,3").DataType("string")))

	container.Add(ws)

}
