package services

import (
	"pandax/apps/device/entity"
	"pandax/pkg/global"
)

type (
	ProductOtaModel interface {
		Insert(data entity.ProductOta) (*entity.ProductOta, error)
		FindOne(id string) (*entity.ProductOta, error)
		FindListPage(page, pageSize int, data entity.ProductOta) (*[]entity.ProductOta, int64, error)
		FindList(data entity.ProductOta) (*[]entity.ProductOta, error)
		Update(data entity.ProductOta) (*entity.ProductOta, error)
		Delete(ids []string) error
	}

	otaModelImpl struct {
		table string
	}
)

var ProductOtaModelDao ProductOtaModel = &otaModelImpl{
	table: `product_ota`,
}

func (m *otaModelImpl) Insert(data entity.ProductOta) (*entity.ProductOta, error) {
	if err := m.checkLatest(data); err != nil {
		return nil, err
	}
	err := global.Db.Table(m.table).Create(&data).Error
	return &data, err
}

func (m *otaModelImpl) checkLatest(data entity.ProductOta) error {
	// 将原来的最新版更改为旧版本
	if data.IsLatest == true {
		latest, err := m.getLatest()
		if err != nil {
			return err
		}
		if err := m.updateLatest(latest.Id, false); err != nil {
			return err
		}
	}
	return nil
}

func (m *otaModelImpl) FindOne(id string) (*entity.ProductOta, error) {
	resData := new(entity.ProductOta)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	return resData, err
}

func (m *otaModelImpl) FindListPage(page, pageSize int, data entity.ProductOta) (*[]entity.ProductOta, int64, error) {
	list := make([]entity.ProductOta, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Description != "" {
		db = db.Where("description = ?", data.Description)
	}
	if data.Pid != "" {
		db = db.Where("pid = ?", data.Pid)
	}
	if data.Id != "" {
		db = db.Where("id = ?", data.Id)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Url != "" {
		db = db.Where("url = ?", data.Url)
	}
	if data.Version != "" {
		db = db.Where("version = ?", data.Version)
	}
	if err := db.Count(&total).Error; err != nil {
		return &list, total, err
	}
	err := db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *otaModelImpl) FindList(data entity.ProductOta) (*[]entity.ProductOta, error) {
	list := make([]entity.ProductOta, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Description != "" {
		db = db.Where("description = ?", data.Description)
	}
	if data.Pid != "" {
		db = db.Where("pid = ?", data.Pid)
	}
	if data.Id != "" {
		db = db.Where("id = ?", data.Id)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Version != "" {
		db = db.Where("version = ?", data.Version)
	}
	err := db.Order("create_time").Find(&list).Error
	return &list, err
}

func (m *otaModelImpl) getLatest() (*entity.ProductOta, error) {
	resData := new(entity.ProductOta)
	db := global.Db.Table(m.table).Where("is_latest = ?", true)
	err := db.First(resData).Error
	return resData, err
}

func (m *otaModelImpl) Update(data entity.ProductOta) (*entity.ProductOta, error) {
	if err := m.checkLatest(data); err != nil {
		return nil, err
	}
	err := global.Db.Table(m.table).Updates(&data).Error
	return &data, err
}
func (m *otaModelImpl) updateLatest(id string, IsLatest bool) error {
	return global.Db.Table(m.table).Where("id = ?", id).Update("is_latest", IsLatest).Error
}

func (m *otaModelImpl) Delete(ids []string) error {
	return global.Db.Table(m.table).Delete(&entity.ProductOta{}, "id in (?)", ids).Error
}
