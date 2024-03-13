package router

import (
	"pandax/apps/rule/api"
	"pandax/apps/rule/services"
	"pandax/kit/model"
	"pandax/kit/restfulx"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

func InitRuleChainMsgLogRouter(container *restful.Container) {
	s := &api.RuleChainMsgLogApi{
		RuleChainMsgLogApp: services.RuleChainMsgLogModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/rule/chain/log").Produces(restful.MIME_JSON)
	tags := []string{"规则链日志"}

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取规则引擎日志分页列表").Handle(s.GetRuleChainMsgLogList)
	}).
		Doc("获取规则引擎日志分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("deviceName", "设备名").Required(false).DataType("string")).
		Param(ws.QueryParameter("msgType", "消息类型").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/delete").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除规则引擎信息").Handle(s.DeleteRuleChainMsgLog)
	}).
		Doc("删除规则链日志信息").
		Param(ws.QueryParameter("deviceName", "设备名").Required(false).DataType("string")).
		Param(ws.QueryParameter("msgType", "消息类型").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	container.Add(ws)

}
