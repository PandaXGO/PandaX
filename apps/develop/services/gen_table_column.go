package services

import (
	"errors"
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
		FindDbTablesColumnListPage(page, pageSize int, data entity.DBColumns) (*[]entity.DBColumns, int64, error)
		FindDbTableColumnList(tableName string) ([]entity.DBColumns, error)

		Insert(data entity.DevGenTableColumn) (*entity.DevGenTableColumn, error)
		FindList(data entity.DevGenTableColumn, exclude bool) (*[]entity.DevGenTableColumn, error)
		Update(data entity.DevGenTableColumn) (*entity.DevGenTableColumn, error)
		Delete(tableId []int64) error
	}

	devTableColumnModelImpl struct {
		table string
	}
)

var DevTableColumnModelDao SysGenTableColumnModel = &devTableColumnModelImpl{
	table: "dev_gen_table_columns",
}

func (m *devTableColumnModelImpl) FindDbTablesColumnListPage(page, pageSize int, data entity.DBColumns) (*[]entity.DBColumns, int64, error) {
	list := make([]entity.DBColumns, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	if global.Conf.Server.DbType != "mysql" && global.Conf.Server.DbType != "postgresql" {
		return nil, 0, errors.New("只支持mysql和postgresql数据库")
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
	return &list, total, err
}

func (m *devTableColumnModelImpl) FindDbTableColumnList(tableName string) ([]entity.DBColumns, error) {

	if global.Conf.Server.DbType != "mysql" && global.Conf.Server.DbType != "postgresql" {
		return nil, errors.New("只支持mysql和postgresql数据库")
	}
	db := global.Db.Table("information_schema.columns")
	if global.Conf.Server.DbType == "mysql" {
		db = db.Where("table_schema= ? ", global.Conf.Gen.Dbname)
	}
	if global.Conf.Server.DbType == "postgresql" {
		db = db.Where("table_schema = ? ", "public")
	}
	if tableName == "" {
		return nil, errors.New("table name cannot be empty！")
	}

	db = db.Where("table_name = ?", tableName)
	resData := make([]entity.DBColumns, 0)
	if global.Conf.Server.DbType == "mysql" {
		err := db.Find(&resData).Error
		return resData, err
	}
	if global.Conf.Server.DbType == "postgresql" {
		pr, err := getPgPR(tableName)
		if err != nil {
			return resData, errors.New("查询PG表主键字段失败")
		}
		resDataP := make([]entity.DBColumnsP, 0)
		err = db.Find(&resDataP).Error
		if err != nil {
			return resData, errors.New("查询表字段失败")
		}
		for _, data := range resDataP {
			dbc := entity.DBColumns{
				TableSchema:            data.TableSchema,
				TableName:              data.TableName,
				ColumnName:             data.ColumnName,
				ColumnDefault:          data.ColumnDefault,
				IsNullable:             data.IsNullable,
				DataType:               data.DataType,
				CharacterMaximumLength: data.CharacterMaximumLength,
				CharacterSetName:       data.CharacterSetName,
				ColumnType:             data.ColumnType,
				ColumnKey:              data.ColumnKey,
				Extra:                  data.Extra,
				ColumnComment:          data.ColumnComment,
			}
			// 设置为主键
			if pr == data.ColumnName {
				dbc.ColumnKey = "PRIMARY KEY"
			}
			resData = append(resData, dbc)
		}

		return resData, nil
	}
	return resData, nil
}

func getPgPR(tableName string) (string, error) {
	sql := `SELECT
    kcu.column_name
FROM
    information_schema.table_constraints AS tc
    JOIN information_schema.key_column_usage AS kcu ON tc.constraint_name = kcu.constraint_name
WHERE
    tc.constraint_type = 'PRIMARY KEY'
    AND tc.table_schema = 'public'
    AND tc.table_name = ?;`
	var pkname string
	err := global.Db.Raw(sql, tableName).Scan(&pkname).Error
	return pkname, err
}

func (m *devTableColumnModelImpl) Insert(dgt entity.DevGenTableColumn) (*entity.DevGenTableColumn, error) {
	err := global.Db.Table(m.table).Create(&dgt).Error
	if err != nil {
		global.Log.Error(err)
	}
	return &dgt, err
}

func (m *devTableColumnModelImpl) FindList(data entity.DevGenTableColumn, exclude bool) (*[]entity.DevGenTableColumn, error) {
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
	return &list, err
}

func (m *devTableColumnModelImpl) Update(data entity.DevGenTableColumn) (*entity.DevGenTableColumn, error) {
	err := global.Db.Table(m.table).Model(&data).Updates(&data).Error
	return &data, err
}

func (m *devTableColumnModelImpl) Delete(tableId []int64) error {
	err := global.Db.Table(m.table).Delete(&entity.DevGenTableColumn{}, "table_id in (?)", tableId).Error
	return err
}
