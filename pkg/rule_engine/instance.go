package rule_engine

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/manifest"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/rule_engine/nodes"
	"strings"
)

// ruleChainInstance is rulechain's runtime instance that manage all nodes in this chain,
type ruleChainInstance struct {
	firstRuleNodeId string
	nodes           map[string]nodes.Node
}

func NewRuleChainInstance(data []byte) (*ruleChainInstance, []error) {
	errors := make([]error, 0)

	manifest, err := manifest.New(data)
	if err != nil {
		errors = append(errors, err)
		logrus.WithError(err).Errorf("invalidi manifest file")
		return nil, errors
	}
	return newInstanceWithManifest(manifest)
}

// newWithManifest create rule chain by user's manifest file
func newInstanceWithManifest(m *manifest.Manifest) (*ruleChainInstance, []error) {
	errs := make([]error, 0)
	r := &ruleChainInstance{
		firstRuleNodeId: m.FirstRuleNodeId,
		nodes:           make(map[string]nodes.Node),
	}
	// Create All nodes
	for _, n := range m.Nodes {
		metadata := nodes.NewMetadataWithValues(n.Properties)
		node, err := nodes.NewNode(n.Type, n.Id, metadata)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		if _, found := r.nodes[n.Id]; found {
			err := fmt.Errorf("node '%s' already exist in rulechain", n.Id)
			errs = append(errs, err)
			continue
		}
		r.nodes[n.Id] = node
	}
	for _, edge := range m.Edges {
		originalNode, found := r.nodes[edge.SourceNodeId]
		if !found {
			err := fmt.Errorf("original node '%s' no exist in", originalNode.Name())
			errs = append(errs, err)
			continue
		}
		targetNode, found := r.nodes[edge.TargetNodeId]
		if !found {
			err := fmt.Errorf("target node '%s' no exist in rulechain", targetNode.Name())
			errs = append(errs, err)
			continue
		}
		//可以有多个类型
		if ty, ok := edge.Properties["lineType"]; ok {
			split := strings.Split(ty.(string), "/")
			for _, ty := range split {
				originalNode.AddLinkedNode(ty, targetNode)
			}
		}
	}
	return r, errs
}

// StartRuleChain
func (c *ruleChainInstance) StartRuleChain(context context.Context, message message.Message) error {
	if node, found := c.nodes[c.firstRuleNodeId]; found {
		go node.Handle(message)
	}
	return nil
}
