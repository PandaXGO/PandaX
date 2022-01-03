package entity

type SysRoleDept struct {
	RoleId int64 `gorm:"type:int"`
	DeptId int64 `gorm:"type:int"`
	Id     int64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`
}
