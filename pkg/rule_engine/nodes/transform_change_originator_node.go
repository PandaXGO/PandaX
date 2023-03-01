package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type transformChangeOriginatorNode struct {
	bareNode
	OriginatorSource string `json:"originatorSource" yaml:"originatorSource"`
	Direction        string `json:"direction" yaml:"direction"`
	MaxRelationLevel int    `json:"maxRelationLevel" yaml:"maxRelationLevel"`
	//RelationFilters  []runtime.RelationFilter `json:"relationFilters" yaml:"relationFilters"`
}

type transformChangeOriginatorNodeFactory struct{}

func (f transformChangeOriginatorNodeFactory) Name() string     { return "TransformChangeOriginatorNode" }
func (f transformChangeOriginatorNodeFactory) Category() string { return NODE_CATEGORY_TRANSFORM }

func (f transformChangeOriginatorNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"Success", "Failure"}
	node := &transformChangeOriginatorNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
		//RelationFilters: []runtime.RelationFilter{},
	}
	return decodePath(meta, node)
}

func (n *transformChangeOriginatorNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	//successLabelNode := n.GetLinkedNode("Sucess")
	failureLabelNode := n.GetLinkedNode("Failure")
	//relationQuery := runtime.NewRelationQuery()

	//entities := relationQuery.QueryEntities(n.Direction, n.MaxRelationLevel, n.RelationFilters)
	/*if len(entities) > 0 && entities[0] == msg.GetOriginator() {
		msg.SetOriginator(entities[0])
		return successLabelNode.Handle(msg)
	}*/
	return failureLabelNode.Handle(msg)
}
