package api

import (
	"github.com/XM-GO/PandaKit/biz"
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/kakuilan/kgo"
	"log"
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

	list, total := p.ProductTemplateApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetProductTemplateListAll Template所有列表数据
func (p *ProductTemplateApi) GetProductTemplateListAll(rc *restfulx.ReqCtx) {
	data := entity.ProductTemplate{}
	data.Pid = restfulx.QueryParam(rc, "pid")
	data.Classify = restfulx.QueryParam(rc, "classify")
	data.Name = restfulx.QueryParam(rc, "name")
	list := p.ProductTemplateApp.FindList(data)
	rc.ResData = list
}

// GetProductTemplate 获取Template
func (p *ProductTemplateApi) GetProductTemplate(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	rc.ResData = p.ProductTemplateApp.FindOne(id)
}

// InsertProductTemplate 添加Template
func (p *ProductTemplateApi) InsertProductTemplate(rc *restfulx.ReqCtx) {
	var data entity.ProductTemplate
	restfulx.BindJsonAndValid(rc, &data)
	data.Id = kgo.KStr.Uniqid("tm_")
	p.ProductTemplateApp.Insert(data)
	// 向超级表中添加字段
	if data.Classify == entity.ATTRIBUTES_TSL {
		err := global.TdDb.AddSTableField(data.Pid+"_"+entity.ATTRIBUTES_TSL, data.Key, data.Type, 0)
		biz.ErrIsNil(err, "添加字段失败")
	}
	if data.Classify == entity.TELEMETRY_TSL {
		len := 0
		if data.Type == "string" {
			len = int(data.Define["length"].(int64))
		}
		err := global.TdDb.AddSTableField(data.Pid+"_"+entity.TELEMETRY_TSL, data.Key, data.Type, len)
		biz.ErrIsNil(err, "添加字段失败")
	}
	// todo 已经创建的表，超级表字段更新时，需要同步设备表字段
}

// UpdateProductTemplate 修改Template
func (p *ProductTemplateApi) UpdateProductTemplate(rc *restfulx.ReqCtx) {
	var data entity.ProductTemplate
	restfulx.BindJsonAndValid(rc, &data)
	one := p.ProductTemplateApp.FindOne(data.Id)
	if data.Classify == entity.ATTRIBUTES_TSL {
		// 删除原来的key
		global.TdDb.DelSTableField(data.Pid+"_"+entity.ATTRIBUTES_TSL, one.Key)
		// 添加修改完成的key
		err := global.TdDb.AddSTableField(data.Pid+"_"+entity.ATTRIBUTES_TSL, data.Key, data.Type, 0)
		biz.ErrIsNil(err, "添加字段失败")
	}
	if data.Classify == entity.TELEMETRY_TSL {
		len := 0
		if data.Type == "string" {
			len = int(data.Define["length"].(int64))
		}
		global.TdDb.DelSTableField(data.Pid+"_"+entity.TELEMETRY_TSL, one.Key)
		err := global.TdDb.AddSTableField(data.Pid+"_"+entity.TELEMETRY_TSL, data.Key, data.Type, len)
		log.Println(err)
		biz.ErrIsNil(err, "添加字段失败")
	}

	p.ProductTemplateApp.Update(data)
}

// DeleteProductTemplate 删除Template
func (p *ProductTemplateApi) DeleteProductTemplate(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	data := p.ProductTemplateApp.FindOne(id)
	// 向超级表中添加字段
	if data.Classify == entity.ATTRIBUTES_TSL {
		err := global.TdDb.DelSTableField(data.Pid+"_"+entity.ATTRIBUTES_TSL, data.Key)
		biz.ErrIsNil(err, "删除字段失败")
	}
	if data.Classify == entity.TELEMETRY_TSL {
		err := global.TdDb.DelSTableField(data.Pid+"_"+entity.TELEMETRY_TSL, data.Key)
		biz.ErrIsNil(err, "删除字段失败")
	}
	ids := strings.Split(id, ",")
	p.ProductTemplateApp.Delete(ids)
}
