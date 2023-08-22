package tdengine

import (
	"database/sql"
	"fmt"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	"time"
)

const (
	TIME_TYPE_PROP  = "telemetry"
	TIME_TYPE_ATRE  = "attributes"
	TIME_TYPE_LOGS  = "logs"
	TIME_TYPE_ALARM = "alarm"
	TIME_TYPE_EVENT = "event"
)

type TdEngine struct {
	db     *sql.DB
	dbName string
}

func NewTdengine(username, password, host, db string) (*TdEngine, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s",
		username, password, "http", host, db)
	open, err := sql.Open("taosRestful", dsn)
	return &TdEngine{
		db:     open,
		dbName: db,
	}, err

}

// GetTdEngineAllDb 获取所有数据库
func (s *TdEngine) GetTdEngineAllDb() (data []string, err error) {
	rows, err := s.db.Query("show databases;")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		data = append(data, name)
	}
	return
}

// GetListTableByDatabases 获取指定数据库下所有的表列表
func (s *TdEngine) GetListTableByDatabases() (data []*TDEngineTablesList, err error) {
	rows, err := s.db.Query("SELECT table_name AS tableName, db_name AS dbName, create_time AS createTime, stable_name AS stableName FROM information_schema.ins_tables WHERE db_name = '" + s.dbName + "'")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tableName, db, stableName string
		var createTime *time.Time
		err = rows.Scan(&tableName, &db, &createTime, &stableName)
		if err != nil {
			return
		}
		var tDEngineTablesList = new(TDEngineTablesList)
		tDEngineTablesList.TableName = tableName
		tDEngineTablesList.DbName = db
		tDEngineTablesList.StableName = stableName
		tDEngineTablesList.CreateTime = createTime
		data = append(data, tDEngineTablesList)
	}
	return
}

// GetTdEngineTableInfoByTable 获取指定数据表结构信息
func (s *TdEngine) GetTdEngineTableInfoByTable(tableName string) (data []*TDEngineTableInfo, err error) {
	rows, err := s.db.Query("DESCRIBE " + s.dbName + "." + tableName + ";")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tDEngineTableInfo = new(TDEngineTableInfo)
		err = rows.Scan(&tDEngineTableInfo.Field, &tDEngineTableInfo.Type, &tDEngineTableInfo.Length, &tDEngineTableInfo.Note)
		if err != nil {
			return
		}
		data = append(data, tDEngineTableInfo)
	}
	return
}

// GetTdEngineTableDataByTable 获取指定数据表数据信息
func (s *TdEngine) GetTdEngineTableDataByTable(tableName string) (data *TableDataInfo, err error) {
	data = new(TableDataInfo)

	rows, err := s.db.Query("SELECT * FROM " + tableName)
	if err != nil {
		return
	}
	defer rows.Close()

	//获取查询结果字段
	columns, _ := rows.Columns()
	//字段数组
	var filed []string
	//封装scanArg
	scanArgs := make([]any, len(columns))
	for i := range columns {
		filed = append(filed, columns[i])
		scanArgs[i] = &columns[i]
	}
	data.Filed = append(data.Filed, filed...)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return
		}
		//封装返回结果
		var resultMap = make(map[string]interface{})
		for i := range columns {
			resultMap[filed[i]] = columns[i]
		}
		data.Info = append(data.Info, resultMap)
	}

	return
}

// GetOne 超级表查询，单条数据
func (s *TdEngine) GetOne(sql string, args ...any) (rs map[string]interface{}, err error) {
	rows, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	values := make([]any, len(columns))
	rs = make(map[string]interface{}, len(columns))
	for i := range values {
		values[i] = new(any)
	}

	for rows.Next() {
		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		for i, c := range columns {
			//rs[c] = s.Time(values[i])
			rs[c] = values[i]
		}
		rows.Close()
	}

	return
}

// GetAll 超级表查询，多条数据
func (s *TdEngine) GetAll(sql string, args ...any) (rs []map[string]interface{}, err error) {

	rows, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, _ := rows.Columns()

	for rows.Next() {
		values := make([]any, len(columns))
		for i := range values {
			values[i] = new(any)
		}

		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		m := make(map[string]interface{}, len(columns))
		for i, c := range columns {
			//m[c] = s.Time(gvar.New(values[i]))
			m[c] = values[i]
		}
		rs = append(rs, m)
	}

	return
}

// REST连接时区处理
func (s *TdEngine) Time(v string) (rs string) {
	if t, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", v); err == nil {
		rs = t.Local().Format("2006-01-02 15:04:05")
		return
	}
	rs = v
	return
}
