package api

import (
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
	"strings"
)

type NoticeApi struct {
	OrganizationApp services.SysOrganizationModel
	NoticeApp       services.SysNoticeModel
}

// GetNoticeList 通知列表数据
func (p *NoticeApi) GetNoticeList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	noticeType := restfulx.QueryParam(rc, "noticeType")
	title := restfulx.QueryParam(rc, "title")

	// 获取组织的子组织id
	one := p.OrganizationApp.FindOne(rc.LoginAccount.OrganizationId)
	split := strings.Split(strings.Trim(one.OrganizationPath, "/"), "/")
	// 获取所有父组织id
	ids := utils.OrganizationPCIds(split, rc.LoginAccount.OrganizationId, true)
	notice := entity.SysNotice{NoticeType: noticeType, Title: title, OrganizationIds: ids}
	list, total := p.NoticeApp.FindListPage(pageNum, pageSize, notice)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

// InsertNotice 添加通知
func (p *NoticeApi) InsertNotice(rc *restfulx.ReqCtx) {
	var notice entity.SysNotice
	restfulx.BindJsonAndValid(rc, &notice)
	notice.UserName = rc.LoginAccount.UserName
	p.NoticeApp.Insert(notice)
}

// UpdateNotice 修改通知
func (p *NoticeApi) UpdateNotice(rc *restfulx.ReqCtx) {
	var notice entity.SysNotice
	restfulx.BindJsonAndValid(rc, &notice)

	p.NoticeApp.Update(notice)
}

// DeleteNotice 删除通知
func (p *NoticeApi) DeleteNotice(rc *restfulx.ReqCtx) {
	noticeId := restfulx.PathParam(rc, "noticeId")
	noticeIds := utils.IdsStrToIdsIntGroup(noticeId)
	p.NoticeApp.Delete(noticeIds)
}
