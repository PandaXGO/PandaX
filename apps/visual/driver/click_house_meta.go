package driver

import (
	"database/sql"
	"fmt"
	"github.com/XM-GO/PandaKit/biz"
	_ "github.com/go-sql-driver/mysql"
	"pandax/apps/visual/entity"
)

func getClickHouseDB(d *entity.VisualDataSource) (*sql.DB, error) {
	// 设置dataSourceName  -> 更多参数参考：https://github.com/go-sql-driver/mysql#dsn-data-source-name
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?timeout=8s",
		d.Db.Username, d.Db.Password, "tcp", d.Db.Host, d.Db.Port, d.Db.Dbname)
	if d.Db.Config != "" {
		dsn = fmt.Sprintf("%s&%s", dsn, d.Db.Config)
	}
	return sql.Open("mysql", dsn)
}

// ---------------------------------- mysql元数据 -----------------------------------
const (
	// 表信息元数据
	ClickHouse_TABLE_MA = `SELECT name tableName, comment tableComment from system.tables WHERE database = '%s'`

	// 表信息
	ClickHouse_TABLE_INFO = `SELECT name tableName, comment tableComment, total_rows tableRows
	FROM system.tables WHERE database = '%s'`

	// 列信息元数据
	ClickHouse_COLUMN_MA = `SELECT table tableName, name columnName, type columnType, comment columnComment from system.columns
	WHERE database='%s' AND table in (%s) ORDER BY table`
)

type ClickHouseMetadata struct {
	di *DbInstance
}

// 获取表基础元信息, 如表名等
func (mm *ClickHouseMetadata) GetTables() []map[string]interface{} {
	res, err := mm.di.innerSelect(fmt.Sprintf(ClickHouse_TABLE_MA, mm.di.Info.Db.Dbname))
	biz.ErrIsNilAppendErr(err, "获取表基本信息失败: %s")
	return res
}

// 获取列元信息, 如列名等
func (mm *ClickHouseMetadata) GetColumns(tableNames ...string) []map[string]interface{} {
	tableName := ""
	for i := 0; i < len(tableNames); i++ {
		if i != 0 {
			tableName = tableName + ", "
		}
		tableName = tableName + "'" + tableNames[i] + "'"
	}
	result, err := mm.di.innerSelect(fmt.Sprintf(ClickHouse_COLUMN_MA, mm.di.Info.Db.Dbname, tableName))
	biz.ErrIsNilAppendErr(err, "获取数据库列信息失败: %s")
	return result
}

// 获取表信息，比GetTableMetedatas获取更详细的表信息
func (mm *ClickHouseMetadata) GetTableInfos() []map[string]interface{} {
	res, err := mm.di.innerSelect(fmt.Sprintf(ClickHouse_TABLE_INFO, mm.di.Info.Db.Dbname))
	biz.ErrIsNilAppendErr(err, "获取表信息失败: %s")
	return res
}
