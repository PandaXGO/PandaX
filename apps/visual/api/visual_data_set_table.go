package api

// ==========================================================================
// 生成日期：2023-04-10 03:05:21 +0800 CST
// 生成路径: apps/visual/api/visual_data_set_table.go
// 生成人：panda
// ==========================================================================
import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"strings"

	"pandax/apps/visual/entity"
	"pandax/apps/visual/services"
)

type VisualDataSetTableApi struct {
	VisualDataSetTableApp services.VisualDataSetTableModel
}

// GetVisualDataSetTableList DataSetTable列表数据
func (p *VisualDataSetTableApi) GetVisualDataSetTableList(rc *restfulx.ReqCtx) {
	data := entity.VisualDataSetTable{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)

	list, total := p.VisualDataSetTableApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetVisualDataSetTable 获取DataSetTable
func (p *VisualDataSetTableApi) GetVisualDataSetTable(rc *restfulx.ReqCtx) {
	tableId := restfulx.PathParam(rc, "tableId")
	rc.ResData = p.VisualDataSetTableApp.FindOne(tableId)
}

// InsertVisualDataSetTable 添加DataSetTable
func (p *VisualDataSetTableApi) InsertVisualDataSetTable(rc *restfulx.ReqCtx) {
	var data entity.VisualDataSetTable
	restfulx.BindQuery(rc, &data)

	p.VisualDataSetTableApp.Insert(data)
}

// UpdateVisualDataSetTable 修改DataSetTable
func (p *VisualDataSetTableApi) UpdateVisualDataSetTable(rc *restfulx.ReqCtx) {
	var data entity.VisualDataSetTable
	restfulx.BindQuery(rc, &data)

	p.VisualDataSetTableApp.Update(data)
}

// DeleteVisualDataSetTable 删除DataSetTable
func (p *VisualDataSetTableApi) DeleteVisualDataSetTable(rc *restfulx.ReqCtx) {
	tableId := restfulx.PathParam(rc, "tableId")
	tableIds := strings.Split(tableId, ",")
	p.VisualDataSetTableApp.Delete(tableIds)
}
