package tdengine

import (
	"fmt"
)

const connectTableName = "events"

type Events struct {
	Ts       string `json:"ts"`
	Name     string `json:"name"`    //标识  connet
	Type     string `json:"type"`    // 事件类型 info alarm fault
	Content  string `json:"content"` // 事件描述
	DeviceId string `json:"deviceId"`
}

// CreateEventTable 创建设备连接事件表
func (s *TdEngine) CreateEventTable() (err error) {
	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s (ts TIMESTAMP,deviceId NCHAR(64),name NCHAR(64),
    type NCHAR(64),content NCHAR(255))`, s.dbName, connectTableName)
	_, err = s.db.Exec(sql)
	return
}

func (s *TdEngine) InsertEvent(data map[string]any) (err error) {
	return s.InsertDevice(connectTableName, data)
}

func (s *TdEngine) GetAllEvents(sql string, args ...any) (list []Events, err error) {
	rows, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Events

		err = rows.Scan(&event.Ts, &event.DeviceId, &event.Name, &event.Type, &event.Content)
		if err != nil {
			return nil, err
		}
		event.Ts = s.Time(event.Ts)

		list = append(list, event)
	}

	return
}
