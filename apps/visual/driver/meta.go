package driver

type DbMetadata interface {
	GetTables() []map[string]interface{}
	GetColumns(tableNames ...string) []map[string]interface{}
	GetPrimaryKey(tableName string) string
	GetTableInfos() []map[string]interface{}
	GetTableRecord(tableName string, pageNum, pageSize int) ([]string, []map[string]interface{}, error)
}
