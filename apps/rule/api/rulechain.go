package api

import (
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"pandax/apps/rule/entity"
	"pandax/apps/rule/services"
	"pandax/pkg/global_model"
	"pandax/pkg/rule_engine"
	"strings"
)

type RuleChainApi struct {
	RuleChainApp services.RuleChainModel
}

func (r *RuleChainApi) GetNodeLabels(rc *restfulx.ReqCtx) {
	rc.ResData = rule_engine.GetCategory()
}

func (r *RuleChainApi) GetNodeDebug(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	ruleId := restfulx.QueryParam(rc, "ruleId")
	nodeId := restfulx.QueryParam(rc, "nodeId")

	total, list := rule_engine.GetDebugDataPage(pageNum, pageSize, ruleId, nodeId)
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// GetRuleChainList WorkInfo列表数据
func (p *RuleChainApi) GetRuleChainList(rc *restfulx.ReqCtx) {
	data := entity.RuleChain{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.RuleName = restfulx.QueryParam(rc, "ruleName")

	data.RoleId = rc.LoginAccount.RoleId
	data.Owner = rc.LoginAccount.UserName

	list, total := p.RuleChainApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

func (p *RuleChainApi) GetRuleChainListLabel(rc *restfulx.ReqCtx) {
	data := entity.RuleChain{}
	data.RuleName = restfulx.QueryParam(rc, "ruleName")
	list := p.RuleChainApp.FindListBaseLabel(data)
	rc.ResData = list
}

// GetRuleChain 获取规则链
func (p *RuleChainApi) GetRuleChain(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	rc.ResData = p.RuleChainApp.FindOne(id)
}

// InsertRuleChain 添加规则链
func (p *RuleChainApi) InsertRuleChain(rc *restfulx.ReqCtx) {
	var data entity.RuleChain
	restfulx.BindJsonAndValid(rc, &data)
	data.Id = global_model.GenerateID()
	data.Owner = rc.LoginAccount.UserName
	data.OrgId = rc.LoginAccount.OrganizationId
	p.RuleChainApp.Insert(data)
}

// UpdateRuleChain 修改规则链
func (p *RuleChainApi) UpdateRuleChain(rc *restfulx.ReqCtx) {
	var data entity.RuleChain
	restfulx.BindJsonAndValid(rc, &data)
	p.RuleChainApp.Update(data)
}

// DeleteRuleChain 删除规则链
func (p *RuleChainApi) DeleteRuleChain(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	one := p.RuleChainApp.FindOne(id)
	biz.IsTrue(!(one.Root == "1"), "主链不可被删除")
	ids := strings.Split(id, ",")
	p.RuleChainApp.Delete(ids)
}

// CloneRuleChain 克隆规则链
func (p *RuleChainApi) CloneRuleChain(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	one := p.RuleChainApp.FindOne(id)
	one.RuleName = one.RuleName + "-克隆"
	one.Id = global_model.GenerateID()
	one.Root = "0"
	p.RuleChainApp.Insert(*one)
}

// UpdateRuleRoot 修改根链
func (p *RuleChainApi) UpdateRuleRoot(rc *restfulx.ReqCtx) {
	var rule entity.RuleChain
	restfulx.BindJsonAndValid(rc, &rule)
	// 修改主链为普通链
	p.RuleChainApp.UpdateByRoot()
	// 修改当前链为主链
	p.RuleChainApp.Update(rule)
}
