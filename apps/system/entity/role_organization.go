package entity

type SysRoleOrganization struct {
	RoleId         int64 `gorm:"type:int"`
	OrganizationId int64 `gorm:"type:int"`
	Id             int64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`
}
