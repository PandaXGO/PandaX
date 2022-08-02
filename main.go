package main

import (
	"github.com/spf13/cobra"
	"os"
	"pandax/apps/job/jobs"
	"pandax/base/config"
	"pandax/base/ginx"
	"pandax/base/logger"
	"pandax/base/starter"
	"pandax/pkg/global"
	"pandax/pkg/initialize"
	"pandax/pkg/middleware"
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
		ginx.UseAfterHandlerInterceptor(middleware.OperationHandler)
		// gin前置 函数
		ginx.UseBeforeHandlerInterceptor(ginx.PermissionHandler)
		// gin后置 函数
		ginx.UseAfterHandlerInterceptor(ginx.LogHandler)
		go func() {
			// 启动系统调度任务
			jobs.InitJob()
			jobs.Setup()
		}()
		starter.RunWebServer(initialize.InitRouter())
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
