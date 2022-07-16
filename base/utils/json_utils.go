package utils

import (
	"encoding/json"
)

func Json2Map(jsonStr string) map[string]any {
	var res map[string]any
	if jsonStr == "" {
		return res
	}
	_ = json.Unmarshal([]byte(jsonStr), &res)
	return res
}
