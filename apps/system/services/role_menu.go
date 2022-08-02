package services

import (
	"fmt"
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysRoleMenuModel interface {
		Insert(roleId int64, menuId []int64) bool
		FindList(data entity.SysRoleMenu) *[]entity.SysRoleMenu
		Update(data entity.SysRoleMenu) *entity.SysRoleMenu
		Delete(RoleId int64, MenuID int64)
		GetPermis(roleId int64) []string
		GetMenuPaths(rm entity.SysRoleMenu) []entity.MenuPath
		DeleteRoleMenu(RoleId int64)
		DeleteRoleMenus(roleIds []int64)
	}

	sysRoleMenuImpl struct {
		table string
	}
)

var SysRoleMenuModelDao SysRoleMenuModel = &sysRoleMenuImpl{
	table: `sys_role_menus`,
}

func (m *sysRoleMenuImpl) Insert(roleId int64, menuId []int64) bool {

	var role entity.SysRole
	biz.ErrIsNil(global.Db.Table("sys_roles").Where("role_id = ?", roleId).First(&role).Error, "查询角色失败")

	var menu []entity.SysMenu
	biz.ErrIsNil(global.Db.Table("sys_menus").Where("menu_id in (?)", menuId).Find(&menu).Error, "查询菜单失败")

	//拼接 sql 串
	sql := "INSERT INTO sys_role_menus (role_id,menu_id,role_name) VALUES "

	for i := 0; i < len(menu); i++ {
		if len(menu)-1 == i {
			//最后一条数据 以分号结尾
			sql += fmt.Sprintf("(%d,%d,'%s');", role.RoleId, menu[i].MenuId, role.RoleKey)
		} else {
			sql += fmt.Sprintf("(%d,%d,'%s'),", role.RoleId, menu[i].MenuId, role.RoleKey)
		}
	}
	biz.ErrIsNil(global.Db.Exec(sql).Error, "新增角色菜单失败")

	return true
}

func (m *sysRoleMenuImpl) FindList(data entity.SysRoleMenu) *[]entity.SysRoleMenu {
	list := make([]entity.SysRoleMenu, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.RoleId != 0 {
		db = db.Where("role_id = ?", data.RoleId)
	}
	biz.ErrIsNil(db.Find(&list).Error, "查询角色菜单失败")
	return &list
}

// 查询权限标识
func (m *sysRoleMenuImpl) GetPermis(roleId int64) []string {
	var r []entity.SysMenu
	db := global.Db.Select("sys_menus.permission").Table("sys_menus").Joins("left join sys_role_menus on sys_menus.menu_id = sys_role_menus.menu_id")

	db = db.Where("role_id = ?", roleId)

	db = db.Where("sys_menus.menu_type in ('F','C')")

	biz.ErrIsNil(db.Find(&r).Error, "查询查询权限标识列表失败")

	var list []string
	for i := 0; i < len(r); i++ {
		list = append(list, r[i].Permission)
	}
	return list
}

func (m *sysRoleMenuImpl) GetMenuPaths(rm entity.SysRoleMenu) []entity.MenuPath {
	var r []entity.MenuPath
	db := global.Db.Select("sys_menus.path").Table(m.table)
	db = db.Joins("left join sys_roles on sys_roles.role_id=sys_role_menus.role_id")
	db = db.Joins("left join sys_menus on sys_menus.id=sys_role_menus.menu_id")
	db = db.Where("sys_roles.role_key = ? and sys_menus.type=1", rm.RoleName)

	biz.ErrIsNil(db.Find(&r).Error, "查询菜单路径失败")
	return r
}

func (m *sysRoleMenuImpl) Update(data entity.SysRoleMenu) *entity.SysRoleMenu {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改菜单失败")
	return &data
}

func (m *sysRoleMenuImpl) DeleteRoleMenu(roleId int64) {
	var rm entity.SysRoleMenu
	if err := global.Db.Table(m.table).Where("role_id = ?", roleId).Delete(&rm).Error; err != nil {
		biz.ErrIsNil(err, "删除角色菜单失败")
	}
	return
}

func (m *sysRoleMenuImpl) DeleteRoleMenus(roleIds []int64) {
	var rm entity.SysRoleMenu
	biz.ErrIsNil(global.Db.Table(m.table).Where("role_id in (?)", roleIds).Delete(&rm).Error, "批量删除角色菜单失败")
}

func (m *sysRoleMenuImpl) Delete(RoleId int64, MenuID int64) {
	var rm entity.SysRoleMenu
	rm.RoleId = RoleId
	db := global.Db.Table(m.table).Where("role_id = ?", RoleId)
	if MenuID != 0 {
		db = db.Where("menu_id = ?", MenuID)
	}

	biz.ErrIsNil(db.Delete(&rm).Error, "删除角色菜单失败")
	return

}
