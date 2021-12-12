package entity

import (
	"pandax/base/casbin"
	"pandax/base/model"
)

type SysRole struct {
	RoleId    int64               `json:"roleId" gorm:"primary_key;AUTO_INCREMENT"`
	RoleName  string              `json:"roleName" gorm:"type:varchar(128);comment:角色名称"`
	Status    string              `json:"status" gorm:"type:varchar(1);comment:状态"`
	RoleKey   string              `json:"roleKey" gorm:"type:varchar(128);comment:角色代码"`
	RoleSort  int64               `json:"roleSort" gorm:"type:int(4);comment:角色排序"`
	DataScope string              `json:"dataScope" gorm:"type:varchar(1);comment:数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）"`
	Flag      string              `json:"flag" gorm:"type:varchar(128);comment:删除标识"`
	CreateBy  string              `json:"createBy" gorm:"type:varchar(128);"`
	UpdateBy  string              `json:"updateBy" gorm:"type:varchar(128);"`
	Remark    string              `json:"remark" gorm:"type:varchar(255);"`
	ApiIds    []casbin.CasbinRule `json:"apiIds" gorm:"-"`
	MenuIds   []int64             `json:"menuIds" gorm:"-"`
	DeptIds   []int64             `json:"deptIds" gorm:"-"`
	model.BaseModel
}

type MenuIdList struct {
	MenuId int64 `json:"menuId"`
}

type DeptIdList struct {
	DeptId int64 `json:"deptId"`
}
