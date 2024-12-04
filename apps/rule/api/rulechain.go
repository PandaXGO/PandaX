package api

import (
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
	"pandax/apps/rule/entity"
	"pandax/apps/rule/services"
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

	total, list, err := rule_engine.RuleEngine.GetDebugDataPage(pageNum, pageSize, ruleId, nodeId)
	biz.ErrIsNil(err, "获取规则测试数据错误")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

func (r *RuleChainApi) ClearNodeDebug(rc *restfulx.ReqCtx) {
	ruleId := restfulx.QueryParam(rc, "ruleId")
	nodeId := restfulx.QueryParam(rc, "nodeId")
	rule_engine.RuleEngine.ClearDebugData(ruleId, nodeId)
}

// GetRuleChainList WorkInfo列表数据
func (p *RuleChainApi) GetRuleChainList(rc *restfulx.ReqCtx) {
	data := entity.RuleChain{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.RuleName = restfulx.QueryParam(rc, "ruleName")

	data.RoleId = rc.LoginAccount.RoleId
	data.Owner = rc.LoginAccount.UserName

	list, total, err := p.RuleChainApp.FindListPage(pageNum, pageSize, data)
	biz.ErrIsNil(err, "获取规则链列表错误")
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
	data.RoleId = rc.LoginAccount.RoleId
	data.Owner = rc.LoginAccount.UserName

	list, err := p.RuleChainApp.FindListBaseLabel(data)
	biz.ErrIsNilAppendErr(err, "获取规则链Label错误")
	rc.ResData = list
}

// GetRuleChain 获取规则链
func (p *RuleChainApi) GetRuleChain(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	data, err := p.RuleChainApp.FindOne(id)
	biz.ErrIsNil(err, "获取规则链信息错误")
	rc.ResData = data
}

// InsertRuleChain 添加规则链
func (p *RuleChainApi) InsertRuleChain(rc *restfulx.ReqCtx) {
	var data entity.RuleChain
	restfulx.BindJsonAndValid(rc, &data)
	data.Id = utils.GenerateID("rc")
	data.Owner = rc.LoginAccount.UserName
	data.OrgId = rc.LoginAccount.OrganizationId
	_, err := p.RuleChainApp.Insert(data)
	biz.ErrIsNil(err, "添加规则链错误")
}

// UpdateRuleChain 修改规则链
func (p *RuleChainApi) UpdateRuleChain(rc *restfulx.ReqCtx) {
	var data entity.RuleChain
	restfulx.BindJsonAndValid(rc, &data)
	_, err := p.RuleChainApp.Update(data)
	biz.ErrIsNil(err, "修改规则链错误")
}

// DeleteRuleChain 删除规则链
func (p *RuleChainApi) DeleteRuleChain(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	one, err := p.RuleChainApp.FindOne(id)
	biz.ErrIsNil(err, "规则链不存在")
	biz.IsTrue(!(one.Root == "1"), "主链不可被删除")
	ids := strings.Split(id, ",")
	biz.ErrIsNil(p.RuleChainApp.Delete(ids), "删除规则链失败")
}

// CloneRuleChain 克隆规则链
func (p *RuleChainApi) CloneRuleChain(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	one, err := p.RuleChainApp.FindOne(id)
	biz.ErrIsNil(err, "规则链不存在")
	one.RuleName = one.RuleName + "-克隆"
	one.Id = utils.GenerateID("rc")
	one.Root = "0"
	_, err = p.RuleChainApp.Insert(*one)
	biz.ErrIsNil(err, "克隆规则链失败")
}

// UpdateRuleRoot 修改根链
func (p *RuleChainApi) UpdateRuleRoot(rc *restfulx.ReqCtx) {
	var rule entity.RuleChain
	restfulx.BindJsonAndValid(rc, &rule)
	// 修改主链为普通链
	err := p.RuleChainApp.UpdateByRoot()
	biz.ErrIsNil(err, "修改主链错误")
	// 修改当前链为主链
	_, err = p.RuleChainApp.Update(rule)
	biz.ErrIsNil(err, "修改当前链为主链错误")
}
