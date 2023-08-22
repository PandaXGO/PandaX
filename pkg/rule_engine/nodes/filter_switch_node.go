package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type switchFilterNode struct {
	bareNode
	Script string `json:"script" yaml:"script"`
}

type switchFilterNodeFactory struct{}

func (f switchFilterNodeFactory) Name() string     { return "SwitchNode" }
func (f switchFilterNodeFactory) Category() string { return NODE_CATEGORY_FILTER }
func (f switchFilterNodeFactory) Labels() []string {
	return []string{
		"True", "False",
		message.RowMes,
		message.AttributesMes,
		message.TelemetryMes,
		message.RpcRequestMes,
		message.AlarmMes,
		message.UpEventMes,
		message.ConnectMes,
		message.DisConnectMes,
	}
}
func (f switchFilterNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &switchFilterNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *switchFilterNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	scriptEngine := NewScriptEngine(msg, "Switch", n.Script)
	SwitchResults, err := scriptEngine.ScriptOnSwitch()
	if err != nil {
		return err
	}
	nodes := n.GetLinkedNodes()
	for label, node := range nodes {
		for _, switchresult := range SwitchResults {
			if label == switchresult {
				go node.Handle(msg)
			}
		}
	}
	return nil
}
