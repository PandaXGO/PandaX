package rule_engine

import (
	"context"
	"github.com/sirupsen/logrus"
	"pandax/pkg/global"
	"pandax/pkg/rule_engine/manifest"
	"pandax/pkg/rule_engine/message"
	"pandax/pkg/rule_engine/nodes"
)

var ruleChainDebugData = message.NewRuleChainDebugData(100)

type ruleChainInstance struct {
	ruleId          string
	firstRuleNodeId string
	nodes           map[string]nodes.Node
}

func NewRuleChainInstance(ruleId string, data []byte) (*ruleChainInstance, []error) {
	errors := make([]error, 0)

	manifest, err := manifest.New(data)
	if err != nil {
		errors = append(errors, err)
		logrus.WithError(err).Errorf("invalidi manifest file")
		return nil, errors
	}
	withManifest, errs := newInstanceWithManifest(manifest)
	if len(errs) > 0 {
		return nil, errs
	}
	withManifest.ruleId = ruleId
	return withManifest, nil
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
func (c *ruleChainInstance) StartRuleChain(context context.Context, message *message.Message) error {
	// 处理debug的通道消息
	go func() {
		for {
			select {
			case debugMsg := <-message.DeBugChan:
				// 保存到tdengine时序数据库中
				ruleChainDebugData.Add(c.ruleId, debugMsg.NodeId, debugMsg)
			case <-message.EndDeBugChan:
				global.Log.Debug("规则链%s,执行结束", message.Id)
				return
			}
		}

	}()
	if node, found := c.nodes[c.firstRuleNodeId]; found {
		err := node.Handle(message)
		message.EndDeBugChan <- struct{}{}
		return err
	}
	return nil
}
