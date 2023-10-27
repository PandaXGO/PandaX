package nodes

import (
	"github.com/nats-io/nats.go"
	"pandax/pkg/rule_engine/message"
)

type externalNatsNode struct {
	bareNode
	Url     string `json:"url"`
	Subject string `json:"subject"`
	Body    string
	client  *nats.Conn
}

type externalNatsNodeFactory struct{}

func (f externalNatsNodeFactory) Name() string     { return "NatsNode" }
func (f externalNatsNodeFactory) Category() string { return NODE_CATEGORY_EXTERNAL }
func (f externalNatsNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f externalNatsNodeFactory) Create(id string, meta Properties) (Node, error) {
	node := &externalNatsNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	_, err := decodePath(meta, node)
	if err != nil {
		return node, err
	}
	connect, err := nats.Connect(node.Url)
	if err != nil {
		return node, err
	}
	node.client = connect
	return node, nil
}

func (n *externalNatsNode) Handle(msg *message.Message) error {
	n.Debug(msg, message.DEBUGIN, "")
	defer n.client.Close()
	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	template, err := ParseTemplate(n.Body, msg.GetAllMap())
	if err != nil {
		return err
	}
	err = n.client.Publish(n.Subject, []byte(template))
	if err != nil {
		n.Debug(msg, message.DEBUGOUT, err.Error())
		if failureLabelNode != nil {
			return failureLabelNode.Handle(msg)
		} else {
			return err
		}
	}
	if successLabelNode != nil {
		n.Debug(msg, message.DEBUGOUT, "")
		return successLabelNode.Handle(msg)
	}
	return nil
}
