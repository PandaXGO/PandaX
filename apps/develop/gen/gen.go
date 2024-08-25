package gen

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"text/template"

	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/utils"
	"pandax/apps/develop/entity"
	"pandax/apps/develop/services"
	sysEntity "pandax/apps/system/entity"
	sysServices "pandax/apps/system/services"
	"pandax/pkg/global"

	"github.com/kakuilan/kgo"
)

var ToolsGenTableColumn = toolsGenTableColumn{}

type toolsGenTableColumn struct{}

var (
	ColumnTypeStr      = []string{"char", "varchar", "narchar", "varchar2", "tinytext", "text", "mediumtext", "longtext"}
	ColumnTypeTime     = []string{"datetime", "time", "date", "timestamp", "timestamptz"}
	ColumnTypeNumber   = []string{"tinyint", "smallint", "mediumint", "int", "int2", "int4", "int8", "number", "integer", "numeric", "bigint", "float", "float4", "float8", "double", "decimal"}
	ColumnNameNotEdit  = []string{"create_by", "update_by", "create_time", "update_time", "delete_time"}
	ColumnNameNotList  = []string{"create_by", "update_by", "update_time", "delete_time"}
	ColumnNameNotQuery = []string{"create_by", "update_by", "create_time", "update_time", "delete_time", "remark"}
)

func (s *toolsGenTableColumn) GetDbType(columnType string) string {
	if strings.Index(columnType, "(") > 0 {
		return columnType[0:strings.Index(columnType, "(")]
	} else {
		return columnType
	}
}

func (s *toolsGenTableColumn) IsExistInArray(value string, array []string) bool {
	return strings.Contains(strings.Join(array, ","), value)
}

func (s *toolsGenTableColumn) GetColumnLength(columnType string) int {
	start := strings.LastIndex(columnType, "(")
	end := strings.LastIndex(columnType, ")")
	if start >= 0 && end >= 0 {
		result := columnType[start+1 : end]
		i, _ := strconv.Atoi(result)
		return i
	}
	return 0
}

func (s *toolsGenTableColumn) IsStringObject(dataType string) bool {
	return s.IsExistInArray(dataType, ColumnTypeStr)
}

func (s *toolsGenTableColumn) IsTimeObject(dataType string) bool {
	return s.IsExistInArray(dataType, ColumnTypeTime)
}

func (s *toolsGenTableColumn) IsNumberObject(dataType string) bool {
	return s.IsExistInArray(dataType, ColumnTypeNumber)
}

func (s *toolsGenTableColumn) IsNotEdit(name string) bool {
	if strings.Contains(name, "id") {
		return true
	}
	return s.IsExistInArray(name, ColumnNameNotEdit)
}

func (s *toolsGenTableColumn) IsNotList(name string) bool {
	return s.IsExistInArray(name, ColumnNameNotList)
}

func (s *toolsGenTableColumn) IsNotQuery(name string) bool {
	return s.IsExistInArray(name, ColumnNameNotQuery)
}

func (s *toolsGenTableColumn) CheckNameColumn(columnName string) bool {
	if len(columnName) >= 4 {
		tmp := columnName[len(columnName)-4:]
		return tmp == "name"
	}
	return false
}

func (s *toolsGenTableColumn) CheckStatusColumn(columnName string) bool {
	if len(columnName) >= 6 {
		tmp := columnName[len(columnName)-6:]
		return tmp == "status"

	}
	return false
}

func (s *toolsGenTableColumn) CheckTypeColumn(columnName string) bool {
	if len(columnName) >= 4 {
		tmp := columnName[len(columnName)-4:]
		return tmp == "type"
	}
	return false
}

func (s *toolsGenTableColumn) CheckSexColumn(columnName string) bool {
	if len(columnName) >= 3 {
		tmp := columnName[len(columnName)-3:]
		return tmp == "sex"
	}
	return false
}

type Generator struct {
	TableName string
}

