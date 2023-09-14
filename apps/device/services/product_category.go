package services

import (
	"errors"
	"github.com/PandaXGO/PandaKit/biz"
	"pandax/apps/device/entity"
	"pandax/pkg/global"
)

type (
	ProductCategoryModel interface {
		Insert(data entity.ProductCategory) *entity.ProductCategory
		FindOne(id string) *entity.ProductCategory
		FindListPage(page, pageSize int, data entity.ProductCategory) (*[]entity.ProductCategory, int64)
		FindList(data entity.ProductCategory) *[]entity.ProductCategory
		Update(data entity.ProductCategory) *entity.ProductCategory
		Delete(ids []string)
		SelectProductCategory(data entity.ProductCategory) []entity.ProductCategory
		SelectProductCategoryLabel(data entity.ProductCategory) []entity.ProductCategoryLabel
	}

	productCategoryModelImpl struct {
		table string
	}
)

var ProductCategoryModelDao ProductCategoryModel = &productCategoryModelImpl{
	table: `product_categories`,
}

func (m *productCategoryModelImpl) Insert(data entity.ProductCategory) *entity.ProductCategory {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加产品分组失败")

	path := "/" + data.Id
	if data.Pid != "0" {
		vsg := m.FindOne(data.Pid)
		path = vsg.Path + path
	} else {
		path = "/0" + path
	}
	data.Path = path
	biz.ErrIsNil(global.Db.Table(m.table).Model(&data).Updates(&data).Error, "修改产品分组信息失败")
	return &data
}

func (m *productCategoryModelImpl) FindOne(id string) *entity.ProductCategory {
	resData := new(entity.ProductCategory)
	db := global.Db.Table(m.table).Where("id = ?", id)
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询产品分组失败")
	return resData
}

func (m *productCategoryModelImpl) FindListPage(page, pageSize int, data entity.ProductCategory) (*[]entity.ProductCategory, int64) {
	list := make([]entity.ProductCategory, 0)
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
	biz.ErrIsNil(err, "查询产品分组分页列表失败")
	return &list, total
}

func (m *productCategoryModelImpl) FindList(data entity.ProductCategory) *[]entity.ProductCategory {
	list := make([]entity.ProductCategory, 0)
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
	biz.ErrIsNil(db.Order("sort").Find(&list).Error, "查询产品分组列表失败")
	return &list
}

func (m *productCategoryModelImpl) Update(data entity.ProductCategory) *entity.ProductCategory {
	one := m.FindOne(data.Id)

	path := "/" + data.Id
	if data.Pid != "0" {
		vsg := m.FindOne(data.Pid)
		path = vsg.Path + path
	} else {
		path = "/0" + path
	}
	data.Path = path

	if data.Path != "" && data.Path != one.Path {
		biz.ErrIsNil(errors.New("上级分组不允许修改！"), "上级分组不允许修改")
	}

	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改产品分组失败")
	return &data
}

func (m *productCategoryModelImpl) Delete(s []string) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.ProductCategory{}, "id in (?)", s).Error, "删除产品分组失败")
}

func (m *productCategoryModelImpl) SelectProductCategory(data entity.ProductCategory) []entity.ProductCategory {
	list := m.FindList(data)
	sd := make([]entity.ProductCategory, 0)
	li := *list
	for i := 0; i < len(li); i++ {
		if li[i].Pid != "0" {
			continue
		}
		info := DiGuiProductCategory(list, li[i])
		sd = append(sd, info)
	}
	return sd
}

func (m *productCategoryModelImpl) SelectProductCategoryLabel(data entity.ProductCategory) []entity.ProductCategoryLabel {
	list := m.FindList(data)

	dl := make([]entity.ProductCategoryLabel, 0)
	organizationl := *list
	for i := 0; i < len(organizationl); i++ {
		if organizationl[i].Pid != "0" {
			continue
		}
		e := entity.ProductCategoryLabel{}
		e.Id = organizationl[i].Id
		e.Name = organizationl[i].Name
		organizationsInfo := DiGuiProductCategoryLabel(list, e)

		dl = append(dl, organizationsInfo)
	}
	return dl
}

func DiGuiProductCategory(sglist *[]entity.ProductCategory, menu entity.ProductCategory) entity.ProductCategory {
	list := *sglist

	min := make([]entity.ProductCategory, 0)
	for j := 0; j < len(list); j++ {

		if menu.Id != list[j].Pid {
			continue
		}
		mi := entity.ProductCategory{}
		mi.Id = list[j].Id
		mi.Pid = list[j].Pid
		mi.Path = list[j].Path
		mi.Name = list[j].Name
		mi.Sort = list[j].Sort
		mi.Status = list[j].Status
		mi.Description = list[j].Description
		ms := DiGuiProductCategory(sglist, mi)
		min = append(min, ms)
	}
	menu.Children = min
	return menu
}
func DiGuiProductCategoryLabel(sglist *[]entity.ProductCategory, organization entity.ProductCategoryLabel) entity.ProductCategoryLabel {
	list := *sglist

	min := make([]entity.ProductCategoryLabel, 0)
	for j := 0; j < len(list); j++ {
		if organization.Id != list[j].Pid {
			continue
		}
		sg := entity.ProductCategoryLabel{}
		sg.Id = list[j].Id
		sg.Name = list[j].Name
		ms := DiGuiProductCategoryLabel(sglist, sg)
		min = append(min, ms)

	}
	organization.Children = min
	return organization
}
