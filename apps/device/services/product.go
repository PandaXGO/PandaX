package services

import (
	"context"
	"github.com/PandaXGO/PandaKit/biz"
	"log"
	"pandax/apps/device/entity"
	"pandax/pkg/global"
)

type (
	ProductModel interface {
		Insert(data entity.Product) *entity.Product
		FindOne(id string) *entity.ProductRes
		FindListPage(page, pageSize int, data entity.Product) (*[]entity.ProductRes, int64)
		FindList(data entity.Product) *[]entity.ProductRes
		Update(data entity.Product) *entity.Product
		Delete(ids []string)
	}

	productModelImpl struct {
		table string
	}
)

var ProductModelDao ProductModel = &productModelImpl{
	table: `products`,
}

func (m *productModelImpl) Insert(data entity.Product) *entity.Product {
	tx := global.Db.Begin()
	// 添加产品及规则链到redis中
	err := tx.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加产品失败")
	// 创建taos数据库超级表
	err = createDeviceStable(data.Id)
	if err != nil {
		tx.Rollback()
		biz.ErrIsNil(err, "添加设备失败，超级表创建失败")
	}
	tx.Commit()
	return &data
}

func (m *productModelImpl) FindOne(id string) *entity.ProductRes {
	resData := new(entity.ProductRes)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.Preload("ProductCategory").First(resData).Error
	biz.ErrIsNil(err, "查询产品失败")
	return resData
}

func (m *productModelImpl) FindListPage(page, pageSize int, data entity.Product) (*[]entity.ProductRes, int64) {
	list := make([]entity.ProductRes, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.DeviceType != "" {
		db = db.Where("device_type = ?", data.DeviceType)
	}
	if data.ProductCategoryId != "" {
		db = db.Where("product_category_id = ?", data.ProductCategoryId)
	}
	if data.Owner != "" {
		db = db.Where("owner = ?", data.Owner)
	}
	if data.ProtocolName != "" {
		db = db.Where("protocol_name like ?", "%"+data.ProtocolName+"%")
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.RuleChainId != "" {
		db = db.Where("rule_chain_id = ?", data.RuleChainId)
	}
	err := db.Count(&total).Error
	err = db.Order("create_time").Preload("ProductCategory").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询产品分页列表失败")
	return &list, total
}

func (m *productModelImpl) FindList(data entity.Product) *[]entity.ProductRes {
	list := make([]entity.ProductRes, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.DeviceType != "" {
		db = db.Where("device_type = ?", data.DeviceType)
	}
	if data.ProductCategoryId != "" {
		db = db.Where("product_category_id = ?", data.ProductCategoryId)
	}
	if data.Owner != "" {
		db = db.Where("owner = ?", data.Owner)
	}
	if data.ProtocolName != "" {
		db = db.Where("protocol_name like ?", "%"+data.ProtocolName+"%")
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.RuleChainId != "" {
		db = db.Where("rule_chain_id = ?", data.RuleChainId)
	}
	biz.ErrIsNil(db.Order("create_time").Preload("ProductCategory").Find(&list).Error, "查询产品列表失败")
	return &list
}

func (m *productModelImpl) Update(data entity.Product) *entity.Product {
	// go的一些默认值 int 0 bool false  保存失败需要先转成map
	err := global.Db.Table(m.table).Where("id = ?", data.Id).Updates(data).Error
	log.Println("update", err)
	//biz.ErrIsNil(, "修改产品失败")
	return &data
}

func (m *productModelImpl) Delete(ids []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.Product{}, "id in (?)", ids).Error, "删除产品失败")
	for _, id := range ids {
		// 删除超级表
		err := deleteDeviceStable(id)
		global.Log.Error("时序数据库超级表删除失败", err)
		// 删除所有缓存
		global.RedisDb.Del(context.Background(), id)
	}
}

func createDeviceStable(productId string) error {
	err := global.TdDb.CreateStable(productId + "_" + entity.ATTRIBUTES_TSL)
	if err != nil {
		return err
	}
	err = global.TdDb.CreateStable(productId + "_" + entity.TELEMETRY_TSL)
	if err != nil {
		return err
	}
	return nil
}

func deleteDeviceStable(productId string) error {
	err := global.TdDb.DropStable(productId + "_" + entity.ATTRIBUTES_TSL)
	if err != nil {
		return err
	}
	err = global.TdDb.DropStable(productId + "_" + entity.TELEMETRY_TSL)
	if err != nil {
		return err
	}
	return nil
}
