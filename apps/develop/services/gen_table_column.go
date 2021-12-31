package services

import (
	"errors"
	"pandax/apps/develop/entity"
	"pandax/base/biz"
	"pandax/base/config"
	"pandax/base/global"
)

/**
 * @Description
 * @Author Panda
 * @Date 2021/12/31 14:44
 **/

type (
	SysGenTableColumnModel interface {
		FindDbTablesColumnListPage(page, pageSize int, data entity.DevGenTableColumn) (*[]entity.DevGenTableColumn, int64)
		FindDbTableColumnList(tableName string) *[]entity.DevGenTableColumn

		Insert(data entity.DevGenTableColumn) *entity.DevGenTableColumn
		FindList(data entity.DevGenTableColumn) *[]entity.DevGenTableColumn
		Update(data entity.DevGenTableColumn) *entity.DevGenTableColumn
	}

	devTableColumnModelImpl struct {
		table              string
		ColumnTypeStr      []string //数据库字符串类型
		ColumnTypeTime     []string //数据库时间类型
		ColumnTypeNumber   []string //数据库数字类型
		ColumnNameNotEdit  []string //页面不需要编辑字段
		ColumnNameNotList  []string //页面不需要显示的列表字段
		ColumnNameNotQuery []string //页面不需要查询字段
	}
)

var SysSysConfigModelDao SysGenTableColumnModel = &devTableColumnModelImpl{
	table:              "dev_gen_table_columns",
	ColumnTypeStr:      []string{"char", "varchar", "narchar", "varchar2", "tinytext", "text", "mediumtext", "longtext"},
	ColumnTypeTime:     []string{"datetime", "time", "date", "timestamp"},
	ColumnTypeNumber:   []string{"tinyint", "smallint", "mediumint", "int", "number", "integer", "bigint", "float", "float", "double", "decimal"},
	ColumnNameNotEdit:  []string{"id", "created_by", "created_at", "updated_by", "updated_at", "deleted_at"},
	ColumnNameNotList:  []string{"id", "created_by", "updated_by", "created_at", "updated_at", "deleted_at"},
	ColumnNameNotQuery: []string{"id", "created_by", "updated_by", "created_at", "updated_at", "deleted_at", "remark"},
}

func (m *devTableColumnModelImpl) FindDbTablesColumnListPage(page, pageSize int, data entity.DevGenTableColumn) (*[]entity.DevGenTableColumn, int64) {
	list := make([]entity.DevGenTableColumn, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	if config.Conf.Server.DbType != "mysql" && config.Conf.Server.DbType == "postgresql" {
		biz.ErrIsNil(errors.New("只支持mysql和postgresql数据库"), "只支持mysql和postgresql数据库")
	}

	db := global.Db.Table("information_schema.COLUMNS")
	db = db.Where("table_schema= ? ", config.Conf.Gen.Dbname)

	if data.TableName != "" {
		db = db.Where("table_name = ?", data.TableName)
	}

	err := db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询生成代码列表信息失败")
	return &list, total
}

func (m *devTableColumnModelImpl) FindDbTableColumnList(tableName string) *[]entity.DevGenTableColumn {
	resData := make([]entity.DevGenTableColumn, 0)
	if config.Conf.Server.DbType != "mysql" && config.Conf.Server.DbType == "postgresql" {
		biz.ErrIsNil(errors.New("只支持mysql和postgresql数据库"), "只支持mysql和postgresql数据库")
	}
	db := global.Db.Table("information_schema.columns")
	db = db.Where("table_schema = ? ", config.Conf.Gen.Dbname)
	biz.IsTrue(tableName != "", "table name cannot be empty！")

	db = db.Where("table_name = ?", tableName)
	err := db.First(&resData).Error
	biz.ErrIsNil(err, err.Error())
	return &resData
}

func (m *devTableColumnModelImpl) Insert(dgt entity.DevGenTableColumn) *entity.DevGenTableColumn {
	err := global.Db.Table(m.table).Create(&dgt).Error
	biz.ErrIsNil(err, "新增生成代码字段表失败")
	return &dgt
}

func (m *devTableColumnModelImpl) FindList(data entity.DevGenTableColumn) *[]entity.DevGenTableColumn {
	list := make([]entity.DevGenTableColumn, 0)
	err := global.Db.Table(m.table).Where("table_id = ?", data.TableId).Find(&list).Error

	biz.ErrIsNil(err, "查询生成代码字段表信息失败")
	return &list
}

func (m *devTableColumnModelImpl) Update(data entity.DevGenTableColumn) *entity.DevGenTableColumn {
	err := global.Db.Table(m.table).Model(&data).Updates(&data).Error
	biz.ErrIsNil(err, "修改生成代码字段表失败")
	return &data
}
