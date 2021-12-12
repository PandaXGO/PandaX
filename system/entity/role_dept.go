package entity

type SysRoleDept struct {
	RoleId int64 `gorm:"type:int(11)"`
	DeptId int64 `gorm:"type:int(11)"`
	Id     int64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id" form:"id"`
}
