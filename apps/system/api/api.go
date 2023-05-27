package api

import (
	"github.com/XM-GO/PandaKit/casbin"
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/utils"
	entity "pandax/apps/system/entity"
	services "pandax/apps/system/services"
	"pandax/pkg/global"
)

type SystemApiApi struct {
	ApiApp services.SysApiModel
}

func (s *SystemApiApi) CreateApi(rc *restfulx.ReqCtx) {
	var api entity.SysApi
	restfulx.BindQuery(rc, &api)
	s.ApiApp.Insert(api)
}

func (s *SystemApiApi) DeleteApi(rc *restfulx.ReqCtx) {
	ids := rc.Request.PathParameter("id")
	s.ApiApp.Delete(utils.IdsStrToIdsIntGroup(ids))
}

func (s *SystemApiApi) GetApiList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	path := rc.Request.QueryParameter("path")
	description := rc.Request.QueryParameter("description")
	method := rc.Request.QueryParameter("method")
	apiGroup := rc.Request.QueryParameter("apiGroup")
	api := entity.SysApi{Path: path, Description: description, Method: method, ApiGroup: apiGroup}
	list, total := s.ApiApp.FindListPage(pageNum, pageSize, api)
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

func (s *SystemApiApi) GetApiById(rc *restfulx.ReqCtx) {
	id := restfulx.QueryInt(rc, "id", 0)
	rc.ResData = s.ApiApp.FindOne(int64(id))

}

func (s *SystemApiApi) UpdateApi(rc *restfulx.ReqCtx) {
	var api entity.SysApi
	restfulx.BindQuery(rc, &api)
	s.ApiApp.Update(api)
}

func (s *SystemApiApi) GetAllApis(rc *restfulx.ReqCtx) {
	rc.ResData = s.ApiApp.FindList(entity.SysApi{})
}

func (s *SystemApiApi) GetPolicyPathByRoleId(rc *restfulx.ReqCtx) {
	roleKey := rc.Request.QueryParameter("roleKey")
	ca := casbin.CasbinS{ModelPath: global.Conf.Casbin.ModelPath}
	rc.ResData = ca.GetPolicyPathByRoleId(roleKey)
}