func (g *Generator) Generate() (entity.DevGenTable, error) {
	var data entity.DevGenTable
	data.TableName = g.TableName
	tableNameList := strings.Split(g.TableName, "_")
	for i := 0; i < len(tableNameList); i++ {
		strStart := string([]byte(tableNameList[i])[:1])
		strEnd := string([]byte(tableNameList[i])[1:])
		data.ClassName += strings.ToUpper(strStart) + strEnd
		if i >= 1 {
			data.BusinessName += strings.ToLower(strStart) + strEnd
			data.FunctionName = strings.ToUpper(strStart) + strEnd
		}
	}
	data.PackageName = "system"
	data.TplCategory = "crud"
	data.ModuleName = strings.Replace(g.TableName, "_", "-", -1)

	dbColumn, err := services.DevTableColumnModelDao.FindDbTableColumnList(g.TableName)
	if err != nil {
		return data, err
	}
	data.TableComment = data.ClassName
	data.FunctionAuthor = "panda"

	var wg sync.WaitGroup
	columnChan := make(chan entity.DevGenTableColumn, len(dbColumn)) // 创建带缓冲的通道

	for x := 0; x < len(dbColumn); x++ {
		index := x
		wg.Add(1)
		go func(wg *sync.WaitGroup, y int) {
			defer wg.Done()
			var column entity.DevGenTableColumn
			column.ColumnComment = dbColumn[y].ColumnComment
			column.ColumnName = dbColumn[y].ColumnName
			column.ColumnType = dbColumn[y].DataType
			column.Sort = y + 1
			column.IsPk = "0"

			nameList := strings.Split(dbColumn[y].ColumnName, "_")
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

			if column.ColumnComment == "" {
				column.ColumnComment = column.GoField
			}

			dataType := strings.ToLower(column.ColumnType)
			if ToolsGenTableColumn.IsStringObject(dataType) {
				column.GoType = "string"
				columnLength := ToolsGenTableColumn.GetColumnLength(column.ColumnType)
				if columnLength >= 500 {
					column.HtmlType = "textarea"
				} else {
					column.HtmlType = "input"
				}
			} else if ToolsGenTableColumn.IsTimeObject(dataType) {
				column.GoType = "Time"
				column.HtmlType = "datetime"
			} else if ToolsGenTableColumn.IsNumberObject(dataType) {
				column.HtmlType = "input"
				t := ""
				if global.Conf.Server.DbType == "postgresql" {
					t = column.ColumnType
				} else {
					t, _ := utils.ReplaceString(`\(.+\)`, "", column.ColumnType)
					t = strings.Split(strings.TrimSpace(t), " ")[0]
					t = strings.ToLower(t)
				}
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

			if strings.Contains(dbColumn[y].ColumnKey, "PR") {
				column.IsPk = "1"
				data.PkColumn = dbColumn[y].ColumnName
				data.PkGoField = column.GoField
				data.PkGoType = column.GoType
				data.PkJsonField = column.JsonField
				if dbColumn[y].Extra == "auto_increment" {
					column.IsIncrement = "1"
				}
			}

			if ToolsGenTableColumn.IsNotEdit(column.ColumnName) {
				column.IsRequired = "0"
				column.IsInsert = "0"
			} else {
				column.IsRequired = "0"
				column.IsInsert = "1"
				if strings.Contains(column.ColumnName, "name") || strings.Contains(column.ColumnName, "status") {
					column.IsRequired = "1"
				}
			}

			if ToolsGenTableColumn.IsNotEdit(column.ColumnName) {
				column.IsEdit = "0"
			} else {
				if column.IsPk == "1" {
					column.IsEdit = "0"
				} else {
					column.IsEdit = "1"
				}
			}

			if ToolsGenTableColumn.IsNotList(column.ColumnName) {
				column.IsList = "0"
			} else {
				column.IsList = "1"
			}

			if ToolsGenTableColumn.IsNotQuery(column.ColumnName) {
				column.IsQuery = "0"
			} else {
				column.IsQuery = "1"
			}

			if ToolsGenTableColumn.CheckNameColumn(column.ColumnName) {
				column.QueryType = "LIKE"
			} else {
				column.QueryType = "EQ"
			}

			if ToolsGenTableColumn.CheckStatusColumn(column.ColumnName) {
				column.HtmlType = "radio"
			} else if ToolsGenTableColumn.CheckTypeColumn(column.ColumnName) || ToolsGenTableColumn.CheckSexColumn(column.ColumnName) {
				column.HtmlType = "select"
			}

			columnChan <- column // 将处理结果放入通道
		}(&wg, index)
	}
	wg.Wait()

	close(columnChan) // 关闭通道

	for column := range columnChan {
		data.Columns = append(data.Columns, column) // 将通道中的结果放入data.Columns
	}

	return data, nil
}

func Preview(tableId int64) map[string]interface{} {
	defer func() {
		if err := recover(); err != nil {
			global.Log.Error(err)
		}
	}()

	t1, err := template.ParseFiles("resource/template/go/entity.template")
	biz.ErrIsNil(err, "entity模版读取失败")

	t2, err := template.ParseFiles("resource/template/go/service.template")
	biz.ErrIsNil(err, "service模版读取失败")

	t3, err := template.ParseFiles("resource/template/go/api.template")
	biz.ErrIsNilAppendErr(err, "api模版读取失败！")

	t4, err := template.ParseFiles("resource/template/go/router.template")
	biz.ErrIsNil(err, "router模版读取失败")

	t5, err := template.ParseFiles("resource/template/js/api.template")
	biz.ErrIsNil(err, "js模版读取失败")

	t6, err := template.ParseFiles("resource/template/vue/list-vue.template")
	biz.ErrIsNil(err, "vue列表模版读取失败！")

	t7, err := template.ParseFiles("resource/template/vue/edit-vue.template")
	biz.ErrIsNil(err, "vue编辑模版读取失败！")

	tab, _ := services.DevGenTableModelDao.FindOne(entity.DevGenTable{TableId: tableId}, false)

	var b1 bytes.Buffer
	t1.Execute(&b1, tab)
	var b2 bytes.Buffer
	t2.Execute(&b2, tab)
	var b3 bytes.Buffer
	t3.Execute(&b3, tab)
	var b4 bytes.Buffer
	t4.Execute(&b4, tab)
	var b5 bytes.Buffer
	t5.Execute(&b5, tab)
	var b6 bytes.Buffer
	t6.Execute(&b6, tab)
	var b7 bytes.Buffer
	t7.Execute(&b7, tab)

	mp := make(map[string]interface{})
	mp["template/entity.template"] = b1.String()
	mp["template/services.template"] = b2.String()
	mp["template/api.template"] = b3.String()
	mp["template/router.template"] = b4.String()
	mp["template/jsApi.template"] = b5.String()
	mp["template/listVue.template"] = b6.String()
	mp["template/editVue.template"] = b7.String()
	return mp
}

func GenCode(tableId int64) {
	defer func() {
		if err := recover(); err != nil {
			global.Log.Error(err)
		}
	}()

	tab, err := services.DevGenTableModelDao.FindOne(entity.DevGenTable{TableId: tableId}, false)
	biz.ErrIsNil(err, "读取表信息失败！")
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
	t1.Execute(&b1, tab)
	var b2 bytes.Buffer
	t2.Execute(&b2, tab)
	var b3 bytes.Buffer
	t3.Execute(&b3, tab)
	var b4 bytes.Buffer
	t4.Execute(&b4, tab)
	var b5 bytes.Buffer
	t5.Execute(&b5, tab)
	var b6 bytes.Buffer
	t6.Execute(&b6, tab)
	var b7 bytes.Buffer
	t7.Execute(&b7, tab)

	kgo.KFile.WriteFile("./apps/"+tab.PackageName+"/entity/"+tab.TableName+".go", b1.Bytes())
	kgo.KFile.WriteFile("./apps/"+tab.PackageName+"/services/"+tab.TableName+".go", b2.Bytes())
	kgo.KFile.WriteFile("./apps/"+tab.PackageName+"/api/"+tab.TableName+".go", b3.Bytes())
	kgo.KFile.WriteFile("./apps/"+tab.PackageName+"/router/"+tab.TableName+".go", b4.Bytes())
	kgo.KFile.WriteFile(global.Conf.Gen.Frontpath+"/api/"+tab.PackageName+"/"+tab.BusinessName+".ts", b5.Bytes())
	kgo.KFile.WriteFile(global.Conf.Gen.Frontpath+"/views/"+tab.PackageName+"/"+tab.BusinessName+"/index.vue", b6.Bytes())
	kgo.KFile.WriteFile(global.Conf.Gen.Frontpath+"/views/"+tab.PackageName+"/"+tab.BusinessName+"/"+"component"+"/editModule.vue", b7.Bytes())
}

func GenConfigure(tableId, parentId int) {
	tab, err := services.DevGenTableModelDao.FindOne(entity.DevGenTable{TableId: int64(tableId)}, false)
	if err != nil {
		return
	}
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
	insert, _ := sysServices.SysMenuModelDao.Insert(menu)
	if err != nil {
		return
	}
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
