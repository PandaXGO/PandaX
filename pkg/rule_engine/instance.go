package rule_engine

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/manifest"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/rule_engine/nodes"
)

// ruleChainInstance is rulechain's runtime instance that manage all nodes in this chain,
type ruleChainInstance struct {
	firstRuleNodeId string
	nodes           map[string]nodes.Node
}

func newRuleChainInstance(data []byte) (*ruleChainInstance, []error) {
	errors := []error{}

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
	errs := []error{}
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
		originalNode.AddLinkedNode(edge.Type, targetNode)
	}
	for name, node := range r.nodes {
		targetNodes := node.GetLinkedNodes()
		mustLabels := node.MustLabels()
		for _, label := range mustLabels {
			if _, found := targetNodes[label]; !found {
				err := fmt.Errorf("the label '%s' in node '%s' no exist'", label, name)
				errs = append(errs, err)
				continue
			}
		}
	}

	return r, errs
}

// StartRuleChain TODO 是否需要添加context
func (c *ruleChainInstance) StartRuleChain(context context.Context, message message.Message) error {
	if node, found := c.nodes[c.firstRuleNodeId]; found {
		go node.Handle(message)
	}
	return nil
}
