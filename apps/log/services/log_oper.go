package services

import (
	"pandax/apps/log/entity"
	"pandax/pkg/global"
)

type (
	LogOperModel interface {
		Insert(data entity.LogOper) (*entity.LogOper, error)
		FindOne(infoId int64) (*entity.LogOper, error)
		FindListPage(page, pageSize int, data entity.LogOper) (*[]entity.LogOper, int64, error)
		Delete(infoId []int64) error
		DeleteAll() error
	}

	logLogOperModelImpl struct {
		table string
	}
)

var LogOperModelDao LogOperModel = &logLogOperModelImpl{
	table: `log_opers`,
}

func (m *logLogOperModelImpl) Insert(data entity.LogOper) (*entity.LogOper, error) {
	err := global.Db.Table(m.table).Create(&data).Error
	return &data, err
}

func (m *logLogOperModelImpl) FindOne(operId int64) (*entity.LogOper, error) {
	resData := new(entity.LogOper)
	err := global.Db.Table(m.table).Where("oper_id = ?", operId).First(resData).Error
	return resData, err
}

func (m *logLogOperModelImpl) FindListPage(page, pageSize int, data entity.LogOper) (*[]entity.LogOper, int64, error) {
	list := make([]entity.LogOper, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.BusinessType != "" {
		db = db.Where("business_type = ?", data.BusinessType)
	}
	if data.OperLocation != "" {
		db = db.Where("oper_location like ?", "%"+data.OperLocation+"%")
	}
	if data.Title != "" {
		db = db.Where("title like ?", "%"+data.Title+"%")
	}
	if data.OperName != "" {
		db = db.Where("oper_name like ?", "%"+data.OperName+"%")
	}
	err := db.Where("delete_time IS NULL").Count(&total).Error
	if err != nil {
		return &list, total, err
	}
	err = db.Order("create_time desc").Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *logLogOperModelImpl) Delete(operIds []int64) error {
	err := global.Db.Table(m.table).Delete(&entity.LogOper{}, "oper_id in (?)", operIds).Error
	return err
}

func (m *logLogOperModelImpl) DeleteAll() error {
	return global.Db.Exec("DELETE FROM log_opers").Error
}
