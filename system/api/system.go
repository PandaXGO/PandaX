package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kakuilan/kgo"
	"net/http"
	"pandax/base/biz"
	"pandax/base/ctx"
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
	osDic := make(map[string]interface{}, 0)
	osDic["goOs"] = runtime.GOOS
	osDic["arch"] = runtime.GOARCH
	osDic["mem"] = runtime.MemProfileRate
	osDic["compiler"] = runtime.Compiler
	osDic["version"] = runtime.Version()
	osDic["numGoroutine"] = runtime.NumGoroutine()

	used, free, total := kgo.KOS.DiskUsage("/")
	diskDic := make(map[string]interface{}, 0)
	diskDic["total"] = total / GB
	diskDic["free"] = free / GB
	diskDic["used"] = used / GB

	used2, free2, total2 := kgo.KOS.MemoryUsage(true)
	memDic := make(map[string]interface{}, 0)
	memDic["total"] = total2 / GB
	memDic["used"] = used2 / GB
	memDic["free"] = free2 / GB

	cpuDic := make(map[string]interface{}, 0)
	used3, idle, total3 := kgo.KOS.CpuUsage()
	cpuDic["total"] = total3 / GB
	cpuDic["used"] = used3 / GB
	cpuDic["free"] = idle / GB

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
		if err := recover(); err != nil {
			wsConn.WriteMessage(websocket.TextMessage, []byte(err.(error).Error()))
			wsConn.Close()
		}
	}()

	if err != nil {
		panic(biz.NewBizErr("升级websocket失败"))
	}
	// 权限校验
	rc := ctx.NewReqCtxWithGin(g)
	if err = ctx.PermissionHandler(rc); err != nil {
		panic(biz.NewBizErr("没有权限"))
	}

	// 登录账号信息
	la := rc.LoginAccount
	ws.Put(uint64(la.UserId), wsConn)
}
