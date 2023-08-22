package router

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/develop/api"
	"pandax/apps/develop/api/vo"
	"pandax/apps/develop/entity"
	"pandax/apps/develop/services"
)

func InitGenTableRouter(container *restful.Container) {
	// 登录日志
	s := &api.GenTableApi{
		GenTableApp: services.DevGenTableModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/develop/code/table").Produces(restful.MIME_JSON)
	tags := []string{"codetable"}

	ws.Route(ws.GET("/db/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取数据库列表").Handle(s.GetDBTableList)
	}).
		Doc("获取数据库列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取表列表").Handle(s.GetTablePage)
	}).
		Doc("获取表列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/info/tableName").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取表信息By tableName").Handle(s.GetTableInfoByName)
	}).
		Doc("获取表信息By tableName").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", vo.TableInfoVo{}))

	ws.Route(ws.GET("/info/{tableId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取表信息").Handle(s.GetTableInfo)
	}).
		Doc("获取表信息").
		Param(ws.PathParameter("tenantId", "租户Id").DataType("int").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", vo.TableInfoVo{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/tableTree").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取表树").Handle(s.GetTableTree)
	}).
		Doc("获取表树").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", []entity.DevGenTable{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("新增表").Handle(s.Insert)
	}).
		Doc("新增表").
		Metadata(restfulspec.KeyOpenAPITags, tags)) // from the request

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改表").Handle(s.Update)
	}).
		Doc("修改表").
		Metadata(restfulspec.KeyOpenAPITags, tags)) // from the request

	ws.Route(ws.DELETE("/{tableId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除表").Handle(s.Delete)
	}).
		Doc("删除表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("tableId", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
