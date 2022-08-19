package gen

import (
	"bytes"
	"fmt"
	"github.com/XM-GO/PandaKit/biz"
	"github.com/XM-GO/PandaKit/utils"
	"github.com/kakuilan/kgo"
	"log"
	"os"
	"pandax/apps/develop/entity"
	"pandax/apps/develop/services"
	sysEntity "pandax/apps/system/entity"
	sysServices "pandax/apps/system/services"
	"pandax/pkg/global"
	"strconv"
	"strings"
	"sync"
	"text/template"
)

var (
	ToolsGenTableColumn = &toolsGenTableColumn{
		ColumnTypeStr:      []string{"char", "varchar", "narchar", "varchar2", "tinytext", "text", "mediumtext", "longtext"},
		ColumnTypeTime:     []string{"datetime", "time", "date", "timestamp", "timestamptz"},
		ColumnTypeNumber:   []string{"tinyint", "smallint", "mediumint", "int", "int2", "int4", "int8", "number", "integer", "numeric", "bigint", "float", "float4", "float8", "double", "decimal"},
		ColumnNameNotEdit:  []string{"create_by", "update_by", "create_time", "update_time", "delete_time"},
		ColumnNameNotList:  []string{"create_by", "update_by", "update_time", "delete_time"},
		ColumnNameNotQuery: []string{"create_by", "update_by", "create_time", "update_time", "delete_time", "remark"},
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
	if strings.Contains(name, "id") {
		return true
	}
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
		// js函数名和权限标识使用
		if i >= 1 {
			data.BusinessName += strings.ToLower(strStart) + strEnd
			data.FunctionName = strings.ToUpper(strStart) + strEnd
		}
	}
	data.PackageName = "system"
	data.TplCategory = "crud"
	// 中横线表名称，接口路径、前端文件夹名称和js名称使用
	data.ModuleName = strings.Replace(tableName, "_", "-", -1)

	dbColumn := services.DevTableColumnModelDao.FindDbTableColumnList(tableName)
	data.TableComment = data.ClassName
	data.FunctionAuthor = "panda"

	wg := sync.WaitGroup{}
	dcs := *dbColumn
	for i := 0; i < len(dcs); i++ {
		index := i
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			defer wg.Done()
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
				global.Log.Info("是否自增主键", dcs[i].Extra)
				if dcs[i].Extra == "auto_increment" {
					column.IsIncrement = "1"
				}
			}

			if column.ColumnComment == "" {
				column.ColumnComment = column.GoField
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
				column.HtmlType = "input"
				t := ""
				if global.Conf.Server.DbType == "postgresql" {
					t = column.ColumnType
				} else {
					t, _ := utils.ReplaceString(`\(.+\)`, "", column.ColumnType)
					t = strings.Split(strings.TrimSpace(t), " ")[0]
					t = strings.ToLower(t)
				}
				// 如果是浮点型
				switch t {
				case "float", "float4", "float8", "double", "decimal":
					column.GoType = "float64"
				case "bit", "int", "int2", "int4", "tinyint", "small_int", "smallint", "medium_int", "mediumint":
					if utils.Contains(column.ColumnType, "unsigned") != -1 {
						column.GoType = "uint"
					} else {
						column.GoType = "int"
					}
				case "big_int", "int8", "bigint", "numeric":
					if utils.Contains(column.ColumnType, "unsigned") != -1 {
						column.GoType = "uint64"
					} else {
						column.GoType = "int64"
					}
				}
			} else {
				switch dataType {
				case "bool":
					column.GoType = "bool"
					column.HtmlType = "switch"
				}
			}
			//新增字段
			if s.IsNotEdit(column.ColumnName) {
				column.IsRequired = "0"
				column.IsInsert = "0"
			} else {
				column.IsRequired = "0"
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
		}(&wg, index)
	}
	wg.Wait()
	return data
}

