package api

// ==========================================================================
// 生成日期：2023-04-10 02:51:27 +0000 UTC
// 生成路径: apps/visual/api/visual_data_source.go
// 生成人：panda
// ==========================================================================
import (
	"github.com/XM-GO/PandaKit/biz"
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/kakuilan/kgo"
	"pandax/apps/visual/driver"
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
	data.SourceType = restfulx.QueryParam(rc, "sourceType")
	data.Status = restfulx.QueryParam(rc, "status")

	list, total := p.VisualDataSourceApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetVisualDataSourceListAll DataSource列表数据
func (p *VisualDataSourceApi) GetVisualDataSourceListAll(rc *restfulx.ReqCtx) {
	data := entity.VisualDataSource{}
	data.SourceName = restfulx.QueryParam(rc, "sourceName")
	data.SourceType = restfulx.QueryParam(rc, "sourceType")
	data.Status = restfulx.QueryParam(rc, "status")
	list := p.VisualDataSourceApp.FindList(data)
	rc.ResData = list
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
	data.SourceId = kgo.KStr.Uniqid("px")
	err := driver.TestConnection(&data)
	if err != nil {
		data.Status = "0"
	} else {
		data.Status = "1"
	}
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

// GetDataSourceTest 校验数据源连接性
func (p *VisualDataSourceApi) GetDataSourceTest(rc *restfulx.ReqCtx) {
	var data entity.VisualDataSource
	restfulx.BindQuery(rc, &data)
	err := driver.TestConnection(&data)
	biz.ErrIsNilAppendErr(err, "数据库连接失败: %s")
}

// GetDataSourceTables 获取数据源下所有表
func (p *VisualDataSourceApi) GetDataSourceTables(rc *restfulx.ReqCtx) {
	sourceId := restfulx.PathParam(rc, "sourceId")
	one := p.VisualDataSourceApp.FindOne(sourceId)
	instance := driver.NewDbInstance(one)
	biz.IsTrue(instance != nil, "获取数据源下所有表失败")
	rc.ResData = instance.GetMeta().GetTableInfos()
}

// GetDataSourceTableDetails 获取表下面的所有字段
func (p *VisualDataSourceApi) GetDataSourceTableDetails(rc *restfulx.ReqCtx) {
	sourceId := restfulx.PathParam(rc, "sourceId")
	tableName := restfulx.QueryParam(rc, "tableName")
	one := p.VisualDataSourceApp.FindOne(sourceId)
	instance := driver.NewDbInstance(one)
	biz.IsTrue(instance != nil, "获取表下所有字段失败")
	rc.ResData = instance.GetMeta().GetColumns(tableName)
}
