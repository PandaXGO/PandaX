package api

// ==========================================================================
// 生成日期：2023-03-29 19:46:55 +0800 CST
// 生成路径: apps/flow/api/flow_work_templates.go
// 生成人：panda
// ==========================================================================
import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"

	"github.com/XM-GO/PandaKit/utils"
	"pandax/apps/flow/entity"
	"pandax/apps/flow/services"
)

type FlowWorkTemplatesApi struct {
	FlowWorkTemplatesApp services.FlowWorkTemplatesModel
}

// GetFlowWorkTemplatesList WorkTemplates列表数据
func (p *FlowWorkTemplatesApi) GetFlowWorkTemplatesList(rc *restfulx.ReqCtx) {
	data := entity.FlowWorkTemplates{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.Name = restfulx.QueryParam(rc, "name")

	list, total := p.FlowWorkTemplatesApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetFlowWorkTemplates 获取WorkTemplates
func (p *FlowWorkTemplatesApi) GetFlowWorkTemplates(rc *restfulx.ReqCtx) {
	id := restfulx.PathParamInt(rc, "id")
	rc.ResData = p.FlowWorkTemplatesApp.FindOne(int64(id))
}

// InsertFlowWorkTemplates 添加WorkTemplates
func (p *FlowWorkTemplatesApi) InsertFlowWorkTemplates(rc *restfulx.ReqCtx) {
	var data entity.FlowWorkTemplates
	restfulx.BindQuery(rc, &data)
	data.Creator = int(rc.LoginAccount.UserId)
	p.FlowWorkTemplatesApp.Insert(data)
}

// UpdateFlowWorkTemplates 修改WorkTemplates
func (p *FlowWorkTemplatesApi) UpdateFlowWorkTemplates(rc *restfulx.ReqCtx) {
	var data entity.FlowWorkTemplates
	restfulx.BindQuery(rc, &data)
	p.FlowWorkTemplatesApp.Update(data)
}

// DeleteFlowWorkTemplates 删除WorkTemplates
func (p *FlowWorkTemplatesApi) DeleteFlowWorkTemplates(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := utils.IdsStrToIdsIntGroup(id)
	p.FlowWorkTemplatesApp.Delete(ids)
}
