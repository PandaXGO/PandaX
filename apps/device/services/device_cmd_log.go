package services

import (
	"github.com/PandaXGO/PandaKit/biz"
	"pandax/apps/device/entity"
	"pandax/pkg/global"
	"time"
)

type (
	DeviceCmdLogModel interface {
		Insert(data entity.DeviceCmdLog) error
		FindOne(id string) *entity.DeviceCmdLog
		FindListPage(page, pageSize int, data entity.DeviceCmdLog) (*[]entity.DeviceCmdLog, int64)
		Update(data entity.DeviceCmdLog) error
		UpdateResp(id, responseContent, responseTime string) error
		Delete(ids []string)
	}

	cmdLogModelImpl struct {
		table string
	}
)

var DeviceCmdLogModelDao DeviceCmdLogModel = &cmdLogModelImpl{
	table: `device_cmd_log`,
}

func (m *cmdLogModelImpl) Insert(data entity.DeviceCmdLog) error {
	err := global.Db.Table(m.table).Create(&data).Error
	return err
}

func (m *cmdLogModelImpl) FindOne(id string) *entity.DeviceCmdLog {
	resData := new(entity.DeviceCmdLog)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询设备指令下发失败")
	return resData
}

func (m *cmdLogModelImpl) FindListPage(page, pageSize int, data entity.DeviceCmdLog) (*[]entity.DeviceCmdLog, int64) {
	list := make([]entity.DeviceCmdLog, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	if data.DeviceId != "" {
		db = db.Where("device_id = ?", data.DeviceId)
	}
	if data.CmdName != "" {
		db = db.Where("cmd_name like ?", "%"+data.CmdName+"%")
	}
	if data.Type != "" {
		db = db.Where("type = ?", data.Type)
	}
	if data.State != "" {
		db = db.Where("state = ?", data.State)
	}
	err := db.Count(&total).Error
	err = db.Order("request_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询设备指令下发分页列表失败")
	return &list, total
}

func (m *cmdLogModelImpl) Update(data entity.DeviceCmdLog) error {
	err := global.Db.Table(m.table).
		Where("id = ?", data.Id).
		Updates(&data).Error
	return err
}

func (m *cmdLogModelImpl) UpdateResp(id, responseContent, state string) error {
	responseTime := time.Now().Format("2006-01-02 15:04:05")
	err := global.Db.Table(m.table).
		Where("id = ?", id).
		Update("response_content", responseContent).
		Update("response_time", responseTime).Update("state", state).Error
	return err
}

func (m *cmdLogModelImpl) Delete(id []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.DeviceCmdLog{}, "id in (?)", id).Error, "删除设备指令下发失败")
}
