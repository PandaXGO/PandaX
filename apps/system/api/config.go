package api

import (
	entity "pandax/apps/system/entity"
	services "pandax/apps/system/services"
	"pandax/base/biz"
	"pandax/base/ginx"
	"pandax/base/utils"
)

type ConfigApi struct {
	ConfigApp services.SysConfigModel
}

// @Summary 配置列表数据
// @Description 获取JSON
// @Tags 配置
// @Param configName query string false "configName"
// @Param configKey query string false "configKey"
// @Param configType query string false "configType"
// @Param pageSize query int false "页条数"
// @Param pageNum query int false "页码"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/config [get]
// @Security
func (p *ConfigApi) GetConfigList(rc *ginx.ReqCtx) {
	pageNum := ginx.QueryInt(rc.GinCtx, "pageNum", 1)
	pageSize := ginx.QueryInt(rc.GinCtx, "pageSize", 10)
	configName := rc.GinCtx.Query("configName")
	configKey := rc.GinCtx.Query("configKey")
	configType := rc.GinCtx.Query("configType")
	config := entity.SysConfig{ConfigName: configName, ConfigKey: configKey, ConfigType: configType}
	list, total := p.ConfigApp.FindListPage(pageNum, pageSize, config)

	rc.ResData = map[string]any{
		"data":     list,
		"total":    total,
		"pageNum":  pageNum,
		"pageSize": pageSize,
	}
}

// @Summary 配置列表数据ByKey
// @Description 获取JSON
// @Tags 配置
// @Param configKey query string false "configKey"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/config/configKey [get]
// @Security
func (p *ConfigApi) GetConfigListByKey(rc *ginx.ReqCtx) {
	configKey := rc.GinCtx.Query("configKey")
	biz.IsTrue(configKey != "", "请传入配置Key")
	rc.ResData = p.ConfigApp.FindList(entity.SysConfig{ConfigKey: configKey})
}

// @Summary 获取配置
// @Description 获取JSON
// @Tags 配置
// @Param configId path int true "configId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Router /system/config/{configId} [get]
// @Security
func (p *ConfigApi) GetConfig(rc *ginx.ReqCtx) {
	id := ginx.PathParamInt(rc.GinCtx, "configId")
	p.ConfigApp.FindOne(int64(id))
}

// @Summary 添加配置
// @Description 获取JSON
// @Tags 配置
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysConfig true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /system/config [post]
// @Security X-TOKEN
func (p *ConfigApi) InsertConfig(rc *ginx.ReqCtx) {
	var config entity.SysConfig
	ginx.BindJsonAndValid(rc.GinCtx, &config)

	p.ConfigApp.Insert(config)
}

// @Summary 修改配置
// @Description 获取JSON
// @Tags 配置
// @Accept  application/json
// @Product application/json
// @Param data body entity.SysConfig true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "添加失败"}"
// @Router /system/config [put]
// @Security X-TOKEN
func (p *ConfigApi) UpdateConfig(rc *ginx.ReqCtx) {
	var post entity.SysConfig
	ginx.BindJsonAndValid(rc.GinCtx, &post)
	p.ConfigApp.Update(post)
}

// @Summary 删除配置
// @Description 删除数据
// @Tags 配置
// @Param configId path string true "configId 多个使用逗号分割"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": 400, "message": "删除失败"}"
// @Router /system/config/{configId} [delete]
func (p *ConfigApi) DeleteConfig(rc *ginx.ReqCtx) {
	configId := rc.GinCtx.Param("configId")
	p.ConfigApp.Delete(utils.IdsStrToIdsIntGroup(configId))
}
