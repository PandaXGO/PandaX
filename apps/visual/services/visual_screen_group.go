// ==========================================================================
// 生成日期：2023-04-10 02:51:27 +0000 UTC
// 生成路径: apps/visual/services/visual_data_set_group.go
// 生成人：panda
// ==========================================================================

package services

import (
	"errors"
	"github.com/XM-GO/PandaKit/biz"
	"github.com/kakuilan/kgo"
	"pandax/apps/visual/entity"
	"pandax/pkg/global"
)

type (
	VisualScreenGroupModel interface {
		Insert(data entity.VisualScreenGroup) *entity.VisualScreenGroup
		FindOne(id int64) *entity.VisualScreenGroup
		FindListPage(page, pageSize int, data entity.VisualScreenGroup) (*[]entity.VisualScreenGroup, int64)
		FindList(data entity.VisualScreenGroup) *[]entity.VisualScreenGroup
		Update(data entity.VisualScreenGroup) *entity.VisualScreenGroup
		Delete(ids []int64)
		SelectScreenGroup(data entity.VisualScreenGroup) []entity.VisualScreenGroup
		SelectScreenGroupLabel(data entity.VisualScreenGroup) []entity.ScreenGroupLabel
	}

	screenGroupModelImpl struct {
		table string
	}
)

var VisualScreenGroupModelDao VisualScreenGroupModel = &screenGroupModelImpl{
	table: `visual_screen_group`,
}

func (m *screenGroupModelImpl) Insert(data entity.VisualScreenGroup) *entity.VisualScreenGroup {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加数据集分组失败")

	path := "/" + kgo.KConv.Int2Str(data.Id)
	if int(data.Pid) != 0 {
		vsg := m.FindOne(data.Pid)
		path = vsg.Path + path
	} else {
		path = "/0" + path
	}
	data.Path = path
	biz.ErrIsNil(global.Db.Table(m.table).Model(&data).Updates(&data).Error, "修改分组信息失败")
	return &data
}

func (m *screenGroupModelImpl) FindOne(id int64) *entity.VisualScreenGroup {
	resData := new(entity.VisualScreenGroup)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询数据集分组失败")
	return resData
}

func (m *screenGroupModelImpl) FindListPage(page, pageSize int, data entity.VisualScreenGroup) (*[]entity.VisualScreenGroup, int64) {
	list := make([]entity.VisualScreenGroup, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Path != "" {
		db = db.Where("path like %?%", "%"+data.Path+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	err := db.Count(&total).Error
	err = db.Order("sort").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询数据集分组分页列表失败")
	return &list, total
}

func (m *screenGroupModelImpl) FindList(data entity.VisualScreenGroup) *[]entity.VisualScreenGroup {
	list := make([]entity.VisualScreenGroup, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.Name != "" {
		db = db.Where("name like ?", "%"+data.Name+"%")
	}
	if data.Path != "" {
		db = db.Where("path like %?%", "%"+data.Path+"%")
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	biz.ErrIsNil(db.Order("sort").Find(&list).Error, "查询数据集分组列表失败")
	return &list
}

func (m *screenGroupModelImpl) Update(data entity.VisualScreenGroup) *entity.VisualScreenGroup {
	one := m.FindOne(data.Id)

	path := "/" + kgo.KConv.Int2Str(data.Id)
	if int(data.Pid) != 0 {
		vsg := m.FindOne(data.Pid)
		path = vsg.Path + path
	} else {
		path = "/0" + path
	}
	data.Path = path

	if data.Path != "" && data.Path != one.Path {
		biz.ErrIsNil(errors.New("上级分组不允许修改！"), "上级分组不允许修改")
	}

	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改数据集分组失败")
	return &data
}

func (m *screenGroupModelImpl) Delete(s []int64) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.VisualScreenGroup{}, "id in (?)", s).Error, "删除数据集分组失败")
}

func (m *screenGroupModelImpl) SelectScreenGroup(data entity.VisualScreenGroup) []entity.VisualScreenGroup {
	list := m.FindList(data)
	sd := make([]entity.VisualScreenGroup, 0)
	li := *list
	for i := 0; i < len(li); i++ {
		if li[i].Pid != 0 {
			continue
		}
		info := Digui(list, li[i])

		sd = append(sd, info)
	}
	return sd
}

func (m *screenGroupModelImpl) SelectScreenGroupLabel(data entity.VisualScreenGroup) []entity.ScreenGroupLabel {
	deptlist := m.FindList(data)

	dl := make([]entity.ScreenGroupLabel, 0)
	deptl := *deptlist
	for i := 0; i < len(deptl); i++ {
		if deptl[i].Pid != 0 {
			continue
		}
		e := entity.ScreenGroupLabel{}
		e.Id = deptl[i].Id
		e.Name = deptl[i].Name
		deptsInfo := DiguiDeptLable(deptlist, e)

		dl = append(dl, deptsInfo)
	}
	return dl
}

func Digui(sglist *[]entity.VisualScreenGroup, menu entity.VisualScreenGroup) entity.VisualScreenGroup {
	list := *sglist

	min := make([]entity.VisualScreenGroup, 0)
	for j := 0; j < len(list); j++ {

		if menu.Id != list[j].Pid {
			continue
		}
		mi := entity.VisualScreenGroup{}
		mi.Id = list[j].Id
		mi.Pid = list[j].Pid
		mi.Path = list[j].Path
		mi.Name = list[j].Name
		mi.Sort = list[j].Sort
		mi.Status = list[j].Status
		ms := Digui(sglist, mi)
		min = append(min, ms)
	}
	menu.Children = min
	return menu
}
func DiguiDeptLable(sglist *[]entity.VisualScreenGroup, dept entity.ScreenGroupLabel) entity.ScreenGroupLabel {
	list := *sglist

	min := make([]entity.ScreenGroupLabel, 0)
	for j := 0; j < len(list); j++ {

		if dept.Id != list[j].Id {
			continue
		}
		sg := entity.ScreenGroupLabel{list[j].Id, list[j].Name, []entity.ScreenGroupLabel{}}
		ms := DiguiDeptLable(sglist, sg)
		min = append(min, ms)

	}
	dept.Children = min
	return dept
}
