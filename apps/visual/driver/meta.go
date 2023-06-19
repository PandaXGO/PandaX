package driver

type DbMetadata interface {
	GetTables() []map[string]interface{}
	GetColumns(tableNames ...string) []map[string]interface{}
	GetTableInfos() []map[string]interface{}
}
