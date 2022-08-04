package api

/**
 * @Description
 * @Author 熊猫
 * @Date 2022/7/14 17:55
 **/
import (
	"github.com/XM-GO/PandaKit/model"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/utils"
	"pandax/apps/system/entity"
	"pandax/apps/system/services"
)

type SysTenantsApi struct {
	SysTenantsApp services.SysTenantsModel
}

func (p *SysTenantsApi) GetSysTenantsList(rc *restfulx.ReqCtx) {
	data := entity.SysTenants{}
	pageNum := restfulx.QueryInt(rc, "pageNum", 1)
	pageSize := restfulx.QueryInt(rc, "pageSize", 10)

	list, total := p.SysTenantsApp.FindListPage(pageNum, pageSize, data)

	rc.ResData = model.ResultPage{
		Total:    total,
		PageNum:  int64(pageNum),
		PageSize: int64(pageNum),
		Data:     list,
	}
}

func (p *SysTenantsApi) GetSysTenantsAll(rc *restfulx.ReqCtx) {
	list := make([]entity.SysTenants, 0)
	if rc.LoginAccount.RoleKey == "admin" {
		data := entity.SysTenants{}
		list = *p.SysTenantsApp.FindList(data)
	} else {
		list = append(list, *p.SysTenantsApp.FindOne(rc.LoginAccount.TenantId))
	}
	rc.ResData = list
}

func (p *SysTenantsApi) GetSysTenants(rc *restfulx.ReqCtx) {
	tenantId := restfulx.PathParamInt(rc, "tenantId")
	p.SysTenantsApp.FindOne(int64(tenantId))
}

func (p *SysTenantsApi) InsertSysTenants(rc *restfulx.ReqCtx) {
	var data entity.SysTenants
	restfulx.BindQuery(rc, &data)

	p.SysTenantsApp.Insert(data)
}

func (p *SysTenantsApi) UpdateSysTenants(rc *restfulx.ReqCtx) {
	var data entity.SysTenants
	restfulx.BindQuery(rc, &data)

	p.SysTenantsApp.Update(data)
}

func (p *SysTenantsApi) DeleteSysTenants(rc *restfulx.ReqCtx) {
	tenantId := rc.Request.PathParameter("tenantId")
	tenantIds := utils.IdsStrToIdsIntGroup(tenantId)
	p.SysTenantsApp.Delete(tenantIds)
}

// IsTenantAdmin 是否为主租户
func IsTenantAdmin(tenantId int64) bool {
	if tenantId == 1 {
		return true
	}
	return false
}
