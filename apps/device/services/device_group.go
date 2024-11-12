package services

import (
	"errors"
	"pandax/apps/device/entity"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
)

type (
	DeviceGroupModel interface {
		Insert(data entity.DeviceGroup) (*entity.DeviceGroup, error)
		FindOne(id string) (*entity.DeviceGroup, error)
		FindListPage(page, pageSize int, data entity.DeviceGroup) (*[]entity.DeviceGroup, int64, error)
		FindList(data entity.DeviceGroup) (*[]entity.DeviceGroup, error)
		Update(data entity.DeviceGroup) (*entity.DeviceGroup, error)
		Delete(ids []string) error
		SelectDeviceGroup(data entity.DeviceGroup) ([]entity.DeviceGroup, error)
		SelectDeviceGroupLabel(data entity.DeviceGroup) ([]entity.DeviceGroupLabel, error)
	}

	deviceGroupModelImpl struct {
		table string
	}
)

var DeviceGroupModelDao DeviceGroupModel = &deviceGroupModelImpl{
	table: `device_groups`,
}

func (m *deviceGroupModelImpl) Insert(data entity.DeviceGroup) (*entity.DeviceGroup, error) {
	if err := global.Db.Table(m.table).Create(&data).Error; err != nil {
		return nil, err
	}
	path := "/" + data.Id
	if data.Pid != "0" {
		vsg, _ := m.FindOne(data.Pid)
		path = vsg.Path + path
	} else {
		path = "/0" + path
	}
	data.Path = path
	err := global.Db.Table(m.table).Model(&data).Updates(&data).Error
	return &data, err
}

func (m *deviceGroupModelImpl) FindOne(id string) (*entity.DeviceGroup, error) {
	resData := new(entity.DeviceGroup)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	return resData, err
}

func (m *deviceGroupModelImpl) FindListPage(page, pageSize int, data entity.DeviceGroup) (*[]entity.DeviceGroup, int64, error) {
	list := make([]entity.DeviceGroup, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Path != "" {
		db = db.Where("path like ?", "%"+data.Path+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	// 组织数据访问权限
	if err := model.OrgAuthSet(db, data.RoleId, data.Owner); err != nil {
		return nil, 0, err
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := db.Order("sort").Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *deviceGroupModelImpl) FindList(data entity.DeviceGroup) (*[]entity.DeviceGroup, error) {
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
	if err := model.OrgAuthSet(db, data.RoleId, data.Owner); err != nil {
		return nil, err
	}
	err := db.Order("sort").Find(&list).Error
	return &list, err
}

func (m *deviceGroupModelImpl) Update(data entity.DeviceGroup) (*entity.DeviceGroup, error) {
	one, err := m.FindOne(data.Id)
	if err != nil {
		return nil, err
	}
	path := "/" + data.Id
	if data.Pid != "0" {
		vsg, _ := m.FindOne(data.Pid)
		path = vsg.Path + path
	} else {
		path = "/0" + path
	}
	data.Path = path

	if data.Path != "" && data.Path != one.Path {
		return nil, errors.New("上级分组不允许修改！")
	}
	err = global.Db.Table(m.table).Updates(&data).Error
	return &data, err
}

func (m *deviceGroupModelImpl) Delete(s []string) error {
	return global.Db.Table(m.table).Delete(&entity.DeviceGroup{}, "id in (?)", s).Error
}

func (m *deviceGroupModelImpl) SelectDeviceGroup(data entity.DeviceGroup) ([]entity.DeviceGroup, error) {
	list, _ := m.FindList(data)
	sd := make([]entity.DeviceGroup, 0)
	li := *list
	for i := 0; i < len(li); i++ {
		if li[i].Pid != "0" {
			continue
		}
		info := DiGuiDeviceGroup(list, li[i])
		sd = append(sd, info)
	}
	return sd, nil
}

func (m *deviceGroupModelImpl) SelectDeviceGroupLabel(data entity.DeviceGroup) ([]entity.DeviceGroupLabel, error) {
	organizationlist, _ := m.FindList(data)
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
	return dl, nil
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
