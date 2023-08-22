package api

import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/utils"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
	"strings"
)

type NoticeApi struct {
	DeptApp   services.SysDeptModel
	NoticeApp services.SysNoticeModel
}

// GetNoticeList 通知列表数据
func (p *NoticeApi) GetNoticeList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	noticeType := restfulx.QueryParam(rc, "noticeType")
	title := restfulx.QueryParam(rc, "title")

	// 获取部门的子部门id
	one := p.DeptApp.FindOne(rc.LoginAccount.DeptId)
	split := strings.Split(strings.Trim(one.DeptPath, "/"), "/")
	// 获取所有父部门id
	ids := utils.DeptPCIds(split, rc.LoginAccount.DeptId, true)
	notice := entity.SysNotice{NoticeType: noticeType, Title: title, DeptIds: ids}
	list, total := p.NoticeApp.FindListPage(pageNum, pageSize, notice)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// InsertNotice 添加通知
func (p *NoticeApi) InsertNotice(rc *restfulx.ReqCtx) {
	var notice entity.SysNotice
	restfulx.BindQuery(rc, &notice)
	notice.UserName = rc.LoginAccount.UserName
	p.NoticeApp.Insert(notice)
}

// UpdateNotice 修改通知
func (p *NoticeApi) UpdateNotice(rc *restfulx.ReqCtx) {
	var notice entity.SysNotice
	restfulx.BindQuery(rc, &notice)

	p.NoticeApp.Update(notice)
}

// DeleteNotice 删除通知
func (p *NoticeApi) DeleteNotice(rc *restfulx.ReqCtx) {
	noticeId := restfulx.PathParam(rc, "noticeId")
	noticeIds := utils.IdsStrToIdsIntGroup(noticeId)
	p.NoticeApp.Delete(noticeIds)
}
