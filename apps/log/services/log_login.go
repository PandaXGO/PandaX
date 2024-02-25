package services

import (
	"pandax/apps/log/entity"
	"pandax/pkg/global"
)

type (
	LogLoginModel interface {
		Insert(data entity.LogLogin) (*entity.LogLogin, error)
		FindOne(infoId int64) (*entity.LogLogin, error)
		FindListPage(page, pageSize int, data entity.LogLogin) (*[]entity.LogLogin, int64, error)
		Update(data entity.LogLogin) (*entity.LogLogin, error)
		Delete(infoId []int64) error
		DeleteAll() error
	}

	logLoginModelImpl struct {
		table string
	}
)

var LogLoginModelDao LogLoginModel = &logLoginModelImpl{
	table: `log_logins`,
}

func (m *logLoginModelImpl) Insert(data entity.LogLogin) (*entity.LogLogin, error) {
	err := global.Db.Table(m.table).Create(&data).Error
	return &data, err
}

func (m *logLoginModelImpl) FindOne(infoId int64) (*entity.LogLogin, error) {
	resData := new(entity.LogLogin)
	err := global.Db.Table(m.table).Where("info_id = ?", infoId).First(resData).Error
	return resData, err
}

func (m *logLoginModelImpl) FindListPage(page, pageSize int, data entity.LogLogin) (*[]entity.LogLogin, int64, error) {
	list := make([]entity.LogLogin, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.LoginLocation != "" {
		db = db.Where("login_location like ?", "%"+data.LoginLocation+"%")
	}
	if data.Username != "" {
		db = db.Where("username like ?", "%"+data.Username+"%")
	}
	err := db.Where("delete_time IS NULL").Count(&total).Error
	if err != nil {
		return &list, total, err
	}
	err = db.Order("info_id desc").Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *logLoginModelImpl) Update(data entity.LogLogin) (*entity.LogLogin, error) {
	err := global.Db.Table(m.table).Updates(&data).Error
	return &data, err
}

func (m *logLoginModelImpl) Delete(infoIds []int64) error {
	err := global.Db.Table(m.table).Delete(&entity.LogLogin{}, "info_id in (?)", infoIds).Error
	return err
}

func (m *logLoginModelImpl) DeleteAll() error {
	return global.Db.Exec("DELETE FROM log_logins").Error
}
