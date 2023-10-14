package services

import (
	"errors"
	"github.com/PandaXGO/PandaKit/biz"
	"pandax/apps/device/entity"
	"pandax/pkg/global"
	"pandax/pkg/global_model"
)

type (
	DeviceGroupModel interface {
		Insert(data entity.DeviceGroup) *entity.DeviceGroup
		FindOne(id string) *entity.DeviceGroup
		FindListPage(page, pageSize int, data entity.DeviceGroup) (*[]entity.DeviceGroup, int64)
		FindList(data entity.DeviceGroup) *[]entity.DeviceGroup
		Update(data entity.DeviceGroup) *entity.DeviceGroup
		Delete(ids []string)
		SelectDeviceGroup(data entity.DeviceGroup) []entity.DeviceGroup
		SelectDeviceGroupLabel(data entity.DeviceGroup) []entity.DeviceGroupLabel
	}

	deviceGroupModelImpl struct {
		table string
	}
)

var DeviceGroupModelDao DeviceGroupModel = &deviceGroupModelImpl{
	table: `device_groups`,
}

func (m *deviceGroupModelImpl) Insert(data entity.DeviceGroup) *entity.DeviceGroup {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加设备分组失败")

	path := "/" + data.Id
	if data.Pid != "0" {
		vsg := m.FindOne(data.Pid)
		path = vsg.Path + path
	} else {
		path = "/0" + path
	}
	data.Path = path
	biz.ErrIsNil(global.Db.Table(m.table).Model(&data).Updates(&data).Error, "修改设备分组信息失败")
	return &data
}

func (m *deviceGroupModelImpl) FindOne(id string) *entity.DeviceGroup {
	resData := new(entity.DeviceGroup)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询设备分组失败")
	return resData
}

func (m *deviceGroupModelImpl) FindListPage(page, pageSize int, data entity.DeviceGroup) (*[]entity.DeviceGroup, int64) {
	list := make([]entity.DeviceGroup, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Path != "" {
		db = db.Where("path like %?%", "%"+data.Path+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	// 组织数据访问权限
	global_model.OrgAuthSet(db, data.RoleId, data.Owner)

	err := db.Count(&total).Error
	err = db.Order("sort").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询设备分组分页列表失败")
	return &list, total
}

func (m *deviceGroupModelImpl) FindList(data entity.DeviceGroup) *[]entity.DeviceGroup {
	list := make([]entity.DeviceGroup, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Path != "" {
		db = db.Where("path like %?%", "%"+data.Path+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	// 组织数据访问权限
	global_model.OrgAuthSet(db, data.RoleId, data.Owner)
	biz.ErrIsNil(db.Order("sort").Find(&list).Error, "查询设备分组列表失败")
	return &list
}

func (m *deviceGroupModelImpl) Update(data entity.DeviceGroup) *entity.DeviceGroup {
	one := m.FindOne(data.Id)

	path := "/" + data.Id
	if data.Pid != "0" {
		vsg := m.FindOne(data.Pid)
		path = vsg.Path + path
	} else {
		path = "/0" + path
	}
	data.Path = path

	if data.Path != "" && data.Path != one.Path {
		biz.ErrIsNil(errors.New("上级分组不允许修改！"), "上级分组不允许修改")
	}

	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改设备分组失败")
	return &data
}

func (m *deviceGroupModelImpl) Delete(s []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.DeviceGroup{}, "id in (?)", s).Error, "删除设备分组失败")
}

func (m *deviceGroupModelImpl) SelectDeviceGroup(data entity.DeviceGroup) []entity.DeviceGroup {
	list := m.FindList(data)
	sd := make([]entity.DeviceGroup, 0)
	li := *list
	for i := 0; i < len(li); i++ {
		if li[i].Pid != "0" {
			continue
		}
		info := DiGuiDeviceGroup(list, li[i])
		sd = append(sd, info)
	}
	return sd
}

func (m *deviceGroupModelImpl) SelectDeviceGroupLabel(data entity.DeviceGroup) []entity.DeviceGroupLabel {
	organizationlist := m.FindList(data)

	dl := make([]entity.DeviceGroupLabel, 0)
	organizationl := *organizationlist
	for i := 0; i < len(organizationl); i++ {
		if organizationl[i].Pid != "0" {
			continue
		}
		e := entity.DeviceGroupLabel{}
		e.Id = organizationl[i].Id
		e.Name = organizationl[i].Name
		organizationsInfo := DiGuiDeviceGroupLabel(organizationlist, e)

		dl = append(dl, organizationsInfo)
	}
	return dl
}

func DiGuiDeviceGroup(sglist *[]entity.DeviceGroup, menu entity.DeviceGroup) entity.DeviceGroup {
	list := *sglist

	min := make([]entity.DeviceGroup, 0)
	for j := 0; j < len(list); j++ {

		if menu.Id != list[j].Pid {
			continue
		}
		mi := entity.DeviceGroup{}
		mi.Id = list[j].Id
		mi.Pid = list[j].Pid
		mi.Path = list[j].Path
		mi.Name = list[j].Name
		mi.Sort = list[j].Sort
		mi.Status = list[j].Status
		mi.Description = list[j].Description
		ms := DiGuiDeviceGroup(sglist, mi)
		min = append(min, ms)
	}
	menu.Children = min
	return menu
}
func DiGuiDeviceGroupLabel(sglist *[]entity.DeviceGroup, organization entity.DeviceGroupLabel) entity.DeviceGroupLabel {
	list := *sglist

	min := make([]entity.DeviceGroupLabel, 0)
	for j := 0; j < len(list); j++ {
		if organization.Id != list[j].Pid {
			continue
		}
		sg := entity.DeviceGroupLabel{list[j].Id, list[j].Name, []entity.DeviceGroupLabel{}}
		ms := DiGuiDeviceGroupLabel(sglist, sg)
		min = append(min, ms)

	}
	organization.Children = min
	return organization
}
