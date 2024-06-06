package services

import (
	"errors"
	"pandax/apps/system/entity"
	"pandax/pkg/global"

	"github.com/kakuilan/kgo"
)

type (
	SysOrganizationModel interface {
		Insert(data entity.SysOrganization) (*entity.SysOrganization, error)
		FindOne(organizationId int64) (*entity.SysOrganization, error)
		FindListPage(page, pageSize int, data entity.SysOrganization) (*[]entity.SysOrganization, int64, error)
		FindList(data entity.SysOrganization) (*[]entity.SysOrganization, error)
		Update(data entity.SysOrganization) error
		Delete(organizationId []int64) error
		SelectOrganization(data entity.SysOrganization) ([]entity.SysOrganization, error)
		SelectOrganizationLable(data entity.SysOrganization) ([]entity.OrganizationLable, error)
		SelectOrganizationIds(data entity.SysOrganization) ([]int64, error)
	}

	sysOrganizationModelImpl struct {
		table string
	}
)

var SysOrganizationModelDao SysOrganizationModel = &sysOrganizationModelImpl{
	table: `sys_organizations`,
}

func (m *sysOrganizationModelImpl) Insert(data entity.SysOrganization) (*entity.SysOrganization, error) {
	err := global.Db.Table(m.table).Create(&data).Error
	if err != nil {
		return nil, err
	}
	organizationPath := "/" + kgo.KConv.Int2Str(data.OrganizationId)
	if int(data.ParentId) != 0 {
		organizationP, err := m.FindOne(data.ParentId)
		if err != nil {
			return nil, err
		}
		organizationPath = organizationP.OrganizationPath + organizationPath
	} else {
		organizationPath = "/0" + organizationPath
	}
	data.OrganizationPath = organizationPath
	err = global.Db.Table(m.table).Model(&data).Updates(&data).Error
	return &data, err
}

func (m *sysOrganizationModelImpl) FindOne(organizationId int64) (*entity.SysOrganization, error) {
	resData := new(entity.SysOrganization)
	err := global.Db.Table(m.table).Where("organization_id = ?", organizationId).First(resData).Error
	return resData, err
}

