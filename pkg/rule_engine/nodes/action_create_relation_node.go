package nodes

import (
	"pandax/pkg/rule_engine/message"
)

type createRelationNode struct {
	bareNode
	Direction                       string
	RelationType                    string
	EntityType                      string
	EntityNamePattern               string
	EntityTypePattern               string
	EntityCacheExpiration           int64
	CreateEntityIfNotExists         bool
	ChangeOriginatorToRelatedEntity bool
	RemoveCurrentRelations          bool
}

type createRelationNodeFactory struct{}

func (f createRelationNodeFactory) Name() string     { return "CreateRelationNode" }
func (f createRelationNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f createRelationNodeFactory) Create(id string, meta Metadata) (Node, error) {
	labels := []string{"Success", "Failure"}
	node := &createRelationNode{
		bareNode: newBareNode(f.Name(), id, meta, labels),
	}
	return decodePath(meta, node)
}

func (n *createRelationNode) Handle(msg message.Message) error {
	return nil
}
