package api

import (
	"pandax/apps/resource/entity"
	"pandax/apps/resource/services"
	"pandax/base/ctx"
	"pandax/base/ginx"
	"pandax/base/utils"
)

/**
 * @Description 添加qq群467890197 交流学习
 * @Author 熊猫
 * @Date 2022/1/13 15:19
 **/

type ResOssesApi struct {
	ResOssesApp services.ResOssesModel
}

// @Summary ResOsses列表数据
// @Tags ResOsses
// @Param pageSize query int false "页条数"
// @Param pageNum query int false "页码"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /admin/resOsses/list [get]
// @Security
func (p *ResOssesApi) GetResOssesList(rc *ctx.ReqCtx) {

	pageNum := ginx.QueryInt(rc.GinCtx, "pageNum", 1)
	pageSize := ginx.QueryInt(rc.GinCtx, "pageSize", 10)
	status := rc.GinCtx.Query("status")
	category := rc.GinCtx.Query("category")
	ossCode := rc.GinCtx.Query("ossCode")

	data := entity.ResOss{Status: status, Category: category, OssCode: ossCode}
	list, total := p.ResOssesApp.FindListPage(pageNum, pageSize, data)
	li := *list
	for i, data := range li {
		data.AccessKey = utils.Ddm(data.AccessKey)
		data.SecretKey = utils.Ddm(data.SecretKey)
		li[i] = data
	}
	rc.ResData = map[string]interface{}{
		"data":     list,
		"total":    total,
		"pageNum":  pageNum,
		"pageSize": pageSize,
	}
}

// @Summary 获取ResOsses
// @Description 获取JSON
// @Tags ResOsses
// @Param ossId path int true "ossId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /admin/resOsses/{ossId }[get]
// @Security
func (p *ResOssesApi) GetResOsses(rc *ctx.ReqCtx) {
	ossId := ginx.PathParamInt(rc.GinCtx, "ossId")
	p.ResOssesApp.FindOne(int64(ossId))
}

// @Summary 添加ResOsses
// @Description 获取JSON
// @Tags ResOsses
// @Accept  application/json
// @Product application/json
// @Param data body entity.ResOsses true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /admin/resOsses [post]
// @Security X-TOKEN
func (p *ResOssesApi) InsertResOsses(rc *ctx.ReqCtx) {
	var data entity.ResOss
	ginx.BindJsonAndValid(rc.GinCtx, &data)

	p.ResOssesApp.Insert(data)
}

// @Summary 修改ResOsses
// @Description 获取JSON
// @Tags ResOsses
// @Accept  application/json
// @Product application/json
// @Param data body entity.ResOsses true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /admin/resOsses [put]
// @Security X-TOKEN
func (p *ResOssesApi) UpdateResOsses(rc *ctx.ReqCtx) {
	var data entity.ResOss
	ginx.BindJsonAndValid(rc.GinCtx, &data)

	p.ResOssesApp.Update(data)
}

// @Summary 删除ResOsses
// @Description 删除数据
// @Tags ResOsses
// @Param ossId path string true "ossId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /admin/resOsses/{ossId } [delete]
func (p *ResOssesApi) DeleteResOsses(rc *ctx.ReqCtx) {

	ossId := rc.GinCtx.Param("ossId")
	ossIds := utils.IdsStrToIdsIntGroup(ossId)
	p.ResOssesApp.Delete(ossIds)
}
