package node

import (
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"pandax/apps/devops/services/k8s/common"
	"pandax/apps/devops/services/k8s/dataselect"
)

type NodeCell v1.Node

func (self NodeCell) GetProperty(name dataselect.PropertyName) dataselect.ComparableValue {
	switch name {
	case dataselect.NameProperty:
		return dataselect.StdComparableString(self.ObjectMeta.Name)
	case dataselect.CreationTimestampProperty:
		return dataselect.StdComparableTime(self.ObjectMeta.CreationTimestamp.Time)
	case dataselect.NamespaceProperty:
		return dataselect.StdComparableString(self.ObjectMeta.Namespace)
	default:
		// if name is not supported then just return a constant dummy value, sort will have no effect.
		return nil
	}
}

func toCells(std []v1.Node) []dataselect.DataCell {
	cells := make([]dataselect.DataCell, len(std))

	for i := range std {
		cells[i] = NodeCell(std[i])

	}
	return cells
}

func fromCells(cells []dataselect.DataCell) []v1.Node {
	std := make([]v1.Node, len(cells))
	for i := range std {
		std[i] = v1.Node(cells[i].(NodeCell))
	}
	return std
}

func getNodeConditions(node v1.Node) []common.Condition {
	var conditions []common.Condition
	for _, condition := range node.Status.Conditions {
		conditions = append(conditions, common.Condition{
			Type:               string(condition.Type),
			Status:             metaV1.ConditionStatus(condition.Status),
			LastProbeTime:      condition.LastHeartbeatTime,
			LastTransitionTime: condition.LastTransitionTime,
			Reason:             condition.Reason,
			Message:            condition.Message,
		})
	}
	return conditions
}

//getContainerImages returns container image strings from the given node.
func getContainerImages(node v1.Node) []string {
	var containerImages []string
	for _, image := range node.Status.Images {
		for _, name := range image.Names {
			containerImages = append(containerImages, name)
		}
	}
	return containerImages
}