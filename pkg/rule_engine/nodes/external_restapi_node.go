package nodes

import (
	"github.com/sirupsen/logrus"
	"pandax/pkg/rule_engine/message"
)

type externalRestapiNode struct {
	bareNode
	RestEndpointUrlPattern         string            `json:"restEndpointUrlPattern" yaml:"restEndpointUrlPattern"`
	RequestMethod                  string            `json:"requestMethod" yaml:"requestMethod"`
	headers                        map[string]string `json:"headers" yaml:"headers"`
	UseSimpleClientHttpFactory     bool              `json:"useSimpleClientHttpFactory" yaml:"useSimpleClientHttpFactory"`
	ReadTimeoutMs                  int               `json:"readTimeoutMs" yaml:"readTimeoutMs"`
	MaxParallelRequestsCount       int               `json:"maxParallelRequestsCount" yaml:"maxParallelRequestsCount"`
	UseRedisQueueForMsgPersistence bool              `json:"useRedisQueueForMsgPersistence" yaml:"useRedisQueueForMsgPersistence"`
	trimQueue                      bool              `json:"trimQueue" yaml:"trimQueue"`
	MaxQueueSize                   int               `json:"maxQueueSize" yaml:"maxQueueSize"`
}

type externalRestapiNodeFactory struct{}

func (f externalRestapiNodeFactory) Name() string     { return "ExternalRestapiNode" }
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
