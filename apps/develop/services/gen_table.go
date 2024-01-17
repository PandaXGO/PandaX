package services

import (
	"errors"
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/utils"
	"pandax/apps/develop/entity"
	"pandax/pkg/global"
	"pandax/pkg/global/model"
)

/**
 * @Description
 * @Author Panda
 * @Date 2021/12/31 8:58
 **/

type (
	SysGenTableModel interface {
		FindDbTablesListPage(page, pageSize int, data entity.DBTables) (*[]entity.DBTables, int64, error)
		FindDbTableOne(tableName string) (*entity.DBTables, error)

		// 导入表数据
		Insert(data entity.DevGenTable) error
		FindOne(data entity.DevGenTable, exclude bool) (*entity.DevGenTable, error)
		FindTree(data entity.DevGenTable) (*[]entity.DevGenTable, error)
		FindListPage(page, pageSize int, data entity.DevGenTable) (*[]entity.DevGenTable, int64, error)
		Update(data entity.DevGenTable) (*entity.DevGenTable, error)
		Delete(tableIds []int64) error
	}

	devGenTableModelImpl struct {
		table string
	}
)

var DevGenTableModelDao SysGenTableModel = &devGenTableModelImpl{
	table: "dev_gen_tables",
}

func (m *devGenTableModelImpl) FindDbTablesListPage(page, pageSize int, data entity.DBTables) (*[]entity.DBTables, int64, error) {
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
	//db = db.Where("table_name NOT LIKE 'dev_%'")
	if data.TableName != "" {
		db = db.Where("table_name like ?", "%"+data.TableName+"%")
	}
	if global.Conf.Server.DbType == "mysql" {
		err := db.Limit(pageSize).Offset(offset).Find(&list).Count(&total).Error
		return &list, total, err
	} else {
		err := db.Limit(pageSize).Offset(offset).Find(&pgdata).Count(&total).Error
		for _, pd := range pgdata {
			list = append(list, entity.DBTables{TableName: utils.B2S(pd["table_name"].([]uint8))})
		}
		return &list, total, err
	}
}

func (m *devGenTableModelImpl) FindDbTableOne(tableName string) (*entity.DBTables, error) {
	resData := new(entity.DBTables)
	if global.Conf.Server.DbType != "mysql" && global.Conf.Server.DbType != "postgresql" {
		return nil, errors.New("只支持mysql和postgresql数据库")
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
	return resData, err
}

func (m *devGenTableModelImpl) Insert(dgt entity.DevGenTable) error {
	err := global.Db.Table(m.table).Create(&dgt).Error
	if err != nil {
		return err
	}
	for i := 0; i < len(dgt.Columns); i++ {
		dgt.Columns[i].TableId = dgt.TableId
		columns := dgt.Columns[i]
		columns.OrgId = dgt.OrgId
		columns.Owner = dgt.Owner
		DevTableColumnModelDao.Insert(columns)
	}
	return nil
}

func (m *devGenTableModelImpl) FindOne(data entity.DevGenTable, exclude bool) (*entity.DevGenTable, error) {
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
	if err != nil {
		return resData, err
	}
	list, err := DevTableColumnModelDao.FindList(entity.DevGenTableColumn{TableId: resData.TableId}, exclude)
	resData.Columns = *list
	return resData, err
}

func (m *devGenTableModelImpl) FindTree(data entity.DevGenTable) (*[]entity.DevGenTable, error) {
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
	// 组织数据访问权限
	if err := model.OrgAuthSet(db, data.RoleId, data.Owner); err != nil {
		return nil, err
	}
	if err := db.Find(&resData).Error; err != nil {
		return nil, err
	}
	for i := 0; i < len(resData); i++ {
		var col entity.DevGenTableColumn
		col.TableId = resData[i].TableId
		col.RoleId = data.RoleId
		columns, _ := DevTableColumnModelDao.FindList(col, false)
		resData[i].Columns = *columns
	}
	return &resData, nil
}

func (m *devGenTableModelImpl) FindListPage(page, pageSize int, data entity.DevGenTable) (*[]entity.DevGenTable, int64, error) {
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
	// 组织数据访问权限
	if err := model.OrgAuthSet(db, data.RoleId, data.Owner); err != nil {
		return &list, total, err
	}
	db.Where("delete_time IS NULL")
	if err := db.Count(&total).Error; err != nil {
		return &list, total, err
	}
	if err := db.Limit(pageSize).Offset(offset).Find(&list).Error; err != nil {
		return &list, total, err
	}
	return &list, total, nil
}

func (m *devGenTableModelImpl) Update(data entity.DevGenTable) (*entity.DevGenTable, error) {
	err := global.Db.Table(m.table).Model(&data).Updates(&data).Error
	if err != nil {
		return nil, err
	}

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
		if err != nil {
			return nil, err
		}
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
	return &data, nil
}

func (e *devGenTableModelImpl) DeleteTables(tableId int64) (bool, error) {
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
		return success, err
	}
	if err = tx.Table("sys_columns").Delete(entity.DevGenTableColumn{}, "table_id = ?", tableId).Error; err != nil {
		return success, err
	}
	success = true
	return success, nil
}

func (m *devGenTableModelImpl) Delete(configIds []int64) error {
	err := global.Db.Table(m.table).Delete(&entity.DevGenTable{}, "table_id in (?)", configIds).Error
	if err != nil {
		return errors.New("删除生成代码信息失败")
	}
	DevTableColumnModelDao.Delete(configIds)
	return nil
}
