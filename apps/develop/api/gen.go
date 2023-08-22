package api

import (
	"github.com/XM-GO/PandaKit/restfulx"
	"pandax/apps/develop/gen"
	"pandax/apps/develop/services"
)

type GenApi struct {
	GenTableApp services.SysGenTableModel
}

// Preview 代码视图
func (e *GenApi) Preview(rc *restfulx.ReqCtx) {
	tableId := restfulx.PathParamInt(rc, "tableId")
	rc.ResData = gen.Preview(int64(tableId))
}

// GenCode 代码生成
func (e *GenApi) GenCode(rc *restfulx.ReqCtx) {
	tableId := restfulx.PathParamInt(rc, "tableId")
	gen.GenCode(int64(tableId))
}

// GenConfigure 配置生成
func (e *GenApi) GenConfigure(rc *restfulx.ReqCtx) {
	tableId := restfulx.PathParamInt(rc, "tableId")
	menuId := restfulx.QueryInt(rc, "menuId", 0)
	gen.GenConfigure(tableId, menuId)
}
