package entity

import "github.com/XM-GO/PandaKit/model"

type SysDept struct {
	DeptId   int64     `json:"deptId" gorm:"primary_key;AUTO_INCREMENT"` //部门编码
	TenantId int64     `json:"tenantId" gorm:"type:int;comment:租户Id"`
	ParentId int64     `json:"parentId" gorm:"type:int;comment:上级部门"`
	DeptPath string    `json:"deptPath" gorm:"type:varchar(255);comment:部门路径"`
	DeptName string    `json:"deptName"  gorm:"type:varchar(128);comment:部门名称"`
	Sort     int64     `json:"sort" gorm:"type:int;comment:排序"`
	Leader   string    `json:"leader" gorm:"type:varchar(64);comment:负责人"` // userId
	Phone    string    `json:"phone" gorm:"type:varchar(11);comment:手机"`
	Email    string    `json:"email" gorm:"type:varchar(64);comment:邮箱"`
	Status   string    `json:"status" gorm:"type:varchar(1);comment:状态"`
	CreateBy string    `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
	UpdateBy string    `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
	Children []SysDept `json:"children" gorm:"-"`
	model.BaseModel
}

type DeptLable struct {
	DeptId   int64       `gorm:"-" json:"deptId"`
	DeptName string      `gorm:"-" json:"deptName"`
	Children []DeptLable `gorm:"-" json:"children"`
}
