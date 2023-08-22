package api

import (
	"github.com/XM-GO/PandaKit/biz"
	email "github.com/XM-GO/PandaKit/mail"
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/utils"
	"pandax/apps/resource/api/from"
	"pandax/apps/resource/entity"
	"pandax/apps/resource/services"
)

/**
 * @Description 添加qq群467890197 交流学习
 * @Author 熊猫
 * @Date 2022/1/14 15:23
 **/

type ResEmailsApi struct {
	ResEmailsApp services.ResEmailsModel
}

// GetResEmailsList ResEmails列表数据
func (p *ResEmailsApi) GetResEmailsList(rc *restfulx.ReqCtx) {

	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	status := restfulx.QueryParam(rc, "status")
	category := restfulx.QueryParam(rc, "category")

	data := entity.ResEmail{Status: status, Category: category}
	list, total := p.ResEmailsApp.FindListPage(pageNum, pageSize, data)
	li := *list
	for i, data := range li {
		data.From = utils.DdmMail(data.From)
		data.Secret = utils.DdmPassword(data.Secret)
		li[i] = data
	}
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

// GetResEmails 获取ResEmails
func (p *ResEmailsApi) GetResEmails(rc *restfulx.ReqCtx) {
	mailId := restfulx.PathParamInt(rc, "mailId")
	p.ResEmailsApp.FindOne(int64(mailId))
}

// InsertResEmails 添加ResEmails
func (p *ResEmailsApi) InsertResEmails(rc *restfulx.ReqCtx) {
	var data entity.ResEmail
	restfulx.BindQuery(rc, &data)

	p.ResEmailsApp.Insert(data)
}

// UpdateResEmails 修改ResEmails
func (p *ResEmailsApi) UpdateResEmails(rc *restfulx.ReqCtx) {
	var data entity.ResEmail
	restfulx.BindQuery(rc, &data)
	if utils.ISDdmMail(data.From) {
		data.From = ""
	}
	if utils.IsDdmPassword(data.Secret) {
		data.Secret = ""
	}
	p.ResEmailsApp.Update(data)
}

// DeleteResEmails 删除ResEmails
func (p *ResEmailsApi) DeleteResEmails(rc *restfulx.ReqCtx) {
	mailId := restfulx.PathParam(rc, "mailId")
	mailIds := utils.IdsStrToIdsIntGroup(mailId)
	p.ResEmailsApp.Delete(mailIds)
}

// UpdateMailStatus 删除ResEmails
func (p *ResEmailsApi) UpdateMailStatus(rc *restfulx.ReqCtx) {
	var data entity.ResEmail
	restfulx.BindQuery(rc, &data)

	p.ResEmailsApp.Update(entity.ResEmail{MailId: data.MailId, Status: data.Status})
}

// DebugMail 测试发邮件
func (p *ResEmailsApi) DebugMail(rc *restfulx.ReqCtx) {
	var data from.SendMail
	restfulx.BindQuery(rc, &data)

	one := p.ResEmailsApp.FindOne(data.MailId)
	ml := email.Mail{
		Host:     one.Host,
		Port:     one.Port,
		From:     one.From,
		Nickname: one.Nickname,
		Secret:   one.Secret,
		IsSSL:    one.IsSSL,
	}
	biz.ErrIsNil(ml.Email(data.To, data.Subject, data.Body), "邮件发送失败")
}
