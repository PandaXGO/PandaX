package rule_engine

import (
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/rule_engine/nodes"
)

func GetCategory() []map[string]interface{} {
	return nodes.GetCategory()
}

func GetDebugData(ruleId, nodeId string) []message.DebugData {
	if data, ok := ruleChainDebugData.Data[ruleId]; ok {
		return data.Get(nodeId).Items
	}
	return nil
}

func GetDebugDataPage(page, pageSize int, ruleId, nodeId string) (int64, []message.DebugData) {
	if page < 1 {
		page = 1
	}
	offset := pageSize * (page - 1)
	if data, ok := ruleChainDebugData.Data[ruleId]; ok {
		total := len(data.Get(nodeId).Items)
		end := offset + pageSize
		if end >= total {
			end = total - 1
		}
		return int64(total), data.Get(nodeId).Items[offset:end]
	}
	return 0, nil
}
