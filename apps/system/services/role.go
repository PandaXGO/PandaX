package services

import (
	"errors"
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysRoleModel interface {
		Insert(data entity.SysRole) *entity.SysRole
		FindOne(roleId int64) *entity.SysRole
		FindListPage(page, pageSize int, data entity.SysRole) (list *[]entity.SysRole, total int64)
		FindList(data entity.SysRole) (list *[]entity.SysRole)
		Update(data entity.SysRole) *entity.SysRole
		Delete(roleId []int64)
		GetRoleMeunId(data entity.SysRole) []int64
		GetRoleDeptId(data entity.SysRole) []int64
	}

	sysRoleModel struct {
		table string
	}
)

var SysRoleModelDao SysRoleModel = &sysRoleModel{
	table: `sys_roles`,
}

func (m *sysRoleModel) Insert(data entity.SysRole) *entity.SysRole {
	var i int64 = 0
	global.Db.Table(m.table).Where("(role_name = ? or role_key = ?) and delete_time IS NULL", data.RoleName, data.RoleKey).Count(&i)
	biz.IsTrue(i == 0, "角色名称或者角色标识已经存在！")

	data.UpdateBy = ""
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加角色失败")

	return &data
}

func (m *sysRoleModel) FindOne(roleId int64) *entity.SysRole {
	resData := new(entity.SysRole)
	biz.ErrIsNil(global.Db.Table(m.table).Where("role_id = ?", roleId).First(resData).Error, "查询角色失败")
	return resData
}

func (m *sysRoleModel) FindListPage(page, pageSize int, data entity.SysRole) (*[]entity.SysRole, int64) {

	list := make([]entity.SysRole, 0)
	var total int64 = 0

	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.RoleId != 0 {
		db = db.Where("role_id = ?", data.RoleId)
	}
	if data.RoleName != "" {
		db = db.Where("role_name like ?", "%"+data.RoleName+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.RoleKey != "" {
		db = db.Where("role_key like ?", "%"+data.RoleKey+"%")
	}
	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Order("role_sort").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询角色分页列表失败")
	return &list, total
}

func (m *sysRoleModel) FindList(data entity.SysRole) *[]entity.SysRole {
	list := make([]entity.SysRole, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.RoleName != "" {
		db = db.Where("role_name like ?", "%"+data.RoleName+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.RoleKey != "" {
		db = db.Where("role_key like ?", "%"+data.RoleKey+"%")
	}
	db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("role_sort").Find(&list).Error, "查询角色列表失败")
	return &list
}

func (m *sysRoleModel) Update(data entity.SysRole) *entity.SysRole {
	update := new(entity.SysRole)
	biz.ErrIsNil(global.Db.Table(m.table).First(update, data.RoleId).Error, "查询角色失败")
	if data.RoleKey != "" && data.RoleKey != update.RoleKey {
		biz.ErrIsNil(errors.New("角色标识不允许修改！"), "角色标识不允许修改！")
	}
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改角色失败")
	return &data
}

func (m *sysRoleModel) Delete(roleIds []int64) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.SysRole{}, "role_id in (?)", roleIds).Error, "删除角色失败")
	return
}

// 获取角色对应的菜单ids
func (m *sysRoleModel) GetRoleMeunId(data entity.SysRole) []int64 {
	menuIds := make([]int64, 0)
	menuList := make([]entity.MenuIdList, 0)
	err := global.Db.Table("sys_menus").Select("sys_menus.menu_id").Joins("LEFT JOIN sys_role_menus on sys_role_menus.menu_id=sys_menus.menu_id").Where("sys_role_menus.role_id = ? ", data.RoleId).Where("sys_menus.menu_id not in (select sys_menus.parent_id from sys_menus LEFT JOIN sys_role_menus on sys_menus.menu_id=sys_role_menus.menu_id where sys_role_menus.role_id =? )", data.RoleId).Find(&menuList).Error

	biz.ErrIsNil(err, "查询角色菜单列表失败")
	for i := 0; i < len(menuList); i++ {
		menuIds = append(menuIds, menuList[i].MenuId)
	}
	return menuIds
}

func (m *sysRoleModel) GetRoleDeptId(data entity.SysRole) []int64 {
	deptIds := make([]int64, 0)
	deptList := make([]entity.DeptIdList, 0)
	err := global.Db.Table("sys_role_depts").Select("sys_role_depts.dept_id").Joins("LEFT JOIN sys_depts on sys_depts.dept_id=sys_role_depts.dept_id").Where("role_id = ? ", data.RoleId).Where(" sys_role_depts.dept_id not in(select sys_depts.parent_id from sys_role_depts LEFT JOIN sys_depts on sys_depts.dept_id=sys_role_depts.dept_id where role_id =? )", data.RoleId).Find(&deptList).Error
	biz.ErrIsNil(err, "查询角色部门列表失败")

	for i := 0; i < len(deptList); i++ {
		deptIds = append(deptIds, deptList[i].DeptId)
	}
	return deptIds
}
