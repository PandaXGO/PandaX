package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/resource/entity"
	"pandax/pkg/global"
)

/**
 * @Description 添加qq群467890197 交流学习
 * @Author 熊猫
 * @Date 2022/1/13 15:17
 **/

// ==========================================================================
// 生成日期：2022-01-13 15:12:14 +0800 CST
// 生成路径: apps/admin/services/res_osses.go
// 生成人：panda
// ==========================================================================

type (
	ResOssesModel interface {
		Insert(data entity.ResOss) *entity.ResOss
		FindOne(ossId int64) *entity.ResOss
		FindListPage(page, pageSize int, data entity.ResOss) (*[]entity.ResOss, int64)
		FindList(data entity.ResOss) *[]entity.ResOss
		Update(data entity.ResOss) *entity.ResOss
		Delete(ossIds []int64)
	}

	resOssesModelImpl struct {
		table string
	}
)

var ResOssesModelDao ResOssesModel = &resOssesModelImpl{
	table: `res_osses`,
}

func (m *resOssesModelImpl) Insert(data entity.ResOss) *entity.ResOss {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加ResOsses失败")
	return &data
}

func (m *resOssesModelImpl) FindOne(ossId int64) *entity.ResOss {
	resData := new(entity.ResOss)
	err := global.Db.Table(m.table).Where("oss_id = ?", ossId).First(resData).Error
	biz.ErrIsNil(err, "查询ResOsses失败")
	return resData
}

func (m *resOssesModelImpl) FindListPage(page, pageSize int, data entity.ResOss) (*[]entity.ResOss, int64) {
	list := make([]entity.ResOss, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.OssId != 0 {
		db = db.Where("oss_id = ?", data.OssId)
	}
	if data.OssCode != "" {
		db = db.Where("oss_code = ?", data.OssCode)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.Category != "" {
		db = db.Where("category = ?", data.Category)
	}

	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询ResOsses分页列表失败")
	return &list, total
}

func (m *resOssesModelImpl) FindList(data entity.ResOss) *[]entity.ResOss {
	list := make([]entity.ResOss, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.OssId != 0 {
		db = db.Where("oss_id = ?", data.OssId)
	}
	if data.OssCode != "" {
		db = db.Where("oss_code = ?", data.OssCode)
	}
	db.Where("status = '0' AND delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询ResOsses列表失败")
	return &list
}

func (m *resOssesModelImpl) Update(data entity.ResOss) *entity.ResOss {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改ResOsses失败")
	return &data
}

func (m *resOssesModelImpl) Delete(ossIds []int64) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.ResOss{}, "oss_id in (?)", ossIds).Error, "删除ResOsses失败")
}
