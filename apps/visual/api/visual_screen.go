package api

// ==========================================================================
// 生成日期：2023-04-10 11:26:49 +0800 CST
// 生成路径: apps/visual/api/visual_screen.go
// 生成人：panda
// ==========================================================================
import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"strings"

	"pandax/apps/visual/entity"
	"pandax/apps/visual/services"
)

type VisualScreenApi struct {
	VisualScreenApp services.VisualScreenModel
}

// GetVisualScreenList Screen列表数据
func (p *VisualScreenApi) GetVisualScreenList(rc *restfulx.ReqCtx) {
	data := entity.VisualScreen{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.ScreenName = restfulx.QueryParam(rc, "screenName")
	data.Status = restfulx.QueryParam(rc, "status")

	list, total := p.VisualScreenApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetVisualScreen 获取Screen
func (p *VisualScreenApi) GetVisualScreen(rc *restfulx.ReqCtx) {
	screenId := restfulx.PathParam(rc, "screenId")
	rc.ResData = p.VisualScreenApp.FindOne(screenId)
}

// InsertVisualScreen 添加Screen
func (p *VisualScreenApi) InsertVisualScreen(rc *restfulx.ReqCtx) {
	var data entity.VisualScreen
	restfulx.BindQuery(rc, &data)

	p.VisualScreenApp.Insert(data)
}

// UpdateVisualScreen 修改Screen
func (p *VisualScreenApi) UpdateVisualScreen(rc *restfulx.ReqCtx) {
	var data entity.VisualScreen
	restfulx.BindQuery(rc, &data)

	p.VisualScreenApp.Update(data)
}

// DeleteVisualScreen 删除Screen
func (p *VisualScreenApi) DeleteVisualScreen(rc *restfulx.ReqCtx) {
	screenId := restfulx.PathParam(rc, "screenId")
	screenIds := strings.Split(screenId, ",")
	p.VisualScreenApp.Delete(screenIds)
}
