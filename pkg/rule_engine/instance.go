package rule_engine

import (
	"context"
	"errors"
	"pandax/pkg/rule_engine/manifest"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/rule_engine/nodes"

	"github.com/sirupsen/logrus"
)

var ruleChainDebugData = message.NewRuleChainDebugData(100)

type RuleChainInstance struct {
	ruleID          string
	firstRuleNodeID string
	nodes           map[string]nodes.Node
}

func NewRuleChainInstance(ruleID string, data []byte) (*RuleChainInstance, error) {
	manifest, err := manifest.New(data)
	if err != nil {
		logrus.WithError(err).Errorf("invalid manifest file")
		return nil, err
	}
	withManifest, err := newInstanceWithManifest(manifest)
	if err != nil {
		return nil, err
	}
	withManifest.ruleID = ruleID
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

func (c *RuleChainInstance) StartRuleChain(ctx context.Context, msg *message.Message) error {
	debugChan := make(chan *message.DebugData, 100)
	endDebugChan := make(chan struct{})

	go func() {
		for {
			select {
			case debugMsg := <-debugChan:
				ruleChainDebugData.Add(c.ruleID, debugMsg.NodeId, *debugMsg)
			case <-endDebugChan:
				logrus.Debugf("规则链%s,执行结束", msg.Id)
				return
			}
		}
	}()

	node, found := c.nodes[c.firstRuleNodeID]
	if !found {
		return errors.New("first rule node not found")
	}

	err := node.Handle(msg)
	endDebugChan <- struct{}{}
	return err
}
