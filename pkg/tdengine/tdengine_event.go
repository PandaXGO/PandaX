package tdengine

import (
	"fmt"
)

const connectTableName = "device_connect"

type ConnectInfo struct {
	Ts         string `json:"ts"`
	ClientID   string `json:"clientId"`
	Type       string `json:"type"` // 连接类型
	PeerHost   string `json:"peerHost"`
	SocketPort string `json:"sockPort"`
	Protocol   string `json:"protocol"`
	DeviceId   string `json:"deviceId"`
}

// CreateEventTable 创建设备连接事件表
func (s *TdEngine) CreateEventTable() (err error) {
	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s (ts TIMESTAMP,deviceId NCHAR(64),
    type NCHAR(64),clientId NCHAR(64),peerHost NCHAR(64),sockPort NCHAR(64),protocol NCHAR(64))`, s.dbName, connectTableName)
	_, err = s.db.Exec(sql)
	return
}

func (s *TdEngine) InsertEvent(data map[string]any) (err error) {
	return s.InsertDevice(connectTableName, data)
}
