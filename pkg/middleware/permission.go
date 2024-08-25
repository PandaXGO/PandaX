package middleware

import (
	"github.com/PandaXGO/PandaKit/biz"
	"github.com/PandaXGO/PandaKit/casbin"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/token"
	"pandax/pkg/global"

	"github.com/dgrijalva/jwt-go"
)

func PermissionHandler(rc *restfulx.ReqCtx) error {
	permission := rc.RequiredPermission
	// 如果需要的权限信息不为空，并且不需要token，则不返回错误，继续后续逻辑
	if permission != nil && !permission.NeedToken {
		return nil
	}
	tokenStr := rc.Request.Request.Header.Get("X-TOKEN")
	// header不存在则从查询参数token中获取
	if tokenStr == "" {
		tokenStr = rc.Request.QueryParameter("token")
	}
	if tokenStr == "" {
		return biz.PermissionErr
	}
	j := token.NewJWT("", []byte(global.Conf.Jwt.Key), jwt.SigningMethodHS256)
	loginAccount, err := j.ParseToken(tokenStr)
	if err != nil || loginAccount == nil {
		return biz.PermissionErr
	}
	rc.LoginAccount = loginAccount

	if !permission.NeedCasbin {
		return nil
	}

	ca := casbin.CasbinService{ModelPath: global.Conf.Casbin.ModelPath}
	e := ca.GetCasbinEnforcer()
	// 判断策略中是否存在
	success, err := e.Enforce(loginAccount.RoleKey, rc.Request.Request.URL.Path, rc.Request.Request.Method)
	if !success || err != nil {
		return biz.CasbinErr
	}

	return nil
}
