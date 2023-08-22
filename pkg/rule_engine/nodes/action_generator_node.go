package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
	"time"
)

type messageGeneratorNode struct {
	bareNode
	Script       string `json:"script" yaml:"script"`
	PeriodSecond int64  `json:"periodSecond" yaml:"periodSecond"` //周期
	MessageCount int64  `json:"messageCount" yaml:"messageCount"`
}

type messageGeneratorNodeFactory struct{}

func (f messageGeneratorNodeFactory) Name() string     { return "MessageGeneratorNode" }
func (f messageGeneratorNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f messageGeneratorNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f messageGeneratorNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &messageGeneratorNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *messageGeneratorNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")

	ticker := time.NewTicker(time.Duration(n.PeriodSecond) * time.Second)
	count := 0

	go func() {
		for {
			<-ticker.C
			count++
			if int64(count) == n.MessageCount {
				ticker.Stop()
				return
			}
			scriptEngine := NewScriptEngine(msg, "Generate", n.Script)
			generate, err := scriptEngine.ScriptGenerate()
			if err != nil {
				if failureLabelNode != nil {
					go failureLabelNode.Handle(msg)
				}
				return
			}
			msg.SetMsg(generate["msg"].(map[string]interface{}))
			msg.SetType(generate["msgType"].(string))
			msg.SetMetadata(message.NewDefaultMetadata(generate["metadata"].(map[string]interface{})))
			if successLabelNode != nil {
				go successLabelNode.Handle(msg)
			}
		}
	}()

	return nil
}
