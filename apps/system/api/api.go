package api

import (
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/casbin"
	"github.com/PandaXGO/PandaKit/model"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/utils"
	entity "pandax/apps/system/entity"
	services "pandax/apps/system/services"
	"pandax/pkg/global"
)

type SystemApiApi struct {
	ApiApp services.SysApiModel
}

func (s *SystemApiApi) CreateApi(rc *restfulx.ReqCtx) {
	var api entity.SysApi
	restfulx.BindJsonAndValid(rc, &api)
	_, err := s.ApiApp.Insert(api)
	biz.ErrIsNil(err, "添加APi失败")
}

func (s *SystemApiApi) DeleteApi(rc *restfulx.ReqCtx) {
	ids := rc.Request.PathParameter("id")
	err := s.ApiApp.Delete(utils.IdsStrToIdsIntGroup(ids))
	biz.ErrIsNil(err, "删除APi失败")
}

func (s *SystemApiApi) GetApiList(rc *restfulx.ReqCtx) {
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)
	path := rc.Request.QueryParameter("path")
	description := rc.Request.QueryParameter("description")
	method := rc.Request.QueryParameter("method")
	apiGroup := rc.Request.QueryParameter("apiGroup")
	api := entity.SysApi{Path: path, Description: description, Method: method, ApiGroup: apiGroup}
	list, total, err := s.ApiApp.FindListPage(pageNum, pageSize, api)
	biz.ErrIsNil(err, "查询APi分页列表失败")
	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageSize),
		Data:     list,
	}
}

func (s *SystemApiApi) GetApiById(rc *restfulx.ReqCtx) {
	id := restfulx.QueryInt(rc, "id", 0)
	data, err := s.ApiApp.FindOne(int64(id))
	biz.ErrIsNil(err, "查询APi失败")
	rc.ResData = data

}

func (s *SystemApiApi) UpdateApi(rc *restfulx.ReqCtx) {
	var api entity.SysApi
	restfulx.BindJsonAndValid(rc, &api)
	err := s.ApiApp.Update(api)
	biz.ErrIsNil(err, "查询APi失败")
}

func (s *SystemApiApi) GetAllApis(rc *restfulx.ReqCtx) {
	data, err := s.ApiApp.FindList(entity.SysApi{})
	biz.ErrIsNil(err, "查询APi列表失败")
	rc.ResData = data
}

func (s *SystemApiApi) GetPolicyPathByRoleId(rc *restfulx.ReqCtx) {
	roleKey := rc.Request.QueryParameter("roleKey")
	ca := casbin.CasbinService{ModelPath: global.Conf.Casbin.ModelPath}
	rc.ResData = ca.GetPolicyPathByRoleId(roleKey)
}