// 视图预览
func Preview(tableId int64) map[string]any {
	t1, err := template.ParseFiles("resource/template/go/entity.template")
	biz.ErrIsNil(err, "entity模版读取失败")

	t2, err := template.ParseFiles("resource/template/go/service.template")
	biz.ErrIsNil(err, "service模版读取失败")

	t3, err := template.ParseFiles("resource/template/go/api.template")
	biz.ErrIsNil(err, "api模版读取失败！")

	t4, err := template.ParseFiles("resource/template/go/router.template")
	biz.ErrIsNil(err, "router模版读取失败")

	t5, err := template.ParseFiles("resource/template/js/api.template")
	biz.ErrIsNil(err, "js模版读取失败")

	t6, err := template.ParseFiles("resource/template/vue/list-vue.template")
	biz.ErrIsNil(err, "vue列表模版读取失败！")

	t7, err := template.ParseFiles("resource/template/vue/edit-vue.template")
	biz.ErrIsNil(err, "vue编辑模版读取失败！")

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

	mp := make(map[string]any)
	mp["template/entity.template"] = b1.String()
	mp["template/service.template"] = b2.String()
	mp["template/api.template"] = b3.String()
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
	biz.ErrIsNil(err, "entity模版读取失败！")

	t2, err := template.ParseFiles("resource/template/go/service.template")
	biz.ErrIsNil(err, "service模版读取失败！")

	t3, err := template.ParseFiles("resource/template/go/api.template")
	biz.ErrIsNil(err, "api模版读取失败！")

	t4, err := template.ParseFiles("resource/template/go/router.template")
	biz.ErrIsNil(err, "router模版读取失败！")

	t5, err := template.ParseFiles("resource/template/js/api.template")
	biz.ErrIsNil(err, "js模版读取失败！")

	t6, err := template.ParseFiles("resource/template/vue/list-vue.template")
	biz.ErrIsNil(err, "vue列表模版读取失败！")
	t7, err := template.ParseFiles("resource/template/vue/edit-vue.template")
	biz.ErrIsNil(err, "vue编辑模版读取失败！")

	kgo.KFile.Mkdir("./apps/"+tab.PackageName+"/api/", os.ModePerm)
	kgo.KFile.Mkdir("./apps/"+tab.PackageName+"/entity/", os.ModePerm)
	kgo.KFile.Mkdir("./apps/"+tab.PackageName+"/router/", os.ModePerm)
	kgo.KFile.Mkdir("./apps/"+tab.PackageName+"/services/", os.ModePerm)
	kgo.KFile.Mkdir(global.Conf.Gen.Frontpath+"/api/"+tab.PackageName+"/", os.ModePerm)
	kgo.KFile.Mkdir(global.Conf.Gen.Frontpath+"/views/"+tab.PackageName+"/"+tab.BusinessName+"/", os.ModePerm)
	kgo.KFile.Mkdir(global.Conf.Gen.Frontpath+"/views/"+tab.PackageName+"/"+tab.BusinessName+"/"+"component"+"/", os.ModePerm)

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
	kgo.KFile.WriteFile("./apps/"+tab.PackageName+"/services/"+tab.TableName+".go", b2.Bytes())
	kgo.KFile.WriteFile("./apps/"+tab.PackageName+"/api/"+tab.TableName+".go", b3.Bytes())
	kgo.KFile.WriteFile("./apps/"+tab.PackageName+"/router/"+tab.TableName+".go", b4.Bytes())
	kgo.KFile.WriteFile(global.Conf.Gen.Frontpath+"/api/"+tab.PackageName+"/"+tab.BusinessName+".js", b5.Bytes())
	kgo.KFile.WriteFile(global.Conf.Gen.Frontpath+"/views/"+tab.PackageName+"/"+tab.BusinessName+"/index.vue", b6.Bytes())
	kgo.KFile.WriteFile(global.Conf.Gen.Frontpath+"/views/"+tab.PackageName+"/"+tab.BusinessName+"/"+"component"+"/editModule.vue", b7.Bytes())
}

