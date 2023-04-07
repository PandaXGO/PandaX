package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type externalRestapiNode struct {
	bareNode
	RestEndpointUrlPattern string            `json:"restEndpointUrlPattern" yaml:"restEndpointUrlPattern"`
	RequestMethod          string            `json:"requestMethod" yaml:"requestMethod"`
	headers                map[string]string `json:"headers" yaml:"headers"`
}

type externalRestapiNodeFactory struct{}

func (f externalRestapiNodeFactory) Name() string     { return "RestapiNode" }
func (f externalRestapiNodeFactory) Category() string { return NODE_CATEGORY_EXTERNAL }
func (f externalRestapiNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f externalRestapiNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &externalRestapiNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *externalRestapiNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())

	return nil
}
