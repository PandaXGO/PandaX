package services

import (
	"fmt"
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysRoleDeptModel interface {
		Insert(roleId int64, deptIds []int64) bool
		Delete(rm entity.SysRoleDept)
	}

	sysRoleDeptImpl struct {
		table string
	}
)

var SysRoleDeptModelDao SysRoleDeptModel = &sysRoleDeptImpl{
	table: `sys_role_depts`,
}

func (m *sysRoleDeptImpl) Insert(roleId int64, deptIds []int64) bool {
	sql := "INSERT INTO sys_role_depts (role_id, dept_id) VALUES "

	for i := 0; i < len(deptIds); i++ {
		if len(deptIds)-1 == i {
			//最后一条数据 以分号结尾
			sql += fmt.Sprintf("(%d,%d);", roleId, deptIds[i])
		} else {
			sql += fmt.Sprintf("(%d,%d),", roleId, deptIds[i])
		}
	}
	global.Db.Exec(sql)
	return true
}

func (m *sysRoleDeptImpl) Delete(rm entity.SysRoleDept) {
	biz.ErrIsNil(global.Db.Table(m.table).Where("role_id = ?", rm.RoleId).Delete(&rm).Error, "删除角色失败")
	return
}
