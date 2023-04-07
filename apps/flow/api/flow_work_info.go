package api

// ==========================================================================
// 生成日期：2023-03-29 20:01:11 +0800 CST
// 生成路径: apps/flow/api/rulechain.go
// 生成人：panda
// ==========================================================================
import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"

	"github.com/XM-GO/PandaKit/utils"
	"pandax/apps/flow/entity"
	"pandax/apps/flow/services"
)

type FlowWorkInfoApi struct {
	FlowWorkInfoApp services.FlowWorkInfoModel
}

// GetFlowWorkInfoList WorkInfo列表数据
func (p *FlowWorkInfoApi) GetFlowWorkInfoList(rc *restfulx.ReqCtx) {
	data := entity.FlowWorkInfo{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.Name = restfulx.QueryParam(rc, "name")
	data.Icon = restfulx.QueryParam(rc, "icon")

	list, total := p.FlowWorkInfoApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetFlowWorkInfo 获取WorkInfo
func (p *FlowWorkInfoApi) GetFlowWorkInfo(rc *restfulx.ReqCtx) {
	id := restfulx.PathParamInt(rc, "id")
	rc.ResData = p.FlowWorkInfoApp.FindOne(int64(id))
}

// InsertFlowWorkInfo 添加WorkInfo
func (p *FlowWorkInfoApi) InsertFlowWorkInfo(rc *restfulx.ReqCtx) {
	var data entity.FlowWorkInfo
	restfulx.BindQuery(rc, &data)
	data.Creator = int(rc.LoginAccount.UserId)
	p.FlowWorkInfoApp.Insert(data)
}

// UpdateFlowWorkInfo 修改WorkInfo
func (p *FlowWorkInfoApi) UpdateFlowWorkInfo(rc *restfulx.ReqCtx) {
	var data entity.FlowWorkInfo
	restfulx.BindQuery(rc, &data)

	p.FlowWorkInfoApp.Update(data)
}

// DeleteFlowWorkInfo 删除WorkInfo
func (p *FlowWorkInfoApi) DeleteFlowWorkInfo(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := utils.IdsStrToIdsIntGroup(id)
	p.FlowWorkInfoApp.Delete(ids)
}
