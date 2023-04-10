package router

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"pandax/apps/visual/api"
	"pandax/apps/visual/entity"
)

func InitRuleChainRouter(container *restful.Container) {
	s := &api.RuleChainApi{}

	ws := new(restful.WebService)
	ws.Path("/visual/rulechain").Produces(restful.MIME_JSON)
	tags := []string{"rulechain"}

	ws.Route(ws.GET("/nodeLabels").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedCasbin(false).WithLog("获取所有节点标签").Handle(s.GetNodeLabels)
	}).
		Doc("获取所有节点标签").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/test").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithNeedCasbin(false).WithLog("测试规则引擎").Handle(s.RuleChainTest)
	}).
		Doc("测试规则引擎").
		Param(ws.QueryParameter("code", "流程代码").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.GET("/list").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取规则引擎分页列表").Handle(s.GetVisualRuleChainList)
	}).
		Doc("获取规则引擎分页列表").
		Param(ws.QueryParameter("pageNum", "页数").Required(true).DataType("int")).
		Param(ws.QueryParameter("pageSize", "每页条数").Required(true).DataType("int")).
		Param(ws.QueryParameter("ruleName", "规则名").Required(false).DataType("string")).
		Param(ws.QueryParameter("status", "状态").Required(false).DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(model.ResultPage{}).
		Returns(200, "OK", model.ResultPage{}))

	ws.Route(ws.GET("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("获取规则引擎信息").Handle(s.GetVisualRuleChain)
	}).
		Doc("获取规则引擎信息").
		Param(ws.PathParameter("id", "Id").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(entity.VisualRuleChain{}). // on the response
		Returns(200, "OK", entity.VisualRuleChain{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.POST("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("添加规则引擎信息").Handle(s.InsertVisualRuleChain)
	}).
		Doc("添加规则引擎信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualRuleChain{}))

	ws.Route(ws.PUT("").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("修改规则引擎信息").Handle(s.UpdateVisualRuleChain)
	}).
		Doc("修改规则引擎信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(entity.VisualRuleChain{}))

	ws.Route(ws.DELETE("/{id}").To(func(request *restful.Request, response *restful.Response) {
		restfulx.NewReqCtx(request, response).WithLog("删除规则引擎信息").Handle(s.DeleteVisualRuleChain)
	}).
		Doc("删除规则引擎信息").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "多id 1,2,3").DataType("string")))

	container.Add(ws)

}
