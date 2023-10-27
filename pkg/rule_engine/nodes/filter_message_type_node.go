package nodes

import (
	"pandax/pkg/rule_engine/message"
)

type messageTypeFilterNode struct {
	bareNode
	MessageTypes []string `json:"messageTypes" yaml:"messageTypes"`
}

type messageTypeFilterNodeFactory struct{}

func (f messageTypeFilterNodeFactory) Name() string     { return "MessageTypeNode" }
func (f messageTypeFilterNodeFactory) Category() string { return NODE_CATEGORY_FILTER }
func (f messageTypeFilterNodeFactory) Labels() []string { return []string{"True", "False"} }

func (f messageTypeFilterNodeFactory) Create(id string, meta Properties) (Node, error) {
	node := &messageTypeFilterNode{
		bareNode:     newBareNode(f.Name(), id, meta, f.Labels()),
		MessageTypes: []string{},
	}
	return decodePath(meta, node)
}

func (n *messageTypeFilterNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")

	trueLabelNode := n.GetLinkedNode("True")
	falseLabelNode := n.GetLinkedNode("False")

	if n.containsType(msg.MsgType) {
		if trueLabelNode != nil {
			n.Debug(msg, message.DEBUGOUT, "")
			return trueLabelNode.Handle(msg)
		}
	} else {
		if falseLabelNode != nil {
			n.Debug(msg, message.DEBUGOUT, "不包含消息类型")
			return falseLabelNode.Handle(msg)
		}
	}
	return nil
}

func (n *messageTypeFilterNode) containsType(messageType string) bool {
	for _, filterType := range n.MessageTypes {
		if filterType == messageType {
			return true
		}
	}
	return false
}
