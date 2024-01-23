package router

import (
	"pandax/apps/device/api"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	"pandax/kit/model"
	"pandax/kit/restfulx"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitDeviceRouter(container *restful.Container) {
	s := &api.DeviceApi{
		DeviceApp:          services.DeviceModelDao,
		DeviceAlarmApp:     services.DeviceAlarmModelDao,
		ProductApp:         services.ProductModelDao,
		ProductTemplateApp: services.ProductTemplateModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/device").Produces(restful.MIME_JSON)
	tags := []string{"device"}

	ws.Route(ws.GET("/panel").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取DevicePanel").Handle(s.GetDevicePanel)
	}).
		Doc("获取DevicePanel").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.DeviceTotalOutput{}).
		Returns(200, "OK", entity.DeviceTotalOutput{}))

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Device分页列表").Handle(s.GetDeviceList)
	}).
		Doc("获取Device分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("status", "状态").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "名称").Required(false).DataType("string")).
		Param(ws.QueryParameter("pid", "产品ID").Required(false).DataType("string")).
		Param(ws.QueryParameter("gid", "分组Id").Required(false).DataType("string")).
		Param(ws.QueryParameter("deviceType", "设备类型").Required(false).DataType("string")).
		Param(ws.QueryParameter("parentId", "父ID").Required(false).DataType("string")).
		Param(ws.QueryParameter("linkStatus", "连接状态").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/list/all").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Device列表").Handle(s.GetDeviceListAll)
	}).
		Doc("获取Device列表").
		Param(ws.QueryParameter("status", "状态").Required(false).DataType("string")).
		Param(ws.QueryParameter("name", "名称").Required(false).DataType("string")).
		Param(ws.QueryParameter("pid", "产品ID").Required(false).DataType("string")).
		Param(ws.QueryParameter("deviceType", "设备类型").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.Device{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Device信息").Handle(s.GetDevice)
	}).
		Doc("获取Device信息").
		Param(ws.PathParameter("id", "Id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.Device{}). // on the response
		Returns(200, "OK", entity.Device{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/{id}/status").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Device状态信息").Handle(s.GetDeviceStatus)
	}).
		Doc("获取Device状态信息").
		Param(ws.PathParameter("id", "Id").DataType("string")).
		Param(ws.QueryParameter("classify", "分类").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.DeviceStatusVo{}). // on the response
		Returns(200, "OK", []entity.DeviceStatusVo{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/{id}/property/history").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取设备属性的遥测历史").Handle(s.GetDeviceTelemetryHistory)
	}).
		Doc("获取设备属性的遥测历史").
		Param(ws.PathParameter("id", "Id").DataType("string")).
		Param(ws.QueryParameter("key", "属性Key").Required(false).DataType("string")).
		Param(ws.QueryParameter("startTime", "开始时间").Required(false).DataType("string")).
		Param(ws.QueryParameter("endTime", "结束时间").Required(false).DataType("string")).
		Param(ws.QueryParameter("limit", "限制条数").Required(false).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags). // on the response
		Returns(200, "OK", []map[string]any{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/{id}/attribute/down").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Device属性下发").Handle(s.DownAttribute)
	}).
		Doc("获取Device属性下发").
		Param(ws.PathParameter("id", "Id").DataType("string")).
		Param(ws.QueryParameter("key", "属性KEY").Required(false).DataType("string")).
		Param(ws.QueryParameter("value", "属性Value").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加Device信息").Handle(s.InsertDevice)
	}).
		Doc("添加Device信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.Device{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改Device信息").Handle(s.UpdateDevice)
	}).
		Doc("修改Device信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.Device{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除Device信息").Handle(s.DeleteDevice)
	}).
		Doc("删除Device信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	ws.Route(ws.GET("/twin").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Device孪生体").Handle(s.ScreenTwinData)
	}).
		Doc("获取Device孪生体").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.VisualClass{}). // on the response
		Returns(200, "OK", []entity.VisualClass{}))

	ws.Route(ws.GET("/{id}/allot/org").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("分配组织").Handle(s.DeviceAllotOrg)
	}).
		Doc("分配组织").
		Metadata(restfulspec.KeyOpenAPITags, tags). // on the response
		Returns(200, "OK", ""))

	container.Add(ws)
}
