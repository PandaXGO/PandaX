package main

import (
	"context"
	"github.com/XM-GO/PandaKit/config"
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
			global.Log = logger.InitLog(global.Conf.Log)
			global.Db = starter.GormInit(global.Conf.Server.DbType)
			initialize.InitTable()
		} else {
			global.Log.Panic("请配置config")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		restfulx.UseAfterHandlerInterceptor(middleware.OperationHandler)
		// gin前置 函数
		restfulx.UseBeforeHandlerInterceptor(restfulx.PermissionHandler)
		// gin后置 函数
		restfulx.UseAfterHandlerInterceptor(restfulx.LogHandler)
		go func() {
			// 启动系统调度任务
			jobs.InitJob()
			jobs.Setup()
		}()

		app := initialize.InitRouter()

		app.Start(context.TODO())
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
		<-stop

		if err := app.Stop(context.TODO()); err != nil {
			log.Fatal("fatal app stop: %s", err)
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
