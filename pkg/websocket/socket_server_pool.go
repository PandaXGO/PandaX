package websocket

import (
	"github.com/gorilla/websocket"
	"pandax/pkg/global"
)

var Wsp = make(map[string]*Websocket)

// GetWebSocketByScreenId 获取WebSocket
func GetWebSocketByScreenId(screenId string) *Websocket {
	return Wsp[screenId]
}

// AddWebSocketByScreenId 添加WebSocket
func AddWebSocketByScreenId(screenId string, webs *Websocket) {
	Wsp[screenId] = webs
}

// RemoveWebSocket 移除WebSocket
func RemoveWebSocket(screenId string) bool {
	if ws, ok := Wsp[screenId]; ok {
		ws.Conn.Close()
		delete(Wsp, screenId)
		global.Log.Info("已经断开websocket：" + screenId)
		return true
	}
	return false
}

// SendMessage 向特定场景id发送消息，同一场景代码有可能在多台客户机上连接 ，这时就会在多台客户机接受到了数据
func SendMessage(message, screenId string) {
	ws := GetWebSocketByScreenId(screenId)
	if ws != nil {
		ws.Conn.WriteMessage(websocket.TextMessage, []byte(message))
	}
}
