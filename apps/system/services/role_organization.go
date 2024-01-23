package services

import (
	"fmt"
	"pandax/apps/system/entity"
	"pandax/kit/biz"
	"pandax/pkg/global"
)

type (
	SysRoleOrganizationModel interface {
		Insert(roleId int64, organizationIds []int64) bool
		FindOrganizationsByRoleId(roleId int64) ([]int64, error)
		Delete(rm entity.SysRoleOrganization)
	}

	sysRoleOrganizationImpl struct {
		table string
	}
)

var SysRoleOrganizationModelDao SysRoleOrganizationModel = &sysRoleOrganizationImpl{
	table: `sys_role_organizations`,
}

func (m *sysRoleOrganizationImpl) Insert(roleId int64, organizationIds []int64) bool {
	sql := "INSERT INTO sys_role_organizations (role_id, organization_id) VALUES "

	for i := 0; i < len(organizationIds); i++ {
		if len(organizationIds)-1 == i {
			//最后一条数据 以分号结尾
			sql += fmt.Sprintf("(%d,%d);", roleId, organizationIds[i])
		} else {
			sql += fmt.Sprintf("(%d,%d),", roleId, organizationIds[i])
		}
	}
	global.Db.Exec(sql)
	return true
}

func (m *sysRoleOrganizationImpl) FindOrganizationsByRoleId(roleId int64) ([]int64, error) {
	var result []int64
	err := global.Db.Table(m.table).Where("role_id = ?", roleId).Pluck("organization_id", &result).Error
	return result, err
}

func (m *sysRoleOrganizationImpl) Delete(rm entity.SysRoleOrganization) {
	biz.ErrIsNil(global.Db.Table(m.table).Where("role_id = ?", rm.RoleId).Delete(&rm).Error, "删除角色失败")
	return
}
