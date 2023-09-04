package router

import (
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/video/api"
	"pandax/pkg/global"
	"pandax/pkg/ys"
)

func InitEzvizRouter(container *restful.Container) {
	// 登录日志
	s := &api.YsApi{
		Ys: &ys.Ys{
			AppKey: global.Conf.Ys.AppKey,
			Secret: global.Conf.Ys.Secret,
		},
	}
	ws := new(restful.WebService)
	ws.Path("/video/ys").Produces(restful.MIME_JSON)
	tags := []string{"ys"}

	ws.Route(ws.GET("/device/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取设备列表").Handle(s.GetDeviceList)
	}).
		Doc("获取设备列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{deviceSerial}/channel").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("设备通道列表").Handle(s.GetDeviceChannelList)
	}).
		Doc("设备通道列表").
		Param(ws.PathParameter("deviceSerial", "deviceSerial").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(ys.Channel{}). // on the response
		Returns(200, "OK", []ys.Channel{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/{deviceSerial}/start/ptz").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("摄像头操作").Handle(s.StartPtz)
	}).
		Doc("摄像头操作").
		Param(ws.PathParameter("deviceSerial", "deviceSerial").DataType("string")).
		Param(ws.QueryParameter("channelNo", "通道序号").Required(true).DataType("int")).
		Param(ws.QueryParameter("direction", "方向").Required(true).DataType("int")).
		Param(ws.QueryParameter("speed", "速度").Required(true).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/{deviceSerial}/stop/ptz").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("摄像头操作停止").Handle(s.StopPtz)
	}).
		Doc("摄像头操作停止").
		Param(ws.PathParameter("deviceSerial", "deviceSerial").DataType("string")).
		Param(ws.QueryParameter("channelNo", "通道序号").Required(true).DataType("int")).
		Param(ws.QueryParameter("direction", "方向").Required(true).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	container.Add(ws)
}
