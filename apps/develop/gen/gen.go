package gen

import (
	"bytes"
	"fmt"
	"github.com/kakuilan/kgo"
	"os"
	"pandax/apps/develop/entity"
	"pandax/apps/develop/services"
	"pandax/base/biz"
	"pandax/base/config"
	"pandax/base/utils"
	"strconv"
	"strings"
	"text/template"
)

var (
	ToolsGenTableColumn = &toolsGenTableColumn{
		ColumnTypeStr:      []string{"char", "varchar", "narchar", "varchar2", "tinytext", "text", "mediumtext", "longtext"},
		ColumnTypeTime:     []string{"datetime", "time", "date", "timestamp"},
		ColumnTypeNumber:   []string{"tinyint", "smallint", "mediumint", "int", "number", "integer", "bigint", "float", "float", "double", "decimal"},
		ColumnNameNotEdit:  []string{"id", "created_by", "created_at", "updated_by", "updated_at", "deleted_at"},
		ColumnNameNotList:  []string{"id", "created_by", "updated_by", "created_at", "updated_at", "deleted_at"},
		ColumnNameNotQuery: []string{"id", "created_by", "updated_by", "created_at", "updated_at", "deleted_at", "remark"},
	}
)

type toolsGenTableColumn struct {
	ColumnTypeStr      []string //数据库字符串类型
	ColumnTypeTime     []string //数据库时间类型
	ColumnTypeNumber   []string //数据库数字类型
	ColumnNameNotEdit  []string //页面不需要编辑字段
	ColumnNameNotList  []string //页面不需要显示的列表字段
	ColumnNameNotQuery []string //页面不需要查询字段
}

// GetDbType 获取数据库类型字段
func (s *toolsGenTableColumn) GetDbType(columnType string) string {
	if strings.Index(columnType, "(") > 0 {
		return columnType[0:strings.Index(columnType, "(")]
	} else {
		return columnType
	}
}

// IsExistInArray 判断 value 是否存在在切片array中
func (s *toolsGenTableColumn) IsExistInArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// GetColumnLength 获取字段长度
func (s *toolsGenTableColumn) GetColumnLength(columnType string) int {
	start := strings.Index(columnType, "(")
	end := strings.Index(columnType, ")")
	result := ""
	if start >= 0 && end >= 0 {
		result = columnType[start+1 : end-1]
	}
	i, _ := strconv.Atoi(result)
	return i
}

// IsStringObject 判断是否是数据库字符串类型
func (s *toolsGenTableColumn) IsStringObject(dataType string) bool {
	return s.IsExistInArray(dataType, s.ColumnTypeStr)
}

// IsTimeObject 判断是否是数据库时间类型
func (s *toolsGenTableColumn) IsTimeObject(dataType string) bool {
	return s.IsExistInArray(dataType, s.ColumnTypeTime)
}

// IsNumberObject 是否数字类型
func (s *toolsGenTableColumn) IsNumberObject(dataType string) bool {
	return s.IsExistInArray(dataType, s.ColumnTypeNumber)
}

// IsNotEdit 是否不可编辑
func (s *toolsGenTableColumn) IsNotEdit(name string) bool {
	return s.IsExistInArray(name, s.ColumnNameNotEdit)
}

// IsNotList 不在列表显示
func (s *toolsGenTableColumn) IsNotList(name string) bool {
	return s.IsExistInArray(name, s.ColumnNameNotList)
}

// IsNotQuery 不可用于查询
func (s *toolsGenTableColumn) IsNotQuery(name string) bool {
	return s.IsExistInArray(name, s.ColumnNameNotQuery)
}

// CheckNameColumn 检查字段名后4位是否是name
func (s *toolsGenTableColumn) CheckNameColumn(columnName string) bool {
	if len(columnName) >= 4 {
		end := len(columnName)
		start := end - 4
		if start <= 0 {
			start = 0
		}
		tmp := columnName[start:end]
		if tmp == "name" {
			return true
		}
	}
	return false
}

// CheckStatusColumn 检查字段名后6位是否是status
func (s *toolsGenTableColumn) CheckStatusColumn(columnName string) bool {
	if len(columnName) >= 6 {
		end := len(columnName)
		start := end - 6

		if start <= 0 {
			start = 0
		}
		tmp := columnName[start:end]

		if tmp == "status" {
			return true
		}
	}

	return false
}

