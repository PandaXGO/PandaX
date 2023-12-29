package api

import (
	"github.com/PandaXGO/PandaKit/biz"
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

	list, total, err := g.GenTableApp.FindDbTablesListPage(pageNum, pageSize, dbt)
	biz.ErrIsNil(err, "查询配置分页列表信息失败")
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

	list, total, err := g.GenTableApp.FindListPage(pageNum, pageSize, dgt)
	biz.ErrIsNil(err, "分页获取表失败")
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
	result, err := g.GenTableApp.FindOne(dgt, true)
	biz.ErrIsNil(err, "分页获取表信息失败")
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
	result, err := g.GenTableApp.FindOne(dgt, true)
	biz.ErrIsNil(err, "分页获取表信息失败")
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
	tree, err := g.GenTableApp.FindTree(dgt)
	biz.ErrIsNil(err, "获取树表信息失败")
	rc.ResData = tree
}

// Insert 添加表结构
func (g *GenTableApi) Insert(rc *restfulx.ReqCtx) {
	tablesList := strings.Split(restfulx.QueryParam(rc, "tables"), ",")

	wg := sync.WaitGroup{}
	for i := 0; i < len(tablesList); i++ {
		index := i
		wg.Add(1)
		go func(wg *sync.WaitGroup, index int) {
			defer wg.Done()
			genTable, err := gen.ToolsGenTableColumn.GenTableInit(tablesList[index])
			if err != nil {
				return
			}
			genTable.OrgId = rc.LoginAccount.OrganizationId
			genTable.Owner = rc.LoginAccount.UserName
			g.GenTableApp.Insert(genTable)
		}(&wg, index)
	}
	wg.Wait()
}

// Update 修改表结构
func (g *GenTableApi) Update(rc *restfulx.ReqCtx) {
	var data entity.DevGenTable
	restfulx.BindJsonAndValid(rc, &data)
	_, err := g.GenTableApp.Update(data)
	biz.ErrIsNil(err, "修改表结构")
}

// Delete 删除表结构
func (g *GenTableApi) Delete(rc *restfulx.ReqCtx) {
	tableIds := restfulx.PathParam(rc, "tableId")
	group := utils.IdsStrToIdsIntGroup(tableIds)
	err := g.GenTableApp.Delete(group)
	biz.ErrIsNil(err, "删除表结构")
}
