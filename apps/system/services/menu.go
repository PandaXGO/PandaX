package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysMenuModel interface {
		Insert(data entity.SysMenu) *entity.SysMenu
		FindOne(menuId int64) *entity.SysMenu
		FindListPage(page, pageSize int, data entity.SysMenu) (*[]entity.SysMenu, int64)
		FindList(data entity.SysMenu) *[]entity.SysMenu
		Update(data entity.SysMenu) *entity.SysMenu
		Delete(menuId []int64)
		SelectMenu(data entity.SysMenu) *[]entity.SysMenu
		SelectMenuLable(data entity.SysMenu) *[]entity.MenuLable
		SelectMenuRole(roleName string) *[]entity.SysMenu
		GetMenuRole(data entity.MenuRole) *[]entity.MenuRole
	}

	sysMenuModelImpl struct {
		table string
	}
)

var SysMenuModelDao SysMenuModel = &sysMenuModelImpl{
	table: `sys_menus`,
}

func (m *sysMenuModelImpl) Insert(data entity.SysMenu) *entity.SysMenu {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加菜单失败")
	return &data
}

func (m *sysMenuModelImpl) FindOne(menuId int64) *entity.SysMenu {
	resData := new(entity.SysMenu)
	err := global.Db.Table(m.table).Where("menu_id = ?", menuId).First(resData).Error
	biz.ErrIsNil(err, "查询菜单失败")
	return resData
}

func (m *sysMenuModelImpl) FindListPage(page, pageSize int, data entity.SysMenu) (*[]entity.SysMenu, int64) {
	list := make([]entity.SysMenu, 0)
	var total int64 = 0

	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询分页菜单失败")
	return &list, total
}

func (m *sysMenuModelImpl) FindList(data entity.SysMenu) *[]entity.SysMenu {
	list := make([]entity.SysMenu, 0)

	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.MenuName != "" {
		db = db.Where("menu_name like ?", "%"+data.MenuName+"%")
	}
	if data.Path != "" {
		db = db.Where("path = ?", data.Path)
	}
	if data.MenuType != "" {
		db = db.Where("menu_type = ?", data.MenuType)
	}
	if data.Title != "" {
		db = db.Where("title like ?", "%"+data.Title+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	db.Where("delete_time IS NULL")
	err := db.Order("sort").Find(&list).Error
	biz.ErrIsNil(err, "查询菜单列表失败")
	return &list
}

func (m *sysMenuModelImpl) Update(data entity.SysMenu) *entity.SysMenu {
	err := global.Db.Table(m.table).Select("*").Updates(data).Error
	biz.ErrIsNil(err, "修改菜单失败")
	return &data
}

func (m *sysMenuModelImpl) Delete(menuIds []int64) {
	err := global.Db.Table(m.table).Delete(&entity.SysMenu{}, "menu_id in (?)", menuIds).Error
	biz.ErrIsNil(err, "修改菜单失败")
	return
}

func (m *sysMenuModelImpl) SelectMenu(data entity.SysMenu) *[]entity.SysMenu {
	menuList := m.FindList(data)

	redData := make([]entity.SysMenu, 0)
	ml := *menuList
	for i := 0; i < len(ml); i++ {
		if ml[i].ParentId != 0 {
			continue
		}
		menusInfo := DiguiMenu(menuList, ml[i])
		redData = append(redData, menusInfo)
	}
	return &redData
}

func (m *sysMenuModelImpl) SelectMenuLable(data entity.SysMenu) *[]entity.MenuLable {
	menuList := m.FindList(data)

	redData := make([]entity.MenuLable, 0)
	ml := *menuList
	for i := 0; i < len(ml); i++ {
		if ml[i].ParentId != 0 {
			continue
		}
		e := entity.MenuLable{}
		e.MenuId = ml[i].MenuId
		e.MenuName = ml[i].MenuName
		menusInfo := DiguiMenuLable(menuList, e)

		redData = append(redData, menusInfo)
	}
	return &redData
}

func (m *sysMenuModelImpl) GetMenuByRoleKey(roleKey string) *[]entity.SysMenu {
	menus := make([]entity.SysMenu, 0)
	db := global.Db.Table(m.table).Select("sys_menus.*").Joins("left join sys_role_menus on sys_role_menus.menu_id=sys_menus.menu_id")
	db = db.Where("sys_role_menus.role_name=? and menu_type in ('M','C')", roleKey)
	db.Where("sys_menus.delete_time IS NULL")
	err := db.Order("sort").Find(&menus).Error
	biz.ErrIsNil(err, "通过角色名查询菜单失败")
	return &menus
}

func (m *sysMenuModelImpl) SelectMenuRole(roleKey string) *[]entity.SysMenu {
	redData := make([]entity.SysMenu, 0)

	menulist := m.GetMenuByRoleKey(roleKey)
	menuList := *menulist
	redData = make([]entity.SysMenu, 0)
	for i := 0; i < len(menuList); i++ {
		if menuList[i].ParentId != 0 {
			continue
		}
		menusInfo := DiguiMenu(&menuList, menuList[i])

		redData = append(redData, menusInfo)
	}
	return &redData
}
func (m *sysMenuModelImpl) GetMenuRole(data entity.MenuRole) *[]entity.MenuRole {
	menus := make([]entity.MenuRole, 0)

	db := global.Db.Table(m.table)
	if data.MenuName != "" {
		db = db.Where("menu_name = ?", data.MenuName)
	}
	db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("sort").Find(&menus).Error, "查询角色菜单失败")
	return &menus
}

func DiguiMenu(menulist *[]entity.SysMenu, menu entity.SysMenu) entity.SysMenu {
	list := *menulist

	min := make([]entity.SysMenu, 0)
	for j := 0; j < len(list); j++ {

		if menu.MenuId != list[j].ParentId {
			continue
		}
		mi := entity.SysMenu{}
		mi.MenuId = list[j].MenuId
		mi.MenuName = list[j].MenuName
		mi.Title = list[j].Title
		mi.Icon = list[j].Icon
		mi.Path = list[j].Path
		mi.MenuType = list[j].MenuType
		mi.IsKeepAlive = list[j].IsKeepAlive
		mi.Permission = list[j].Permission
		mi.ParentId = list[j].ParentId
		mi.IsAffix = list[j].IsAffix
		mi.IsIframe = list[j].IsIframe
		mi.IsLink = list[j].IsLink
		mi.Component = list[j].Component
		mi.Sort = list[j].Sort
		mi.Status = list[j].Status
		mi.IsHide = list[j].IsHide
		mi.CreatedAt = list[j].CreatedAt
		mi.UpdatedAt = list[j].UpdatedAt
		mi.Children = []entity.SysMenu{}

		if mi.MenuType != "F" {
			ms := DiguiMenu(menulist, mi)
			min = append(min, ms)

		} else {
			min = append(min, mi)
		}

	}
	menu.Children = min
	return menu
}

func DiguiMenuLable(menulist *[]entity.SysMenu, menu entity.MenuLable) entity.MenuLable {
	list := *menulist

	min := make([]entity.MenuLable, 0)
	for j := 0; j < len(list); j++ {

		if menu.MenuId != list[j].ParentId {
			continue
		}
		mi := entity.MenuLable{}
		mi.MenuId = list[j].MenuId
		mi.MenuName = list[j].MenuName
		mi.Children = []entity.MenuLable{}
		if list[j].MenuType != "F" {
			ms := DiguiMenuLable(menulist, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}

	}
	menu.Children = min
	return menu
}
