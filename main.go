package main

import (
	"pandax/base/config"
	"pandax/base/ctx"
	"pandax/base/global"
	"pandax/base/starter"
	"pandax/initialize"
)

func main() {
	global.Db = starter.GormInit(config.Conf.Server.DbType)
	initialize.InitTable()
	ctx.UseBeforeHandlerInterceptor(ctx.PermissionHandler)
	ctx.UseAfterHandlerInterceptor(ctx.LogHandler)
	starter.RunWebServer(initialize.InitRouter())
}
