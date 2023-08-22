package services

import (
	"errors"
	"github.com/XM-GO/PandaKit/biz"
	"github.com/XM-GO/PandaKit/utils"
	"pandax/apps/develop/entity"
	"pandax/pkg/global"
)

/**
 * @Description
 * @Author Panda
 * @Date 2021/12/31 8:58
 **/

type (
	SysGenTableModel interface {
		FindDbTablesListPage(page, pageSize int, data entity.DBTables) (*[]entity.DBTables, int64)
		FindDbTableOne(tableName string) *entity.DBTables

		// 导入表数据
		Insert(data entity.DevGenTable)
		FindOne(data entity.DevGenTable, exclude bool) *entity.DevGenTable
		FindTree(data entity.DevGenTable) *[]entity.DevGenTable
		FindListPage(page, pageSize int, data entity.DevGenTable) (*[]entity.DevGenTable, int64)
		Update(data entity.DevGenTable) *entity.DevGenTable
		Delete(tableIds []int64)
	}

	devGenTableModelImpl struct {
		table string
	}
)

var DevGenTableModelDao SysGenTableModel = &devGenTableModelImpl{
	table: "dev_gen_tables",
}

func (m *devGenTableModelImpl) FindDbTablesListPage(page, pageSize int, data entity.DBTables) (*[]entity.DBTables, int64) {
	list := make([]entity.DBTables, 0)
	pgdata := make([]map[string]any, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	if global.Conf.Server.DbType != "mysql" && global.Conf.Server.DbType != "postgresql" {
		biz.ErrIsNil(errors.New("只支持mysql和postgresql数据库"), "只支持mysql和postgresql数据库")
	}

	db := global.Db.Table("information_schema.tables")
	if global.Conf.Server.DbType == "mysql" {
		db = db.Where("table_schema= ? ", global.Conf.Gen.Dbname)
	}
	if global.Conf.Server.DbType == "postgresql" {
		db = db.Where("table_schema = ? ", "public")
	}
	db = db.Where("table_name NOT LIKE 'dev_%'")
	if data.TableName != "" {
		db = db.Where("table_name like ?", "%"+data.TableName+"%")
	}
	if global.Conf.Server.DbType == "mysql" {
		err := db.Limit(pageSize).Offset(offset).Find(&list).Offset(-1).Limit(-1).Count(&total).Error
		biz.ErrIsNil(err, "查询配置分页列表信息失败")
		return &list, total
	} else {
		err := db.Limit(pageSize).Offset(offset).Find(&pgdata).Offset(-1).Limit(-1).Count(&total).Error
		biz.ErrIsNil(err, "查询配置分页列表信息失败")
		for _, pd := range pgdata {
			list = append(list, entity.DBTables{TableName: utils.B2S(pd["table_name"].([]uint8))})
		}
		return &list, total
	}
}

func (m *devGenTableModelImpl) FindDbTableOne(tableName string) *entity.DBTables {
	resData := new(entity.DBTables)
	if global.Conf.Server.DbType != "mysql" && global.Conf.Server.DbType != "postgresql" {
		biz.ErrIsNil(errors.New("只支持mysql和postgresql数据库"), "只支持mysql和postgresql数据库")
	}
	db := global.Db.Table("information_schema.tables")
	if global.Conf.Server.DbType == "mysql" {
		db = db.Where("table_schema= ? ", global.Conf.Gen.Dbname)
	}
	if global.Conf.Server.DbType == "postgresql" {
		db = db.Where("table_schema= ? ", "public")
	}
	db = db.Where("table_name = ?", tableName)
	err := db.First(&resData).Error
	biz.ErrIsNil(err, err.Error())
	return resData
}

func (m *devGenTableModelImpl) Insert(dgt entity.DevGenTable) {
	err := global.Db.Table(m.table).Create(&dgt).Error
	biz.ErrIsNil(err, "新增生成代码表失败")
	for i := 0; i < len(dgt.Columns); i++ {
		dgt.Columns[i].TableId = dgt.TableId
		DevTableColumnModelDao.Insert(dgt.Columns[i])
	}
}

func (m *devGenTableModelImpl) FindOne(data entity.DevGenTable, exclude bool) *entity.DevGenTable {
	resData := new(entity.DevGenTable)
	db := global.Db.Table(m.table)
	if data.TableName != "" {
		db = db.Where("table_name = ?", data.TableName)
	}
	if data.TableId != 0 {
		db = db.Where("table_id = ?", data.TableId)
	}
	if data.TableComment != "" {
		db = db.Where("table_comment = ?", data.TableComment)
	}
	err := db.First(resData).Error
	biz.ErrIsNil(err, "查询配置信息失败")
	list := DevTableColumnModelDao.FindList(entity.DevGenTableColumn{TableId: resData.TableId}, exclude)
	resData.Columns = *list
	return resData
}

func (m *devGenTableModelImpl) FindTree(data entity.DevGenTable) *[]entity.DevGenTable {
	resData := make([]entity.DevGenTable, 0)
	db := global.Db.Table(m.table)

	if data.TableName != "" {
		db = db.Where("table_name = ?", data.TableName)
	}
	if data.TableId != 0 {
		db = db.Where("table_id = ?", data.TableId)
	}
	if data.TableComment != "" {
		db = db.Where("table_comment = ?", data.TableComment)
	}
	err := db.Find(&resData).Error
	biz.ErrIsNil(err, "获取TableTree失败")
	for i := 0; i < len(resData); i++ {
		var col entity.DevGenTableColumn
		col.TableId = resData[i].TableId
		columns := DevTableColumnModelDao.FindList(col, false)
		resData[i].Columns = *columns
	}
	return &resData
}

func (m *devGenTableModelImpl) FindListPage(page, pageSize int, data entity.DevGenTable) (*[]entity.DevGenTable, int64) {
	list := make([]entity.DevGenTable, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)

	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.TableName != "" {
		db = db.Where("table_name = ?", data.TableName)
	}
	if data.TableComment != "" {
		db = db.Where("table_comment = ?", data.TableComment)
	}
	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询生成代码列表信息失败")
	return &list, total
}

func (m *devGenTableModelImpl) Update(data entity.DevGenTable) *entity.DevGenTable {
	err := global.Db.Table(m.table).Model(&data).Updates(&data).Error
	biz.ErrIsNil(err, "修改生成代码信息失败")

	tableNames := make([]string, 0)
	for i := range data.Columns {
		if data.Columns[i].LinkTableName != "" {
			tableNames = append(tableNames, data.Columns[i].LinkTableName)
		}
	}

	tables := make([]entity.DevGenTable, 0)
	tableMap := make(map[string]*entity.DevGenTable)
	if len(tableNames) > 0 {
		err = global.Db.Table(m.table).Where("table_name in (?)", tableNames).Find(&tables).Error
		biz.ErrIsNil(err, "关联表不存在")
		for i := range tables {
			tableMap[tables[i].TableName] = &tables[i]
		}
	}

	for i := 0; i < len(data.Columns); i++ {
		if data.Columns[i].LinkTableName != "" {
			t, ok := tableMap[data.Columns[i].LinkTableName]
			if ok {
				data.Columns[i].LinkTableClass = t.ClassName
				data.Columns[i].LinkTablePackage = t.BusinessName
				data.Columns[i].LinkLabelId = t.PkColumn
				data.Columns[i].LinkLabelName = t.PkGoField
			}
		}
		DevTableColumnModelDao.Update(data.Columns[i])
	}
	return &data
}

func (e *devGenTableModelImpl) DeleteTables(tableId int64) bool {
	var err error
	success := false
	tx := global.Db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err = tx.Table("sys_tables").Delete(entity.DevGenTable{}, "table_id = ?", tableId).Error; err != nil {
		return success
	}
	if err = tx.Table("sys_columns").Delete(entity.DevGenTableColumn{}, "table_id = ?", tableId).Error; err != nil {
		return success
	}
	success = true
	return success
}

func (m *devGenTableModelImpl) Delete(configIds []int64) {
	err := global.Db.Table(m.table).Delete(&entity.DevGenTable{}, "table_id in (?)", configIds).Error
	biz.ErrIsNil(err, "删除生成代码信息失败")
	DevTableColumnModelDao.Delete(configIds)
	return
}
