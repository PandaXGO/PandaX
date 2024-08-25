package api

import (
	"github.com/PandaXGO/PandaKit/biz"
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
	one, err := p.OrganizationApp.FindOne(rc.LoginAccount.OrganizationId)
	biz.ErrIsNil(err, "组织查询失败")
	split := strings.Split(strings.Trim(one.OrganizationPath, "/"), "/")
	// 获取所有父组织id
	ids := utils.OrganizationPCIds(split, rc.LoginAccount.OrganizationId, true)
	notice := entity.SysNotice{NoticeType: noticeType, Title: title, OrganizationIds: ids}
	list, total, err := p.NoticeApp.FindListPage(pageNum, pageSize, notice)
	biz.ErrIsNil(err, "查询通知列表失败")
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
	_, err := p.NoticeApp.Insert(notice)
	biz.ErrIsNil(err, "添加通知失败")
}

// UpdateNotice 修改通知
func (p *NoticeApi) UpdateNotice(rc *restfulx.ReqCtx) {
	var notice entity.SysNotice
	restfulx.BindJsonAndValid(rc, &notice)

	err := p.NoticeApp.Update(notice)
	biz.ErrIsNil(err, "修改通知失败")
}

// DeleteNotice 删除通知
func (p *NoticeApi) DeleteNotice(rc *restfulx.ReqCtx) {
	noticeId := restfulx.PathParam(rc, "noticeId")
	noticeIds := utils.IdsStrToIdsIntGroup(noticeId)
	err := p.NoticeApp.Delete(noticeIds)
	biz.ErrIsNil(err, "删除通知失败")
}
