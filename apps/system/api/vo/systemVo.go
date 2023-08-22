package vo

import "pandax/apps/system/entity"

/**
 * @Description
 * @Author 熊猫
 * @Date 2022/8/4 15:25
 **/

type DeptTreeVo struct {
	Depts       []entity.DeptLable `json:"depts"`
	CheckedKeys []int64            `json:"checkedKeys"`
}

type MenuTreeVo struct {
	Menus       []entity.MenuLable `json:"menus"`
	CheckedKeys []int64            `json:"checkedKeys"`
}

type MenuPermisVo struct {
	Menus       []RouterVo `json:"menus"`
	Permissions []string   `json:"permissions"`
}

type CaptchaVo struct {
	Base64Captcha string `json:"base64Captcha"`
	CaptchaId     string `json:"captchaId"`
}

type TokenVo struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

type AuthVo struct {
	User        entity.SysUserView `json:"user"`
	Role        entity.SysRole     `json:"role"`
	Permissions []string           `json:"permissions"`
	Menus       []RouterVo         `json:"menus"`
}

type UserProfileVo struct {
	Data    any              `json:"data"`
	PostIds []int64          `json:"postIds"`
	RoleIds []int64          `json:"roleIds"`
	Roles   []entity.SysRole `json:"roles"`
	Posts   []entity.SysPost `json:"posts"`
	Dept    []entity.SysDept `json:"dept"`
}

type UserVo struct {
	Data    any              `json:"data"`
	PostIds string           `json:"postIds"`
	RoleIds string           `json:"roleIds"`
	Roles   []entity.SysRole `json:"roles"`
	Posts   []entity.SysPost `json:"posts"`
	Depts   []entity.SysDept `json:"depts"`
}

type UserRolePost struct {
	Roles []entity.SysRole `json:"roles"`
	Posts []entity.SysPost `json:"posts"`
}
