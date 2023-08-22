package middleware

import (
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/utils"
	"net/http"
	"pandax/apps/log/entity"
	"pandax/apps/log/services"
)

func OperationHandler(rc *restfulx.ReqCtx) error {
	c := rc.Request
	// 请求操作不做记录
	if c.Request.Method == http.MethodGet || rc.LoginAccount == nil {
		return nil
	}
	if rc.RequiredPermission == nil || !rc.RequiredPermission.NeedToken {
		return nil
	}
	go func() {
		oper := entity.LogOper{
			Title:        rc.LogInfo.Description,
			BusinessType: "0",
			Method:       c.Request.Method,
			OperName:     rc.LoginAccount.UserName,
			OperUrl:      c.Request.URL.Path,
			OperIp:       c.Request.RemoteAddr,
			OperLocation: utils.GetRealAddressByIP(c.Request.RemoteAddr),
			OperParam:    "",
			Status:       "0",
		}
		if c.Request.Method == "POST" {
			oper.BusinessType = "1"
		} else if c.Request.Method == "PUT" {
			oper.BusinessType = "2"
		} else if c.Request.Method == "DELETE" {
			oper.BusinessType = "3"
		}
		services.LogOperModelDao.Insert(oper)
	}()

	return nil
}
