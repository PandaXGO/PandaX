package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kakuilan/kgo"
	"net/http"
	"pandax/base/biz"
	"pandax/base/ginx"
	"pandax/base/ws"
	"runtime"
)

type System struct{}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func (s *System) ServerInfo(g *gin.Context) {
	osDic := make(map[string]any, 0)
	osDic["goOs"] = runtime.GOOS
	osDic["arch"] = runtime.GOARCH
	osDic["mem"] = runtime.MemProfileRate
	osDic["compiler"] = runtime.Compiler
	osDic["version"] = runtime.Version()
	osDic["numGoroutine"] = runtime.NumGoroutine()

	info := kgo.KOS.GetSystemInfo()
	diskDic := make(map[string]any, 0)
	diskDic["total"] = info.DiskTotal / GB
	diskDic["free"] = info.DiskFree / GB
	diskDic["used"] = info.DiskUsed / GB
	diskDic["progress"] = int64((float64(info.DiskUsed) / float64(info.DiskTotal)) * 100)

	memDic := make(map[string]any, 0)
	memDic["total"] = info.MemTotal / GB
	memDic["used"] = info.MemUsed / GB
	memDic["free"] = info.MemFree / GB
	memDic["progress"] = int64((float64(info.MemUsed) / float64(info.MemTotal)) * 100)

	cpuDic := make(map[string]any, 0)
	cpuDic["num"] = info.CpuNum
	cpuDic["used"] = fmt.Sprintf("%.2f", info.CpuUser*100)
	cpuDic["free"] = fmt.Sprintf("%.2f", info.CpuFree*100)

	g.JSON(http.StatusOK, gin.H{
		"code": 200,
		"os":   osDic,
		"mem":  memDic,
		"cpu":  cpuDic,
		"disk": diskDic,
	})
}

// 连接websocket
func (s *System) ConnectWs(g *gin.Context) {
	wsConn, err := ws.Upgrader.Upgrade(g.Writer, g.Request, nil)
	defer func() {
		if err := recover(); &err != nil {
			wsConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("websocket 失败： %v", err)))
			wsConn.Close()
		}
	}()

	if err != nil {
		panic(any(biz.NewBizErr("升级websocket失败")))
	}
	// 权限校验
	rc := ginx.NewReqCtx(g)
	if err = ginx.PermissionHandler(rc); err != nil {
		panic(any(biz.NewBizErr("没有权限")))
	}

	// 登录账号信息
	la := rc.LoginAccount
	ws.Put(uint64(la.UserId), wsConn)
}
