package main

import (
	"context"
	"github.com/PandaXGO/PandaKit/logger"
	"github.com/PandaXGO/PandaKit/rediscli"
	"github.com/PandaXGO/PandaKit/restfulx"
	"github.com/PandaXGO/PandaKit/starter"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"pandax/iothub"
	"pandax/pkg/cache"
	"pandax/pkg/config"
	"pandax/pkg/global"
	"pandax/pkg/initialize"
	"pandax/pkg/middleware"
	"pandax/pkg/tdengine"
	"syscall"
)

var (
	configFile string
)

var rootCmd = &cobra.Command{
	Use:   "pandax is the main component in the panda.",
	Short: `pandax is go go-restful frame`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if configFile != "" {
			global.Conf = config.InitConfig(configFile)
			global.Log = logger.InitLog(global.Conf.Log.File.GetFilename(), global.Conf.Log.Level)
			dbGorm := starter.DbGorm{Type: global.Conf.Server.DbType}
			if global.Conf.Server.DbType == "mysql" {
				dbGorm.Dsn = global.Conf.Mysql.Dsn()
				dbGorm.MaxIdleConns = global.Conf.Mysql.MaxIdleConns
				dbGorm.MaxOpenConns = global.Conf.Mysql.MaxOpenConns
			}
			global.Db = dbGorm.GormInit()
			global.Log.Infof("%s连接成功", global.Conf.Server.DbType)
			client, err := rediscli.NewRedisClient(global.Conf.Redis.Host, global.Conf.Redis.Password, global.Conf.Redis.Port, global.Conf.Redis.Db)
			if err != nil {
				global.Log.Panic("Redis连接错误")
			} else {
				global.Log.Info("Redis连接成功")
			}
			cache.RedisDb = client
			tDengine, err := tdengine.NewTdengine(global.Conf.Taos.Username, global.Conf.Taos.Password, global.Conf.Taos.Host, global.Conf.Taos.Database)
			if err != nil {
				global.Log.Panic("Tdengine连接错误")
			} else {
				global.Log.Info("Tdengine连接成功")
			}
			global.TdDb = tDengine
			initialize.InitTable()
			// 初始化事件监听
			go initialize.InitEvents()
		} else {
			global.Log.Panic("请配置config")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// 前置 函数
		restfulx.UseBeforeHandlerInterceptor(middleware.PermissionHandler)
		// 后置 函数
		restfulx.UseAfterHandlerInterceptor(middleware.LogHandler)
		restfulx.UseAfterHandlerInterceptor(middleware.OperationHandler)

		app := initialize.InitRouter()
		global.Log.Info("路由初始化完成")
		app.Start(context.TODO())
		//开启IOTHUB
		go iothub.InitIothub()
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
