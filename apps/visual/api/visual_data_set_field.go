package api

// ==========================================================================
// 生成日期：2023-04-10 03:05:21 +0800 CST
// 生成路径: apps/visual/api/visual_data_set_field.go
// 生成人：panda
// ==========================================================================
import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"strings"

	"pandax/apps/visual/entity"
	"pandax/apps/visual/services"
)

type VisualDataSetFieldApi struct {
	VisualDataSetFieldApp services.VisualDataSetFieldModel
}

// GetVisualDataSetFieldList DataSetField列表数据
func (p *VisualDataSetFieldApi) GetVisualDataSetFieldList(rc *restfulx.ReqCtx) {
	data := entity.VisualDataSetField{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.Name = restfulx.QueryParam(rc, "name")

	list, total := p.VisualDataSetFieldApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetVisualDataSetField 获取DataSetField
func (p *VisualDataSetFieldApi) GetVisualDataSetField(rc *restfulx.ReqCtx) {
	fieldId := restfulx.PathParam(rc, "fieldId")
	rc.ResData = p.VisualDataSetFieldApp.FindOne(fieldId)
}

// InsertVisualDataSetField 添加DataSetField
func (p *VisualDataSetFieldApi) InsertVisualDataSetField(rc *restfulx.ReqCtx) {
	var data entity.VisualDataSetField
	restfulx.BindQuery(rc, &data)

	p.VisualDataSetFieldApp.Insert(data)
}

// UpdateVisualDataSetField 修改DataSetField
func (p *VisualDataSetFieldApi) UpdateVisualDataSetField(rc *restfulx.ReqCtx) {
	var data entity.VisualDataSetField
	restfulx.BindQuery(rc, &data)

	p.VisualDataSetFieldApp.Update(data)
}

// DeleteVisualDataSetField 删除DataSetField
func (p *VisualDataSetFieldApi) DeleteVisualDataSetField(rc *restfulx.ReqCtx) {
	fieldId := restfulx.PathParam(rc, "fieldId")
	fieldIds := strings.Split(fieldId, ",")
	p.VisualDataSetFieldApp.Delete(fieldIds)
}
