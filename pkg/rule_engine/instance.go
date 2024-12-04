package rule_engine

import (
	"pandax/pkg/rule_engine/manifest"
	"pandax/pkg/rule_engine/nodes"

	"github.com/sirupsen/logrus"
)

type RuleChainInstance struct {
	ruleId          string
	firstRuleNodeID string
	nodes           map[string]nodes.Node
}

func NewRuleChainInstance(ruleId string, data []byte) (*RuleChainInstance, error) {
	manifest, err := manifest.New(data)
	if err != nil {
		logrus.WithError(err).Errorf("invalid manifest file")
		return nil, err
	}
	withManifest, err := newInstanceWithManifest(manifest)
	if err != nil {
		return nil, err
	}
	withManifest.ruleId = ruleId
	return withManifest, nil
}

func newInstanceWithManifest(m *manifest.Manifest) (*RuleChainInstance, error) {
	nodes, err := nodes.GetNodes(m)
	if err != nil {
		return nil, err
	}
	r := &RuleChainInstance{
		firstRuleNodeID: m.FirstRuleNodeId,
		nodes:           nodes,
	}
	return r, nil
}
