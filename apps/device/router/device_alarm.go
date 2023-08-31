package router

import (
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"pandax/apps/device/api"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitDeviceAlarmRouter(container *restful.Container) {
	s := &api.DeviceAlarmApi{
		DeviceAlarmApp: services.DeviceAlarmModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/device/alarm").Produces(restful.MIME_JSON)
	tags := []string{"alarm"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取告警分页列表").Handle(s.GetDeviceAlarmList)
	}).
		Doc("获取告警分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("deviceId", "设备Id").Required(false).DataType("string")).
		Param(ws.QueryParameter("type", "告警类型").Required(false).DataType("string")).
		Param(ws.QueryParameter("level", "告警等级").Required(false).DataType("string")).
		Param(ws.QueryParameter("state", "告警状态").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改告警信息").Handle(s.UpdateDeviceAlarm)
	}).
		Doc("修改告警信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DeviceAlarm{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除告警信息").Handle(s.DeleteDeviceAlarm)
	}).
		Doc("删除告警信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
