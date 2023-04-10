package api

// ==========================================================================
// 生成日期：2023-04-10 02:51:27 +0000 UTC
// 生成路径: apps/visual/api/visual_data_set_group.go
// 生成人：panda
// ==========================================================================
import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"strings"

	"pandax/apps/visual/entity"
	"pandax/apps/visual/services"
)

type VisualDataSetGroupApi struct {
	VisualDataSetGroupApp services.VisualDataSetGroupModel
}

// GetVisualDataSetGroupList DataSetGroup列表数据
func (p *VisualDataSetGroupApi) GetVisualDataSetGroupList(rc *restfulx.ReqCtx) {
	data := entity.VisualDataSetGroup{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.Name = restfulx.QueryParam(rc, "name")

	list, total := p.VisualDataSetGroupApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetVisualDataSetGroup 获取DataSetGroup
func (p *VisualDataSetGroupApi) GetVisualDataSetGroup(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	rc.ResData = p.VisualDataSetGroupApp.FindOne(id)
}

// InsertVisualDataSetGroup 添加DataSetGroup
func (p *VisualDataSetGroupApi) InsertVisualDataSetGroup(rc *restfulx.ReqCtx) {
	var data entity.VisualDataSetGroup
	restfulx.BindQuery(rc, &data)

	p.VisualDataSetGroupApp.Insert(data)
}

// UpdateVisualDataSetGroup 修改DataSetGroup
func (p *VisualDataSetGroupApi) UpdateVisualDataSetGroup(rc *restfulx.ReqCtx) {
	var data entity.VisualDataSetGroup
	restfulx.BindQuery(rc, &data)

	p.VisualDataSetGroupApp.Update(data)
}

// DeleteVisualDataSetGroup 删除DataSetGroup
func (p *VisualDataSetGroupApi) DeleteVisualDataSetGroup(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	p.VisualDataSetGroupApp.Delete(ids)
}
