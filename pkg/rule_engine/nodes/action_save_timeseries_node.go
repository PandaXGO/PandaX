package nodes

import (
	"pandax/pkg/rule_engine/message"
)

type SaveTimeSeriesNode struct {
	bareNode
}

type saveTimeSeriesNodeFactory struct{}

func (f saveTimeSeriesNodeFactory) Name() string     { return "SaveTimeSeriesNode" }
func (f saveTimeSeriesNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f saveTimeSeriesNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f saveTimeSeriesNodeFactory) Create(id string, meta Metadata) (Node, error) {
	return nil, nil
}

func (n *SaveTimeSeriesNode) Handle(msg message.Message) error {
	successLableNode := n.GetLinkedNode("Success")
	failureLableNode := n.GetLinkedNode("Failure")
	if msg.GetType() != message.EventTelemetryType {
		if failureLableNode != nil {
			return failureLableNode.Handle(msg)
		} else {
			return nil
		}
	}
	/*deviceName := msg.GetMetadata().GetValues()["deviceName"].(string)
	namespace := msg.GetMetadata().GetValues()["namespace"].(string)
	marshal, err := json.Marshal(msg.GetMsg())*/

	// todo 添加设备上报遥测

	if successLableNode != nil {
		return successLableNode.Handle(msg)
	} else {
		return nil
	}

}
