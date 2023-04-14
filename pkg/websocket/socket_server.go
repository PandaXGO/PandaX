package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
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
//发送消息消息类型 01:发送的设备数据  02:收到指令回复  03: 心跳回复
func OnMessage(message string) {
	log.Println(message)
	//画布离开
	if message != "" && strings.Index(message, "LEAVE") != -1 {
		RemoveWebSocket(strings.Split(message, "LEAVE")[0])
	}
	//客户端传来了控制命令 格式   场景控制代码CONTROLCMD控制命令CONTROLCMD传感器id
	if message != "" && strings.Index(message, "CONTROLCMD") != -1 {
		split := strings.Split(message, "CONTROLCMD")
		if len(split) < 2 {
			return
		}
		screenId, controlCMD := split[0], split[1] //指令cmd : {key: '', value: 3} k:v形式
		log.Println(screenId, controlCMD)
		//TODO 在这里编写代码，将命令发送到现场设备 这里已经拿到了 按钮命令和画布id
		//1. 根据组态Id查询设备Id，及设备模型
		//2. 根据设备下发CMD指令

		sendMessages("02", "'命令发送成功'", screenId)
	}
	//心跳处理
	if message != "" && strings.Index(message, "ping") != -1 {
		sendMessages("03", "'心跳正常'", "")
	}

}

func sendMessages(messageType, messageContent, screenId string) {
	msg := fmt.Sprintf(`{'MESSAGETYPE':'%s','MESSAGECONTENT':%s}`, messageType, messageContent)
	SendMessage(msg, screenId)
}