// CheckTypeColumn 检查字段名后4位是否是type
func (s *toolsGenTableColumn) CheckTypeColumn(columnName string) bool {
	if len(columnName) >= 4 {
		end := len(columnName)
		start := end - 4

		if start <= 0 {
			start = 0
		}

		if columnName[start:end] == "type" {
			return true
		}
	}
	return false
}

// CheckSexColumn 检查字段名后3位是否是sex
func (s *toolsGenTableColumn) CheckSexColumn(columnName string) bool {
	if len(columnName) >= 3 {
		end := len(columnName)
		start := end - 3
		if start <= 0 {
			start = 0
		}
		if columnName[start:end] == "sex" {
			return true
		}
	}
	return false
}

func (s *toolsGenTableColumn) GenTableInit(tableName string) entity.DevGenTable {
	var data entity.DevGenTable
	data.TableName = tableName
	tableNameList := strings.Split(tableName, "_")
	for i := 0; i < len(tableNameList); i++ {
		strStart := string([]byte(tableNameList[i])[:1])
		strEnd := string([]byte(tableNameList[i])[1:])
		// 大驼峰表名 结构体使用
		data.ClassName += strings.ToUpper(strStart) + strEnd
		// 小驼峰表名 js函数名和权限标识使用
		if i == 0 {
			data.BusinessName += strings.ToLower(strStart) + strEnd
		} else {
			data.BusinessName += strings.ToUpper(strStart) + strEnd
		}
	}
	data.PackageName = "admin"
	data.TplCategory = "crud"
	// 中横线表名称，接口路径、前端文件夹名称和js名称使用
	data.ModuleName = strings.Replace(tableName, "_", "-", -1)

	dbTable := services.DevGenTableModelDao.FindDbTableOne(tableName)
	dbColumn := services.DevTableColumnModelDao.FindDbTableColumnList(tableName)

	data.TableComment = dbTable.TableComment
	if dbTable.TableComment == "" {
		data.TableComment = data.ClassName
	}
	data.FunctionName = data.TableComment
	data.FunctionAuthor = "panda"

	dcs := *dbColumn
	for i := 0; i < len(dcs); i++ {
		var column entity.DevGenTableColumn
		column.ColumnComment = dcs[i].ColumnComment
		column.ColumnName = dcs[i].ColumnName
		column.ColumnType = dcs[i].ColumnType
		column.Sort = i + 1
		column.IsPk = "0"

		nameList := strings.Split(dcs[i].ColumnName, "_")
		for i := 0; i < len(nameList); i++ {
			strStart := string([]byte(nameList[i])[:1])
			strend := string([]byte(nameList[i])[1:])
			column.GoField += strings.ToUpper(strStart) + strend
			if i == 0 {
				column.JsonField = strings.ToLower(strStart) + strend
			} else {
				column.JsonField += strings.ToUpper(strStart) + strend
			}
		}
		if strings.Contains(dcs[i].ColumnKey, "PR") {
			column.IsPk = "1"
			data.PkColumn = dcs[i].ColumnName
			data.PkGoField = column.GoField
			data.PkJsonField = column.JsonField
		}

		dataType := s.GetDbType(column.ColumnType)
		if s.IsStringObject(dataType) {
			//字段为字符串类型
			column.GoType = "string"
			columnLength := s.GetColumnLength(column.ColumnType)
			if columnLength >= 500 {
				column.HtmlType = "textarea"
			} else {
				column.HtmlType = "input"
			}
		} else if s.IsTimeObject(dataType) {
			//字段为时间类型
			column.GoType = "Time"
			column.HtmlType = "datetime"
		} else if s.IsNumberObject(dataType) {
			//字段为数字类型
			column.HtmlType = "input"
			t, _ := utils.ReplaceString(`\(.+\)`, "", column.ColumnType)
			t = strings.Split(strings.TrimSpace(t), " ")[0]
			t = strings.ToLower(t)
			// 如果是浮点型
			switch t {
			case "float", "double", "decimal":
				column.GoType = "float64"
			case "bit", "int", "tinyint", "small_int", "smallint", "medium_int", "mediumint":
				if utils.Contains(column.ColumnType, "unsigned") != -1 {
					column.GoType = "uint"
				} else {
					column.GoType = "int"
				}
			case "big_int", "bigint":
				if utils.Contains(column.ColumnType, "unsigned") != -1 {
					column.GoType = "uint64"
				} else {
					column.GoType = "int64"
				}
			}
		}
		//新增字段
		if s.IsNotEdit(column.ColumnName) {
			column.IsRequired = "0"
			column.IsInsert = "0"
		} else {
			column.IsInsert = "1"
			if strings.Index(column.ColumnName, "name") >= 0 || strings.Index(column.ColumnName, "status") >= 0 {
				column.IsRequired = "1"
			}
		}

		// 编辑字段
		if s.IsNotEdit(column.ColumnName) {
			column.IsEdit = "0"
		} else {
			if column.IsPk == "1" {
				column.IsEdit = "0"
			} else {
				column.IsEdit = "1"
			}
		}
		// 列表字段
		if s.IsNotList(column.ColumnName) {
			column.IsList = "0"
		} else {
			column.IsList = "1"
		}
		// 查询字段
		if s.IsNotQuery(column.ColumnName) {
			column.IsQuery = "0"
		} else {
			column.IsQuery = "1"
		}

		// 查询字段类型
		if s.CheckNameColumn(column.ColumnName) {
			column.QueryType = "LIKE"
		} else {
			column.QueryType = "EQ"
		}
		// 状态字段设置单选框
		if s.CheckStatusColumn(column.ColumnName) {
			column.HtmlType = "radio"
		} else if s.CheckTypeColumn(column.ColumnName) || s.CheckSexColumn(column.ColumnName) {
			// 类型&性别字段设置下拉框
			column.HtmlType = "select"
		}
		data.Columns = append(data.Columns, column)
	}
	return data
}

