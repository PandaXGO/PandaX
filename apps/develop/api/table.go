package api

import (
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
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
	dbt := entity.DBTables{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	dbt.TableName = restfulx.QueryParam(rc, "tableName")

	list, total := g.GenTableApp.FindDbTablesListPage(pageNum, pageSize, dbt)
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// GetTablePage 分页列表数据
func (g *GenTableApi) GetTablePage(rc *restfulx.ReqCtx) {
	dgt := entity.DevGenTable{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	dgt.TableName = restfulx.QueryParam(rc, "tableName")
	dgt.TableComment = restfulx.QueryParam(rc, "tableComment")
	dgt.RoleId = rc.LoginAccount.RoleId
	dgt.Owner = rc.LoginAccount.UserName

	list, total := g.GenTableApp.FindListPage(pageNum, pageSize, dgt)
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// GetTableInfo 获取表信息
func (g *GenTableApi) GetTableInfo(rc *restfulx.ReqCtx) {
	dgt := entity.DevGenTable{}
	dgt.TableId = int64(restfulx.PathParamInt(rc, "tableId"))
	dgt.RoleId = rc.LoginAccount.RoleId
	result := g.GenTableApp.FindOne(dgt, true)
	rc.ResData = vo.TableInfoVo{
		List: result.Columns,
		Info: *result,
	}
}

// GetTableInfoByName 获取表信息
func (g *GenTableApi) GetTableInfoByName(rc *restfulx.ReqCtx) {
	dgt := entity.DevGenTable{}
	dgt.TableName = restfulx.QueryParam(rc, "tableName")
	dgt.RoleId = rc.LoginAccount.RoleId
	result := g.GenTableApp.FindOne(dgt, true)
	rc.ResData = vo.TableInfoVo{
		List: result.Columns,
		Info: *result,
	}
}

// GetTableTree 获取树表信息
func (g *GenTableApi) GetTableTree(rc *restfulx.ReqCtx) {
	dgt := entity.DevGenTable{}
	dgt.RoleId = rc.LoginAccount.RoleId
	dgt.Owner = rc.LoginAccount.UserName
	rc.ResData = g.GenTableApp.FindTree(dgt)
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
			genTable.OrgId = rc.LoginAccount.OrganizationId
			genTable.Owner = rc.LoginAccount.UserName
			g.GenTableApp.Insert(genTable)
			wg.Done()
		}(&wg, index)
	}
	wg.Wait()
}

// Update 修改表结构
func (g *GenTableApi) Update(rc *restfulx.ReqCtx) {
	var data entity.DevGenTable
	restfulx.BindJsonAndValid(rc, &data)
	g.GenTableApp.Update(data)
}

// Delete 删除表结构
func (g *GenTableApi) Delete(rc *restfulx.ReqCtx) {
	tableIds := restfulx.PathParam(rc, "tableId")
	group := utils.IdsStrToIdsIntGroup(tableIds)
	g.GenTableApp.Delete(group)
}
