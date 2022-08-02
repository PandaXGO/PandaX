package services

/**
 * @Description
 * @Author 熊猫
 * @Date 2022/7/14 17:49
 **/
import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/system/entity"
	"pandax/pkg/global"
)

type (
	SysTenantsModel interface {
		Insert(data entity.SysTenants) *entity.SysTenants
		FindOne(tenantId int64) *entity.SysTenants
		FindListPage(page, pageSize int, data entity.SysTenants) (*[]entity.SysTenants, int64)
		FindList(data entity.SysTenants) *[]entity.SysTenants
		Update(data entity.SysTenants) *entity.SysTenants
		Delete(tenantIds []int64)
	}

	SysTenantModelImpl struct {
		table string
	}
)

var SysTenantModelDao SysTenantsModel = &SysTenantModelImpl{
	table: `sys_tenants`,
}

func (m *SysTenantModelImpl) Insert(data entity.SysTenants) *entity.SysTenants {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加SysTenant失败")
	return &data
}

func (m *SysTenantModelImpl) FindOne(tenantId int64) *entity.SysTenants {
	resData := new(entity.SysTenants)
	err := global.Db.Table(m.table).Where("id = ?", tenantId).First(resData).Error
	biz.ErrIsNil(err, "查询SysTenant失败")
	return resData
}

func (m *SysTenantModelImpl) FindListPage(page, pageSize int, data entity.SysTenants) (*[]entity.SysTenants, int64) {
	list := make([]entity.SysTenants, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	if data.TenantName != "" {
		db = db.Where("tenant_name like ?", "%"+data.TenantName+"%")
	}
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询SysTenant分页列表失败")
	return &list, total
}

func (m *SysTenantModelImpl) FindList(data entity.SysTenants) *[]entity.SysTenants {
	list := make([]entity.SysTenants, 0)
	db := global.Db.Table(m.table)
	if data.TenantName != "" {
		db = db.Where("tenant_name like ?", "%"+data.TenantName+"%")
	}
	if data.Id != 0 {
		db = db.Where("id = ?", data.Id)
	}
	db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询SysTenant列表失败")
	return &list
}

func (m *SysTenantModelImpl) Update(data entity.SysTenants) *entity.SysTenants {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改SysTenant失败")
	return &data
}

func (m *SysTenantModelImpl) Delete(tenantIds []int64) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.SysTenants{}, "id in (?)", tenantIds).Error, "删除SysTenant失败")
}
