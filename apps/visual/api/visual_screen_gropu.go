package api

// ==========================================================================
// 生成日期：2023-04-10 02:51:27 +0000 UTC
// 生成路径: apps/visual/api/visual_screen_group.go
// 生成人：panda
// ==========================================================================
import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"strings"

	"pandax/apps/visual/entity"
	"pandax/apps/visual/services"
)

type VisualScreenGroupApi struct {
	VisualScreenGroupApp services.VisualScreenGroupModel
}

// GetVisualScreenGroupList DataSetGroup列表数据
func (p *VisualScreenGroupApi) GetVisualScreenGroupList(rc *restfulx.ReqCtx) {
	data := entity.VisualScreenGroup{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.Name = restfulx.QueryParam(rc, "name")

	list, total := p.VisualScreenGroupApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetVisualScreenGroup 获取DataSetGroup
func (p *VisualScreenGroupApi) GetVisualScreenGroup(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	rc.ResData = p.VisualScreenGroupApp.FindOne(id)
}

// InsertVisualScreenGroup 添加DataSetGroup
func (p *VisualScreenGroupApi) InsertVisualScreenGroup(rc *restfulx.ReqCtx) {
	var data entity.VisualScreenGroup
	restfulx.BindQuery(rc, &data)

	p.VisualScreenGroupApp.Insert(data)
}

// UpdateVisualScreenGroup 修改DataSetGroup
func (p *VisualScreenGroupApi) UpdateVisualScreenGroup(rc *restfulx.ReqCtx) {
	var data entity.VisualScreenGroup
	restfulx.BindQuery(rc, &data)

	p.VisualScreenGroupApp.Update(data)
}

// DeleteVisualScreenGroup 删除DataSetGroup
func (p *VisualScreenGroupApi) DeleteVisualScreenGroup(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	p.VisualScreenGroupApp.Delete(ids)
}