// GenConfigure 生成菜单，api
func GenConfigure(tableId, parentId int) {
	tab := services.DevGenTableModelDao.FindOne(entity.DevGenTable{TableId: int64(tableId)}, false)

	//生成菜单 一个菜单 三个按钮
	component := "Layout"
	if parentId != 0 {
		component = fmt.Sprintf("/%s/%s/index", tab.PackageName, tab.BusinessName)
	}
	menu := sysEntity.SysMenu{
		ParentId:    int64(parentId),
		MenuName:    tab.TableComment,
		MenuType:    "C",
		Sort:        1,
		Icon:        "elementSetting",
		Path:        fmt.Sprintf("/%s/%s", tab.PackageName, tab.BusinessName),
		Component:   component,
		IsIframe:    "1",
		IsHide:      "0",
		IsKeepAlive: "1",
		IsAffix:     "1",
		Permission:  fmt.Sprintf("%s:%s:list", tab.PackageName, tab.BusinessName),
		Status:      "0",
		CreateBy:    "admin",
	}
	insert := sysServices.SysMenuModelDao.Insert(menu)
	log.Println("insert", insert.MenuId)
	//新增按钮
	menuA := sysEntity.SysMenu{
		ParentId:   insert.MenuId,
		MenuName:   "新增" + tab.TableComment,
		MenuType:   "F",
		Sort:       1,
		Permission: fmt.Sprintf("%s:%s:add", tab.PackageName, tab.BusinessName),
		Status:     "0",
		CreateBy:   "admin",
	}
	go sysServices.SysMenuModelDao.Insert(menuA)
	//修改按钮
	menuE := sysEntity.SysMenu{
		ParentId:   insert.MenuId,
		MenuName:   "修改" + tab.TableComment,
		MenuType:   "F",
		Sort:       2,
		Permission: fmt.Sprintf("%s:%s:edit", tab.PackageName, tab.BusinessName),
		Status:     "0",
		CreateBy:   "admin",
	}
	go sysServices.SysMenuModelDao.Insert(menuE)
	//删除按钮
	menuD := sysEntity.SysMenu{
		ParentId:   insert.MenuId,
		MenuName:   "删除" + tab.TableComment,
		MenuType:   "F",
		Sort:       3,
		Permission: fmt.Sprintf("%s:%s:delete", tab.PackageName, tab.BusinessName),
		Status:     "0",
		CreateBy:   "admin",
	}
	go sysServices.SysMenuModelDao.Insert(menuD)
	//生成api
	apiL := sysEntity.SysApi{
		Path:        fmt.Sprintf("/%s/%s/list", tab.PackageName, tab.BusinessName),
		Description: fmt.Sprintf("查询%s列表（分页）", tab.TableComment),
		ApiGroup:    tab.BusinessName,
		Method:      "GET",
	}
	go sysServices.SysApiModelDao.Insert(apiL)
	apiG := sysEntity.SysApi{
		Path:        fmt.Sprintf("/%s/%s/:%s", tab.PackageName, tab.BusinessName, tab.PkJsonField),
		Description: fmt.Sprintf("获取%s信息", tab.TableComment),
		ApiGroup:    tab.BusinessName,
		Method:      "GET",
	}
	go sysServices.SysApiModelDao.Insert(apiG)
	apiA := sysEntity.SysApi{
		Path:        fmt.Sprintf("/%s/%s", tab.PackageName, tab.BusinessName),
		Description: fmt.Sprintf("添加%s信息", tab.TableComment),
		ApiGroup:    tab.BusinessName,
		Method:      "POST",
	}
	go sysServices.SysApiModelDao.Insert(apiA)
	apiE := sysEntity.SysApi{
		Path:        fmt.Sprintf("/%s/%s", tab.PackageName, tab.BusinessName),
		Description: fmt.Sprintf("修改%s信息", tab.TableComment),
		ApiGroup:    tab.BusinessName,
		Method:      "PUT",
	}
	go sysServices.SysApiModelDao.Insert(apiE)
	apiD := sysEntity.SysApi{
		Path:        fmt.Sprintf("/%s/%s/:%s", tab.PackageName, tab.BusinessName, tab.PkJsonField),
		Description: fmt.Sprintf("删除%s信息", tab.TableComment),
		ApiGroup:    tab.BusinessName,
		Method:      "DELETE",
	}
	go sysServices.SysApiModelDao.Insert(apiD)
}
