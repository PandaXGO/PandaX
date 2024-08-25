package entity

import "github.com/PandaXGO/PandaKit/model"

type LoginM struct {
	Username string `gorm:"type:varchar(64)" json:"username"`
	Password string `gorm:"type:varchar(128)" json:"password"`
}

type SysUserId struct {
	UserId int64 `gorm:"primary_key;AUTO_INCREMENT"  json:"userId"` // 编码
}

type SysUserB struct {
	NickName       string `gorm:"type:varchar(128)" json:"nickName"` // 昵称
	Phone          string `gorm:"type:varchar(11)" json:"phone"`     // 手机号
	RoleId         int64  `gorm:"type:int" json:"roleId"`            // 角色编码
	Salt           string `gorm:"type:varchar(255)" json:"salt"`     //盐
	Avatar         string `gorm:"type:varchar(255)" json:"avatar"`   //头像
	Sex            string `gorm:"type:varchar(255)" json:"sex"`      //性别
	Email          string `gorm:"type:varchar(128)" json:"email"`    //邮箱
	OrganizationId int64  `gorm:"type:int" json:"organizationId"`    //组织编码
	PostId         int64  `gorm:"type:int" json:"postId"`            //职位编码
	RoleIds        string `gorm:"type:varchar(255)" json:"roleIds"`  //多角色
	PostIds        string `gorm:"type:varchar(255)" json:"postIds"`  // 多岗位
	CreateBy       string `gorm:"type:varchar(128)" json:"createBy"` //
	UpdateBy       string `gorm:"type:varchar(128)" json:"updateBy"` //
	Remark         string `gorm:"type:varchar(255)" json:"remark"`   //备注
	Status         string `gorm:"type:varchar(1);" json:"status"`
	model.BaseModel
}

type SysUser struct {
	SysUserId
	SysUserB
	LoginM
}

type SysUserPwd struct {
	OldPassword string `json:"oldPassword" form:"oldPassword"`
	NewPassword string `json:"newPassword" form:"newPassword"`
}

type SysUserPage struct {
	SysUserId
	SysUserB
	LoginM
	OrganizationName string `gorm:"-" json:"organizationName"`
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
	UUID     string `form:"UUID" json:"uuid" binding:"required"`
}

type SysUserView struct {
	SysUserId
	SysUserB
	LoginM
	RoleName string `gorm:"column:role_name"  json:"role_name"`
}
