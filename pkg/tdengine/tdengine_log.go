package tdengine

import (
	"fmt"
	"time"

	"github.com/kakuilan/kgo"
)

const logTableName = "logs"

// 日志 TDengine
type TdLog struct {
	Ts       string `json:"ts" dc:"时间"`
	DeviceId string `json:"deviceId" dc:"设备标识"`
	TraceId  string `json:"traceId" dc:"追踪"`
	Type     string `json:"type" dc:"日志类型"` // 命令调用  上行 下行
	Content  string `json:"content" dc:"日志内容"`
}

// CreateLogStable 添加LOG超级表
func (s *TdEngine) CreateLogStable() (err error) {
	sql := "CREATE STABLE IF NOT EXISTS ? (ts TIMESTAMP,deviceId NCHAR(64),traceId NCHAR(64),type NCHAR(20), content VARCHAR(1000))"
	_, err = s.db.Exec(sql, logTableName)
	return
}

// InsertLog 写入数据
func (s *TdEngine) InsertLog(log *TdLog) (err error) {
	logs, err := kgo.KConv.Struct2Map(*log, "")
	if err != nil {
		return err
	}
	err = s.InsertDevice(logTableName, logs)
	return
}

// ClearLog 清理过期数据
func (s *TdEngine) ClearLog() (err error) {
	ts := time.Now().Add(-7 * 24 * time.Hour).Format("2006-01-02")

	sql := fmt.Sprintf("DELETE FROM %s WHERE ts < ?", logTableName)
	_, err = s.db.Exec(sql, ts)

	return
}

// GetAllLog 超级表查询，多条数据
func (s *TdEngine) GetAllLog(sql string, args ...any) (list []TdLog, err error) {
	rows, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var log TdLog

		err = rows.Scan(&log.Ts, &log.DeviceId, &log.TraceId, &log.Type, &log.Content)
		if err != nil {
			return nil, err
		}
		log.Ts = s.Time(log.Ts)

		list = append(list, log)
	}

	return
}
