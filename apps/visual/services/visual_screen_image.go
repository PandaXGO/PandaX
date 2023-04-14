package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/visual/entity"
	"pandax/pkg/global"
)

type (
	VisualScreenImageModel interface {
		Insert(data entity.VisualScreenImage) *entity.VisualScreenImage
		FindOne(fileName string) *entity.VisualScreenImage
		FindList(data entity.VisualScreenImage) *[]entity.VisualScreenImage
		Delete(fileName string)
	}

	screenImageModelImpl struct {
		table string
	}
)

var VisualScreenImageModelDao VisualScreenImageModel = &screenImageModelImpl{
	table: `visual_screen_image`,
}

func (m *screenImageModelImpl) Insert(data entity.VisualScreenImage) *entity.VisualScreenImage {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加图片失败")
	return &data
}

func (m *screenImageModelImpl) FindOne(fileName string) *entity.VisualScreenImage {
	resData := new(entity.VisualScreenImage)
	db := global.Db.Table(m.table).Where("file_name = ?", fileName)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询图片失败")
	return resData
}

func (m *screenImageModelImpl) FindList(data entity.VisualScreenImage) *[]entity.VisualScreenImage {
	list := make([]entity.VisualScreenImage, 0)
	db := global.Db.Table(m.table)
	biz.ErrIsNil(db.Find(&list).Error, "查询图片列表失败")
	return &list
}

func (m *screenImageModelImpl) Delete(fileName string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.VisualScreenImage{}, "file_name = ?", fileName).Error, "删除图片失败")
}
