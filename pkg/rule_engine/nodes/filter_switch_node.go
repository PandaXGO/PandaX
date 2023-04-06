//  Licensed under the Apache License, Version 2.0 (the "License"); you may
//  not use p file except in compliance with the License. You may obtain
//  a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//  WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//  License for the specific language governing permissions and limitations
//  under the License.
package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type switchFilterNode struct {
	bareNode
	Scripts string `json:"scripts" yaml:"scripts"`
}

type switchFilterNodeFactory struct{}

func (f switchFilterNodeFactory) Name() string     { return "SwitchNode" }
func (f switchFilterNodeFactory) Category() string { return NODE_CATEGORY_FILTER }
func (f switchFilterNodeFactory) Labels() []string {
	return []string{
		"Failure", "True", "False", message.MessageTypePostTelemetryRequest,
		message.MessageTypeConnectEvent,
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

	scriptEngine := NewScriptEngine()
	SwitchResults, err := scriptEngine.ScriptOnSwitch(msg, n.Scripts)
	if err != nil {
		return nil
	}
	nodes := n.GetLinkedNodes()
	for label, node := range nodes {
		for _, switchresult := range SwitchResults {
			if label == switchresult {
				return node.Handle(msg)
			}
		}
	}
	return nil
}
