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

func InitNoticeRouter(container *restful.Container) {
	s := &api.NoticeApi{
		DeptApp:   services.SysDeptModelDao,
		NoticeApp: services.SysNoticeModelDao,
	}
	ws := new(restful.WebService)
	ws.Path("/system/notice").Produces(restful.MIME_JSON)
	tags := []string{"notice"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取通知分页列表").Handle(s.GetNoticeList)
	}).
		Doc("获取通知分页列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("noticeType", "noticeType").DataType("string")).
		Param(ws.QueryParameter("title", "title").DataType("string")).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加通知信息").Handle(s.InsertNotice)
	}).
		Doc("添加通知信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysNotice{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改通知信息").Handle(s.UpdateNotice)
	}).
		Doc("修改通知信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.SysNotice{}))

	ws.Route(ws.DELETE("/{noticeId}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除通知信息").Handle(s.DeleteNotice)
	}).
		Doc("删除通知信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("noticeId", "多id 1,2,3").DataType("string")))

	container.Add(ws)

}
