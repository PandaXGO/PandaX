package api

import (
	"pandax/apps/develop/gen"
	"pandax/apps/develop/services"
	"pandax/base/ctx"
	"pandax/base/ginx"
)

type GenApi struct {
	GenTableApp services.SysGenTableModel
}

// @Summary 代码视图
// @Description 获取JSON
// @Tags 工具 / 生成工具
// @Param tableId path int true "tableId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /develop/code/gen/preview/{tableId} [get]
// @Security X-TOKEN
func (e *GenApi) Preview(rc *ctx.ReqCtx) {
	tableId := ginx.PathParamInt(rc.GinCtx, "tableId")
	rc.ResData = gen.Preview(int64(tableId))
}

// @Summary 代码生成
// @Description 获取JSON
// @Tags 工具 / 生成工具
// @Param tableId path int true "tableId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /develop/code/gen/code/{tableId} [get]
// @Security X-TOKEN
func (e *GenApi) GenCode(rc *ctx.ReqCtx) {
	tableId := ginx.PathParamInt(rc.GinCtx, "tableId")
	gen.GenCode(int64(tableId))
}
