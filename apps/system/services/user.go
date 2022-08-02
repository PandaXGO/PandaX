package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"github.com/kakuilan/kgo"
	"golang.org/x/crypto/bcrypt"
	"pandax/apps/system/entity"
	"pandax/pkg/global"
	"time"
)

type (
	SysUserModel interface {
		Login(u entity.Login) *entity.SysUser
		Insert(data entity.SysUser) *entity.SysUser
		FindOne(data entity.SysUser) (resData *entity.SysUserView)
		FindListPage(page, pageSize int, data entity.SysUser) (list *[]entity.SysUserPage, total int64)
		FindList(data entity.SysUser) (list *[]entity.SysUserView)
		Update(data entity.SysUser) *entity.SysUser
		Delete(userId []int64)
		SetPwd(data entity.SysUser, pwd entity.SysUserPwd) bool
	}

	sysUserModelImpl struct {
		table string
	}
)

var SysUserModelDao SysUserModel = &sysUserModelImpl{
	table: `sys_users`,
}

func (m *sysUserModelImpl) Login(u entity.Login) *entity.SysUser {
	user := new(entity.SysUser)

	err := global.Db.Table(m.table).Where("username = ? ", u.Username).Find(user)
	biz.ErrIsNil(err.Error, "查询用户信息失败")

	// 验证密码
	b := kgo.KEncr.PasswordVerify([]byte(u.Password), []byte(user.Password))
	biz.IsTrue(b, "密码错误")

	//验证租户
	if SysTenantModelDao.FindOne(user.TenantId).ExpireTime.Unix() < time.Now().Unix() {
		biz.IsTrue(b, "租户已经过期")
	}

	return user
}

func (m *sysUserModelImpl) Insert(data entity.SysUser) *entity.SysUser {
	bytes, _ := kgo.KEncr.PasswordHash([]byte(data.Password), bcrypt.DefaultCost)
	data.Password = string(bytes)

	// check 用户名
	var count int64
	global.Db.Table(m.table).Where("username = ? and delete_time IS NULL", data.Username).Count(&count)
	biz.IsTrue(count == 0, "账户已存在！")

	biz.ErrIsNil(global.Db.Table(m.table).Create(&data).Error, "添加用户失败")
	return &data
}

func (m *sysUserModelImpl) FindOne(data entity.SysUser) *entity.SysUserView {
	resData := new(entity.SysUserView)

	db := global.Db.Table(m.table).Select([]string{"sys_users.*", "sys_roles.role_name"})
	db = db.Joins("left join sys_roles on sys_users.role_id=sys_roles.role_id")
	if data.UserId != 0 {
		db = db.Where("user_id = ?", data.UserId)
	}
	if data.TenantId != 0 {
		db = db.Where("tenant_id = ?", data.TenantId)
	}
	if data.Username != "" {
		db = db.Where("username = ?", data.Username)
	}
	if data.Password != "" {
		db = db.Where("password = ?", data.Password)
	}
	if data.RoleId != 0 {
		db = db.Where("role_id = ?", data.RoleId)
	}
	if data.DeptId != 0 {
		db = db.Where("dept_id = ?", data.DeptId)
	}
	if data.PostId != 0 {
		db = db.Where("post_id = ?", data.PostId)
	}
	biz.ErrIsNil(db.Preload("SysTenants").First(resData).Error, "查询用户失败")

	return resData
}

func (m *sysUserModelImpl) FindListPage(page, pageSize int, data entity.SysUser) (*[]entity.SysUserPage, int64) {
	list := make([]entity.SysUserPage, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table).Select("sys_users.*,sys_depts.dept_name")
	db = db.Joins("left join sys_depts on sys_depts.dept_id = sys_users.dept_id")
	// 此处填写 where参数判断
	if data.Username != "" {
		db = db.Where("sys_users.username = ?", data.Username)
	}
	if data.TenantId != 0 {
		db = db.Where("sys_users.tenant_id = ?", data.TenantId)
	}
	if data.NickName != "" {
		db = db.Where("sys_users.nick_name like ?", "%"+data.NickName+"%")
	}

	if data.Status != "" {
		db = db.Where("sys_users.status = ?", data.Status)
	}

	if data.Phone != "" {
		db = db.Where("sys_users.phone like ?", "%"+data.Phone+"%")
	}
	if data.DeptId != 0 {
		db = db.Where("sys_users.dept_id = ?", data.DeptId)
	}
	db.Where("sys_users.delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Preload("SysTenants").Find(&list).Error
	biz.ErrIsNil(err, "查询用户分页列表失败")
	return &list, total
}

func (m *sysUserModelImpl) FindList(data entity.SysUser) *[]entity.SysUserView {
	list := make([]entity.SysUserView, 0)
	// 此处填写 where参数判断
	db := global.Db.Table(m.table).Select([]string{"sys_users.*", "sys_roles.role_name"})
	db = db.Joins("left join sys_roles on sys_users.role_id=sys_roles.role_id")
	if data.UserId != 0 {
		db = db.Where("user_id = ?", data.UserId)
	}
	if data.TenantId != 0 {
		db = db.Where("tenant_id = ?", data.TenantId)
	}
	if data.Username != "" {
		db = db.Where("username = ?", data.Username)
	}

	if data.Password != "" {
		db = db.Where("password = ?", data.Password)
	}

	if data.RoleId != 0 {
		db = db.Where("role_id = ?", data.RoleId)
	}

	if data.DeptId != 0 {
		db = db.Where("dept_id = ?", data.DeptId)
	}

	if data.PostId != 0 {
		db = db.Where("post_id = ?", data.PostId)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	db.Where("sys_users.delete_time IS NULL")
	biz.ErrIsNil(db.Find(&list).Error, "查询用户列表失败")

	return &list
}

func (m *sysUserModelImpl) Update(data entity.SysUser) *entity.SysUser {
	if data.Password != "" {
		bytes, _ := kgo.KEncr.PasswordHash([]byte(data.Password), bcrypt.DefaultCost)
		data.Password = string(bytes)
	}
	update := new(entity.SysUser)
	biz.ErrIsNil(global.Db.Table(m.table).First(update, data.UserId).Error, "查询用户失败")

	if data.RoleId == 0 {
		data.RoleId = update.RoleId
	}

	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改用户失败")
	return &data
}

func (m *sysUserModelImpl) Delete(userIds []int64) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.SysUser{}, "user_id in (?)", userIds).Error, "删除用户失败")
}

func (m *sysUserModelImpl) SetPwd(data entity.SysUser, pwd entity.SysUserPwd) bool {
	user := m.FindOne(data)
	bl := kgo.KEncr.PasswordVerify([]byte(pwd.OldPassword), []byte(user.Password))
	biz.IsTrue(bl, "旧密码输入错误")

	data.Password = pwd.NewPassword
	m.Update(data)

	return true
}
