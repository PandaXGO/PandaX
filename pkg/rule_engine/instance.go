package rule_engine

import (
	"context"
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/manifest"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/rule_engine/nodes"
)

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
	nodes, err := nodes.GetNodes(m)
	if err != nil {
		errs = append(errs, err)
		return nil, errs
	}
	r := &ruleChainInstance{
		firstRuleNodeId: m.FirstRuleNodeId,
		nodes:           nodes,
	}
	return r, errs
}

// StartRuleChain TODO 是否需要添加context
func (c *ruleChainInstance) StartRuleChain(context context.Context, message message.Message) error {
	if node, found := c.nodes[c.firstRuleNodeId]; found {
		node.Handle(message)
	}
	return nil
}
