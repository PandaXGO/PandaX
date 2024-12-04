package rule_engine

import (
	"errors"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/rule_engine/nodes"
	"sync"

	"github.com/sirupsen/logrus"
)

var RuleEngine = (*RuleChainEngine)(nil)

type RuleChainEngine struct {
	pool               sync.Map
	ruleChainDebugData *message.RuleChainDebugData
}

func init() {
	RuleEngine = &RuleChainEngine{
		pool:               sync.Map{},
		ruleChainDebugData: message.NewRuleChainDebugData(100),
	}
}

func (en *RuleChainEngine) GetRuleInstance(ruleID string) *RuleChainInstance {
	instance, ok := en.pool.Load(ruleID)
	if ok {
		return instance.(*RuleChainInstance)
	}
	// TODO 没有就去数据库查询并返回

	return nil
}

// iD为产品Id, 一个产品对应一个规则实体
func (en *RuleChainEngine) SaveRuleInstance(id string, instance *RuleChainInstance) (*RuleChainInstance, error) {
	en.pool.Store(id, instance)
	return instance, nil
}

func (en *RuleChainEngine) DeletRuleInstance(id string) {
	en.pool.Delete(id)
}

func (en *RuleChainEngine) StartRuleInstance(instance *RuleChainInstance, msg *message.Message) error {
	go func() {
		for {
			select {
			case debugMsg := <-msg.DeBugChan:
				en.ruleChainDebugData.Add(instance.ruleId, debugMsg.NodeId, debugMsg)
			case <-msg.EndDeBugChan:
				logrus.Debugf("规则链%s,执行结束", msg.Id)
				return
			}
		}
	}()
	node, found := instance.nodes[instance.firstRuleNodeID]
	if !found {
		return errors.New("first rule node not found")
	}

	err := node.Handle(msg)
	msg.EndDeBugChan <- struct{}{}
	return err
}

func (en *RuleChainEngine) GetDebugData(ruleId, nodeId string) []message.DebugData {
	if data, ok := en.ruleChainDebugData.Data[ruleId]; ok {
		return data.Get(nodeId).Items
	}
	return nil
}

func (en *RuleChainEngine) ClearDebugData(ruleId, nodeId string) {
	if data, ok := en.ruleChainDebugData.Data[ruleId]; ok {
		data.Clear(nodeId)
	}
}

func (en *RuleChainEngine) GetDebugDataPage(page, pageSize int, ruleId, nodeId string) (int64, []message.DebugData, error) {
	if page < 1 {
		page = 1
	}
	offset := pageSize * (page - 1)
	if data, ok := en.ruleChainDebugData.Data[ruleId]; ok {
		nodeData := data.Get(nodeId)
		if nodeData != nil {
			total := len(nodeData.Items)
			end := offset + pageSize
			if end >= total {
				end = total
			}
			return int64(total), nodeData.Items[offset:end], nil
		}
	}
	return 0, nil, errors.New("规则不存在")
}

func GetCategory() []map[string]interface{} {
	return nodes.GetCategory()
}
