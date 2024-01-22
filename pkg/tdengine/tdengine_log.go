package tdengine

import "time"

const logTableName = "device_log"

// CreateLogStable 添加LOG超级表
func (s *TdEngine) CreateLogStable() (err error) {
	sql := "CREATE STABLE IF NOT EXISTS device_log (ts TIMESTAMP, type VARCHAR(20), content VARCHAR(1000)) TAGS (device VARCHAR(255))"
	_, err = s.db.Exec(sql)
	return
}

// InsertLog 写入数据
func (s *TdEngine) InsertLog(log *TdLog) (err error) {
	sql := "INSERT INTO log_? USING device_log TAGS (?) VALUES (?, ?, ?)"
	_, err = s.db.Exec(sql, log.Device, log.Device, log.Ts, log.Type, log.Content)

	return
}

// ClearLog 清理过期数据
func (s *TdEngine) ClearLog() (err error) {
	ts := time.Now().Add(-7 * 24 * time.Hour).Format("2006-01-02")

	sql := "DELETE FROM device_log WHERE ts < ?"
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

		err = rows.Scan(&log.Ts, &log.Type, &log.Content, &log.Device)
		if err != nil {
			return nil, err
		}
		log.Ts = s.Time(log.Ts)

		list = append(list, log)
	}

	return
}
