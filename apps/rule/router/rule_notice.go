// ==========================================================================
// 生成日期：2022-09-02 15:49:39 +0800 CST
// 生成路径: apps/rule/router/rule_notice.go
// 生成人：panda
// ==========================================================================
package router

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"pandax/apps/rule/api"
	"pandax/apps/rule/entity"
	"pandax/apps/rule/services"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitRuleNoticeRouter(container *restful.Container) {
	s := &api.RuleNoticeApi{
		RuleNoticeApp: services.RuleNoticeModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/rule/notice").Produces(restful.MIME_JSON)
	tags := []string{"notice"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Notice分页列表").Handle(s.GetRuleNoticeList)
	}).
		Doc("获取Notice分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取Notice信息").Handle(s.GetRuleNotice)
	}).
		Doc("获取Notice信息").
		Param(ws.PathParameter("id", "Id").DataType("int")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.RuleNotice{}). // on the response
		Returns(200, "OK", entity.RuleNotice{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加Notice信息").Handle(s.InsertRuleNotice)
	}).
		Doc("添加Notice信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.RuleNotice{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改Notice信息").Handle(s.UpdateRuleNotice)
	}).
		Doc("修改Notice信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.RuleNotice{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除Notice信息").Handle(s.DeleteRuleNotice)
	}).
		Doc("删除Notice信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	container.Add(ws)
}
