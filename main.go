package main

import (
	"context"
	"pandax/apps/devops/services/k8s/Init"
	"pandax/apps/job/jobs"
	"pandax/base/config"
	"pandax/base/ctx"
	"pandax/base/global"
	"pandax/base/starter"
	"pandax/initialize"
	"pandax/middleware"
)

func main() {
	global.Db = starter.GormInit(config.Conf.Server.DbType)
	initialize.InitTable()
	Init.GetK8sClient(context.Background(), "")
	// gin后置 函数
	ctx.UseAfterHandlerInterceptor(middleware.OperationHandler)
	// gin前置 函数
	ctx.UseBeforeHandlerInterceptor(ctx.PermissionHandler)
	// gin后置 函数
	ctx.UseAfterHandlerInterceptor(ctx.LogHandler)
	go func() {
		// 启动系统调度任务
		jobs.InitJob()
		jobs.Setup()
	}()
	starter.RunWebServer(initialize.InitRouter())
}
