package services

import (
	"pandax/apps/device/entity"
	"pandax/pkg/global"
)

type (
	ProductModel interface {
		Insert(data entity.Product) (*entity.Product, error)
		FindOne(id string) (*entity.ProductRes, error)
		FindListPage(page, pageSize int, data entity.Product) (*[]entity.ProductRes, int64, error)
		FindList(data entity.Product) (*[]entity.ProductRes, error)
		Update(data entity.Product) (*entity.Product, error)
		Delete(ids []string) error
		FindProductCount() (entity.DeviceCount, error)
	}

	productModelImpl struct {
		table string
	}
)

var ProductModelDao ProductModel = &productModelImpl{
	table: `products`,
}

func (m *productModelImpl) Insert(data entity.Product) (*entity.Product, error) {
	tx := global.Db.Begin()
	// 添加产品及规则链到redis中
	if err := tx.Table(m.table).Create(&data).Error; err != nil {
		return nil, err
	}
	// 创建taos数据库超级表
	err := createDeviceStable(data.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &data, nil
}

func (m *productModelImpl) FindOne(id string) (*entity.ProductRes, error) {
	resData := new(entity.ProductRes)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.Preload("ProductCategory").First(resData).Error
	return resData, err
}

func (m *productModelImpl) FindListPage(page, pageSize int, data entity.Product) (*[]entity.ProductRes, int64, error) {
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
	if data.ProtocolName != "" {
		db = db.Where("protocol_name like ?", "%"+data.ProtocolName+"%")
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.RuleChainId != "" {
		db = db.Where("rule_chain_id = ?", data.RuleChainId)
	}

	if err := db.Count(&total).Error; err != nil {
		return &list, total, err
	}
	err := db.Order("create_time").Preload("ProductCategory").Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *productModelImpl) FindList(data entity.Product) (*[]entity.ProductRes, error) {
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
	if data.ProtocolName != "" {
		db = db.Where("protocol_name like ?", "%"+data.ProtocolName+"%")
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.RuleChainId != "" {
		db = db.Where("rule_chain_id = ?", data.RuleChainId)
	}
	err := db.Order("create_time").Preload("ProductCategory").Find(&list).Error
	return &list, err
}

func (m *productModelImpl) Update(data entity.Product) (*entity.Product, error) {
	// go的一些默认值 int 0 bool false  保存失败需要先转成map
	err := global.Db.Table(m.table).Where("id = ?", data.Id).Updates(data).Error
	return &data, err
}

func (m *productModelImpl) Delete(ids []string) error {
	if err := global.Db.Table(m.table).Delete(&entity.Product{}, "id in (?)", ids).Error; err != nil {
		return err
	}
	return nil
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

// 获取产品数量统计
func (m *productModelImpl) FindProductCount() (count entity.DeviceCount, err error) {
	sql := `SELECT COUNT(*) AS total, (SELECT COUNT(*) FROM products WHERE DATE(create_time) = CURDATE()) AS today  FROM products`
	if global.Conf.Server.DbType == "postgresql" {
		sql = `SELECT COUNT(*) AS total, (SELECT COUNT(*) FROM products WHERE DATE(create_time) = current_date) AS today  FROM products`
	}
	err = global.Db.Raw(sql).Scan(&count).Error
	return
}
