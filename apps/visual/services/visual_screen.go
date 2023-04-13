// ==========================================================================
// 生成日期：2023-04-10 11:26:49 +0800 CST
// 生成路径: apps/visual/services/visual_screen.go
// 生成人：panda
// ==========================================================================

package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/visual/entity"
	"pandax/pkg/global"
)

type (
	VisualScreenModel interface {
		Insert(data entity.VisualScreen) *entity.VisualScreen
		FindOne(screenId string) *entity.VisualScreen
		FindListPage(page, pageSize int, data entity.VisualScreen) (*[]entity.VisualScreen, int64)
		FindList(data entity.VisualScreen) *[]entity.VisualScreen
		Update(data entity.VisualScreen) *entity.VisualScreen
		Delete(screenIds []string)
	}

	screenModelImpl struct {
		table string
	}
)

var VisualScreenModelDao VisualScreenModel = &screenModelImpl{
	table: `visual_screen`,
}

func (m *screenModelImpl) Insert(data entity.VisualScreen) *entity.VisualScreen {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加bi大屏失败")
	return &data
}

func (m *screenModelImpl) FindOne(screenId string) *entity.VisualScreen {
	resData := new(entity.VisualScreen)
	db := global.Db.Table(m.table).Where("screen_id = ?", screenId)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询bi大屏失败")
	return resData
}

func (m *screenModelImpl) FindListPage(page, pageSize int, data entity.VisualScreen) (*[]entity.VisualScreen, int64) {
	list := make([]entity.VisualScreen, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.UserId != 0 {
		db = db.Where("user_id = ?", data.UserId)
	}
	db.Where("delete_time IS NULL")
	if data.ScreenName != "" {
		db = db.Where("screen_name like ?", "%"+data.ScreenName+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.GroupId != 0 {
		db = db.Where("group_id = ?", data.GroupId)
	}
	if data.ScreenRemark != "" {
		db = db.Where("screen_remark like ?", "%"+data.ScreenRemark+"%")
	}
	if data.Creator != "" {
		db = db.Where("creator = ?", data.Creator)
	}

	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询bi大屏分页列表失败")
	return &list, total
}

func (m *screenModelImpl) FindList(data entity.VisualScreen) *[]entity.VisualScreen {
	list := make([]entity.VisualScreen, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.UserId != 0 {
		db = db.Where("user_id = ?", data.UserId)
	}
	db.Where("delete_time IS NULL")
	if data.ScreenName != "" {
		db = db.Where("screen_name like ?", "%"+data.ScreenName+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.ScreenRemark != "" {
		db = db.Where("screen_remark = ?", data.ScreenRemark)
	}
	if data.Creator != "" {
		db = db.Where("screen_remark like ?", "%"+data.ScreenRemark+"%")
	}
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询bi大屏列表失败")
	return &list
}

func (m *screenModelImpl) Update(data entity.VisualScreen) *entity.VisualScreen {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改bi大屏失败")
	return &data
}

func (m *screenModelImpl) Delete(screenIds []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.VisualScreen{}, "screen_id in (?)", screenIds).Error, "删除bi大屏失败")
}
