package router

import (
	"pandax/apps/device/api"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	"pandax/kit/restfulx"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitDeviceGroupRouter(container *restful.Container) {
	s := &api.DeviceGroupApi{
		DeviceGroupApp: services.DeviceGroupModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/device/group").Produces(restful.MIME_JSON)
	tags := []string{"设备分组"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取DeviceGroup列表").Handle(s.GetDeviceGroupList)
	}).
		Doc("获取DeviceGroup列表").
		Param(ws.QueryParameter("name", "名称").Required(false).DataType("string")).
		Param(ws.QueryParameter("status", "状态").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", []entity.DeviceGroup{}))

	ws.Route(ws.GET("/list/all").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取DeviceGroup所有列表").Handle(s.GetDeviceGroupAllList)
	}).
		Doc("获取DeviceGroup分页列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", []entity.DeviceGroup{}))

	ws.Route(ws.GET("/list/tree").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取DeviceGroup树").Handle(s.GetDeviceGroupTree)
	}).
		Doc("获取DeviceGroup树").
		Param(ws.QueryParameter("name", "名称").Required(false).DataType("string")).
		Param(ws.QueryParameter("status", "状态").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", []entity.DeviceGroup{}))

	ws.Route(ws.GET("/list/tree/label").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取DeviceGroup树").Handle(s.GetDeviceGroupTreeLabel)
	}).
		Doc("获取DeviceGroup树").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", []entity.DeviceGroupLabel{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取DeviceGroup信息").Handle(s.GetDeviceGroup)
	}).
		Doc("获取DeviceGroup信息").
		Param(ws.PathParameter("", "Id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.DeviceGroup{}). // on the response
		Returns(200, "OK", entity.DeviceGroup{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加DeviceGroup信息").Handle(s.InsertDeviceGroup)
	}).
		Doc("添加DeviceGroup信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DeviceGroup{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改DeviceGroup信息").Handle(s.UpdateDeviceGroup)
	}).
		Doc("修改DeviceGroup信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.DeviceGroup{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除DeviceGroup信息").Handle(s.DeleteDeviceGroup)
	}).
		Doc("删除DeviceGroup信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
