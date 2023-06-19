package driver

import (
	"database/sql"
	"fmt"
	"github.com/XM-GO/PandaKit/biz"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	"pandax/apps/visual/entity"
)

func getTdDB(d *entity.VisualDataSource) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",
		d.Db.Username, d.Db.Password, "http", d.Db.Host, d.Db.Port, d.Db.Dbname)
	if d.Db.Config != "" {
		dsn = fmt.Sprintf("%s&%s", dsn, d.Db.Config)
	}
	return sql.Open("taosRestful", dsn)
}

// ---------------------------------- mysql元数据 -----------------------------------
const (
	// 表信息元数据
	TD_TABLE_MA = `SELECT table_name table_name,table_comment table_comment from information_schema.INS_TABLES WHERE db_name = '%s'`

	// 表信息
	TD_TABLE_INFO = `SELECT table_name table_name, table_comment table_comment, columns table_rows, create_time create_time 
	FROM information_schema.INS_TABLES WHERE db_name = '%s'`

	//列信息元数据
	TD_COLUMN_MA = `SELECT table_name table_name, col_name column_name, col_type column_type,col_nullable nullable from information_schema.INS_COLUMNS
	WHERE db_name = '%s' AND table_name in (%s) ORDER BY table_name`
)

type TDMetadata struct {
	di *DbInstance
}

// 获取表基础元信息, 如表名等
func (mm *TDMetadata) GetTables() []map[string]interface{} {
	res, err := mm.di.innerSelect(fmt.Sprintf(TD_TABLE_MA, mm.di.Info.Db.Dbname))
	biz.ErrIsNilAppendErr(err, "获取表基本信息失败: %s")
	return res
}

// 获取列元信息, 如列名等
func (mm *TDMetadata) GetColumns(tableNames ...string) []map[string]interface{} {
	tableName := ""
	for i := 0; i < len(tableNames); i++ {
		if i != 0 {
			tableName = tableName + ", "
		}
		tableName = tableName + "'" + tableNames[i] + "'"
	}
	result, err := mm.di.innerSelect(fmt.Sprintf(TD_COLUMN_MA, mm.di.Info.Db.Dbname, tableName))
	biz.ErrIsNilAppendErr(err, "获取数据库列信息失败: %s")
	return result
}

// 获取表信息，比GetTableMetedatas获取更详细的表信息
func (mm *TDMetadata) GetTableInfos() []map[string]interface{} {
	res, err := mm.di.innerSelect(fmt.Sprintf(TD_TABLE_INFO, mm.di.Info.Db.Dbname))
	biz.ErrIsNilAppendErr(err, "获取表信息失败: %s")
	return res
}
