package api

// ==========================================================================
// 生成日期：2023-04-10 02:51:27 +0000 UTC
// 生成路径: apps/visual/api/visual_data_source.go
// 生成人：panda
// ==========================================================================
import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"strings"

	"pandax/apps/visual/entity"
	"pandax/apps/visual/services"
)

type VisualDataSourceApi struct {
	VisualDataSourceApp services.VisualDataSourceModel
}

// GetVisualDataSourceList DataSource列表数据
func (p *VisualDataSourceApi) GetVisualDataSourceList(rc *restfulx.ReqCtx) {
	data := entity.VisualDataSource{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.SourceName = restfulx.QueryParam(rc, "sourceName")
	data.Status = restfulx.QueryParam(rc, "status")

	list, total := p.VisualDataSourceApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetVisualDataSource 获取DataSource
func (p *VisualDataSourceApi) GetVisualDataSource(rc *restfulx.ReqCtx) {
	sourceId := restfulx.PathParam(rc, "sourceId")
	rc.ResData = p.VisualDataSourceApp.FindOne(sourceId)
}

// InsertVisualDataSource 添加DataSource
func (p *VisualDataSourceApi) InsertVisualDataSource(rc *restfulx.ReqCtx) {
	var data entity.VisualDataSource
	restfulx.BindQuery(rc, &data)

	p.VisualDataSourceApp.Insert(data)
}

// UpdateVisualDataSource 修改DataSource
func (p *VisualDataSourceApi) UpdateVisualDataSource(rc *restfulx.ReqCtx) {
	var data entity.VisualDataSource
	restfulx.BindQuery(rc, &data)

	p.VisualDataSourceApp.Update(data)
}

// DeleteVisualDataSource 删除DataSource
func (p *VisualDataSourceApi) DeleteVisualDataSource(rc *restfulx.ReqCtx) {
	sourceId := restfulx.PathParam(rc, "sourceId")
	sourceIds := strings.Split(sourceId, ",")
	p.VisualDataSourceApp.Delete(sourceIds)
}
