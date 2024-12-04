package entity

import (
	"encoding/json"
)

type RuleDataJson struct {
	Id          string
	GlobalColor string                 `json:"globalColor"`
	DataCode    map[string]interface{} `json:"dataCode"`
	OpenRule    bool                   `json:"openRule"`
	Setting     map[string]interface{} `json:"setting"`
}

// 序列化
func (m *RuleDataJson) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

// 反序列化
func (m *RuleDataJson) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}
