package services

import (
	"errors"
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/develop/entity"
	"pandax/pkg/global"
)

/**
 * @Description
 * @Author Panda
 * @Date 2021/12/31 14:44
 **/

type (
	SysGenTableColumnModel interface {
		FindDbTablesColumnListPage(page, pageSize int, data entity.DBColumns) (*[]entity.DBColumns, int64)
		FindDbTableColumnList(tableName string) *[]entity.DBColumns

		Insert(data entity.DevGenTableColumn) *entity.DevGenTableColumn
		FindList(data entity.DevGenTableColumn, exclude bool) *[]entity.DevGenTableColumn
		Update(data entity.DevGenTableColumn) *entity.DevGenTableColumn
		Delete(tableId []int64)
	}

	devTableColumnModelImpl struct {
		table string
	}
)

var DevTableColumnModelDao SysGenTableColumnModel = &devTableColumnModelImpl{
	table: "dev_gen_table_columns",
}

func (m *devTableColumnModelImpl) FindDbTablesColumnListPage(page, pageSize int, data entity.DBColumns) (*[]entity.DBColumns, int64) {
	list := make([]entity.DBColumns, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	if global.Conf.Server.DbType != "mysql" && global.Conf.Server.DbType != "postgresql" {
		biz.ErrIsNil(errors.New("只支持mysql和postgresql数据库"), "只支持mysql和postgresql数据库")
	}
	db := global.Db.Table("information_schema.COLUMNS")
	if global.Conf.Server.DbType == "mysql" {
		db = db.Where("table_schema= ? ", global.Conf.Gen.Dbname)
	}
	if global.Conf.Server.DbType == "postgresql" {
		db = db.Where("table_schema = ? ", "public")
	}

	if data.TableName != "" {
		db = db.Where("table_name = ?", data.TableName)
	}

	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询生成代码列表信息失败")
	return &list, total
}

func (m *devTableColumnModelImpl) FindDbTableColumnList(tableName string) *[]entity.DBColumns {
	resData := make([]entity.DBColumns, 0)
	if global.Conf.Server.DbType != "mysql" && global.Conf.Server.DbType != "postgresql" {
		biz.ErrIsNil(errors.New("只支持mysql和postgresql数据库"), "只支持mysql和postgresql数据库")
	}
	db := global.Db.Table("information_schema.columns")
	if global.Conf.Server.DbType == "mysql" {
		db = db.Where("table_schema= ? ", global.Conf.Gen.Dbname)
	}
	if global.Conf.Server.DbType == "postgresql" {
		db = db.Where("table_schema = ? ", "public")
	}
	biz.IsTrue(tableName != "", "table name cannot be empty！")

	db = db.Where("table_name = ?", tableName)
	err := db.Find(&resData).Error
	biz.ErrIsNil(err, "查询表字段失败")
	return &resData
}

func (m *devTableColumnModelImpl) Insert(dgt entity.DevGenTableColumn) *entity.DevGenTableColumn {
	err := global.Db.Table(m.table).Create(&dgt).Error
	biz.ErrIsNil(err, "新增生成代码字段表失败")
	return &dgt
}

func (m *devTableColumnModelImpl) FindList(data entity.DevGenTableColumn, exclude bool) *[]entity.DevGenTableColumn {
	list := make([]entity.DevGenTableColumn, 0)
	db := global.Db.Table(m.table).Where("table_id = ?", data.TableId)
	if exclude {
		notIn := make([]string, 6)
		notIn = append(notIn, "id")
		notIn = append(notIn, "create_by")
		notIn = append(notIn, "update_by")
		notIn = append(notIn, "create_time")
		notIn = append(notIn, "update_time")
		notIn = append(notIn, "delete_time")
		db = db.Where("column_name not in(?)", notIn)
	}
	err := db.Find(&list).Error
	biz.ErrIsNil(err, "查询生成代码字段表信息失败")
	return &list
}

func (m *devTableColumnModelImpl) Update(data entity.DevGenTableColumn) *entity.DevGenTableColumn {
	err := global.Db.Table(m.table).Model(&data).Updates(&data).Error
	biz.ErrIsNil(err, "修改生成代码字段表失败")
	return &data
}

func (m *devTableColumnModelImpl) Delete(tableId []int64) {
	err := global.Db.Table(m.table).Delete(&entity.DevGenTableColumn{}, "table_id in (?)", tableId).Error
	biz.ErrIsNil(err, "删除生成代码字段表失败")
	return
}
