package driver

import (
	"database/sql"
	"fmt"
	"github.com/XM-GO/PandaKit/biz"
	_ "github.com/go-sql-driver/mysql"
	"pandax/apps/visual/entity"
)

func getMysqlDB(d *entity.VisualDataSource) (*sql.DB, error) {
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
	// mysql 表信息元数据
	MYSQL_TABLE_MA = `SELECT table_name tableName, table_comment tableComment from information_schema.tables WHERE table_schema = (SELECT database())`

	// mysql 表信息
	MYSQL_TABLE_INFO = `SELECT table_name tableName, table_comment tableComment, table_rows tableRows,
	data_length dataLength, index_length indexLength, create_time createTime 
	FROM information_schema.tables 
    WHERE table_schema = (SELECT database())`

	// mysql 列信息元数据
	MYSQL_COLUMN_MA = `SELECT table_name tableName, column_name columnName, column_type columnType,column_comment columnComment, 
       column_key columnKey, is_nullable nullable from information_schema.columns
	WHERE table_schema = (SELECT database()) AND table_name in (%s) ORDER BY tableName, ordinal_position`
)

type MysqlMetadata struct {
	di *DbInstance
}

// 获取表基础元信息, 如表名等
func (mm *MysqlMetadata) GetTables() []map[string]interface{} {
	res, err := mm.di.innerSelect(MYSQL_TABLE_MA)
	biz.ErrIsNilAppendErr(err, "获取表基本信息失败: %s")
	return res
}

// 获取列元信息, 如列名等
func (mm *MysqlMetadata) GetColumns(tableNames ...string) []map[string]interface{} {
	tableName := ""
	for i := 0; i < len(tableNames); i++ {
		if i != 0 {
			tableName = tableName + ", "
		}
		tableName = tableName + "'" + tableNames[i] + "'"
	}
	result, err := mm.di.innerSelect(fmt.Sprintf(MYSQL_COLUMN_MA, tableName))
	biz.ErrIsNilAppendErr(err, "获取数据库列信息失败: %s")
	return result
}

// 获取表信息，比GetTableMetedatas获取更详细的表信息
func (mm *MysqlMetadata) GetTableInfos() []map[string]interface{} {
	res, err := mm.di.innerSelect(MYSQL_TABLE_INFO)
	biz.ErrIsNilAppendErr(err, "获取表信息失败: %s")
	return res
}
