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

func InitDeviceCmdLogRouter(container *restful.Container) {
	s := &api.DeviceCmdLogApi{
		DeviceCmdLogApp: services.DeviceCmdLogModelDao,
		DeviceApp:       services.DeviceModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/device/cmd").Produces(restful.MIME_JSON)
	tags := []string{"设备命令"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取命令下发分页列表").Handle(s.GetDeviceCmdLogList)
	}).
		Doc("获取命令下发分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("deviceId", "设备Id").Required(false).DataType("string")).
		Param(ws.QueryParameter("type", "命令下发分类").Required(false).DataType("string")).
		Param(ws.QueryParameter("state", "命令下发状态").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("命令下发").Handle(s.InsertDeviceCmdLog)
	}).
		Doc("命令下发").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DeviceCmdLog{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除命令下发信息").Handle(s.DeleteDeviceCmdLog)
	}).
		Doc("删除命令下发信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
