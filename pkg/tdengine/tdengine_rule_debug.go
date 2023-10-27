package tdengine

import (
	"fmt"
)

const debugTableName = "device_rule_debug"

func (s *TdEngine) CreateDeviceRuleDebugTable() (err error) {
	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s (ts TIMESTAMP,nodeId NCHAR(64),msgd NCHAR(64),debugType NCHAR(64),
    deviceName NCHAR(64),msgType NCHAR(64),msg VARCHAR,metadata VARCHAR,error VARCHAR)`, s.dbName, debugTableName)
	_, err = s.db.Exec(sql)
	return
}

func (s *TdEngine) InsertRuleDebug(data map[string]any) (err error) {
	return s.InsertDevice(debugTableName, data)
}
