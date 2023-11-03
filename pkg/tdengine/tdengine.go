package tdengine

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/PandaXGO/PandaKit/httpclient"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	"strings"
	"time"
)

type TdEngine struct {
	db     *sql.DB
	dbName string
}

func InitTd(username, password, host, db string) (*TdEngine, error) {
	_, err := CreateDataBase(username, password, host, db)
	if err != nil {
		return nil, err
	}
	return NewTdengine(username, password, host, db)
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

// 创建数据库
func CreateDataBase(username, password, host, dbname string) (float64, error) {
	sql := "CREATE DATABASE if not exists " + dbname + " KEEP 365 VGROUPS 10"
	url := fmt.Sprintf("http://%s/rest/sql", host)
	token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
	data, err := httpclient.NewRequest(url).Header("Authorization", "Basic "+token).PostText(sql).BodyToMap()
	if err != nil {
		return 0, err
	}
	return data["rows"].(float64), nil
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

// GetListTableByStableName 获取指定超级表下所有的子表
func (s *TdEngine) GetListTableByStableName(stableName string) (data []*TDEngineTablesList, err error) {
	sql := `SELECT table_name AS tableName, db_name AS dbName, create_time AS createTime, stable_name AS stableName FROM information_schema.ins_tables WHERE db_name = ? and stable_name = ?`

	rows, err := s.db.Query(sql, s.db, strings.ToLower(stableName))
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
	rows, err := s.db.Query("DESCRIBE " + s.dbName + "." + strings.ToLower(tableName) + ";")
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

	rows, err := s.db.Query("SELECT * FROM " + strings.ToLower(tableName))
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
