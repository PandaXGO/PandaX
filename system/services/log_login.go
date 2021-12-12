package services

import (
	"pandax/base/biz"
	"pandax/base/global"
	"pandax/system/entity"
)

type (
	LogLoginModel interface {
		Insert(data entity.LogLogin) *entity.LogLogin
		FindOne(infoId int64) *entity.LogLogin
		FindListPage(page, pageSize int, data entity.LogLogin) (*[]entity.LogLogin, int64)
		Update(data entity.LogLogin) *entity.LogLogin
		Delete(infoId []int64)
	}

	logLoginModelImpl struct {
		table string
	}
)

var LogLoginModelDao LogLoginModel = &logLoginModelImpl{
	table: `log_logins`,
}

func (m *logLoginModelImpl) Insert(data entity.LogLogin) *entity.LogLogin {
	data.CreateBy = "0"
	data.UpdateBy = "0"
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加登录日志信息失败")

	return &data
}

func (m *logLoginModelImpl) FindOne(infoId int64) *entity.LogLogin {
	resData := new(entity.LogLogin)
	err := global.Db.Table(m.table).Where("`info_id` = ?", infoId).First(resData).Error
	biz.ErrIsNil(err, "查询登录日志信息失败")
	return resData
}

func (m *logLoginModelImpl) FindListPage(page, pageSize int, data entity.LogLogin) (*[]entity.LogLogin, int64) {
	list := make([]entity.LogLogin, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.Username != "" {
		db = db.Where("username like ?", "%"+data.Username+"%")
	}
	err := db.Where("`delete_time` IS NULL").Count(&total).Error
	err = db.Order("info_id desc").Limit(pageSize).Offset(offset).Find(&list).Error

	biz.ErrIsNil(err, "查询登录分页日志信息失败")
	return &list, total
}

func (m *logLoginModelImpl) Update(data entity.LogLogin) *entity.LogLogin {
	err := global.Db.Table(m.table).Updates(&data).Error
	biz.ErrIsNil(err, "修改登录日志信息失败")
	return &data
}

func (m *logLoginModelImpl) Delete(infoIds []int64) {
	err := global.Db.Table(m.table).Delete(&entity.LogLogin{}, "`info_id` in (?)", infoIds).Error
	biz.ErrIsNil(err, "删除登录日志信息失败")
	return
}
