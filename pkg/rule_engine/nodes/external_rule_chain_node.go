package nodes

import (
	"pandax/apps/rule/services"
	"pandax/pkg/rule_engine/message"

	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

type externalRuleChainNode struct {
	bareNode
	RuleId string `json:"ruleId" yaml:"ruleId"`
}

type externalRuleChainNodeFactory struct{}

func (f externalRuleChainNodeFactory) Name() string     { return "RuleChainNode" }
func (f externalRuleChainNodeFactory) Category() string { return NODE_CATEGORY_FLOWS }
func (f externalRuleChainNodeFactory) Labels() []string { return []string{} }
func (f externalRuleChainNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &externalRuleChainNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}

	return decodePath(meta, node)
}

func (n *externalRuleChainNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())
	data := services.RuleChainModelDao.FindOne(n.RuleId)
	if data == nil {
		return errors.New(fmt.Sprintf("节点 %s ,获取规则链失败", n.Name()))
	}

	/*code, _ := json.Marshal(data.RuleDataJson.LfData.DataCode)
	m, err := manifest.New(code)
	if err != nil {
		logrus.WithError(err).Errorf("invalidi manifest file")
		return err
	}
	nodes, err := GetNodes(m)
	if err != nil {
		return errors.New(fmt.Sprintf("节点 %s ,构建节点失败", n.Name()))
	}
	if node, found := nodes[m.FirstRuleNodeId]; found {
		go node.Handle(msg)
	}*/
	return nil
}
