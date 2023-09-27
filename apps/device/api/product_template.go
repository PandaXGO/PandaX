package api

import (
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"pandax/pkg/global"
	"pandax/pkg/tool"
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
	data.Id = tool.GenerateID()
	data.OrgId = rc.LoginAccount.OrganizationId
	if data.Classify == entity.ATTRIBUTES_TSL || data.Classify == entity.TELEMETRY_TSL {
		// 向超级表及子表中添加字段
		len := 0
		stable := data.Pid + "_" + data.Classify
		if data.Classify == entity.TELEMETRY_TSL {
			if data.Type == "string" {
				len = int(data.Define["length"].(int64))
			}
		}
		err := global.TdDb.AddSTableField(stable, data.Key, data.Type, len)
		biz.ErrIsNil(err, "添加字段失败")
		// 向子表中添加字段
		go func() {
			tables, err := global.TdDb.GetListTableByStableName(stable)
			if err != nil {
				global.Log.Error(err)
				return
			}
			for _, table := range tables {
				err = global.TdDb.AddTableField(table.TableName, data.Key, data.Type, 0)
				global.Log.Error(err)
			}
		}()
	}
	p.ProductTemplateApp.Insert(data)
}

// UpdateProductTemplate 修改Template
func (p *ProductTemplateApi) UpdateProductTemplate(rc *restfulx.ReqCtx) {
	var data entity.ProductTemplate
	restfulx.BindJsonAndValid(rc, &data)
	one := p.ProductTemplateApp.FindOne(data.Id)
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
			len = int(data.Define["length"].(int64))
		}
		global.TdDb.DelSTableField(stable, one.Key)
		err := global.TdDb.AddSTableField(stable, data.Key, data.Type, len)
		biz.ErrIsNil(err, "添加字段失败")
	}
	// 向子表中添加字段
	go func() {
		tables, err := global.TdDb.GetListTableByStableName(stable)
		if err != nil {
			global.Log.Error(err)
			return
		}
		for _, table := range tables {
			global.TdDb.DelTableField(table.TableName, data.Key)
			err = global.TdDb.AddTableField(table.TableName, data.Key, data.Type, len)
			global.Log.Error(err)
		}
	}()
	p.ProductTemplateApp.Update(data)

}

// DeleteProductTemplate 删除Template
func (p *ProductTemplateApi) DeleteProductTemplate(rc *restfulx.ReqCtx) {
	id := restfulx.PathParam(rc, "id")
	data := p.ProductTemplateApp.FindOne(id)
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
	// 删除子表
	go func() {
		tables, err := global.TdDb.GetListTableByStableName(stable)
		if err != nil {
			global.Log.Error(err)
			return
		}
		for _, table := range tables {
			global.TdDb.DelTableField(table.TableName, data.Key)
			global.Log.Error(err)
		}
	}()
	ids := strings.Split(id, ",")
	p.ProductTemplateApp.Delete(ids)
}
