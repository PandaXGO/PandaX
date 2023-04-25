package driver

import (
	"database/sql"
	"fmt"
	"github.com/XM-GO/PandaKit/cache"
	"pandax/apps/visual/entity"
	"pandax/pkg/global"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// db实例
type DbInstance struct {
	Id   string
	Info *entity.VisualDataSource
	db   *sql.DB
}

func NewDbInstance(source *entity.VisualDataSource) *DbInstance {
	byCache := GetDbInstanceByCache(source.SourceId)
	if byCache != nil {
		return byCache
	}
	conn, err := GetDbConn(source)
	if err != nil {
		return nil
	}
	di := &DbInstance{
		Id:   source.SourceId,
		Info: source,
		db:   conn,
	}
	AddDbInstanceToCache(source.SourceId, di)
	return di
}

// 执行查询语句
// 依次返回 列名数组，结果map，错误
func (d *DbInstance) SelectData(execSql string) ([]string, []map[string]interface{}, error) {
	return SelectDataByDb(d.db, execSql)
}

// 将查询结果映射至struct，可具体参考sqlx库
func (d *DbInstance) SelectData2Struct(execSql string, dest interface{}) error {
	return Select2StructByDb(d.db, execSql, dest)
}

// 执行内部查询语句，不返回列名以及不限制行数
// 依次返回 结果map，错误
func (d *DbInstance) innerSelect(execSql string) ([]map[string]interface{}, error) {
	_, res, err := SelectDataByDb(d.db, execSql)
	return res, err
}

// 执行 update, insert, delete，建表等sql
// 返回影响条数和错误
func (d *DbInstance) Exec(sql string) (int64, error) {
	res, err := d.db.Exec(sql)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// 获取数据库元信息实现接口
func (di *DbInstance) GetMeta() DbMetadata {
	dbType := di.Info.SourceType
	if dbType == entity.DbTypeMysql {
		return &MysqlMetadata{di: di}
	}
	if dbType == entity.DbTypePostgres {
		return &PgsqlMetadata{di: di}
	}
	return nil
}

// 关闭连接
func (d *DbInstance) Close() {
	if d.db != nil {
		if err := d.db.Close(); err != nil {
			global.Log.Errorf("关闭数据库实例[%s]连接失败: %s", d.Id, err.Error())
		}
		d.db = nil
	}
}

const DbConnExpireTime = 45 * time.Minute

//------------------------------------------------------------------------------

// 客户端连接缓存，指定时间内没有访问则会被关闭, key为数据库实例id:数据库
var dbCache = cache.NewTimedCache(DbConnExpireTime, 5*time.Second).
	WithUpdateAccessTime(true).
	OnEvicted(func(key interface{}, value interface{}) {
		global.Log.Info(fmt.Sprintf("删除db连接缓存 id = %s", key))
		value.(*DbInstance).Close()
	})

func GetDbCacheKey(dbId uint64, db string) string {
	return fmt.Sprintf("%d:%s", dbId, db)
}

// 删除db缓存并关闭该数据库所有连接
func CloseDb(id string) {
	dbCache.Delete(id)
}
func GetDbInstanceByCache(id string) *DbInstance {
	if load, ok := dbCache.Get(id); ok {
		return load.(*DbInstance)
	}
	return nil
}

// 将实体添加到缓存中
func AddDbInstanceToCache(id string, di *DbInstance) {
	dbCache.AddIfAbsent(id, di)
}

func TestConnection(d *entity.VisualDataSource) error {
	// 验证第一个库是否可以连接即可
	DB, err := GetDbConn(d)
	if err != nil {
		return err
	} else {
		DB.Close()
		return nil
	}
}

// 获取数据库连接
func GetDbConn(d *entity.VisualDataSource) (*sql.DB, error) {
	var DB *sql.DB
	var err error
	if d.SourceType == entity.DbTypeMysql {
		DB, err = getMysqlDB(d)
	} else if d.SourceType == entity.DbTypePostgres {
		DB, err = getPgsqlDB(d)
	}

	if err != nil {
		return nil, err
	}
	err = DB.Ping()
	if err != nil {
		DB.Close()
		return nil, err
	}

	return DB, nil
}

func SelectDataByDb(db *sql.DB, selectSql string) ([]string, []map[string]interface{}, error) {
	rows, err := db.Query(selectSql)
	if err != nil {
		return nil, nil, err
	}
	// rows对象一定要close掉，如果出错，不关掉则会很迅速的达到设置最大连接数，
	// 后面的链接过来直接报错或拒绝，实际上也没有起效果
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	colTypes, _ := rows.ColumnTypes()
	// 这里表示一行填充数据
	scans := make([]interface{}, len(colTypes))
	// 这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(colTypes))
	// 这里scans引用vals，把数据填充到[]byte里
	for k := range vals {
		scans[k] = &vals[k]
	}

	result := make([]map[string]interface{}, 0)
	// 列名用于前端表头名称按照数据库与查询字段顺序显示
	colNames := make([]string, 0)
	// 是否第一次遍历，列名数组只需第一次遍历时加入
	isFirst := true
	for rows.Next() {
		// 不Scan也会导致等待，该链接实际处于未工作的状态，然后也会导致连接数迅速达到最大
		err := rows.Scan(scans...)
		if err != nil {
			return nil, nil, err
		}
		// 每行数据
		rowData := make(map[string]interface{})
		// 把vals中的数据复制到row中
		for i, v := range vals {
			colType := colTypes[i]
			colName := colType.Name()
			// 如果是第一行，则将列名加入到列信息中，由于map是无序的，所有需要返回列名的有序数组
			if isFirst {
				colNames = append(colNames, colName)
			}
			rowData[colName] = valueConvert(v, colType)
		}
		// 放入结果集
		result = append(result, rowData)
		isFirst = false
	}
	return colNames, result, nil
}

// 将查询的值转为对应列类型的实际值，不全部转为字符串
func valueConvert(data []byte, colType *sql.ColumnType) interface{} {
	if data == nil {
		return nil
	}
	// 列的数据库类型名
	colDatabaseTypeName := strings.ToLower(colType.DatabaseTypeName())

	// 如果类型是bit，则直接返回第一个字节即可
	if strings.Contains(colDatabaseTypeName, "bit") {
		return data[0]
	}

	// 这里把[]byte数据转成string
	stringV := string(data)
	if stringV == "" {
		return ""
	}
	colScanType := strings.ToLower(colType.ScanType().Name())

	if strings.Contains(colScanType, "int") {
		// 如果长度超过16位，则返回字符串，因为前端js长度大于16会丢失精度
		if len(stringV) > 16 {
			return stringV
		}
		intV, _ := strconv.Atoi(stringV)
		switch colType.ScanType().Kind() {
		case reflect.Int8:
			return int8(intV)
		case reflect.Uint8:
			return uint8(intV)
		case reflect.Int64:
			return int64(intV)
		case reflect.Uint64:
			return uint64(intV)
		case reflect.Uint:
			return uint(intV)
		default:
			return intV
		}
	}
	if strings.Contains(colScanType, "float") || strings.Contains(colDatabaseTypeName, "decimal") {
		floatV, _ := strconv.ParseFloat(stringV, 64)
		return floatV
	}

	return stringV
}

// 查询数据结果映射至struct。可参考sqlx库
func Select2StructByDb(db *sql.DB, selectSql string, dest interface{}) error {
	rows, err := db.Query(selectSql)
	if err != nil {
		return err
	}
	// rows对象一定要close掉，如果出错，不关掉则会很迅速的达到设置最大连接数，
	// 后面的链接过来直接报错或拒绝，实际上也没有起效果
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	return scanAll(rows, dest, false)
}
