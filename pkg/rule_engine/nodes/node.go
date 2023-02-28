package nodes

import (
	"errors"
	"pandax/pkg/rule_engine/message"

	"github.com/sirupsen/logrus"
)

type Node interface {
	Name() string
	Id() string
	Metadata() Metadata
	MustLabels() []string
	Handle(message.Message) error

	AddLinkedNode(label string, node Node)
	GetLinkedNode(label string) Node
	GetLinkedNodes() map[string]Node
}

type bareNode struct {
	name   string
	id     string
	nodes  map[string]Node
	meta   Metadata
	labels []string
}

func newBareNode(name string, id string, meta Metadata, labels []string) bareNode {
	return bareNode{
		name:   name,
		id:     id,
		nodes:  make(map[string]Node),
		meta:   meta,
		labels: labels,
	}
}

func (n *bareNode) Name() string                          { return n.name }
func (n *bareNode) WithId(id string)                      { n.id = id }
func (n *bareNode) Id() string                            { return n.id }
func (n *bareNode) MustLabels() []string                  { return n.labels }
func (n *bareNode) AddLinkedNode(label string, node Node) { n.nodes[label] = node }

func (n *bareNode) GetLinkedNode(label string) Node {
	if node, found := n.nodes[label]; found {
		return node
	}
	logrus.Errorf("no label '%s' in node '%s'", label, n.name)
	return nil
}

func (n *bareNode) GetLinkedNodes() map[string]Node { return n.nodes }

func (n *bareNode) Metadata() Metadata { return n.meta }

func (n *bareNode) Handle(message.Message) error { return errors.New("not implemented") }

func decodePath(meta Metadata, n Node) (Node, error) {
	if err := meta.DecodePath(n); err != nil {
		return n, err
	}
	return n, nil
}
