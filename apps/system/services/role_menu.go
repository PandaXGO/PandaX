package services

import (
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysRoleMenuModel interface {
		Insert(roleId int64, menuId []int64) error
		FindList(data entity.SysRoleMenu) (*[]entity.SysRoleMenu, error)
		Update(data entity.SysRoleMenu) error
		Delete(RoleId int64, MenuID int64) error
		GetPermis(roleId int64) ([]string, error)
		GetMenuPaths(rm entity.SysRoleMenu) ([]entity.MenuPath, error)
		DeleteRoleMenu(RoleId int64) error
		DeleteRoleMenus(roleIds []int64) error
	}

	sysRoleMenuImpl struct {
		table string
	}
)

var SysRoleMenuModelDao SysRoleMenuModel = &sysRoleMenuImpl{
	table: `sys_role_menus`,
}

func (m *sysRoleMenuImpl) Insert(roleId int64, menuId []int64) error {

	var role entity.SysRole
	err := global.Db.Table("sys_roles").Where("role_id = ?", roleId).First(&role).Error
	if err != nil {
		return err
	}
	var menu []entity.SysMenu
	err = global.Db.Table("sys_menus").Where("menu_id in (?)", menuId).Find(&menu).Error
	if err != nil {
		return err
	}
	menus := make([]entity.SysRoleMenu, 0)
	for i := 0; i < len(menu); i++ {
		menus = append(menus, entity.SysRoleMenu{RoleId: role.RoleId, MenuId: menu[i].MenuId, RoleName: role.RoleKey})
	}
	return global.Db.CreateInBatches(&menus, len(menus)).Error
}

func (m *sysRoleMenuImpl) FindList(data entity.SysRoleMenu) (*[]entity.SysRoleMenu, error) {
	list := make([]entity.SysRoleMenu, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.RoleId != 0 {
		db = db.Where("role_id = ?", data.RoleId)
	}
	err := db.Find(&list).Error
	return &list, err
}

// 查询权限标识
func (m *sysRoleMenuImpl) GetPermis(roleId int64) ([]string, error) {
	var r []entity.SysMenu
	db := global.Db.Select("sys_menus.permission").Table("sys_menus").Joins("left join sys_role_menus on sys_menus.menu_id = sys_role_menus.menu_id")

	db = db.Where("role_id = ?", roleId)

	db = db.Where("sys_menus.menu_type in ('F','C')")

	err := db.Find(&r).Error
	if err != nil {
		return nil, err
	}
	var list []string
	for i := 0; i < len(r); i++ {
		list = append(list, r[i].Permission)
	}
	return list, nil
}

func (m *sysRoleMenuImpl) GetMenuPaths(rm entity.SysRoleMenu) ([]entity.MenuPath, error) {
	var r []entity.MenuPath
	db := global.Db.Select("sys_menus.path").Table(m.table)
	db = db.Joins("left join sys_roles on sys_roles.role_id=sys_role_menus.role_id")
	db = db.Joins("left join sys_menus on sys_menus.id=sys_role_menus.menu_id")
	db = db.Where("sys_roles.role_key = ? and sys_menus.type=1", rm.RoleName)

	err := db.Find(&r).Error
	return r, err
}

func (m *sysRoleMenuImpl) Update(data entity.SysRoleMenu) error {
	return global.Db.Table(m.table).Updates(&data).Error
}

func (m *sysRoleMenuImpl) DeleteRoleMenu(roleId int64) error {
	var rm entity.SysRoleMenu
	return global.Db.Table(m.table).Where("role_id = ?", roleId).Delete(&rm).Error
}

func (m *sysRoleMenuImpl) DeleteRoleMenus(roleIds []int64) error {
	var rm entity.SysRoleMenu
	return global.Db.Table(m.table).Where("role_id in (?)", roleIds).Delete(&rm).Error
}

func (m *sysRoleMenuImpl) Delete(RoleId int64, MenuID int64) error {
	var rm entity.SysRoleMenu
	rm.RoleId = RoleId
	db := global.Db.Table(m.table).Where("role_id = ?", RoleId)
	if MenuID != 0 {
		db = db.Where("menu_id = ?", MenuID)
	}

	return db.Delete(&rm).Error

}
