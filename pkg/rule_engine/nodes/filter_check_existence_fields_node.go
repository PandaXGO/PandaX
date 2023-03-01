package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

//检查关联关系
//该消息来自与哪个实体或到那个实体
type checkExistenceFieldsNode struct {
	bareNode
}

type checkExistenceFieldsNodeFactory struct{}

func (f checkExistenceFieldsNodeFactory) Name() string     { return "CheckExistenceFieldsNode" }
func (f checkExistenceFieldsNodeFactory) Category() string { return NODE_CATEGORY_FILTER }

func (f checkExistenceFieldsNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"True", "False"}
	node := &checkExistenceFieldsNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}

func (n *checkExistenceFieldsNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	//trueLabelNode := n.GetLinkedNode("True")
	//falseLabelNode := n.GetLinkedNode("False")

	return nil
}
