package api

import (
	"context"
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/kakuilan/kgo"
	"pandax/apps/visual/entity"
	"pandax/apps/visual/services"
	"pandax/pkg/rule_engine"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/rule_engine/nodes"
	"strings"
)

type RuleChainApi struct {
	VisualRuleChainApp services.VisualRuleChainModel
}

func (r *RuleChainApi) GetNodeLabels(rc *restfulx.ReqCtx) {
	rc.ResData = nodes.GetCategory()
}
func (r *RuleChainApi) RuleChainTest(rc *restfulx.ReqCtx) {
	code := restfulx.QueryParam(rc, "code")
	instance, _ := rule_engine.NewRuleChainInstance([]byte(code))
	newMessage := message.NewMessage()
	newMessage.SetMetadata(message.NewMetadata())
	instance.StartRuleChain(context.Background(), newMessage)
	rc.ResData = []map[string]interface{}{}
}

// GetVisualRuleChainList WorkInfo列表数据
func (p *RuleChainApi) GetVisualRuleChainList(rc *restfulx.ReqCtx) {
	data := entity.VisualRuleChain{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.RuleName = restfulx.QueryParam(rc, "ruleName")
	data.Status = restfulx.QueryParam(rc, "status")
	list, total := p.VisualRuleChainApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetVisualRuleChain 获取规则链
func (p *RuleChainApi) GetVisualRuleChain(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	rc.ResData = p.VisualRuleChainApp.FindOne(id)
}

// InsertVisualRuleChain 添加规则链
func (p *RuleChainApi) InsertVisualRuleChain(rc *restfulx.ReqCtx) {
	var data entity.VisualRuleChain
	restfulx.BindQuery(rc, &data)
	data.RuleId = kgo.KStr.Uniqid("px")
	data.Creator = rc.LoginAccount.UserName
	p.VisualRuleChainApp.Insert(data)
}

// UpdateVisualRuleChain 修改规则链
func (p *RuleChainApi) UpdateVisualRuleChain(rc *restfulx.ReqCtx) {
	var data entity.VisualRuleChain
	restfulx.BindQuery(rc, &data)

	p.VisualRuleChainApp.Update(data)
}

// DeleteVisualRuleChain 删除规则链
func (p *RuleChainApi) DeleteVisualRuleChain(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	p.VisualRuleChainApp.Delete(ids)
}

// UpdateRuleStatus 修改状态
func (p *RuleChainApi) UpdateRuleStatus(rc *restfulx.ReqCtx) {
	var rule entity.VisualRuleChain
	restfulx.BindQuery(rc, &rule)
	p.VisualRuleChainApp.Update(rule)
}