// 视图预览
func Preview(tableId int64) map[string]interface{} {
	t1, err := template.ParseFiles("resource/template/go/entity.template")
	biz.ErrIsNil(err, fmt.Sprintf("entity模版读取失败！错误详情：%s", err.Error()))

	t2, err := template.ParseFiles("resource/template/go/api.template")
	biz.ErrIsNil(err, fmt.Sprintf("api模版读取失败！错误详情：%s", err.Error()))

	t3, err := template.ParseFiles("resource/template/go/service.template")
	biz.ErrIsNil(err, fmt.Sprintf("service模版读取失败！错误详情：%s", err.Error()))

	t4, err := template.ParseFiles("resource/template/go/router.template")
	biz.ErrIsNil(err, fmt.Sprintf("router模版读取失败！错误详情：%s", err.Error()))

	t5, err := template.ParseFiles("resource/template/js/api.template")
	biz.ErrIsNil(err, fmt.Sprintf("js模版读取失败！错误详情：%s", err.Error()))

	t6, err := template.ParseFiles("resource/template/vue/list-vue.template")
	biz.ErrIsNil(err, fmt.Sprintf("vue列表模版读取失败！错误详情：%s", err.Error()))

	t7, err := template.ParseFiles("resource/template/vue/edit-vue.template")
	biz.ErrIsNil(err, fmt.Sprintf("vue编辑模版读取失败！错误详情：%s", err.Error()))

	tab := services.DevGenTableModelDao.FindOne(entity.DevGenTable{TableId: tableId}, false)

	var b1 bytes.Buffer
	err = t1.Execute(&b1, tab)
	var b2 bytes.Buffer
	err = t2.Execute(&b2, tab)
	var b3 bytes.Buffer
	err = t3.Execute(&b3, tab)
	var b4 bytes.Buffer
	err = t4.Execute(&b4, tab)
	var b5 bytes.Buffer
	err = t5.Execute(&b5, tab)
	var b6 bytes.Buffer
	err = t6.Execute(&b6, tab)
	var b7 bytes.Buffer
	err = t7.Execute(&b7, tab)

	mp := make(map[string]interface{})
	mp["template/entity.template"] = b1.String()
	mp["template/api.template"] = b2.String()
	mp["template/service.template"] = b3.String()
	mp["template/router.template"] = b4.String()
	mp["template/jsApi.template"] = b5.String()
	mp["template/listVue.template"] = b6.String()
	mp["template/editVue.template"] = b7.String()
	return mp
}

