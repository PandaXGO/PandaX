package services

import (
	"fmt"
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysRoleOrganizationModel interface {
		Insert(roleId int64, organizationIds []int64) error
		FindOrganizationsByRoleId(roleId int64) ([]int64, error)
		Delete(rm entity.SysRoleOrganization) error
	}

	sysRoleOrganizationImpl struct {
		table string
	}
)

var SysRoleOrganizationModelDao SysRoleOrganizationModel = &sysRoleOrganizationImpl{
	table: `sys_role_organizations`,
}

func (m *sysRoleOrganizationImpl) Insert(roleId int64, organizationIds []int64) error {
	sql := "INSERT INTO sys_role_organizations (role_id, organization_id) VALUES "

	for i := 0; i < len(organizationIds); i++ {
		if len(organizationIds)-1 == i {
			//最后一条数据 以分号结尾
			sql += fmt.Sprintf("(%d,%d);", roleId, organizationIds[i])
		} else {
			sql += fmt.Sprintf("(%d,%d),", roleId, organizationIds[i])
		}
	}

	return global.Db.Exec(sql).Error
}

func (m *sysRoleOrganizationImpl) FindOrganizationsByRoleId(roleId int64) ([]int64, error) {
	var result []int64
	err := global.Db.Table(m.table).Where("role_id = ?", roleId).Pluck("organization_id", &result).Error
	return result, err
}

func (m *sysRoleOrganizationImpl) Delete(rm entity.SysRoleOrganization) error {
	return global.Db.Table(m.table).Where("role_id = ?", rm.RoleId).Delete(&rm).Error
}
