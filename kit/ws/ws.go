package ws

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"pandax/kit/logger"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Connection struct {
	conn *websocket.Conn
	lock sync.Mutex
}

var connections = make(map[uint64]*Connection)
var connLock sync.Mutex

func init() {
	go checkConn()
}

// 放置ws连接
func Put(userId uint64, conn *websocket.Conn) {
	connLock.Lock()
	defer connLock.Unlock()

	Delete(userId)
	connections[userId] = &Connection{
		conn: conn,
	}
}

func checkConn() {
	heartbeat := time.Duration(20) * time.Second
	tick := time.NewTicker(heartbeat)
	for range tick.C {
		connLock.Lock()
		for uid, conn := range connections {
			conn.lock.Lock()
			err := conn.conn.WriteControl(websocket.PingMessage, []byte("ping"), time.Now().Add(heartbeat/2))
			conn.lock.Unlock()
			if err != nil {
				logger.Log.Info("删除ping失败的websocket连接：uid = ", uid)
				Delete(uid)
			}
		}
		connLock.Unlock()
	}
}

// 删除ws连接
func Delete(userid uint64) {
	connLock.Lock()
	defer connLock.Unlock()

	conn := connections[userid]
	if conn != nil {
		conn.lock.Lock()
		conn.conn.Close()
		conn.lock.Unlock()
		delete(connections, userid)
	}
}

// 对指定用户发送消息
func SendMsg(userId uint64, msg *Msg) {
	connLock.Lock()
	defer connLock.Unlock()

	conn := connections[userId]
	if conn != nil {
		conn.lock.Lock()
		defer conn.lock.Unlock()

		bytes, err := json.Marshal(msg)
		if err != nil {
			logger.Log.Error("发送消息失败：", err)
			return
		}
		err = conn.conn.WriteMessage(websocket.TextMessage, bytes)
		if err != nil {
			logger.Log.Error("发送消息失败：", err)
		}
	}
}
