package api

import (
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"pandax/apps/rule/entity"
	"pandax/apps/rule/services"
	"pandax/pkg/rule_engine/nodes"
	"strings"
)

type RuleChainMsgLogApi struct {
	RuleChainMsgLogApp services.RuleChainMsgLogModel
}

func (r *RuleChainMsgLogApi) GetNodeLabels(rc *restfulx.ReqCtx) {
	rc.ResData = nodes.GetCategory()
}

// GetRuleChainMsgLogList 列表数据
func (p *RuleChainMsgLogApi) GetRuleChainMsgLogList(rc *restfulx.ReqCtx) {
	data := entity.RuleChainMsgLog{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.DeviceName = restfulx.QueryParam(rc, "deviceName")
	list, total := p.RuleChainMsgLogApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// DeleteRuleChainMsgLog 删除规则链
func (p *RuleChainMsgLogApi) DeleteRuleChainMsgLog(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	p.RuleChainMsgLogApp.Delete(ids)
}
