package services

import (
	"errors"
	"pandax/apps/system/entity"
	"pandax/kit/biz"
	"pandax/pkg/global"

	"github.com/kakuilan/kgo"
)

type (
	SysOrganizationModel interface {
		Insert(data entity.SysOrganization) *entity.SysOrganization
		FindOne(organizationId int64) *entity.SysOrganization
		FindListPage(page, pageSize int, data entity.SysOrganization) (*[]entity.SysOrganization, int64)
		FindList(data entity.SysOrganization) *[]entity.SysOrganization
		Update(data entity.SysOrganization) *entity.SysOrganization
		Delete(organizationId []int64)
		SelectOrganization(data entity.SysOrganization) []entity.SysOrganization
		SelectOrganizationLable(data entity.SysOrganization) []entity.OrganizationLable
		SelectOrganizationIds(data entity.SysOrganization) []int64
	}

	sysOrganizationModelImpl struct {
		table string
	}
)

var SysOrganizationModelDao SysOrganizationModel = &sysOrganizationModelImpl{
	table: `sys_organizations`,
}

func (m *sysOrganizationModelImpl) Insert(data entity.SysOrganization) *entity.SysOrganization {
	biz.ErrIsNil(global.Db.Table(m.table).Create(&data).Error, "新增组织信息失败")
	organizationPath := "/" + kgo.KConv.Int2Str(data.OrganizationId)
	if int(data.ParentId) != 0 {
		organizationP := m.FindOne(data.ParentId)
		organizationPath = organizationP.OrganizationPath + organizationPath
	} else {
		organizationPath = "/0" + organizationPath
	}
	data.OrganizationPath = organizationPath
	biz.ErrIsNil(global.Db.Table(m.table).Model(&data).Updates(&data).Error, "修改组织信息失败")
	return &data
}

func (m *sysOrganizationModelImpl) FindOne(organizationId int64) *entity.SysOrganization {
	resData := new(entity.SysOrganization)
	err := global.Db.Table(m.table).Where("organization_id = ?", organizationId).First(resData).Error
	biz.ErrIsNil(err, "查询组织信息失败")
	return resData
}

func (m *sysOrganizationModelImpl) FindListPage(page, pageSize int, data entity.SysOrganization) (*[]entity.SysOrganization, int64) {
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
	biz.ErrIsNil(err, "查询组织分页列表信息失败")
	return &list, total
}

func (m *sysOrganizationModelImpl) FindList(data entity.SysOrganization) *[]entity.SysOrganization {
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
	biz.ErrIsNil(err, "查询组织列表信息失败")
	return &list
}

func (m *sysOrganizationModelImpl) Update(data entity.SysOrganization) *entity.SysOrganization {
	one := m.FindOne(data.OrganizationId)

	organizationPath := "/" + kgo.KConv.Int2Str(data.OrganizationId)
	if int(data.ParentId) != 0 {
		organizationP := m.FindOne(data.ParentId)
		organizationPath = organizationP.OrganizationPath + organizationPath
	} else {
		organizationPath = "/0" + organizationPath
	}
	data.OrganizationPath = organizationPath

	if data.OrganizationPath != "" && data.OrganizationPath != one.OrganizationPath {
		biz.ErrIsNil(errors.New("上级组织不允许修改！"), "上级组织不允许修改")
	}
	biz.ErrIsNil(global.Db.Table(m.table).Model(&data).Updates(&data).Error, "修改组织信息失败")
	return &data
}

func (m *sysOrganizationModelImpl) Delete(organizationIds []int64) {
	err := global.Db.Table(m.table).Delete(&entity.SysOrganization{}, "organization_id in (?)", organizationIds).Error
	biz.ErrIsNil(err, "删除组织信息失败")
	return
}

func (m *sysOrganizationModelImpl) SelectOrganization(data entity.SysOrganization) []entity.SysOrganization {
	list := m.FindList(data)

	sd := make([]entity.SysOrganization, 0)
	li := *list
	for i := 0; i < len(li); i++ {
		if li[i].ParentId != 0 {
			continue
		}
		info := Digui(list, li[i])

		sd = append(sd, info)
	}
	return sd
}

func (m *sysOrganizationModelImpl) SelectOrganizationLable(data entity.SysOrganization) []entity.OrganizationLable {
	organizationlist := m.FindList(data)

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
	return dl
}

func (m *sysOrganizationModelImpl) SelectOrganizationIds(data entity.SysOrganization) []int64 {
	organizationlist := m.FindList(data)
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
	return dl
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
