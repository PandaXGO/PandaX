package entity

type SysRoleMenu struct {
	RoleId   int64  `gorm:"type:int"`
	MenuId   int64  `gorm:"type:int"`
	RoleName string `gorm:"type:varchar(128)"`
	Id       int64  `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`
}

type MenuPath struct {
	Path string `json:"path"`
}
