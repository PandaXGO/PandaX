package form

// User register structure
type Register struct {
	Username    string `json:"userName"`
	Password    string `json:"passWord"`
	NickName    string `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg   string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	AuthorityId string `json:"authorityId" gorm:"default:888"`
}

// User login structure
type Login struct {
	Username  string `json:"username" ` // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// Modify password structure
type ChangePasswordStruct struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

type UserSearch struct {
	Username       string `json:"username"`       // 用户UUID
	NickName       string `json:"nickName"`       // 角色ID
	Status         int64  `json:"status"`         // 角色ID
	Phone          string `json:"phone"`          // 角色ID
	PostId         int64  `json:"postId"`         // 角色ID
	OrganizationId int64  `json:"organizationId"` // 角色ID
}
