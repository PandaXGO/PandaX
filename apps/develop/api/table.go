package api

import (
	"pandax/apps/develop/api/vo"
	"pandax/apps/develop/entity"
	"pandax/apps/develop/gen"
	"pandax/apps/develop/services"
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
	"strings"
	"sync"
)

type GenTableApi struct {
	GenTableApp services.SysGenTableModel
}

// GetDBTableList 获取数据库表列表
func (g *GenTableApi) GetDBTableList(rc *restfulx.ReqCtx) {
	dbTables := entity.DBTables{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	dbTables.TableName = restfulx.QueryParam(rc, "tableName")

	list, total, err := g.GenTableApp.FindDbTablesListPage(pageNum, pageSize, dbTables)
	biz.ErrIsNil(err, "查询配置分页列表信息失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// GetTablePage 获取表页列表数据
func (g *GenTableApi) GetTablePage(rc *restfulx.ReqCtx) {
	devGenTable := entity.DevGenTable{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	devGenTable.TableName = restfulx.QueryParam(rc, "tableName")
	devGenTable.TableComment = restfulx.QueryParam(rc, "tableComment")
	devGenTable.RoleId = rc.LoginAccount.RoleId
	devGenTable.Owner = rc.LoginAccount.UserName

	list, total, err := g.GenTableApp.FindListPage(pageNum, pageSize, devGenTable)
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
	devGenTable := entity.DevGenTable{}
	devGenTable.TableId = int64(restfulx.PathParamInt(rc, "tableId"))
	devGenTable.RoleId = rc.LoginAccount.RoleId
	result, err := g.GenTableApp.FindOne(devGenTable, true)
	biz.ErrIsNil(err, "分页获取表信息失败")
	rc.ResData = vo.TableInfoVo{
		List: result.Columns,
		Info: *result,
	}
}

// GetTableInfoByName 获取表信息
func (g *GenTableApi) GetTableInfoByName(rc *restfulx.ReqCtx) {
	devGenTable := entity.DevGenTable{}
	devGenTable.TableName = restfulx.QueryParam(rc, "tableName")
	devGenTable.RoleId = rc.LoginAccount.RoleId
	result, err := g.GenTableApp.FindOne(devGenTable, true)
	biz.ErrIsNil(err, "分页获取表信息失败")
	rc.ResData = vo.TableInfoVo{
		List: result.Columns,
		Info: *result,
	}
}

// GetTableTree 获取树表信息
func (g *GenTableApi) GetTableTree(rc *restfulx.ReqCtx) {
	devGenTable := entity.DevGenTable{}
	devGenTable.RoleId = rc.LoginAccount.RoleId
	devGenTable.Owner = rc.LoginAccount.UserName
	tree, err := g.GenTableApp.FindTree(devGenTable)
	biz.ErrIsNil(err, "获取树表信息失败")
	rc.ResData = tree
}

// Insert 添加表结构
func (g *GenTableApi) Insert(rc *restfulx.ReqCtx) {
	tables := strings.Split(restfulx.QueryParam(rc, "tables"), ",")

	var wg sync.WaitGroup
	for _, table := range tables {
		wg.Add(1)
		go func(table string) {
			defer wg.Done()
			var tg = gen.Generator{
				TableName: table,
			}
			genTable, err := tg.Generate()
			if err != nil {
				biz.ErrIsNil(err, "创建表结构")
			}
			genTable.OrgId = rc.LoginAccount.OrganizationId
			genTable.Owner = rc.LoginAccount.UserName
			g.GenTableApp.Insert(genTable)
		}(table)
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