func (m *sysOrganizationModelImpl) FindListPage(page, pageSize int, data entity.SysOrganization) (*[]entity.SysOrganization, int64, error) {
	list := make([]entity.SysOrganization, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)

	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.OrganizationId != 0 {
		db = db.Where("organization_id = ?", data.OrganizationId)
	}
	if data.OrganizationName != "" {
		db = db.Where("organization_name like ?", "%"+data.OrganizationName+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.OrganizationPath != "" {
		db = db.Where("organizationPath like %?%", data.OrganizationPath)
	}
	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *sysOrganizationModelImpl) FindList(data entity.SysOrganization) (*[]entity.SysOrganization, error) {
	list := make([]entity.SysOrganization, 0)

	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.OrganizationId != 0 {
		db = db.Where("organization_id = ?", data.OrganizationId)
	}
	if data.OrganizationName != "" {
		db = db.Where("organization_name like ?", "%"+data.OrganizationName+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	db.Where("delete_time IS NULL")
	err := db.Order("sort").Find(&list).Error
	return &list, err
}

func (m *sysOrganizationModelImpl) Update(data entity.SysOrganization) error {
	one, err := m.FindOne(data.OrganizationId)
	if err != nil {
		return err
	}
	organizationPath := "/" + kgo.KConv.Int2Str(data.OrganizationId)
	if int(data.ParentId) != 0 {
		organizationP, err := m.FindOne(data.ParentId)
		if err != nil {
			return err
		}
		organizationPath = organizationP.OrganizationPath + organizationPath
	} else {
		organizationPath = "/0" + organizationPath
	}
	data.OrganizationPath = organizationPath

	if data.OrganizationPath != "" && data.OrganizationPath != one.OrganizationPath {
		return errors.New("上级组织不允许修改！")
	}
	err = global.Db.Table(m.table).Model(&data).Updates(&data).Error
	return err
}

func (m *sysOrganizationModelImpl) Delete(organizationIds []int64) error {
	return global.Db.Table(m.table).Delete(&entity.SysOrganization{}, "organization_id in (?)", organizationIds).Error
}

func (m *sysOrganizationModelImpl) SelectOrganization(data entity.SysOrganization) ([]entity.SysOrganization, error) {
	list, err := m.FindList(data)
	if err != nil {
		return nil, err
	}
	sd := make([]entity.SysOrganization, 0)
	li := *list
	for i := 0; i < len(li); i++ {
		if li[i].ParentId != 0 {
			continue
		}
		info := Digui(list, li[i])

		sd = append(sd, info)
	}
	return sd, nil
}

func (m *sysOrganizationModelImpl) SelectOrganizationLable(data entity.SysOrganization) ([]entity.OrganizationLable, error) {
	organizationlist, err := m.FindList(data)
	if err != nil {
		return nil, err
	}
	dl := make([]entity.OrganizationLable, 0)
	organizationl := *organizationlist
	for i := 0; i < len(organizationl); i++ {
		if organizationl[i].ParentId != 0 {
			continue
		}
		e := entity.OrganizationLable{}
		e.OrganizationId = organizationl[i].OrganizationId
		e.OrganizationName = organizationl[i].OrganizationName
		organizationsInfo := DiguiOrganizationLable(organizationlist, e)

		dl = append(dl, organizationsInfo)
	}
	return dl, nil
}

func (m *sysOrganizationModelImpl) SelectOrganizationIds(data entity.SysOrganization) ([]int64, error) {
	organizationlist, err := m.FindList(data)
	if err != nil {
		return nil, err
	}
	dl := make([]int64, 0)
	organizationl := *organizationlist
	for i := 0; i < len(organizationl); i++ {
		if organizationl[i].ParentId != 0 {
			continue
		}
		dl = append(dl, organizationl[i].OrganizationId)
		e := entity.OrganizationLable{}
		e.OrganizationId = organizationl[i].OrganizationId
		e.OrganizationName = organizationl[i].OrganizationName
		id := DiguiOrganizationId(organizationlist, e)
		dl = append(dl, id...)
	}
	return dl, nil
}

func Digui(organizationlist *[]entity.SysOrganization, menu entity.SysOrganization) entity.SysOrganization {
	list := *organizationlist

	min := make([]entity.SysOrganization, 0)
	for j := 0; j < len(list); j++ {

		if menu.OrganizationId != list[j].ParentId {
			continue
		}
		mi := entity.SysOrganization{}
		mi.OrganizationId = list[j].OrganizationId
		mi.ParentId = list[j].ParentId
		mi.OrganizationPath = list[j].OrganizationPath
		mi.OrganizationName = list[j].OrganizationName
		mi.Sort = list[j].Sort
		mi.Leader = list[j].Leader
		mi.Phone = list[j].Phone
		mi.Email = list[j].Email
		mi.Status = list[j].Status
		mi.CreatedAt = list[j].CreatedAt
		mi.UpdatedAt = list[j].UpdatedAt
		mi.Children = []entity.SysOrganization{}
		ms := Digui(organizationlist, mi)
		min = append(min, ms)
	}
	menu.Children = min
	return menu
}
func DiguiOrganizationLable(organizationlist *[]entity.SysOrganization, organization entity.OrganizationLable) entity.OrganizationLable {
	list := *organizationlist

	min := make([]entity.OrganizationLable, 0)
	for j := 0; j < len(list); j++ {

		if organization.OrganizationId != list[j].ParentId {
			continue
		}
		mi := entity.OrganizationLable{list[j].OrganizationId, list[j].OrganizationName, []entity.OrganizationLable{}}
		ms := DiguiOrganizationLable(organizationlist, mi)
		min = append(min, ms)

	}
	organization.Children = min
	return organization
}

func DiguiOrganizationId(organizationlist *[]entity.SysOrganization, organization entity.OrganizationLable) []int64 {
	list := *organizationlist
	min := make([]int64, 0)
	for j := 0; j < len(list); j++ {
		if organization.OrganizationId != list[j].ParentId {
			continue
		}
		min = append(min, list[j].OrganizationId)
		mi := entity.OrganizationLable{list[j].OrganizationId, list[j].OrganizationName, []entity.OrganizationLable{}}
		id := DiguiOrganizationId(organizationlist, mi)
		min = append(min, id...)
	}
	return min
}
