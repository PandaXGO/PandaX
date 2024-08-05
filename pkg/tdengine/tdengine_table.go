package tdengine

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/kakuilan/kgo"
)

// RunSql 运行
func (s *TdEngine) RunSql(sql string) error {
	_, err := s.db.Exec(sql)
	return err
}

// InsertDevice 数据入库
func (s *TdEngine) InsertDevice(deviceKey string, data map[string]interface{}) error {
	if len(data) == 0 {
		return nil
	}

	var (
		field        []string
		value        []interface{}
		placeholders []string
	)

	for k, v := range data {
		field = append(field, k)
		value = append(value, fmt.Sprintf("'%s'", kgo.KConv.ToStr(v)))
		placeholders = append(placeholders, "?")
	}

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", strings.ToLower(deviceKey), strings.Join(field, ","), strings.Join(placeholders, ","))
	_, err := s.db.Exec(sql, value...)

	return err
}

// CreateStable 创建超级表 rowdata 源数据，在需要数据解析时使用
func (s *TdEngine) CreateStable(table string) error {
	columns := []string{"ts TIMESTAMP,rowdata NCHAR(255)"}
	sql := fmt.Sprintf("CREATE STABLE IF NOT EXISTS %s.%s (%s) TAGS (device nchar(64))", s.dbName, table, strings.Join(columns, ","))
	_, err := s.db.Exec(sql)
	return err
}

// CreateTable 添加子表
func (s *TdEngine) CreateTable(stable, table string) error {
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS  %s USING %s TAGS ('%s')", table, stable, table)
	_, err := s.db.Exec(sql)
	return err
}

func (s *TdEngine) column(dataType, key, name string, maxLength int) string {
	column := ""
	comment := ""
	if name != "" {
		comment = "COMMENT '" + name + "'"
	}
	tdType := ""
	switch dataType {
	case "int64":
		tdType = "INT"
	case "long":
		tdType = "BIGINT"
	case "float64":
		tdType = "FLOAT"
	case "double":
		tdType = "DOUBLE"
	case "string":
		if maxLength == 0 {
			maxLength = 255
		}
		tdType = "NCHAR(" + strconv.Itoa(maxLength) + ")"
	case "boolean":
		tdType = "BOOL"
	case "date":
		tdType = "TIMESTAMP"
	default:
		if maxLength == 0 {
			maxLength = 255
		}
		tdType = "VARCHAR(" + strconv.Itoa(maxLength*10) + ")"
	}
	column = fmt.Sprintf("%s %s %s", key, tdType, comment)
	return column
}

// 删除超级表
func (s *TdEngine) DropStable(table string) error {
	sql := fmt.Sprintf("DROP STABLE IF EXISTS %s.%s", s.dbName, strings.ToLower(table))
	_, err := s.db.Exec(sql)
	return err
}

// 删除子表
func (s *TdEngine) DropTable(table string) error {
	sql := fmt.Sprintf("DROP TABLE IF EXISTS %s.%s", s.dbName, strings.ToLower(table))
	_, err := s.db.Exec(sql)
	return err
}

// AddSTableField 添加数据库超级表字段
func (s *TdEngine) AddSTableField(tableName, fieldName string, dataType string, len int) error {
	sql := fmt.Sprintf("ALTER STABLE %s.%s ADD COLUMN %s", s.dbName, strings.ToLower(tableName), s.column(dataType, fieldName, "", len))
	_, err := s.db.Exec(sql)
	return err
}

// AddTableField 添加数据库表字段
func (s *TdEngine) AddTableField(tableName, fieldName string, dataType string, len int) error {
	sql := fmt.Sprintf("ALTER TABLE %s.%s ADD COLUMN %s", s.dbName, strings.ToLower(tableName), s.column(dataType, fieldName, "", len))
	_, err := s.db.Exec(sql)
	return err
}

// DelTableField 删除数据库表字段
func (s *TdEngine) DelTableField(tableName, fieldName string) error {
	sql := fmt.Sprintf("ALTER TABLE %s.%s DROP COLUMN %s", s.dbName, strings.ToLower(tableName), fieldName)
	_, err := s.db.Exec(sql)
	return err
}

// DelSTableField 删除数据库超级表字段
func (s *TdEngine) DelSTableField(tableName, fieldName string) error {
	sql := fmt.Sprintf("ALTER STABLE %s.%s DROP COLUMN %s", s.dbName, strings.ToLower(tableName), fieldName)
	_, err := s.db.Exec(sql)
	return err
}

// ModifyDatabaseField 修改数据库指定字段长度
func (s *TdEngine) ModifyDatabaseField(tableName, fieldName string, dataType string, len int) error {
	sql := fmt.Sprintf("ALTER STABLE %s.%s MODIFY COLUMN %s", s.dbName, strings.ToLower(tableName), s.column(dataType, fieldName, "", len))
	_, err := s.db.Exec(sql)
	if err != nil {
		return errors.New("设置字段长度失败,长度只能增大不能缩小")
	}
	return nil
}

// AddTag 添加标签
func (s *TdEngine) AddTag(tableName, tagName string, dataType string, len int) error {
	sql := fmt.Sprintf("ALTER STABLE %s.%s ADD TAG %s", s.dbName, tableName, s.column(dataType, tagName, "", len))
	_, err := s.db.Exec(sql)
	return err
}

// DelTag 删除标签
func (s *TdEngine) DelTag(tableName, tagName string) error {
	sql := fmt.Sprintf("ALTER STABLE %s.%s DROP TAG %s", s.dbName, tableName, tagName)
	_, err := s.db.Exec(sql)
	return err
}

// ModifyTag 修改标签
func (s *TdEngine) ModifyTag(tableName, tagName string, dataType string, len int) error {
	sql := fmt.Sprintf("ALTER STABLE %s.%s MODIFY TAG %s", s.dbName, tableName, s.column(dataType, tagName, "", len))
	_, err := s.db.Exec(sql)
	if err != nil {
		return errors.New("设置标签长度失败,长度只能增大不能缩小")
	}
	return nil
}
