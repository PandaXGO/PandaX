package router

import (
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/rule/api"
	"pandax/apps/rule/entity"
	"pandax/apps/rule/services"
)

func InitRuleChainRouter(container *restful.Container) {
	s := &api.RuleChainApi{
		RuleChainApp: services.RuleChainModelDao,
	}

	ws := new(restful.WebService)
	ws.Path("/rule/chain").Produces(restful.MIME_JSON)
	tags := []string{"rulechain"}

	ws.Route(ws.GET("/nodeLabels").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedCasbin(false).WithLog("获取所有节点标签").Handle(s.GetNodeLabels)
	}).
		Doc("获取所有节点标签").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/node/debug").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedCasbin(false).WithLog("获取规则链节点日志").Handle(s.GetNodeDebug)
	}).
		Doc("获取规则链节点日志").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("ruleId", "规则ID").Required(false).DataType("string")).
		Param(ws.QueryParameter("nodeId", "节点ID").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/node/debug/clear").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedCasbin(false).WithLog("清除规则链节点日志").Handle(s.ClearNodeDebug)
	}).
		Doc("清除规则链节点日志").
		Param(ws.QueryParameter("ruleId", "规则ID").Required(false).DataType("string")).
		Param(ws.QueryParameter("nodeId", "节点ID").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取规则引擎分页列表").Handle(s.GetRuleChainList)
	}).
		Doc("获取规则引擎分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("ruleName", "规则名").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/list/label").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取规则引擎Label列表").Handle(s.GetRuleChainListLabel)
	}).
		Doc("获取规则引擎Label列表").
		Param(ws.QueryParameter("ruleName", "规则名").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", entity.RuleChainBaseLabel{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取规则引擎信息").Handle(s.GetRuleChain)
	}).
		Doc("获取规则引擎信息").
		Param(ws.PathParameter("id", "Id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.RuleChain{}). // on the response
		Returns(200, "OK", entity.RuleChain{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加规则引擎信息").Handle(s.InsertRuleChain)
	}).
		Doc("添加规则引擎信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.RuleChain{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改规则引擎信息").Handle(s.UpdateRuleChain)
	}).
		Doc("修改规则引擎信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.RuleChain{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除规则引擎信息").Handle(s.DeleteRuleChain)
	}).
		Doc("删除规则引擎信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	ws.Route(ws.POST("/clone/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("克隆规则引擎").Handle(s.CloneRuleChain)
	}).
		Doc("克隆规则引擎").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "id").DataType("string")))

	ws.Route(ws.PUT("/changeRoot").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改规则链").Handle(s.UpdateRuleRoot)
	}).
		Doc("修改规则链").
		Metadata(restfulspec.KeyOpenAPITags, tags))

	container.Add(ws)

}
