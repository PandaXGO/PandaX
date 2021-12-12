package api

import (
	"pandax/base/biz"
	"pandax/base/ctx"
	"pandax/base/ginx"
	"pandax/system/entity"
	"pandax/system/services"
)

type SettingApi struct {
	SettingApp services.SysSettingsModel
}

// 获取信息
func (s *SettingApi) GetSettingsInfo(rc *ctx.ReqCtx) {
	key := rc.GinCtx.Query("key")
	biz.IsTrue(key != "", "请传入字典类型")
	rc.ResData = s.SettingApp.FindOne(key)
}

func (s *SettingApi) SetSettingsInfo(rc *ctx.ReqCtx) {
	var set entity.SysSettings
	ginx.BindJsonAndValid(rc.GinCtx, &set)
	one := s.SettingApp.FindOne(set.Key)
	if one.Id != 0 {
		s.SettingApp.Update(set)
	} else {
		s.SettingApp.Insert(set)
	}
}
