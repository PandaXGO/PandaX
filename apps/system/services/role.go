package services

import (
	"errors"
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysRoleModel interface {
		Insert(data entity.SysRole) (*entity.SysRole, error)
		FindOne(roleId int64) (*entity.SysRole, error)
		FindListPage(page, pageSize int, data entity.SysRole) (list *[]entity.SysRole, total int64, err error)
		FindList(data entity.SysRole) (list *[]entity.SysRole, err error)
		Update(data entity.SysRole) (*entity.SysRole, error)
		Delete(roleId []int64) error
		GetRoleMeunId(data entity.SysRole) ([]int64, error)
		GetRoleOrganizationId(data entity.SysRole) ([]int64, error)
		FindOrganizationsByRoleId(roleId int64) (entity.SysRoleAuth, error)
	}

	sysRoleModel struct {
		table string
	}
)

var SysRoleModelDao SysRoleModel = &sysRoleModel{
	table: `sys_roles`,
}

func (m *sysRoleModel) Insert(data entity.SysRole) (*entity.SysRole, error) {
	var i int64 = 0
	global.Db.Table(m.table).Where("(role_name = ? or role_key = ?) and delete_time IS NULL", data.RoleName, data.RoleKey).Count(&i)
	if i != 0 {
		return nil, errors.New("角色名称或者角色标识已经存在！")
	}

	data.UpdateBy = ""
	err := global.Db.Table(m.table).Create(&data).Error
	return &data, err
}

func (m *sysRoleModel) FindOne(roleId int64) (*entity.SysRole, error) {
	resData := new(entity.SysRole)
	err := global.Db.Table(m.table).Where("role_id = ?", roleId).First(resData).Error
	return resData, err
}

func (m *sysRoleModel) FindListPage(page, pageSize int, data entity.SysRole) (*[]entity.SysRole, int64, error) {

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
	return &list, total, err
}

func (m *sysRoleModel) FindList(data entity.SysRole) (*[]entity.SysRole, error) {
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
	err := db.Order("role_sort").Find(&list).Error
	return &list, err
}

func (m *sysRoleModel) Update(data entity.SysRole) (*entity.SysRole, error) {
	update := new(entity.SysRole)
	err := global.Db.Table(m.table).First(update, data.RoleId).Error
	if err != nil {
		return nil, err
	}
	if data.RoleKey != "" && data.RoleKey != update.RoleKey {
		return nil, errors.New("角色标识不允许修改！")
	}
	err = global.Db.Table(m.table).Updates(&data).Error
	return &data, err
}

func (m *sysRoleModel) Delete(roleIds []int64) error {
	return global.Db.Table(m.table).Delete(&entity.SysRole{}, "role_id in (?)", roleIds).Error

}

// 获取角色对应的菜单ids
func (m *sysRoleModel) GetRoleMeunId(data entity.SysRole) ([]int64, error) {
	menuIds := make([]int64, 0)
	menuList := make([]entity.MenuIdList, 0)
	err := global.Db.Table("sys_menus").Select("sys_menus.menu_id").Joins("LEFT JOIN sys_role_menus on sys_role_menus.menu_id=sys_menus.menu_id").Where("sys_role_menus.role_id = ? ", data.RoleId).Where("sys_menus.menu_id not in (select sys_menus.parent_id from sys_menus LEFT JOIN sys_role_menus on sys_menus.menu_id=sys_role_menus.menu_id where sys_role_menus.role_id =? )", data.RoleId).Find(&menuList).Error
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(menuList); i++ {
		menuIds = append(menuIds, menuList[i].MenuId)
	}
	return menuIds, nil
}

func (m *sysRoleModel) GetRoleOrganizationId(data entity.SysRole) ([]int64, error) {
	organizationIds := make([]int64, 0)
	organizationList := make([]entity.OrganizationIdList, 0)
	err := global.Db.Table("sys_role_organizations").Select("sys_role_organizations.organization_id").Joins("LEFT JOIN sys_organizations on sys_organizations.organization_id=sys_role_organizations.organization_id").Where("role_id = ? ", data.RoleId).Where(" sys_role_organizations.organization_id not in(select sys_organizations.parent_id from sys_role_organizations LEFT JOIN sys_organizations on sys_organizations.organization_id=sys_role_organizations.organization_id where role_id =? )", data.RoleId).Find(&organizationList).Error
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(organizationList); i++ {
		organizationIds = append(organizationIds, organizationList[i].OrganizationId)
	}
	return organizationIds, nil
}

func (m *sysRoleModel) FindOrganizationsByRoleId(roleId int64) (entity.SysRoleAuth, error) {
	var roleData entity.SysRoleAuth
	GROUP_CONCAT := "GROUP_CONCAT(sys_role_organizations.organization_id) as org"
	if global.Conf.Server.DbType == "postgresql" {
		GROUP_CONCAT = "string_agg(CAST(sys_role_organizations.organization_id AS VARCHAR), ',')  as org"
	}
	err := global.Db.Raw("SELECT sys_roles.data_scope, "+GROUP_CONCAT+" FROM sys_roles LEFT JOIN sys_role_organizations ON sys_roles.role_id =  sys_role_organizations.role_id WHERE sys_roles.role_id = ? GROUP BY sys_roles.role_id", roleId).Scan(&roleData).Error
	return roleData, err
}
