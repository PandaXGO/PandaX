package api

import (
	"github.com/XM-GO/PandaKit/biz"
	email "github.com/XM-GO/PandaKit/mail"
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

// @Summary ResEmails列表数据
// @Tags ResEmails
// @Param pageSize query int false "页条数"
// @Param pageNum query int false "页码"
// @Param status query string false "状态"
// @Param category query string false "分类"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /resource/email/list [get]
// @Security
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
	rc.ResData = map[string]any{
		"data":     list,
		"total":    total,
		"pageNum":  pageNum,
		"pageSize": pageSize,
	}
}

// @Summary 获取ResEmails
// @Description 获取JSON
// @Tags ResEmails
// @Param mailId path int true "mailId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /resource/email/{mailId} [get]
// @Security
func (p *ResEmailsApi) GetResEmails(rc *restfulx.ReqCtx) {
	mailId := restfulx.PathParamInt(rc, "mailId")
	p.ResEmailsApp.FindOne(int64(mailId))
}

// @Summary 添加ResEmails
// @Description 获取JSON
// @Tags ResEmails
// @Accept  application/json
// @Product application/json
// @Param data body entity.ResEmail true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /resource/email [post]
// @Security X-TOKEN
func (p *ResEmailsApi) InsertResEmails(rc *restfulx.ReqCtx) {
	var data entity.ResEmail
	restfulx.BindQuery(rc, &data)

	p.ResEmailsApp.Insert(data)
}

// @Summary 修改ResEmails
// @Description 获取JSON
// @Tags ResEmails
// @Accept  application/json
// @Product application/json
// @Param data body entity.ResEmail true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /resource/email [put]
// @Security X-TOKEN
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

// @Summary 删除ResEmails
// @Description 删除数据
// @Tags ResEmails
// @Param mailId path string true "mailId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /resource/email/{mailId} [delete]
func (p *ResEmailsApi) DeleteResEmails(rc *restfulx.ReqCtx) {
	mailId := restfulx.PathParam(rc, "mailId")
	mailIds := utils.IdsStrToIdsIntGroup(mailId)
	p.ResEmailsApp.Delete(mailIds)
}

// @Summary 删除ResEmails
// @Description 获取JSON
// @Tags ResOsses
// @Accept  application/json
// @Product application/json
// @Param data body entity.ResEmail true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /resource/oss [put]
// @Security X-TOKEN
func (p *ResEmailsApi) UpdateMailStatus(rc *restfulx.ReqCtx) {
	var data entity.ResEmail
	restfulx.BindQuery(rc, &data)

	p.ResEmailsApp.Update(entity.ResEmail{MailId: data.MailId, Status: data.Status})
}

// @Summary 测试发邮件
// @Description 获取JSON
// @Tags ResEmails
// @Accept  application/json
// @Product application/json
// @Param data body from.SendMail true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /resource/email/debugMail [post]
// @Security X-TOKEN
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
