package services

import (
	"pandax/apps/device/entity"
	"pandax/pkg/global"
	"strings"
)

type (
	ProductTemplateModel interface {
		Insert(data entity.ProductTemplate) (*entity.ProductTemplate, error)
		FindOne(id string) (*entity.ProductTemplate, error)
		FindOneByKey(deviceId, key string) (*entity.ProductTemplate, error)
		FindListPage(page, pageSize int, data entity.ProductTemplate) (*[]entity.ProductTemplate, int64, error)
		FindListAttrs(data entity.ProductTemplate) (*[]entity.ProductTemplate, error)
		FindList(data entity.ProductTemplate) (*[]entity.ProductTemplate, error)
		Update(data entity.ProductTemplate) (*entity.ProductTemplate, error)
		Delete(ids []string) error
	}

	templateModelImpl struct {
		table string
	}
)

var ProductTemplateModelDao ProductTemplateModel = &templateModelImpl{
	table: `product_templates`,
}

func (m *templateModelImpl) Insert(data entity.ProductTemplate) (*entity.ProductTemplate, error) {
	err := global.Db.Table(m.table).Create(&data).Error
	return &data, err
}

func (m *templateModelImpl) FindOne(id string) (*entity.ProductTemplate, error) {
	resData := new(entity.ProductTemplate)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	return resData, err
}

func (m *templateModelImpl) FindOneByKey(deviceId, key string) (*entity.ProductTemplate, error) {
	resData := new(entity.ProductTemplate)
	db := global.Db.Table(m.table).Where("pid = ?", deviceId).Where("key = ?", key)
	err := db.First(resData).Error
	return resData, err
}

func (m *templateModelImpl) FindListPage(page, pageSize int, data entity.ProductTemplate) (*[]entity.ProductTemplate, int64, error) {
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
	if err := db.Count(&total).Error; err != nil {
		return &list, total, err
	}
	err := db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *templateModelImpl) FindListAttrs(data entity.ProductTemplate) (*[]entity.ProductTemplate, error) {
	list := make([]entity.ProductTemplate, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Pid != "" {
		db = db.Where("pid = ?", data.Pid)
	}
	db = db.Where("classify in (?)", []string{"attributes", "telemetry"})
	err := db.Order("create_time").Find(&list).Error
	return &list, err
}

func (m *templateModelImpl) FindList(data entity.ProductTemplate) (*[]entity.ProductTemplate, error) {
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
	err := db.Order("create_time").Find(&list).Error
	return &list, err
}

func (m *templateModelImpl) Update(data entity.ProductTemplate) (*entity.ProductTemplate, error) {
	err := global.Db.Table(m.table).Updates(&data).Error
	return &data, err
}

func (m *templateModelImpl) Delete(ids []string) error {
	return global.Db.Table(m.table).Delete(&entity.ProductTemplate{}, "id in (?)", ids).Error
}
