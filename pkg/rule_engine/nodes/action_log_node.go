package nodes

import (
	"pandax/apps/visual/entity"
	"pandax/apps/visual/services"
	"pandax/pkg/rule_engine/message"
)

type logNode struct {
	bareNode
	Script string `json:"script"`
}

type logNodeFactory struct{}

func (f logNodeFactory) Name() string     { return "LogNode" }
func (f logNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f logNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f logNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &logNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *logNode) Handle(msg message.Message) error {
	successLableNode := n.GetLinkedNode("Success")
	failureLableNode := n.GetLinkedNode("Failure")

	scriptEngine := NewScriptEngine(msg, "ToString", n.Script)
	logMessage, err := scriptEngine.ScriptToString()
	if err != nil {
		if failureLableNode != nil {
			return failureLableNode.Handle(msg)
		} else {
			return err
		}
	}
	services.VisualRuleChainMsgLogModelDao.Insert(entity.VisualRuleChainMsgLog{
		MessageId:  msg.GetId(),
		MsgType:    msg.GetType(),
		DeviceName: msg.GetMetadata().GetValues()["deviceName"].(string),
		Ts:         msg.GetTs(),
		Content:    logMessage,
	})
	if err != nil {
		if failureLableNode != nil {
			return failureLableNode.Handle(msg)
		} else {
			return err
		}
	}
	if successLableNode != nil {
		return successLableNode.Handle(msg)
	}
	return nil
}
