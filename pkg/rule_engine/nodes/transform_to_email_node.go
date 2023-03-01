package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type transformToEmailNode struct {
	bareNode
	From    string `json:"from" yaml:"from"`
	To      string `json:"to" yaml:"to"`
	Cc      string `json:"cc" yaml:"cc"`
	Bcc     string `json:"bcc" yaml:"bcc"`
	Subject string `json:"subject" yaml:"subject"`
	Body    string `json:"body" yaml:"body"`
}

type transformToEmailNodeFactory struct{}

func (f transformToEmailNodeFactory) Name() string     { return "TransformToEmailNode" }
func (f transformToEmailNodeFactory) Category() string { return NODE_CATEGORY_TRANSFORM }

func (f transformToEmailNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"Success", "Failure"}

	node := &transformToEmailNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}

func (n *transformToEmailNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	successLabelNode := n.GetLinkedNode("Success")
	//failureLabelNode := n.GetLinkedNode("Failure")

	/*dialer := runtime.NewDialer(runtime.EMAIL)
	variables := map[string]string{
		"from":    n.From,
		"to":      n.To,
		"cc":      n.Cc,
		"bcc":     n.Bcc,
		"subject": n.Subject,
		"body":    n.Body,
	}
	if err := dialer.DialAndSend(msg.GetMetadata(), variables); err != nil {
		return failureLabelNode.Handle(msg)
	}*/
	return successLabelNode.Handle(msg)
}
