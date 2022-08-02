package services

import (
	"errors"
	"github.com/XM-GO/PandaKit/biz"
	"github.com/kakuilan/kgo"
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysDeptModel interface {
		Insert(data entity.SysDept) *entity.SysDept
		FindOne(deptId int64) *entity.SysDept
		FindListPage(page, pageSize int, data entity.SysDept) (*[]entity.SysDept, int64)
		FindList(data entity.SysDept) *[]entity.SysDept
		Update(data entity.SysDept) *entity.SysDept
		Delete(deptId []int64)
		SelectDept(data entity.SysDept) []entity.SysDept
		SelectDeptLable(data entity.SysDept) []entity.DeptLable
	}

	sysDeptModelImpl struct {
		table string
	}
)

var SysDeptModelDao SysDeptModel = &sysDeptModelImpl{
	table: `sys_depts`,
}

func (m *sysDeptModelImpl) Insert(data entity.SysDept) *entity.SysDept {
	biz.ErrIsNil(global.Db.Table(m.table).Create(&data).Error, "新增部门信息失败")
	deptPath := "/" + kgo.KConv.Int2Str(data.DeptId)
	if int(data.ParentId) != 0 {
		deptP := m.FindOne(data.ParentId)
		deptPath = deptP.DeptPath + deptPath
	} else {
		deptPath = "/0" + deptPath
	}
	data.DeptPath = deptPath
	biz.ErrIsNil(global.Db.Table(m.table).Model(&data).Updates(&data).Error, "修改部门信息失败")
	return &data
}

func (m *sysDeptModelImpl) FindOne(deptId int64) *entity.SysDept {
	resData := new(entity.SysDept)
	err := global.Db.Table(m.table).Where("dept_id = ?", deptId).First(resData).Error
	biz.ErrIsNil(err, "查询部门信息失败")
	return resData
}

func (m *sysDeptModelImpl) FindListPage(page, pageSize int, data entity.SysDept) (*[]entity.SysDept, int64) {
	list := make([]entity.SysDept, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)

	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.DeptId != 0 {
		db = db.Where("dept_id = ?", data.DeptId)
	}
	if data.DeptName != "" {
		db = db.Where("dept_name like ?", "%"+data.DeptName+"%")
	}
	if data.TenantId != 0 {
		db = db.Where("tenant_id = ?", data.TenantId)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.DeptPath != "" {
		db = db.Where("deptPath like %?%", data.DeptPath)
	}
	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询部门分页列表信息失败")
	return &list, total
}

func (m *sysDeptModelImpl) FindList(data entity.SysDept) *[]entity.SysDept {
	list := make([]entity.SysDept, 0)

	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.DeptId != 0 {
		db = db.Where("dept_id = ?", data.DeptId)
	}
	if data.TenantId != 0 {
		db = db.Where("tenant_id = ?", data.TenantId)
	}
	if data.DeptName != "" {
		db = db.Where("dept_name like ?", "%"+data.DeptName+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	db.Where("delete_time IS NULL")
	err := db.Order("sort").Find(&list).Error
	biz.ErrIsNil(err, "查询部门列表信息失败")
	return &list
}

func (m *sysDeptModelImpl) Update(data entity.SysDept) *entity.SysDept {
	one := m.FindOne(data.DeptId)

	deptPath := "/" + kgo.KConv.Int2Str(data.DeptId)
	if int(data.ParentId) != 0 {
		deptP := m.FindOne(data.ParentId)
		deptPath = deptP.DeptPath + deptPath
	} else {
		deptPath = "/0" + deptPath
	}
	data.DeptPath = deptPath

	if data.DeptPath != "" && data.DeptPath != one.DeptPath {
		biz.ErrIsNil(errors.New("上级部门不允许修改！"), "上级部门不允许修改")
	}
	biz.ErrIsNil(global.Db.Table(m.table).Model(&data).Updates(&data).Error, "修改部门信息失败")
	return &data
}

func (m *sysDeptModelImpl) Delete(deptIds []int64) {
	err := global.Db.Table(m.table).Delete(&entity.SysDept{}, "dept_id in (?)", deptIds).Error
	biz.ErrIsNil(err, "删除部门信息失败")
	return
}

func (m *sysDeptModelImpl) SelectDept(data entity.SysDept) []entity.SysDept {
	list := m.FindList(data)

	sd := make([]entity.SysDept, 0)
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

func (m *sysDeptModelImpl) SelectDeptLable(data entity.SysDept) []entity.DeptLable {
	deptlist := m.FindList(data)

	dl := make([]entity.DeptLable, 0)
	deptl := *deptlist
	for i := 0; i < len(deptl); i++ {
		if deptl[i].ParentId != 0 {
			continue
		}
		e := entity.DeptLable{}
		e.DeptId = deptl[i].DeptId
		e.DeptName = deptl[i].DeptName
		deptsInfo := DiguiDeptLable(deptlist, e)

		dl = append(dl, deptsInfo)
	}
	return dl
}

func Digui(deptlist *[]entity.SysDept, menu entity.SysDept) entity.SysDept {
	list := *deptlist

	min := make([]entity.SysDept, 0)
	for j := 0; j < len(list); j++ {

		if menu.DeptId != list[j].ParentId {
			continue
		}
		mi := entity.SysDept{}
		mi.DeptId = list[j].DeptId
		mi.ParentId = list[j].ParentId
		mi.DeptPath = list[j].DeptPath
		mi.DeptName = list[j].DeptName
		mi.Sort = list[j].Sort
		mi.Leader = list[j].Leader
		mi.Phone = list[j].Phone
		mi.Email = list[j].Email
		mi.Status = list[j].Status
		mi.CreatedAt = list[j].CreatedAt
		mi.UpdatedAt = list[j].UpdatedAt
		mi.Children = []entity.SysDept{}
		ms := Digui(deptlist, mi)
		min = append(min, ms)
	}
	menu.Children = min
	return menu
}
func DiguiDeptLable(deptlist *[]entity.SysDept, dept entity.DeptLable) entity.DeptLable {
	list := *deptlist

	min := make([]entity.DeptLable, 0)
	for j := 0; j < len(list); j++ {

		if dept.DeptId != list[j].ParentId {
			continue
		}
		mi := entity.DeptLable{list[j].DeptId, list[j].DeptName, []entity.DeptLable{}}
		ms := DiguiDeptLable(deptlist, mi)
		min = append(min, ms)

	}
	dept.Children = min
	return dept
}
