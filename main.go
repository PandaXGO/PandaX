package main

import (
	"context"
	"pandax/pkg/config"

	"github.com/XM-GO/PandaKit/logger"
	"github.com/XM-GO/PandaKit/restfulx"
	"github.com/XM-GO/PandaKit/starter"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"pandax/apps/job/jobs"
	"pandax/pkg/global"
	"pandax/pkg/initialize"
	"pandax/pkg/middleware"
	"syscall"
)

var (
	configFile string
)

var rootCmd = &cobra.Command{
	Use:   "panda is the main component in the panda.",
	Short: `panda is go gin frame`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if configFile != "" {
			global.Conf = config.InitConfig(configFile)
			global.Log = logger.InitLog(global.Conf.Log.File.GetFilename(), global.Conf.Log.Level)
			dbGorm := starter.DbGorm{Type: global.Conf.Server.DbType}
			if global.Conf.Server.DbType == "mysql" {
				dbGorm.Dsn = global.Conf.Mysql.Dsn()
				dbGorm.MaxIdleConns = global.Conf.Mysql.MaxIdleConns
				dbGorm.MaxOpenConns = global.Conf.Mysql.MaxOpenConns
			} else {
				dbGorm.Dsn = global.Conf.Postgresql.PgDsn()
				dbGorm.MaxIdleConns = global.Conf.Postgresql.MaxIdleConns
				dbGorm.MaxOpenConns = global.Conf.Postgresql.MaxOpenConns
			}
			global.Db = dbGorm.GormInit()
			initialize.InitTable()
		} else {
			global.Log.Panic("请配置config")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		restfulx.UseAfterHandlerInterceptor(middleware.OperationHandler)
		// 前置 函数
		restfulx.UseBeforeHandlerInterceptor(middleware.PermissionHandler)
		// 后置 函数
		restfulx.UseAfterHandlerInterceptor(middleware.LogHandler)
		go func() {
			// 启动系统调度任务
			jobs.InitJob()
			jobs.Setup()
		}()

		app := initialize.InitRouter()
		global.Log.Info("路由初始化完成")
		app.Start(context.TODO())
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
		<-stop

		if err := app.Stop(context.TODO()); err != nil {
			log.Fatalf("fatal app stop: %s", err)
			os.Exit(-3)
		}
	},
}

func init() {
	rootCmd.Flags().StringVar(&configFile, "config", getEnvStr("PANDA_CONFIG", "./config.yml"), "panda config file path.")
}

func getEnvStr(env string, defaultValue string) string {
	v := os.Getenv(env)
	if v == "" {
		return defaultValue
	}
	return v
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErrf("panda root cmd execute: %s", err)
		os.Exit(1)
	}
}
