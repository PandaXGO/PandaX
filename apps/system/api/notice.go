package api

import (
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
	"pandax/base/ginx"
	"pandax/base/utils"
	"strings"
)

type NoticeApi struct {
	DeptApp   services.SysDeptModel
	NoticeApp services.SysNoticeModel
}

// @Summary 通知列表数据
// @Description 获取JSON
// @Tags 通知
// @Param noticeType query string false "noticeType"
// @Param title query string false "title"
// @Param pageSize query int false "页条数"
// @Param pageNum query int false "页码"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/post [get]
// @Security
func (p *NoticeApi) GetNoticeList(rc *ginx.ReqCtx) {
	pageNum := ginx.QueryInt(rc.GinCtx, "pageNum", 1)
	pageSize := ginx.QueryInt(rc.GinCtx, "pageSize", 10)
	noticeType := rc.GinCtx.Query("noticeType")
	title := rc.GinCtx.Query("title")

	// 获取部门的子部门id
	one := p.DeptApp.FindOne(rc.LoginAccount.DeptId)
	split := strings.Split(strings.Trim(one.DeptPath, "/"), "/")
	// 获取所有父部门id
	ids := utils.DeptPCIds(split, rc.LoginAccount.DeptId, true)
	notice := entity.SysNotice{NoticeType: noticeType, Title: title, DeptIds: ids}
	list, total := p.NoticeApp.FindListPage(pageNum, pageSize, notice)

	rc.ResData = map[string]any{
		"data":     list,
		"total":    total,
		"pageNum":  pageNum,
		"pageSize": pageSize,
	}
}

// @Summary 添加通知
// @Description 获取JSON
// @Tags 通知
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysNotice true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /system/notice [post]
// @Security X-TOKEN
func (p *NoticeApi) InsertNotice(rc *ginx.ReqCtx) {
	var notice entity.SysNotice
	ginx.BindJsonAndValid(rc.GinCtx, &notice)
	notice.UserName = rc.LoginAccount.UserName
	p.NoticeApp.Insert(notice)
}

// @Summary 修改通知
// @Description 获取JSON
// @Tags 通知
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysNotice true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /system/notice [put]
// @Security X-TOKEN
func (p *NoticeApi) UpdateNotice(rc *ginx.ReqCtx) {
	var notice entity.SysNotice
	ginx.BindJsonAndValid(rc.GinCtx, &notice)

	p.NoticeApp.Update(notice)
}

// @Summary 删除通知
// @Description 删除数据
// @Tags 通知
// @Param noticeId path string true "noticeId "
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /system/notice/{noticeId} [delete]
func (p *NoticeApi) DeleteNotice(rc *ginx.ReqCtx) {
	noticeId := rc.GinCtx.Param("noticeId")
	noticeIds := utils.IdsStrToIdsIntGroup(noticeId)
	p.NoticeApp.Delete(noticeIds)
}
