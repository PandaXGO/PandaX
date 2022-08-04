package api

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/utils"
	"pandax/apps/develop/api/vo"
	"pandax/apps/develop/entity"
	"pandax/apps/develop/gen"
	"pandax/apps/develop/services"
	"strings"
	"sync"
)

type GenTableApi struct {
	GenTableApp services.SysGenTableModel
}

// GetDBTableList 分页列表数据 / page list data
func (g *GenTableApi) GetDBTableList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	tableName := restfulx.QueryParam(rc, "tableName")

	list, total := g.GenTableApp.FindDbTablesListPage(pageNum, pageSize, entity.DBTables{TableName: tableName})
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetTablePage 分页列表数据
func (g *GenTableApi) GetTablePage(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	tableName := restfulx.QueryParam(rc, "tableName")
	tableComment := restfulx.QueryParam(rc, "tableComment")

	list, total := g.GenTableApp.FindListPage(pageNum, pageSize, entity.DevGenTable{TableName: tableName, TableComment: tableComment})
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetTableInfo 获取表信息
func (g *GenTableApi) GetTableInfo(rc *restfulx.ReqCtx) {
	tableId := restfulx.PathParamInt(rc, "tableId")
	result := g.GenTableApp.FindOne(entity.DevGenTable{TableId: int64(tableId)}, true)
	rc.ResData = vo.TableInfoVo{
		List: result.Columns,
		Info: *result,
	}
}

// GetTableInfoByName 获取表信息
func (g *GenTableApi) GetTableInfoByName(rc *restfulx.ReqCtx) {
	tableName := restfulx.QueryParam(rc, "tableName")
	result := g.GenTableApp.FindOne(entity.DevGenTable{TableName: tableName}, true)
	rc.ResData = vo.TableInfoVo{
		List: result.Columns,
		Info: *result,
	}
}

// GetTableTree 获取树表信息
func (g *GenTableApi) GetTableTree(rc *restfulx.ReqCtx) {
	rc.ResData = g.GenTableApp.FindTree(entity.DevGenTable{})
}

// Insert 添加表结构
func (g *GenTableApi) Insert(rc *restfulx.ReqCtx) {
	tablesList := strings.Split(restfulx.QueryParam(rc, "tables"), ",")

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

// Update 修改表结构
func (g *GenTableApi) Update(rc *restfulx.ReqCtx) {
	var data entity.DevGenTable
	restfulx.BindQuery(rc, &data)
	g.GenTableApp.Update(data)
}

// Delete 删除表结构
func (g *GenTableApi) Delete(rc *restfulx.ReqCtx) {
	tableIds := restfulx.PathParam(rc, "tableId")
	group := utils.IdsStrToIdsIntGroup(tableIds)
	g.GenTableApp.Delete(group)
}
