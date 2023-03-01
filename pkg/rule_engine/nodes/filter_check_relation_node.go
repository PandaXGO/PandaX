package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

const (
	RelationTypeContains    = "Contains"
	RelationTypeNotContains = "NotContains"
)

//检查关联关系
//该消息来自与哪个实体或到那个实体
type checkRelationFilterNode struct {
	bareNode
	Direction    string   `json:"direction" yaml:"direction"`
	RelationType string   `json:"relationType" yaml:"relationType"`
	InstanceType string   `json:"instanceType" yaml:"instanceType"`
	Values       []string `json:"values" yaml:"values"`
}

type checkRelationFilterNodeFactory struct{}

func (f checkRelationFilterNodeFactory) Name() string     { return "CheckRelationFilterNode" }
func (f checkRelationFilterNodeFactory) Category() string { return NODE_CATEGORY_FILTER }

func (f checkRelationFilterNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"True", "False"}
	node := &checkRelationFilterNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
		Values:   []string{},
	}
	return decodePath(meta, node)
}

func (n *checkRelationFilterNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	trueLabelNode := n.GetLinkedNode("True")
	falseLabelNode := n.GetLinkedNode("False")

	//direction := msg.GetDirection()
	attr := msg.GetMetadata().GetKeyValue(n.InstanceType)
	switch n.RelationType {
	case RelationTypeContains:
		for _, val := range n.Values {
			// specified attribute exist in names
			if attr == val {
				return trueLabelNode.Handle(msg)
			}
		}
		// not found
		return falseLabelNode.Handle(msg)

	case RelationTypeNotContains:
		for _, val := range n.Values {
			// specified attribute exist in names
			if attr == val {
				return falseLabelNode.Handle(msg)
			}
		}
		// not found
		return trueLabelNode.Handle(msg)
	}
	return nil
}
