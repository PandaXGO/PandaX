package entity

import "github.com/PandaXGO/PandaKit/model"

// 组织组织
type SysOrganization struct {
	OrganizationId   int64             `json:"organizationId" gorm:"primary_key;AUTO_INCREMENT"` //组织编码
	ParentId         int64             `json:"parentId" gorm:"type:int;comment:上级组织"`
	OrganizationPath string            `json:"organizationPath" gorm:"type:varchar(255);comment:组织路径"`
	OrganizationName string            `json:"organizationName"  gorm:"type:varchar(128);comment:组织名称"`
	Sort             int64             `json:"sort" gorm:"type:int;comment:排序"`
	Leader           string            `json:"leader" gorm:"type:varchar(64);comment:负责人"` // userId
	Phone            string            `json:"phone" gorm:"type:varchar(11);comment:手机"`
	Email            string            `json:"email" gorm:"type:varchar(64);comment:邮箱"`
	Status           string            `json:"status" gorm:"type:varchar(1);comment:状态"`
	CreateBy         string            `json:"createBy" gorm:"type:varchar(64);comment:创建人"`
	UpdateBy         string            `json:"updateBy" gorm:"type:varchar(64);comment:修改人"`
	Children         []SysOrganization `json:"children" gorm:"-"`
	model.BaseModel
}

type OrganizationLable struct {
	OrganizationId   int64               `gorm:"-" json:"organizationId"`
	OrganizationName string              `gorm:"-" json:"organizationName"`
	Children         []OrganizationLable `gorm:"-" json:"children"`
}