// 生成 代码
func GenCode(tableId int64) {

	tab := services.DevGenTableModelDao.FindOne(entity.DevGenTable{TableId: tableId}, false)
	tab.ModuleName = strings.Replace(tab.TableName, "_", "-", -1)

	t1, err := template.ParseFiles("resource/template/go/entity.template")
	biz.ErrIsNil(err, fmt.Sprintf("entity模版读取失败！错误详情：%s", err.Error()))

	t2, err := template.ParseFiles("resource/template/go/api.template")
	biz.ErrIsNil(err, fmt.Sprintf("api模版读取失败！错误详情：%s", err.Error()))

	t3, err := template.ParseFiles("resource/template/go/service.template")
	biz.ErrIsNil(err, fmt.Sprintf("service模版读取失败！错误详情：%s", err.Error()))

	t4, err := template.ParseFiles("resource/template/go/router.template")
	biz.ErrIsNil(err, fmt.Sprintf("router模版读取失败！错误详情：%s", err.Error()))

	t5, err := template.ParseFiles("resource/template/js/api.template")
	biz.ErrIsNil(err, fmt.Sprintf("js模版读取失败！错误详情：%s", err.Error()))

	t6, err := template.ParseFiles("resource/template/vue/list-vue.template")
	biz.ErrIsNil(err, fmt.Sprintf("vue列表模版读取失败！错误详情：%s", err.Error()))
	t7, err := template.ParseFiles("resource/template/vue/edit-vue.template")
	biz.ErrIsNil(err, fmt.Sprintf("vue编辑模版读取失败！错误详情：%s", err.Error()))

	kgo.KFile.Mkdir("./apps/"+tab.PackageName+"/api/", os.ModePerm)
	kgo.KFile.Mkdir("./apps/"+tab.PackageName+"/entity/", os.ModePerm)
	kgo.KFile.Mkdir("./apps/"+tab.PackageName+"/router/", os.ModePerm)
	kgo.KFile.Mkdir("./apps/"+tab.PackageName+"/services/", os.ModePerm)
	kgo.KFile.Mkdir(config.Conf.Gen.Frontpath+"/api/"+tab.PackageName+"/", os.ModePerm)
	kgo.KFile.Mkdir(config.Conf.Gen.Frontpath+"/views/"+tab.PackageName+"/"+tab.ModuleName+"/", os.ModePerm)
	kgo.KFile.Mkdir(config.Conf.Gen.Frontpath+"/views/"+tab.PackageName+"/"+tab.ModuleName+"/"+"component"+"/", os.ModePerm)

	var b1 bytes.Buffer
	err = t1.Execute(&b1, tab)
	var b2 bytes.Buffer
	err = t2.Execute(&b2, tab)
	var b3 bytes.Buffer
	err = t3.Execute(&b3, tab)
	var b4 bytes.Buffer
	err = t4.Execute(&b4, tab)
	var b5 bytes.Buffer
	err = t5.Execute(&b5, tab)
	var b6 bytes.Buffer
	err = t6.Execute(&b6, tab)
	var b7 bytes.Buffer
	err = t7.Execute(&b7, tab)

	kgo.KFile.WriteFile("./apps/"+tab.PackageName+"/entity/"+tab.TableName+".go", b1.Bytes())
	kgo.KFile.WriteFile("./apps/"+tab.PackageName+"/api/"+tab.TableName+".go", b2.Bytes())
	kgo.KFile.WriteFile("./apps/"+tab.PackageName+"/services/"+tab.TableName+".go", b3.Bytes())
	kgo.KFile.WriteFile("./apps/"+tab.PackageName+"/router/"+tab.TableName+".go", b4.Bytes())
	kgo.KFile.WriteFile(config.Conf.Gen.Frontpath+"/api/"+tab.PackageName+"/"+tab.ModuleName+".js", b5.Bytes())
	kgo.KFile.WriteFile(config.Conf.Gen.Frontpath+"/views/"+tab.PackageName+"/"+tab.ModuleName+"/index.vue", b6.Bytes())
	kgo.KFile.WriteFile(config.Conf.Gen.Frontpath+"/views/"+tab.PackageName+"/"+tab.ModuleName+"/"+"component"+"/index.vue", b7.Bytes())
}
