package router

import (
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/resource/api"
	"pandax/apps/resource/entity"
	"pandax/apps/resource/services"
)

/**
 * @Description 添加qq群467890197 交流学习
 * @Author 熊猫
 * @Date 2022/1/14 15:24
 **/

func InitResEmailsRouter(container *restful.Container) {
	s := &api.ResEmailsApi{
		ResEmailsApp: services.ResEmailsModelDao,
	}
	ws := new(restful.WebService)
	ws.Path("/resource/email").Produces(restful.MIME_JSON)
	tags := []string{"email"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取ResEmails分页列表").Handle(s.GetResEmailsList)
	}).
		Doc("获取ResEmails分页列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]entity.ResEmail{}).
		Returns(200, "OK", []entity.ResEmail{}))

	ws.Route(ws.GET("/{mailId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取ResEmails信息").Handle(s.GetResEmails)
	}).
		Doc("获取ResEmails信息").
		Param(ws.PathParameter("mailId", "Id").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.ResEmail{}). // on the response
		Returns(200, "OK", entity.ResEmail{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加ResEmails信息").Handle(s.InsertResEmails)
	}).
		Doc("添加ResEmails信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.ResEmail{})) // from the request

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改ResEmails信息").Handle(s.UpdateResEmails)
	}).
		Doc("修改ResEmails信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.ResEmail{})) // from the request

	ws.Route(ws.DELETE("/{mailId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除ResEmails信息").Handle(s.DeleteResEmails)
	}).
		Doc("删除ResEmails信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("mailId", "多id 1,2,3").DataType("string")))

	ws.Route(ws.POST("/debugMail").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("调试").Handle(s.DebugMail)
	}).
		Doc("调试").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.PUT("/changeStatus").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改状态").Handle(s.UpdateMailStatus)
	}).
		Doc("修改状态").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	container.Add(ws)
}
