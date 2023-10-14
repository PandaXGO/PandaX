package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"pandax/apps/device/entity"
	"pandax/apps/device/util"
	"pandax/pkg/global"
	"strings"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Websocket struct {
	Conn *websocket.Conn
}

func NewWebsocket(writer http.ResponseWriter, r *http.Request, header http.Header) (*Websocket, error) {
	ws, err := upGrader.Upgrade(writer, r, header)
	if err != nil {
		return nil, err
	}
	ws.SetCloseHandler(func(code int, text string) error {
		global.Log.Info(fmt.Sprintf("websocket 连接关闭,code: %d, text: %s", code, text))
		return ws.Close()
	})

	webs := &Websocket{Conn: ws}
	return webs, nil
}

// OnMessage 消息
// 发送消息消息类型 01:发送的设备数据  02:收到指令回复  03: 心跳回复
func OnMessage(ws *Websocket, msg string) {
	if msg != "" && strings.Index(msg, "ONLINE") != -1 {
		screenId := strings.Split(msg, "ONLINE")[0]
		AddWebSocketByScreenId(screenId, ws)
	}
	//画布离开
	if msg != "" && strings.Index(msg, "LEAVE") != -1 {
		RemoveWebSocket(strings.Split(msg, "LEAVE")[0])
	}
	//客户端传来了控制命令 格式   场景控制代码CONTROLCMD控制命令CONTROLCMD传感器id
	if msg != "" && strings.Index(msg, "CONTROLCMD") != -1 {
		split := strings.Split(msg, "CONTROLCMD")
		if len(split) < 2 {
			return
		}
		screenId, controlCMD := split[0], split[1]

		vtsa := new(entity.VisualTwinSendAttrs)
		//1. 获取设备孪生
		err := json.Unmarshal([]byte(controlCMD), vtsa)
		if err != nil {
			global.Log.Error("设备参数下发，孪生体参数解析失败", err)
			sendMessages("02", "下发失败,设备参数结构错误", screenId)
			return
		}
		//2. 根据设备下发属性更改
		err = util.BuildRunDeviceRpc(vtsa.TwinId, "single", map[string]interface{}{
			"method": "setAttributes",
			"params": vtsa.Attrs,
		})
		if err != nil {
			global.Log.Error("命令发送失败", err)
			sendMessages("02", "命令发送失败", screenId)
		} else {
			sendMessages("02", "命令发送成功", screenId)
		}
	}
	//心跳处理
	if msg != "" && strings.Index(msg, "HEARTCMD") != -1 {
		split := strings.Split(msg, "HEARTCMD")
		if len(split) < 1 {
			return
		}
		screenId := split[0]
		sendMessages("03", "心跳正常", screenId)
	}

}
func sendMessages(messageType, messageContent, screenId string) {
	msg := fmt.Sprintf(`{"MESSAGETYPE":"%s","MESSAGECONTENT": "%s"}`, messageType, messageContent)
	SendMessage(msg, screenId)
}
