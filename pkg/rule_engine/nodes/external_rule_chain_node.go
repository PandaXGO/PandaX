package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type externalRuleChainNode struct {
	bareNode
	RuleId string `json:"ruleId" yaml:"ruleId"`
}

type externalRuleChainNodeFactory struct{}

func (f externalRuleChainNodeFactory) Name() string     { return "RuleChainNode" }
func (f externalRuleChainNodeFactory) Category() string { return NODE_CATEGORY_EXTERNAL }
func (f externalRuleChainNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f externalRuleChainNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"Success", "Failure"}

	node := &externalRuleChainNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}

	return decodePath(meta, node)
}

func (n *externalRuleChainNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	return nil
}
