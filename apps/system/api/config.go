package api

import (
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
	entity "pandax/apps/system/entity"
	services "pandax/apps/system/services"
)

type ConfigApi struct {
	ConfigApp services.SysConfigModel
}

func (p *ConfigApi) GetConfigList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	configName := rc.Request.QueryParameter("configName")
	configKey := rc.Request.QueryParameter("configKey")
	configType := rc.Request.QueryParameter("configType")
	config := entity.SysConfig{ConfigName: configName, ConfigKey: configKey, ConfigType: configType}
	list, total, err := p.ConfigApp.FindListPage(pageNum, pageSize, config)
	biz.ErrIsNil(err, "查询配置分页列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

func (p *ConfigApi) GetConfigListByKey(rc *restfulx.ReqCtx) {
	configKey := rc.Request.QueryParameter("configKey")
	biz.IsTrue(configKey != "", "请传入配置Key")
	data, err := p.ConfigApp.FindList(entity.SysConfig{ConfigKey: configKey})
	biz.ErrIsNil(err, "查询配置列表失败")
	rc.ResData = data
}

func (p *ConfigApi) GetConfig(rc *restfulx.ReqCtx) {
	id := restfulx.PathParamInt(rc, "configId")
	data, err := p.ConfigApp.FindOne(int64(id))
	biz.ErrIsNil(err, "查询配置失败")
	rc.ResData = data
}

func (p *ConfigApi) InsertConfig(rc *restfulx.ReqCtx) {
	var config entity.SysConfig
	restfulx.BindJsonAndValid(rc, &config)

	_, err := p.ConfigApp.Insert(config)
	biz.ErrIsNil(err, "添加配置失败")
}

func (p *ConfigApi) UpdateConfig(rc *restfulx.ReqCtx) {
	var post entity.SysConfig
	restfulx.BindJsonAndValid(rc, &post)
	err := p.ConfigApp.Update(post)
	biz.ErrIsNil(err, "修改配置失败")
}

func (p *ConfigApi) DeleteConfig(rc *restfulx.ReqCtx) {
	configId := rc.Request.PathParameter("configId")
	err := p.ConfigApp.Delete(utils.IdsStrToIdsIntGroup(configId))
	biz.ErrIsNil(err, "删除配置失败")
}
