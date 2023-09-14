package api

// ==========================================================================
import (
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/kakuilan/kgo"
	"strings"

	"pandax/apps/device/entity"
	"pandax/apps/device/services"
)

type ProductOtaApi struct {
	ProductOtaApp services.ProductOtaModel
}

// GetProductOtaList Ota列表数据
func (p *ProductOtaApi) GetProductOtaList(rc *restfulx.ReqCtx) {
	data := entity.ProductOta{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.Name = restfulx.QueryParam(rc, "name")
	data.Pid = restfulx.QueryParam(rc, "pid")

	list, total := p.ProductOtaApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// GetProductOta 获取Ota
func (p *ProductOtaApi) GetProductOta(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	rc.ResData = p.ProductOtaApp.FindOne(id)
}

// InsertProductOta 添加Ota
func (p *ProductOtaApi) InsertProductOta(rc *restfulx.ReqCtx) {
	var data entity.ProductOta
	restfulx.BindJsonAndValid(rc, &data)
	data.Id = kgo.KStr.Uniqid("ota_")
	p.ProductOtaApp.Insert(data)
}

// UpdateProductOta 修改Ota
func (p *ProductOtaApi) UpdateProductOta(rc *restfulx.ReqCtx) {
	var data entity.ProductOta
	restfulx.BindJsonAndValid(rc, &data)

	p.ProductOtaApp.Update(data)
}

// DeleteProductOta 删除Ota
func (p *ProductOtaApi) DeleteProductOta(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	p.ProductOtaApp.Delete(ids)
}
