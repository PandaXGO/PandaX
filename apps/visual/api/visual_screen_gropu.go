package api

// ==========================================================================
// 生成日期：2023-04-10 02:51:27 +0000 UTC
// 生成路径: apps/visual/api/visual_screen_group.go
// 生成人：panda
// ==========================================================================
import (
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/utils"
	"pandax/apps/visual/entity"
	"pandax/apps/visual/services"
)

type VisualScreenGroupApi struct {
	VisualScreenGroupApp services.VisualScreenGroupModel
}

// GetScreenGroupTree ScreenGroup 树
func (p *VisualScreenGroupApi) GetScreenGroupTree(rc *restfulx.ReqCtx) {
	name := restfulx.QueryParam(rc, "name")
	status := restfulx.QueryParam(rc, "status")
	id := restfulx.QueryInt(rc, "id", 0)
	sg := entity.VisualScreenGroup{Name: name, Status: status, Id: int64(id)}
	rc.ResData = p.VisualScreenGroupApp.SelectScreenGroup(sg)
}

func (p *VisualScreenGroupApi) GetScreenGroupList(rc *restfulx.ReqCtx) {
	name := restfulx.QueryParam(rc, "name")
	status := restfulx.QueryParam(rc, "status")
	id := restfulx.QueryInt(rc, "id", 0)
	sg := entity.VisualScreenGroup{Name: name, Status: status, Id: int64(id)}

	if sg.Name == "" {
		rc.ResData = p.VisualScreenGroupApp.SelectScreenGroup(sg)
	} else {
		rc.ResData = p.VisualScreenGroupApp.FindList(sg)
	}
}

// GetScreenGroupAllList 查询所有
func (p *VisualScreenGroupApi) GetScreenGroupAllList(rc *restfulx.ReqCtx) {
	var vsg entity.VisualScreenGroup
	rc.ResData = p.VisualScreenGroupApp.FindList(vsg)
}

// GetVisualScreenGroup 获取ScreenGroup
func (p *VisualScreenGroupApi) GetVisualScreenGroup(rc *restfulx.ReqCtx) {
	id := restfulx.PathParamInt(rc, "id")
	rc.ResData = p.VisualScreenGroupApp.FindOne(int64(id))
}

// InsertVisualScreenGroup 添加ScreenGroup
func (p *VisualScreenGroupApi) InsertVisualScreenGroup(rc *restfulx.ReqCtx) {
	var data entity.VisualScreenGroup
	restfulx.BindQuery(rc, &data)

	p.VisualScreenGroupApp.Insert(data)
}

// UpdateVisualScreenGroup 修改ScreenGroup
func (p *VisualScreenGroupApi) UpdateVisualScreenGroup(rc *restfulx.ReqCtx) {
	var data entity.VisualScreenGroup
	restfulx.BindQuery(rc, &data)

	p.VisualScreenGroupApp.Update(data)
}

// DeleteVisualScreenGroup 删除ScreenGroup
func (p *VisualScreenGroupApi) DeleteVisualScreenGroup(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := utils.IdsStrToIdsIntGroup(id)
	p.VisualScreenGroupApp.Delete(ids)
}
