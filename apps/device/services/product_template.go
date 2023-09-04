package services

import (
	"github.com/PandaXGO/PandaKit/biz"
	"pandax/apps/device/entity"
	"pandax/pkg/global"
	"strings"
)

type (
	ProductTemplateModel interface {
		Insert(data entity.ProductTemplate) *entity.ProductTemplate
		FindOne(id string) *entity.ProductTemplate
		FindListPage(page, pageSize int, data entity.ProductTemplate) (*[]entity.ProductTemplate, int64)
		FindListAttrs(data entity.ProductTemplate) *[]entity.ProductTemplate
		FindList(data entity.ProductTemplate) *[]entity.ProductTemplate
		Update(data entity.ProductTemplate) *entity.ProductTemplate
		Delete(ids []string)
	}

	templateModelImpl struct {
		table string
	}
)

var ProductTemplateModelDao ProductTemplateModel = &templateModelImpl{
	table: `product_templates`,
}

func (m *templateModelImpl) Insert(data entity.ProductTemplate) *entity.ProductTemplate {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加产品模型失败")
	return &data
}

func (m *templateModelImpl) FindOne(id string) *entity.ProductTemplate {
	resData := new(entity.ProductTemplate)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询产品模型失败")
	return resData
}

func (m *templateModelImpl) FindListPage(page, pageSize int, data entity.ProductTemplate) (*[]entity.ProductTemplate, int64) {
	list := make([]entity.ProductTemplate, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Classify != "" {
		db = db.Where("classify in (?)", strings.Split(data.Classify, ","))
	}
	if data.Key != "" {
		db = db.Where("key = ?", data.Key)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Type != "" {
		db = db.Where("type = ?", data.Type)
	}
	if data.Pid != "" {
		db = db.Where("pid = ?", data.Pid)
	}
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询产品模型分页列表失败")
	return &list, total
}

func (m *templateModelImpl) FindListAttrs(data entity.ProductTemplate) *[]entity.ProductTemplate {
	list := make([]entity.ProductTemplate, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Pid != "" {
		db = db.Where("pid = ?", data.Pid)
	}
	db = db.Where("classify in (?)", []string{"attributes", "telemetry"})
	err := db.Order("create_time").Find(&list).Error
	biz.ErrIsNil(err, "查询产品模型分页列表失败")
	return &list
}

func (m *templateModelImpl) FindList(data entity.ProductTemplate) *[]entity.ProductTemplate {
	list := make([]entity.ProductTemplate, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Classify != "" {
		db = db.Where("classify in (?)", strings.Split(data.Classify, ","))
	}
	if data.Key != "" {
		db = db.Where("key = ?", data.Key)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Type != "" {
		db = db.Where("type = ?", data.Type)
	}
	if data.Pid != "" {
		db = db.Where("pid = ?", data.Pid)
	}
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询产品模型列表失败")
	return &list
}

func (m *templateModelImpl) Update(data entity.ProductTemplate) *entity.ProductTemplate {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改产品模型失败")
	return &data
}

func (m *templateModelImpl) Delete(ids []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.ProductTemplate{}, "id in (?)", ids).Error, "删除产品模型失败")
}
