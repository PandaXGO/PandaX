package api

import (
	"encoding/json"
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
	"pandax/pkg/global"
	"strings"

	"pandax/apps/device/entity"
	"pandax/apps/device/services"
)

type ProductTemplateApi struct {
	ProductTemplateApp services.ProductTemplateModel
}

// GetProductTemplateList Template列表数据
func (p *ProductTemplateApi) GetProductTemplateList(rc *restfulx.ReqCtx) {
	data := entity.ProductTemplate{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	data.Pid = restfulx.QueryParam(rc, "pid")
	data.Classify = restfulx.QueryParam(rc, "classify")
	data.Name = restfulx.QueryParam(rc, "name")

	list, total, err := p.ProductTemplateApp.FindListPage(pageNum, pageSize, data)
	biz.ErrIsNil(err, "查询产品模板列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// GetProductTemplateListAll Template所有列表数据
func (p *ProductTemplateApi) GetProductTemplateListAll(rc *restfulx.ReqCtx) {
	data := entity.ProductTemplate{}
	data.Pid = restfulx.QueryParam(rc, "pid")
	data.Classify = restfulx.QueryParam(rc, "classify")
	data.Name = restfulx.QueryParam(rc, "name")
	list, err := p.ProductTemplateApp.FindList(data)
	biz.ErrIsNil(err, "查询产品模板列表失败")
	rc.ResData = list
}

// GetProductTemplate 获取Template
func (p *ProductTemplateApi) GetProductTemplate(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	one, err := p.ProductTemplateApp.FindOne(id)
	biz.ErrIsNil(err, "查询产品模板失败")
	rc.ResData = one
}

// InsertProductTemplate 添加Template
func (p *ProductTemplateApi) InsertProductTemplate(rc *restfulx.ReqCtx) {
	var data entity.ProductTemplate
	restfulx.BindJsonAndValid(rc, &data)
	data.Id = utils.GenerateID("pt")
	if data.Classify == entity.ATTRIBUTES_TSL || data.Classify == entity.TELEMETRY_TSL {
		// 向超级表及子表中添加字段
		len := 100
		stable := data.Pid + "_" + data.Classify
		if data.Classify == entity.TELEMETRY_TSL {
			if data.Type == "string" {
				if maxLength, ok := data.Define["maxLength"].(json.Number); ok {
					length, _ := maxLength.Int64()
					len = int(length)
				}
			}
		}
		err := global.TdDb.AddSTableField(stable, data.Key, data.Type, len)
		biz.ErrIsNil(err, "添加字段失败")
	}
	p.ProductTemplateApp.Insert(data)
}

// UpdateProductTemplate 修改Template
func (p *ProductTemplateApi) UpdateProductTemplate(rc *restfulx.ReqCtx) {
	var data entity.ProductTemplate
	restfulx.BindJsonAndValid(rc, &data)
	one, err := p.ProductTemplateApp.FindOne(data.Id)
	biz.ErrIsNil(err, "查询产品模板失败")
	stable := ""
	len := 0
	if data.Classify == entity.ATTRIBUTES_TSL {
		stable = data.Pid + "_" + entity.ATTRIBUTES_TSL
		// 删除原来的key
		global.TdDb.DelSTableField(stable, one.Key)
		// 添加修改完成的key
		err := global.TdDb.AddSTableField(stable, data.Key, data.Type, len)
		biz.ErrIsNil(err, "添加字段失败")
	}
	if data.Classify == entity.TELEMETRY_TSL {
		stable = data.Pid + "_" + entity.TELEMETRY_TSL
		if data.Type == "string" {
			if maxLength, ok := data.Define["maxLength"].(json.Number); ok {
				length, _ := maxLength.Int64()
				len = int(length)
			}
		}
		global.TdDb.DelSTableField(stable, one.Key)
		err := global.TdDb.AddSTableField(stable, data.Key, data.Type, len)
		biz.ErrIsNil(err, "添加字段失败")
	}
	p.ProductTemplateApp.Update(data)

}

// DeleteProductTemplate 删除Template
func (p *ProductTemplateApi) DeleteProductTemplate(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	data, err := p.ProductTemplateApp.FindOne(id)
	biz.ErrIsNil(err, "查询产品模板失败")
	// 向超级表中删除字段
	stable := ""
	if data.Classify == entity.ATTRIBUTES_TSL {
		stable = data.Pid + "_" + entity.ATTRIBUTES_TSL
		err := global.TdDb.DelSTableField(stable, data.Key)
		biz.ErrIsNil(err, "删除字段失败")
	}
	if data.Classify == entity.TELEMETRY_TSL {
		stable = data.Pid + "_" + entity.TELEMETRY_TSL
		err := global.TdDb.DelSTableField(data.Pid+"_"+entity.TELEMETRY_TSL, data.Key)
		biz.ErrIsNil(err, "删除字段失败")
	}
	ids := strings.Split(id, ",")
	p.ProductTemplateApp.Delete(ids)
}
