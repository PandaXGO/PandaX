package api

// ==========================================================================
// 生成日期：2023-04-10 11:31:34 +0800 CST
// 生成路径: apps/visual/api/visual_data_set_table_task_log.go
// 生成人：panda
// ==========================================================================
import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"strings"

	"pandax/apps/visual/entity"
	"pandax/apps/visual/services"
)

type VisualDataSetTableTaskLogApi struct {
	VisualDataSetTableTaskLogApp services.VisualDataSetTableTaskLogModel
}

// GetVisualDataSetTableTaskLogList DataSetTableTaskLog列表数据
func (p *VisualDataSetTableTaskLogApi) GetVisualDataSetTableTaskLogList(rc *restfulx.ReqCtx) {
	data := entity.VisualDataSetTableTaskLog{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.Status = restfulx.QueryParam(rc, "status")

	list, total := p.VisualDataSetTableTaskLogApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetVisualDataSetTableTaskLog 获取DataSetTableTaskLog
func (p *VisualDataSetTableTaskLogApi) GetVisualDataSetTableTaskLog(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	rc.ResData = p.VisualDataSetTableTaskLogApp.FindOne(id)
}

// InsertVisualDataSetTableTaskLog 添加DataSetTableTaskLog
func (p *VisualDataSetTableTaskLogApi) InsertVisualDataSetTableTaskLog(rc *restfulx.ReqCtx) {
	var data entity.VisualDataSetTableTaskLog
	restfulx.BindQuery(rc, &data)

	p.VisualDataSetTableTaskLogApp.Insert(data)
}

// UpdateVisualDataSetTableTaskLog 修改DataSetTableTaskLog
func (p *VisualDataSetTableTaskLogApi) UpdateVisualDataSetTableTaskLog(rc *restfulx.ReqCtx) {
	var data entity.VisualDataSetTableTaskLog
	restfulx.BindQuery(rc, &data)

	p.VisualDataSetTableTaskLogApp.Update(data)
}

// DeleteVisualDataSetTableTaskLog 删除DataSetTableTaskLog
func (p *VisualDataSetTableTaskLogApi) DeleteVisualDataSetTableTaskLog(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	p.VisualDataSetTableTaskLogApp.Delete(ids)
}
