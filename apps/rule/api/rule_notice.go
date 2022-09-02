package api

// ==========================================================================
// 生成日期：2022-09-02 15:49:39 +0800 CST
// 生成路径: apps/rule/api/rule_notice.go
// 生成人：panda
// ==========================================================================
import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"

	"github.com/XM-GO/PandaKit/utils"
	"pandax/apps/rule/entity"
	"pandax/apps/rule/services"
)

type RuleNoticeApi struct {
	RuleNoticeApp services.RuleNoticeModel
}

// GetRuleNoticeList Notice列表数据
func (p *RuleNoticeApi) GetRuleNoticeList(rc *restfulx.ReqCtx) {
	data := entity.RuleNotice{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.Name = restfulx.QueryParam(rc, "name")

	list, total := p.RuleNoticeApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetRuleNotice 获取Notice
func (p *RuleNoticeApi) GetRuleNotice(rc *restfulx.ReqCtx) {
	id := restfulx.PathParamInt(rc, "id")
	rc.ResData = p.RuleNoticeApp.FindOne(int64(id))
}

// InsertRuleNotice 添加Notice
func (p *RuleNoticeApi) InsertRuleNotice(rc *restfulx.ReqCtx) {
	var data entity.RuleNotice
	restfulx.BindQuery(rc, &data)

	p.RuleNoticeApp.Insert(data)
}

// UpdateRuleNotice 修改Notice
func (p *RuleNoticeApi) UpdateRuleNotice(rc *restfulx.ReqCtx) {
	var data entity.RuleNotice
	restfulx.BindQuery(rc, &data)

	p.RuleNoticeApp.Update(data)
}

// DeleteRuleNotice 删除Notice
func (p *RuleNoticeApi) DeleteRuleNotice(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := utils.IdsStrToIdsIntGroup(id)
	p.RuleNoticeApp.Delete(ids)
}
