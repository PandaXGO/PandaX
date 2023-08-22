package nodes

import (
	"github.com/sirupsen/logrus"
	"log"
	"pandax/pkg/global"
	"pandax/pkg/rule_engine/message"
)

type saveTimeSeriesNode struct {
	bareNode
}

type saveTimeSeriesNodeFactory struct{}

func (f saveTimeSeriesNodeFactory) Name() string     { return "SaveTimeSeriesNode" }
func (f saveTimeSeriesNodeFactory) Category() string { return NODE_CATEGORY_ACTION }
func (f saveTimeSeriesNodeFactory) Labels() []string { return []string{"Success", "Failure"} }
func (f saveTimeSeriesNodeFactory) Create(id string, meta Metadata) (Node, error) {
	node := &saveTimeSeriesNode{
		bareNode: newBareNode(f.Name(), id, meta, f.Labels()),
	}
	return decodePath(meta, node)
}

func (n *saveTimeSeriesNode) Handle(msg message.Message) error {
	logrus.Infof("%s handle message '%s'", n.Name(), msg.GetType())
	successLabelNode := n.GetLinkedNode("Success")
	failureLabelNode := n.GetLinkedNode("Failure")
	if msg.GetType() != message.TelemetryMes {
		if failureLabelNode != nil {
			return failureLabelNode.Handle(msg)
		} else {
			return nil
		}
	}
	//deviceId := msg.GetMetadata().GetValues()["deviceId"].(string)
	deviceName := msg.GetMetadata().GetValues()["deviceName"].(string)
	log.Println("telemetry", msg.GetMsg())
	err := global.TdDb.InsertDevice(deviceName+"_telemetry", msg.GetMsg())
	log.Println(err)
	if err != nil {
		if failureLabelNode != nil {
			return failureLabelNode.Handle(msg)
		}
	}
	if successLabelNode != nil {
		return successLabelNode.Handle(msg)
	}
	return nil
}
