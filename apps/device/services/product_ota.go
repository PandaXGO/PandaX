package services

import (
	"github.com/PandaXGO/PandaKit/biz"
	"log"
	"pandax/apps/device/entity"
	"pandax/pkg/global"
)

type (
	ProductOtaModel interface {
		Insert(data entity.ProductOta) *entity.ProductOta
		FindOne(id string) *entity.ProductOta
		FindListPage(page, pageSize int, data entity.ProductOta) (*[]entity.ProductOta, int64)
		FindList(data entity.ProductOta) *[]entity.ProductOta
		Update(data entity.ProductOta) *entity.ProductOta
		Delete(ids []string)
	}

	otaModelImpl struct {
		table string
	}
)

var ProductOtaModelDao ProductOtaModel = &otaModelImpl{
	table: `product_ota`,
}

func (m *otaModelImpl) Insert(data entity.ProductOta) *entity.ProductOta {
	m.checkLatest(data)

	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加产品固件失败")
	return &data
}

func (m *otaModelImpl) checkLatest(data entity.ProductOta) {
	// 将原来的最新版更改为旧版本
	if data.IsLatest == true {
		latest := m.getLatest()
		m.updateLatest(latest.Id, false)
	}
}

func (m *otaModelImpl) FindOne(id string) *entity.ProductOta {
	resData := new(entity.ProductOta)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询产品固件失败")
	return resData
}

func (m *otaModelImpl) FindListPage(page, pageSize int, data entity.ProductOta) (*[]entity.ProductOta, int64) {
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
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询产品固件分页列表失败")
	return &list, total
}

func (m *otaModelImpl) FindList(data entity.ProductOta) *[]entity.ProductOta {
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
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询产品固件列表失败")
	return &list
}

func (m *otaModelImpl) getLatest() *entity.ProductOta {
	resData := new(entity.ProductOta)
	db := global.Db.Table(m.table).Where("is_latest = ?", true)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询产品固件失败")
	return resData
}

func (m *otaModelImpl) Update(data entity.ProductOta) *entity.ProductOta {
	m.checkLatest(data)
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改产品固件失败")
	return &data
}
func (m *otaModelImpl) updateLatest(id string, IsLatest bool) {
	log.Println(id, IsLatest)
	err := global.Db.Table(m.table).Where("id = ?", id).Update("is_latest", IsLatest).Error
	global.Log.Error("更新失败", err)
}

func (m *otaModelImpl) Delete(ids []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.ProductOta{}, "id in (?)", ids).Error, "删除产品固件失败")
}
