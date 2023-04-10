package api

// ==========================================================================
// 生成日期：2023-04-10 11:31:34 +0800 CST
// 生成路径: apps/visual/api/visual_data_set_table_task.go
// 生成人：panda
// ==========================================================================
import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"strings"

	"pandax/apps/visual/entity"
	"pandax/apps/visual/services"
)

type VisualDataSetTableTaskApi struct {
	VisualDataSetTableTaskApp services.VisualDataSetTableTaskModel
}

// GetVisualDataSetTableTaskList DataSetTableTask列表数据
func (p *VisualDataSetTableTaskApi) GetVisualDataSetTableTaskList(rc *restfulx.ReqCtx) {
	data := entity.VisualDataSetTableTask{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.Name = restfulx.QueryParam(rc, "name")

	list, total := p.VisualDataSetTableTaskApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetVisualDataSetTableTask 获取DataSetTableTask
func (p *VisualDataSetTableTaskApi) GetVisualDataSetTableTask(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	rc.ResData = p.VisualDataSetTableTaskApp.FindOne(id)
}

// InsertVisualDataSetTableTask 添加DataSetTableTask
func (p *VisualDataSetTableTaskApi) InsertVisualDataSetTableTask(rc *restfulx.ReqCtx) {
	var data entity.VisualDataSetTableTask
	restfulx.BindQuery(rc, &data)

	p.VisualDataSetTableTaskApp.Insert(data)
}

// UpdateVisualDataSetTableTask 修改DataSetTableTask
func (p *VisualDataSetTableTaskApi) UpdateVisualDataSetTableTask(rc *restfulx.ReqCtx) {
	var data entity.VisualDataSetTableTask
	restfulx.BindQuery(rc, &data)

	p.VisualDataSetTableTaskApp.Update(data)
}

// DeleteVisualDataSetTableTask 删除DataSetTableTask
func (p *VisualDataSetTableTaskApi) DeleteVisualDataSetTableTask(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	p.VisualDataSetTableTaskApp.Delete(ids)
}
