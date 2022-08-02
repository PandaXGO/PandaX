package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/log/entity"
	"pandax/pkg/global"
)

type (
	LogOperModel interface {
		Insert(data entity.LogOper) *entity.LogOper
		FindOne(infoId int64) *entity.LogOper
		FindListPage(page, pageSize int, data entity.LogOper) (*[]entity.LogOper, int64)
		Delete(infoId []int64)
		DeleteAll()
	}

	logLogOperModelImpl struct {
		table string
	}
)

var LogOperModelDao LogOperModel = &logLogOperModelImpl{
	table: `log_opers`,
}

func (m *logLogOperModelImpl) Insert(data entity.LogOper) *entity.LogOper {
	global.Db.Table(m.table).Create(&data)
	return &data
}

func (m *logLogOperModelImpl) FindOne(operId int64) *entity.LogOper {
	resData := new(entity.LogOper)
	err := global.Db.Table(m.table).Where("oper_id = ?", operId).First(resData).Error
	biz.ErrIsNil(err, "查询操作日志信息失败")
	return resData
}

func (m *logLogOperModelImpl) FindListPage(page, pageSize int, data entity.LogOper) (*[]entity.LogOper, int64) {
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
	err = db.Order("create_time desc").Limit(pageSize).Offset(offset).Find(&list).Error

	biz.ErrIsNil(err, "查询操作分页日志信息失败")
	return &list, total
}

func (m *logLogOperModelImpl) Delete(operIds []int64) {
	err := global.Db.Table(m.table).Delete(&entity.LogOper{}, "oper_id in (?)", operIds).Error
	biz.ErrIsNil(err, "删除操作日志信息失败")
	return
}

func (m *logLogOperModelImpl) DeleteAll() {
	global.Db.Exec("DELETE FROM log_opers")
}
