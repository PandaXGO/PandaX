package services

import (
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysMenuModel interface {
		Insert(data entity.SysMenu) (*entity.SysMenu, error)
		FindOne(menuId int64) (*entity.SysMenu, error)
		FindListPage(page, pageSize int, data entity.SysMenu) (*[]entity.SysMenu, int64, error)
		FindList(data entity.SysMenu) (*[]entity.SysMenu, error)
		Update(data entity.SysMenu) error
		Delete(menuId []int64) error
		SelectMenu(data entity.SysMenu) (*[]entity.SysMenu, error)
		SelectMenuLable(data entity.SysMenu) (*[]entity.MenuLable, error)
		SelectMenuRole(roleName string) (*[]entity.SysMenu, error)
		GetMenuRole(data entity.MenuRole) (*[]entity.MenuRole, error)
	}

	sysMenuModelImpl struct {
		table string
	}
)

var SysMenuModelDao SysMenuModel = &sysMenuModelImpl{
	table: `sys_menus`,
}

func (m *sysMenuModelImpl) Insert(data entity.SysMenu) (*entity.SysMenu, error) {
	err := global.Db.Table(m.table).Create(&data).Error
	return &data, err
}

func (m *sysMenuModelImpl) FindOne(menuId int64) (*entity.SysMenu, error) {
	resData := new(entity.SysMenu)
	err := global.Db.Table(m.table).Where("menu_id = ?", menuId).First(resData).Error
	return resData, err
}

func (m *sysMenuModelImpl) FindListPage(page, pageSize int, data entity.SysMenu) (*[]entity.SysMenu, int64, error) {
	list := make([]entity.SysMenu, 0)
	var total int64 = 0

	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *sysMenuModelImpl) FindList(data entity.SysMenu) (*[]entity.SysMenu, error) {
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
	return &list, err
}

func (m *sysMenuModelImpl) Update(data entity.SysMenu) error {
	return global.Db.Table(m.table).Select("*").Updates(data).Error
}

func (m *sysMenuModelImpl) Delete(menuIds []int64) error {
	return global.Db.Table(m.table).Delete(&entity.SysMenu{}, "menu_id in (?)", menuIds).Error
}

func (m *sysMenuModelImpl) SelectMenu(data entity.SysMenu) (*[]entity.SysMenu, error) {
	menuList, err := m.FindList(data)
	if err != nil {
		return nil, err
	}
	redData := make([]entity.SysMenu, 0)
	ml := *menuList
	for i := 0; i < len(ml); i++ {
		if ml[i].ParentId != 0 {
			continue
		}
		menusInfo := DiguiMenu(menuList, ml[i])
		redData = append(redData, menusInfo)
	}
	return &redData, nil
}

func (m *sysMenuModelImpl) SelectMenuLable(data entity.SysMenu) (*[]entity.MenuLable, error) {
	menuList, err := m.FindList(data)
	if err != nil {
		return nil, err
	}
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
	return &redData, nil
}

func (m *sysMenuModelImpl) GetMenuByRoleKey(roleKey string) (*[]entity.SysMenu, error) {
	menus := make([]entity.SysMenu, 0)
	db := global.Db.Table(m.table).Select("sys_menus.*").Joins("left join sys_role_menus on sys_role_menus.menu_id=sys_menus.menu_id")
	db = db.Where("sys_role_menus.role_name=? and menu_type in ('M','C')", roleKey)
	db.Where("sys_menus.delete_time IS NULL")
	err := db.Order("sort").Find(&menus).Error
	return &menus, err
}

func (m *sysMenuModelImpl) SelectMenuRole(roleKey string) (*[]entity.SysMenu, error) {
	redData := make([]entity.SysMenu, 0)

	menulist, err := m.GetMenuByRoleKey(roleKey)
	if err != nil {
		return nil, err
	}
	menuList := *menulist
	redData = make([]entity.SysMenu, 0)
	for i := 0; i < len(menuList); i++ {
		if menuList[i].ParentId != 0 {
			continue
		}
		menusInfo := DiguiMenu(&menuList, menuList[i])

		redData = append(redData, menusInfo)
	}
	return &redData, nil
}
func (m *sysMenuModelImpl) GetMenuRole(data entity.MenuRole) (*[]entity.MenuRole, error) {
	menus := make([]entity.MenuRole, 0)

	db := global.Db.Table(m.table)
	if data.MenuName != "" {
		db = db.Where("menu_name = ?", data.MenuName)
	}
	db.Where("delete_time IS NULL")
	err := db.Order("sort").Find(&menus).Error
	return &menus, err
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
