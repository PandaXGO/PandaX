package nodes

import (
	"errors"
	"github.com/sirupsen/logrus"
	"log"
	"pandax/pkg/rule_engine/manifest"
	"pandax/pkg/rule_engine/message"
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

func GetNodes(m *manifest.Manifest) (map[string]Node, error) {
	nodes := make(map[string]Node)
	// Create All nodes
	for _, n := range m.Nodes {
		metadata := NewMetadataWithValues(n.Properties)
		node, err := NewNode(n.Type, n.Id, metadata)
		if err != nil {
			log.Println(err)
			logrus.Errorf("new node '%s' failure", n.Id)
			continue
		}
		if _, found := nodes[n.Id]; found {
			logrus.Errorf("node '%s' already exist in rulechain", n.Id)
			continue
		}
		nodes[n.Id] = node
	}
	for _, edge := range m.Edges {
		originalNode, found := nodes[edge.SourceNodeId]
		if !found {
			logrus.Errorf("original node '%s' no exist in", originalNode.Name())
			continue
		}
		targetNode, found := nodes[edge.TargetNodeId]
		if !found {
			logrus.Errorf("target node '%s' no exist in rulechain", targetNode.Name())
			continue
		}
		//可以有多个类型
		types := make([]interface{}, 0)
		if _, ok := edge.Properties["lineType"]; !ok {
			types = append(types, "True")
		} else {
			types = edge.Properties["lineType"].([]interface{})
		}
		for _, ty := range types {
			originalNode.AddLinkedNode(ty.(string), targetNode)
		}

	}

	return nodes, nil
}
