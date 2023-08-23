package router

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/rule/api"
	"pandax/apps/rule/services"
)

func InitRuleChainMsgLogRouter(container *restful.Container) {
	s := &api.RuleChainMsgLogApi{
		RuleChainMsgLogApp: services.RuleChainMsgLogModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/rule/chain/log").Produces(restful.MIME_JSON)
	tags := []string{"RuleChainMsgLog"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取规则引擎分页列表").Handle(s.GetRuleChainMsgLogList)
	}).
		Doc("获取规则引擎分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("ruleName", "规则名").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除规则引擎信息").Handle(s.DeleteRuleChainMsgLog)
	}).
		Doc("删除规则链日志信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	container.Add(ws)

}
