package services

import (
	"errors"
	"github.com/XM-GO/PandaKit/biz"
	"github.com/XM-GO/PandaKit/casbin"
	"gorm.io/gorm"
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysApiModel interface {
		Insert(data entity.SysApi) *entity.SysApi
		FindOne(id int64) *entity.SysApi
		FindListPage(page, pageSize int, data entity.SysApi) (*[]entity.SysApi, int64)
		FindList(data entity.SysApi) *[]entity.SysApi
		Update(data entity.SysApi) *entity.SysApi
		Delete(ids []int64)
	}

	sysApiModelImpl struct {
		table string
	}
)

var SysApiModelDao SysApiModel = &sysApiModelImpl{
	table: `sys_apis`,
}

func (m *sysApiModelImpl) Insert(api entity.SysApi) *entity.SysApi {
	err := global.Db.Table(m.table).Where("path = ? AND method = ?", api.Path, api.Method).First(&entity.SysApi{}).Error
	biz.IsTrue(errors.Is(err, gorm.ErrRecordNotFound), "存在相同api")
	err = global.Db.Table(m.table).Create(&api).Error
	biz.ErrIsNil(err, "新增Api失败")
	return &api
}

func (m *sysApiModelImpl) FindOne(id int64) (resData *entity.SysApi) {
	resData = new(entity.SysApi)
	err := global.Db.Table(m.table).Where("id = ?", id).First(&resData).Error
	biz.ErrIsNil(err, "查询Api失败")
	return
}

func (m *sysApiModelImpl) FindListPage(page, pageSize int, data entity.SysApi) (*[]entity.SysApi, int64) {
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
	err = db.Order("api_group").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询配置分页列表信息失败")
	return &list, total
}

func (m *sysApiModelImpl) FindList(data entity.SysApi) *[]entity.SysApi {
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
	biz.ErrIsNil(err, "查询Api列表信息失败")
	return &list
}

func (m *sysApiModelImpl) Update(api entity.SysApi) *entity.SysApi {
	var oldA entity.SysApi
	err := global.Db.Table(m.table).Where("id = ?", api.Id).First(&oldA).Error
	biz.ErrIsNil(err, "【修改api】查询api失败")
	if oldA.Path != api.Path || oldA.Method != api.Method {
		err := global.Db.Table(m.table).Where("path = ? AND method = ?", api.Path, api.Method).First(&entity.SysApi{}).Error
		biz.IsTrue(errors.Is(err, gorm.ErrRecordNotFound), "存在相同api路径")
	}
	// 异常直接抛错误
	ca := casbin.CasbinS{ModelPath: global.Conf.Casbin.ModelPath}
	ca.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
	err = global.Db.Table(m.table).Model(&api).Updates(&api).Error
	biz.ErrIsNil(err, "修改api信息失败")
	return &api
}

func (m *sysApiModelImpl) Delete(ids []int64) {
	err := global.Db.Table(m.table).Delete(&entity.SysApi{}, "id in (?)", ids).Error
	biz.ErrIsNil(err, "删除配置信息失败")
}
