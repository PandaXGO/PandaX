package services

import (
	"github.com/XM-GO/PandaKit/biz"
	"pandax/apps/resource/entity"
	"pandax/pkg/global"
)

/**
 * @Description 添加qq群467890197 交流学习
 * @Author 熊猫
 * @Date 2022/1/14 15:21
 **/

type (
	ResEmailsModel interface {
		Insert(data entity.ResEmail) *entity.ResEmail
		FindOne(mailId int64) *entity.ResEmail
		FindListPage(page, pageSize int, data entity.ResEmail) (*[]entity.ResEmail, int64)
		FindList(data entity.ResEmail) *[]entity.ResEmail
		Update(data entity.ResEmail) *entity.ResEmail
		Delete(mailIds []int64)
	}

	emailModelImpl struct {
		table string
	}
)

var ResEmailsModelDao ResEmailsModel = &emailModelImpl{
	table: `res_emails`,
}

func (m *emailModelImpl) Insert(data entity.ResEmail) *entity.ResEmail {
	err := global.Db.Table(m.table).Create(&data).Error
	biz.ErrIsNil(err, "添加ResEmails失败")
	return &data
}

func (m *emailModelImpl) FindOne(mailId int64) *entity.ResEmail {
	resData := new(entity.ResEmail)
	err := global.Db.Table(m.table).Where("mail_id = ?", mailId).First(resData).Error
	biz.ErrIsNil(err, "查询ResEmails失败")
	return resData
}

func (m *emailModelImpl) FindListPage(page, pageSize int, data entity.ResEmail) (*[]entity.ResEmail, int64) {
	list := make([]entity.ResEmail, 0)
	var total int64 = 0
	offset := pageSize * (page - 1)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.MailId != 0 {
		db = db.Where("mail_id = ?", data.MailId)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.Category != "" {
		db = db.Where("category = ?", data.Category)
	}
	db.Where("delete_time IS NULL")
	err := db.Count(&total).Error
	err = db.Order("create_time").Limit(pageSize).Offset(offset).Find(&list).Error
	biz.ErrIsNil(err, "查询ResEmails分页列表失败")
	return &list, total
}

func (m *emailModelImpl) FindList(data entity.ResEmail) *[]entity.ResEmail {
	list := make([]entity.ResEmail, 0)
	db := global.Db.Table(m.table)
	// 此处填写 where参数判断
	if data.MailId != 0 {
		db = db.Where("mail_id = ?", data.MailId)
	}
	if data.Status != "" {
		db = db.Where("status = ?", data.Status)
	}
	if data.Category != "" {
		db = db.Where("category = ?", data.Category)
	}
	db.Where("delete_time IS NULL")
	biz.ErrIsNil(db.Order("create_time").Find(&list).Error, "查询ResEmails列表失败")
	return &list
}

func (m *emailModelImpl) Update(data entity.ResEmail) *entity.ResEmail {
	biz.ErrIsNil(global.Db.Table(m.table).Updates(&data).Error, "修改ResEmails失败")
	return &data
}

func (m *emailModelImpl) Delete(mailIds []int64) {
	biz.ErrIsNil(global.Db.Table(m.table).Delete(&entity.ResEmail{}, "mail_id in (?)", mailIds).Error, "删除ResEmails失败")
}
