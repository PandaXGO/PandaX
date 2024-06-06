package services

import (
	"errors"
	"pandax/apps/system/entity"
	"pandax/kit/casbin"
	"pandax/pkg/global"

	"gorm.io/gorm"
)

type (
	SysApiModel interface {
		Insert(data entity.SysApi) (*entity.SysApi, error)
		FindOne(id int64) (*entity.SysApi, error)
		FindListPage(page, pageSize int, data entity.SysApi) (*[]entity.SysApi, int64, error)
		FindList(data entity.SysApi) (*[]entity.SysApi, error)
		Update(data entity.SysApi) error
		Delete(ids []int64) error
	}

	sysApiModelImpl struct {
		table string
	}
)

var SysApiModelDao SysApiModel = &sysApiModelImpl{
	table: `sys_apis`,
}

func (m *sysApiModelImpl) Insert(api entity.SysApi) (*entity.SysApi, error) {
	err := global.Db.Table(m.table).Where("path = ? AND method = ?", api.Path, api.Method).First(&entity.SysApi{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("存在相同api")
	}
	err = global.Db.Table(m.table).Create(&api).Error
	return &api, err
}

func (m *sysApiModelImpl) FindOne(id int64) (resData *entity.SysApi, err error) {
	resData = new(entity.SysApi)
	err = global.Db.Table(m.table).Where("id = ?", id).First(&resData).Error
	return
}

func (m *sysApiModelImpl) FindListPage(page, pageSize int, data entity.SysApi) (*[]entity.SysApi, int64, error) {
	list := make([]entity.SysApi, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)

	db := global.Db.Table(m.table)

	if data.Path != "" {
		db = db.Where("path LIKE ?", "%"+data.Path+"%")
	}

	if data.Description != "" {
		db = db.Where("description LIKE ?", "%"+data.Description+"%")
	}

	if data.Method != "" {
		db = db.Where("method = ?", data.Method)
	}

	if data.ApiGroup != "" {
		db = db.Where("api_group = ?", data.ApiGroup)
	}

	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Order("api_group").Limit(pageSize).Offset(offset).Find(&list).Error
	return &list, total, err
}

func (m *sysApiModelImpl) FindList(data entity.SysApi) (*[]entity.SysApi, error) {
	list := make([]entity.SysApi, 0)
	db := global.Db.Table(m.table)

	if data.Path != "" {
		db = db.Where("path LIKE ?", "%"+data.Path+"%")
	}

	if data.Description != "" {
		db = db.Where("description LIKE ?", "%"+data.Description+"%")
	}

	if data.Method != "" {
		db = db.Where("method = ?", data.Method)
	}

	if data.ApiGroup != "" {
		db = db.Where("api_group = ?", data.ApiGroup)
	}
	db.Where("delete_time IS NULL")
	err := db.Order("api_group").Find(&list).Error
	return &list, err
}

func (m *sysApiModelImpl) Update(api entity.SysApi) error {
	var oldA entity.SysApi
	err := global.Db.Table(m.table).Where("id = ?", api.Id).First(&oldA).Error
	if err != nil {
		return err
	}
	if oldA.Path != api.Path || oldA.Method != api.Method {
		err := global.Db.Table(m.table).Where("path = ? AND method = ?", api.Path, api.Method).First(&entity.SysApi{}).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	// 异常直接抛错误
	ca := casbin.CasbinService{ModelPath: global.Conf.Casbin.ModelPath}
	err = ca.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
	if err != nil {
		return err
	}
	err = global.Db.Table(m.table).Model(&api).Updates(&api).Error
	return err
}

func (m *sysApiModelImpl) Delete(ids []int64) error {
	return global.Db.Table(m.table).Delete(&entity.SysApi{}, "id in (?)", ids).Error
}
