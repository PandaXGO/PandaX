package services

import (
	"errors"
	"pandax/apps/device/entity"
	"pandax/pkg/cache"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
	"time"
)

type (
	DeviceModel interface {
		Insert(data entity.Device) (*entity.Device, error)
		FindOneByToken(token string) (*entity.DeviceRes, error)
		FindOneByName(name string) (*entity.DeviceRes, error)
		FindOne(id string) (*entity.DeviceRes, error)
		FindListPage(page, pageSize int, data entity.Device) (*[]entity.DeviceRes, int64, error)
		FindList(data entity.Device) (*[]entity.DeviceRes, error)
		Update(data entity.Device) (*entity.Device, error)
		UpdateStatus(id, linkStatus string) error
		Delete(ids []string) error
		FindDeviceCount() (entity.DeviceCount, error)
		FindDeviceCountGroupByLinkStatus() ([]entity.DeviceCountLinkStatus, error)
		FindDeviceCountGroupByType() ([]entity.DeviceCountType, error)
	}

	deviceModelImpl struct {
		table string
	}
)

var DeviceModelDao DeviceModel = &deviceModelImpl{
	table: `devices`,
}

func (m *deviceModelImpl) Insert(data entity.Device) (*entity.Device, error) {
	tx := global.Db.Begin()
	//1 检查设备名称是否存在
	list, _ := m.FindList(entity.Device{Name: data.Name})
	if list != nil && len(*list) > 0 {
		return nil, errors.New("设备名称已经存在")
	}
	//2 创建认证TOKEN IOTHUB使用
	token := GetDeviceToken(&data)
	// 子网关不需要设置token
	if data.DeviceType == global.GATEWAYS {
		err := cache.SetDeviceEtoken(data.Name, token.GetMarshal(), time.Hour*24*365)
		if err != nil {
			return nil, err
		}
	} else {
		data.Token = token.MD5ID()
		err := cache.SetDeviceEtoken(data.Token, token.GetMarshal(), time.Hour*24*365)
		if err != nil {
			return nil, err
		}
	}
	//3 添加设备
	err := tx.Table(m.table).Create(&data).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// 创建超级表 失败就
	if data.Pid != "" {
		err = createDeviceTable(data.Pid, data.Name)
		if err != nil {
			tx.Rollback()
			return nil, errors.New("时序数据设备表创建失败")
		}
	}
	tx.Commit()
	return &data, nil
}

func (m *deviceModelImpl) FindOne(id string) (*entity.DeviceRes, error) {
	resData := new(entity.DeviceRes)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.Preload("Product").Preload("DeviceGroup").First(resData).Error
	return resData, err
}

func (m *deviceModelImpl) FindOneByName(name string) (*entity.DeviceRes, error) {
	resData := new(entity.DeviceRes)
	db := global.Db.Table(m.table).Where("name = ?", name)
	err := db.First(resData).Error
	return resData, err
}

func (m *deviceModelImpl) FindOneByToken(token string) (*entity.DeviceRes, error) {
	resData := new(entity.DeviceRes)
	db := global.Db.Table(m.table).Preload("Product").Where("token = ?", token)
	err := db.First(resData).Error
	return resData, err
}

