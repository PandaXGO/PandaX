package services

import (
	"github.com/PandaXGO/PandaKit/biz"
	"pandax/apps/device/entity"
	"pandax/pkg/cache"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
	"time"
)

type (
	DeviceModel interface {
		Insert(data entity.Device) *entity.Device
		FindOneByToken(token string) (*entity.DeviceRes, error)
		FindOneByName(name string) (*entity.DeviceRes, error)
		FindOne(id string) (*entity.DeviceRes, error)
		FindListPage(page, pageSize int, data entity.Device) (*[]entity.DeviceRes, int64)
		FindList(data entity.Device) *[]entity.DeviceRes
		Update(data entity.Device) *entity.Device
		UpdateStatus(id, linkStatus string) error
		Delete(ids []string)
		FindDeviceCount() entity.DeviceCount
		FindDeviceCountGroupByLinkStatus() []entity.DeviceCountLinkStatus
		FindDeviceCountGroupByType() []entity.DeviceCountType
	}

	deviceModelImpl struct {
		table string
	}
)

var DeviceModelDao DeviceModel = &deviceModelImpl{
	table: `devices`,
}

func (m *deviceModelImpl) Insert(data entity.Device) *entity.Device {
	tx := global.Db.Begin()
	//1 检查设备名称是否存在
	list := m.FindList(entity.Device{Name: data.Name})
	biz.IsTrue(list != nil && len(*list) == 0, "设备名称已经存在")
	//2 创建认证TOKEN IOTHUB使用
	token := GetDeviceToken(&data)
	// 子网关不需要设置token
	if data.DeviceType == global.GATEWAYS {
		err := cache.SetDeviceEtoken(data.Name, token.GetMarshal(), time.Hour*24*365)
		biz.ErrIsNil(err, "设备缓存失败")
	} else {
		data.Token = token.MD5ID()
		err := cache.SetDeviceEtoken(data.Token, token.GetMarshal(), time.Hour*24*365)
		biz.ErrIsNil(err, "设备缓存失败")
	}
	//3 添加设备
	err := tx.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加设备失败")
	// 创建超级表 失败就
	if data.Pid != "" {
		err = createDeviceTable(data.Pid, data.Name)
		if err != nil {
			tx.Rollback()
			biz.ErrIsNil(err, "添加设备失败，设备表创建失败")
		}
	}
	tx.Commit()
	return &data
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

func (m *deviceModelImpl) FindListPage(page, pageSize int, data entity.Device) (*[]entity.DeviceRes, int64) {
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
	model.OrgAuthSet(db, data.RoleId, data.Owner)

	err := db.Count(&total).Error
	err = db.Order("create_time").Preload("Product").Preload("DeviceGroup").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询设备分页列表失败")
	return &list, total
}

func (m *deviceModelImpl) FindList(data entity.Device) *[]entity.DeviceRes {
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
	model.OrgAuthSet(db, data.RoleId, data.Owner)
	db.Preload("Product").Preload("DeviceGroup")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询设备列表失败")
	return &list
}

func (m *deviceModelImpl) Update(data entity.Device) *entity.Device {
	token := GetDeviceToken(&data)
	if data.DeviceType == global.GATEWAYS {
		err := cache.SetDeviceEtoken(data.Name, token.GetMarshal(), time.Hour*24*365)
		biz.ErrIsNil(err, "设备更改缓存失败")
	} else {
		err := cache.SetDeviceEtoken(data.Token, token.GetMarshal(), time.Hour*24*365)
		biz.ErrIsNil(err, "设备更改缓存失败")
	}
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改设备失败")
	return &data
}
func (m *deviceModelImpl) UpdateStatus(id, linkStatus string) error {
	return global.Db.Table(m.table).Where("id", id).Update("link_status", linkStatus).Update("last_time", time.Now()).Error
}

func (m *deviceModelImpl) Delete(ids []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.Device{}, "id in (?)", ids).Error, "删除设备失败")
	for _, id := range ids {
		device, err := m.FindOne(id)
		if err != nil {
			continue
		}
		// 删除表
		err = deleteDeviceTable(device.Name)
		global.Log.Error("设备时序表删除失败", err)
		// 删除所有缓存
		if device.DeviceType == global.GATEWAYS {
			cache.DelDeviceEtoken(device.Name)
		} else {
			cache.DelDeviceEtoken(device.Token)
		}
	}
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
func (m *deviceModelImpl) FindDeviceCount() (result entity.DeviceCount) {
	sql := `SELECT COUNT(*) AS total, (SELECT COUNT(*) FROM devices WHERE DATE(create_time) = CURDATE()) AS today  FROM devices`
	if global.Conf.Server.DbType == "postgresql" {
		sql = `SELECT COUNT(*) AS total, (SELECT COUNT(*) FROM devices WHERE DATE(create_time) = current_date) AS today  FROM devices`
	}
	err := global.Db.Raw(sql).Scan(&result).Error
	biz.ErrIsNil(err, "获取设备统计总数失败")
	return result
}

// 获取设备类型数量统计
func (m *deviceModelImpl) FindDeviceCountGroupByLinkStatus() (count []entity.DeviceCountLinkStatus) {
	sql := `SELECT link_status, COUNT(*) AS total  FROM devices GROUP BY link_status`
	err := global.Db.Raw(sql, m.table).Scan(&count).Error
	biz.ErrIsNil(err, "获取通过设备在线状态的设备统计总数失败")
	return
}

func (m *deviceModelImpl) FindDeviceCountGroupByType() (count []entity.DeviceCountType) {
	sql := `SELECT device_type, COUNT(*) AS total  FROM devices GROUP BY device_type`
	err := global.Db.Raw(sql, m.table).Scan(&count).Error
	biz.ErrIsNil(err, "获取通过设备类型的设备统计总数失败")
	return
}
