package api

// ==========================================================================
import (
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"pandax/pkg/tool"
	"path"
	"strings"

	"pandax/apps/device/entity"
	"pandax/apps/device/services"
)

type ProductOtaApi struct {
	ProductOtaApp services.ProductOtaModel
}

const filePath = "uploads/file"

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
	// 生成文件MD5值
	md5, err := tool.GetFileMd5(path.Join(filePath, data.Url))
	biz.ErrIsNil(err, "读取文件md5校验值错误")
	data.Id = tool.GenerateID()
	data.Check = md5
	p.ProductOtaApp.Insert(data)
}

// UpdateProductOta 修改Ota
func (p *ProductOtaApi) UpdateProductOta(rc *restfulx.ReqCtx) {
	var data entity.ProductOta
	restfulx.BindJsonAndValid(rc, &data)
	md5, err := tool.GetFileMd5(path.Join(filePath, data.Url))
	biz.ErrIsNil(err, "读取文件md5校验值错误")
	data.Check = md5
	p.ProductOtaApp.Update(data)
}

// DeleteProductOta 删除Ota
func (p *ProductOtaApi) DeleteProductOta(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	ids := strings.Split(id, ",")
	p.ProductOtaApp.Delete(ids)
}