func (m *deviceModelImpl) FindListPage(page, pageSize int, data entity.Device) (*[]entity.DeviceRes, int64, error) {
	list := make([]entity.DeviceRes, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Alias != "" {
		db = db.Where("alias = ?", data.Alias)
	}
	if data.Gid != "" {
		db = db.Where("gid = ?", data.Gid)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.LinkStatus != "" {
		db = db.Where("Link_status = ?", data.LinkStatus)
	}
	if data.Pid != "" {
		db = db.Where("pid = ?", data.Pid)
	}
	if data.ParentId != "" {
		db = db.Where("parent_id = ?", data.ParentId)
	}
	// 组织数据访问权限
	if err := model.OrgAuthSet(db, data.RoleId, data.Owner); err != nil {
		return &list, total, err
	}
	if err := db.Count(&total).Error; err != nil {
		return &list, total, err
	}
	err := db.Order("create_time").Preload("Product").Preload("DeviceGroup").Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *deviceModelImpl) FindList(data entity.Device) (*[]entity.DeviceRes, error) {
	list := make([]entity.DeviceRes, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Alias != "" {
		db = db.Where("alias = ?", data.Alias)
	}
	if data.Gid != "" {
		db = db.Where("gid = ?", data.Gid)
	}
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.DeviceType != "" {
		db = db.Where("device_type = ?", data.DeviceType)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.LinkStatus != "" {
		db = db.Where("Link_status = ?", data.LinkStatus)
	}
	if data.Pid != "" {
		db = db.Where("pid = ?", data.Pid)
	}
	if data.ParentId != "" {
		db = db.Where("parent_id = ?", data.ParentId)
	}
	if err := model.OrgAuthSet(db, data.RoleId, data.Owner); err != nil {
		return nil, err
	}
	db.Preload("Product").Preload("DeviceGroup")
	err := db.Order("create_time").Find(&list).Error
	return &list, err
}

func (m *deviceModelImpl) Update(data entity.Device) (*entity.Device, error) {
	token := GetDeviceToken(&data)
	if data.DeviceType == global.GATEWAYS {
		err := cache.SetDeviceEtoken(data.Name, token.GetMarshal(), time.Hour*24*365)
		if err != nil {
			return nil, errors.New("设备更改缓存失败")
		}
	} else {
		err := cache.SetDeviceEtoken(data.Token, token.GetMarshal(), time.Hour*24*365)
		if err != nil {
			return nil, errors.New("设备更改缓存失败")
		}
	}
	err := global.Db.Table(m.table).Updates(&data).Error
	return &data, err
}
func (m *deviceModelImpl) UpdateStatus(id, linkStatus string) error {
	return global.Db.Table(m.table).Where("id", id).Update("link_status", linkStatus).Update("last_time", time.Now()).Error
}

func (m *deviceModelImpl) Delete(ids []string) error {
	if err := global.Db.Table(m.table).Delete(&entity.Device{}, "id in (?)", ids).Error; err != nil {
		return err
	}
	for _, id := range ids {
		device, err := m.FindOne(id)
		if err != nil {
			continue
		}
		// 删除表
		err = deleteDeviceTable(device.Name)
		// 删除所有缓存
		if device.DeviceType == global.GATEWAYS {
			cache.DelDeviceEtoken(device.Name)
		} else {
			cache.DelDeviceEtoken(device.Token)
		}
	}
	return nil
}

// 创建Tdengine时序数据
func createDeviceTable(productId, device string) error {
	err := global.TdDb.CreateTable(productId+"_"+entity.ATTRIBUTES_TSL, device+"_"+entity.ATTRIBUTES_TSL)
	if err != nil {
		return err
	}
	err = global.TdDb.CreateTable(productId+"_"+entity.TELEMETRY_TSL, device+"_"+entity.TELEMETRY_TSL)
	if err != nil {
		return err
	}
	return nil
}

// 删除Tdengine时序数据
func deleteDeviceTable(device string) error {
	err := global.TdDb.DropTable(device + "_" + entity.ATTRIBUTES_TSL)
	if err != nil {
		return err
	}
	err = global.TdDb.DropTable(device + "_" + entity.TELEMETRY_TSL)
	if err != nil {
		return err
	}
	return nil
}

func GetDeviceToken(data *entity.Device) *model.DeviceAuth {
	now := time.Now()
	etoken := &model.DeviceAuth{
		DeviceId:       data.Id,
		OrgId:          data.OrgId,
		Owner:          data.Owner,
		Name:           data.Name,
		DeviceType:     data.DeviceType,
		ProductId:      data.Pid,
		DeviceProtocol: data.Protocol,
	}
	//设备有效期360天
	etoken.CreatedAt = now.Unix()
	etoken.ExpiredAt = now.Add(time.Hour * 24 * 365).Unix()
	return etoken
}

// 获取设备数量统计
func (m *deviceModelImpl) FindDeviceCount() (result entity.DeviceCount, err error) {
	sql := `SELECT COUNT(*) AS total, (SELECT COUNT(*) FROM devices WHERE DATE(create_time) = CURDATE()) AS today  FROM devices`
	if global.Conf.Server.DbType == "postgresql" {
		sql = `SELECT COUNT(*) AS total, (SELECT COUNT(*) FROM devices WHERE DATE(create_time) = current_date) AS today  FROM devices`
	}
	err = global.Db.Raw(sql).Scan(&result).Error
	return
}

// 获取设备类型数量统计
func (m *deviceModelImpl) FindDeviceCountGroupByLinkStatus() (count []entity.DeviceCountLinkStatus, err error) {
	sql := `SELECT link_status, COUNT(*) AS total  FROM devices GROUP BY link_status`
	err = global.Db.Raw(sql, m.table).Scan(&count).Error
	return
}

func (m *deviceModelImpl) FindDeviceCountGroupByType() (count []entity.DeviceCountType, err error) {
	sql := `SELECT device_type, COUNT(*) AS total  FROM devices GROUP BY device_type`
	err = global.Db.Raw(sql, m.table).Scan(&count).Error
	return
}
