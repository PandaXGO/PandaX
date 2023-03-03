package api

// ==========================================================================
// 生成日期：2022-08-24 22:02:33 +0800 CST
// 生成路径: apps/flow/api/flow_work_classify.go
// 生成人：panda
// ==========================================================================
import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"

	"github.com/XM-GO/PandaKit/utils"
	"pandax/apps/flow/entity"
	"pandax/apps/flow/services"
)

type FlowWorkClassifyApi struct {
	FlowWorkClassifyApp services.FlowWorkClassifyModel
}

// GetFlowWorkClassifyList Classify列表数据
func (p *FlowWorkClassifyApi) GetFlowWorkClassifyList(rc *restfulx.ReqCtx) {
	data := entity.FlowWorkClassify{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.Name = restfulx.QueryParam(rc, "name")

	list, total := p.FlowWorkClassifyApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetFlowWorkClassify 获取Classify
func (p *FlowWorkClassifyApi) GetFlowWorkClassify(rc *restfulx.ReqCtx) {
	id := restfulx.PathParamInt(rc, "id")
	p.FlowWorkClassifyApp.FindOne(int64(id))
}

// InsertFlowWorkClassify 添加Classify
func (p *FlowWorkClassifyApi) InsertFlowWorkClassify(rc *restfulx.ReqCtx) {
	var data entity.FlowWorkClassify
	restfulx.BindQuery(rc, &data)

	p.FlowWorkClassifyApp.Insert(data)
}

// UpdateFlowWorkClassify 修改Classify
func (p *FlowWorkClassifyApi) UpdateFlowWorkClassify(rc *restfulx.ReqCtx) {
	var data entity.FlowWorkClassify
	restfulx.BindQuery(rc, &data)

	p.FlowWorkClassifyApp.Update(data)
}

// DeleteFlowWorkClassify 删除Classify
func (p *FlowWorkClassifyApi) DeleteFlowWorkClassify(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := utils.IdsStrToIdsIntGroup(id)
	p.FlowWorkClassifyApp.Delete(ids)
}
