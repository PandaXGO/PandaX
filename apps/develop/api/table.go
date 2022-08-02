package api

import (
	"pandax/apps/develop/entity"
	"pandax/apps/develop/gen"
	"pandax/apps/develop/services"
	"pandax/base/ginx"
	"pandax/base/utils"
	"strings"
	"sync"
)

type GenTableApi struct {
	GenTableApp services.SysGenTableModel
}

// @Summary 分页列表数据 / page list data
// @Description 数据库表列分页列表 / database table column page list
// @Tags 工具 / 生成工具
// @Param tableName query string false "tableName / 数据表名称"
// @Param pageSize query int false "pageSize / 页条数"
// @Param pageNum query int false "pageNum / 页码"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /develop/code/table/db/list [get]
func (g *GenTableApi) GetDBTableList(rc *ginx.ReqCtx) {
	pageNum := ginx.QueryInt(rc.GinCtx, "pageNum", 1)
	pageSize := ginx.QueryInt(rc.GinCtx, "pageSize", 10)
	tableName := rc.GinCtx.Query("tableName")
	list, total := g.GenTableApp.FindDbTablesListPage(pageNum, pageSize, entity.DBTables{TableName: tableName})
	rc.ResData = map[string]any{
		"data":     list,
		"total":    total,
		"pageNum":  pageNum,
		"pageSize": pageSize,
	}
}

// @Summary 分页列表数据
// @Description 生成表分页列表
// @Tags 工具 / 生成工具
// @Param tableName query string false "tableName / 数据表名称"
// @Param tableComment query string false "tableComment / 数据表描述"
// @Param pageSize query int false "pageSize / 页条数"
// @Param pageIndex query int false "pageIndex / 页码"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /develop/code/table/list [get]
func (g *GenTableApi) GetTablePage(rc *ginx.ReqCtx) {
	pageNum := ginx.QueryInt(rc.GinCtx, "pageNum", 1)
	pageSize := ginx.QueryInt(rc.GinCtx, "pageSize", 10)
	tableName := rc.GinCtx.Query("tableName")
	tableComment := rc.GinCtx.Query("tableComment")
	list, total := g.GenTableApp.FindListPage(pageNum, pageSize, entity.DevGenTable{TableName: tableName, TableComment: tableComment})
	rc.ResData = map[string]any{
		"data":     list,
		"total":    total,
		"pageNum":  pageNum,
		"pageSize": pageSize,
	}
}

// @Summary 获取表信息
// @Description 获取JSON
// @Tags 工具 / 生成工具
// @Param tableId path int true "tableId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /develop/code/table/info/{tableId} [get]
// @Security Bearer
func (g *GenTableApi) GetTableInfo(rc *ginx.ReqCtx) {
	tableId := ginx.PathParamInt(rc.GinCtx, "tableId")
	result := g.GenTableApp.FindOne(entity.DevGenTable{TableId: int64(tableId)}, true)
	mp := make(map[string]any)
	mp["list"] = result.Columns
	mp["info"] = result
	rc.ResData = mp
}

// @Summary 获取表信息
// @Description 获取JSON
// @Tags 工具 / 生成工具
// @Param tableName query string false "tableName / 数据表名称"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /develop/code/table/info/tableName [get]
// @Security X-TOKEN
func (g *GenTableApi) GetTableInfoByName(rc *ginx.ReqCtx) {
	tableName := rc.GinCtx.Query("tableName")
	result := g.GenTableApp.FindOne(entity.DevGenTable{TableName: tableName}, true)
	mp := make(map[string]any)
	mp["list"] = result.Columns
	mp["info"] = result
	rc.ResData = mp
}

// @Summary 获取树表信息
// @Description 获取JSON
// @Tags 工具 / 生成工具
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /develop/code/table/tableTree [get]
// @Security X-TOKEN
func (g *GenTableApi) GetTableTree(rc *ginx.ReqCtx) {
	rc.ResData = g.GenTableApp.FindTree(entity.DevGenTable{})
}

// @Summary 添加表结构
// @Description 添加表结构
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Param tableName query string false "tableName / 数据表名称"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /develop/code/table [post]
// @Security Bearer
func (g *GenTableApi) Insert(rc *ginx.ReqCtx) {
	tablesList := strings.Split(rc.GinCtx.Query("tables"), ",")

	wg := sync.WaitGroup{}
	for i := 0; i < len(tablesList); i++ {
		index := i
		wg.Add(1)
		go func(wg *sync.WaitGroup, index int) {
			genTable := gen.ToolsGenTableColumn.GenTableInit(tablesList[index])
			g.GenTableApp.Insert(genTable)
			wg.Done()
		}(&wg, index)
	}
	wg.Wait()
}

// @Summary 修改表结构
// @Description 修改表结构
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Param data body entity.DevGenTable true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /develop/code/table [put]
// @Security Bearer
func (g *GenTableApi) Update(rc *ginx.ReqCtx) {
	var data entity.DevGenTable
	ginx.BindJsonAndValid(rc.GinCtx, &data)
	g.GenTableApp.Update(data)
}

// Delete
// @Summary 删除表结构
// @Description 删除表结构
// @Tags 工具 / 生成工具
// @Param tableId path int true "tableId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /develop/code/table/{tableId} [delete]
func (g *GenTableApi) Delete(rc *ginx.ReqCtx) {
	tableIds := rc.GinCtx.Param("tableId")
	group := utils.IdsStrToIdsIntGroup(tableIds)
	g.GenTableApp.Delete(group)
}
