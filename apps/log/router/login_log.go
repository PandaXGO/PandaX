package router

import (
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/log/api"
	"pandax/apps/log/entity"
	"pandax/apps/log/services"
)

func InitLoginLogRouter(container *restful.Container) {
	// 登录日志
	s := &api.LogLoginApi{
		LogLoginApp: services.LogLoginModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/log/logLogin").Produces(restful.MIME_JSON)
	tags := []string{"logLogin"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取登录日志列表").Handle(s.GetLoginLogList)
	}).
		Doc("获取登录日志列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.LogLogin{}).
		Returns(200, "OK", []entity.LogLogin{}))

	ws.Route(ws.GET("/{infoId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取登录日志信息").Handle(s.GetLoginLog)
	}).
		Doc("获取登录日志信息").
		Param(ws.PathParameter("infoId", "Id").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.LogLogin{}). // on the response
		Returns(200, "OK", entity.LogLogin{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改登录日志信息").Handle(s.UpdateLoginLog)
	}).
		Doc("修改登录日志信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.LogLogin{}))

	ws.Route(ws.DELETE("/{infoId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除登录日志信息").Handle(s.DeleteLoginLog)
	}).
		Doc("删除登录日志信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("infoId", "多id 1,2,3").DataType("string")))

	ws.Route(ws.DELETE("/all").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("清空登录日志信息").Handle(s.DeleteAll)
	}).
		Doc("清空登录日志信息").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	container.Add(ws)

}
