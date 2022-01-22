package service

import (
	"fmt"
	"pandax/base/global"

	client "k8s.io/client-go/kubernetes"
	"pandax/apps/devops/entity/k8s"
	k8scommon "pandax/apps/devops/services/k8s/common"
	"pandax/apps/devops/services/k8s/dataselect"
	"pandax/apps/devops/services/k8s/event"
)

// GetServiceEvents returns model events for a service with the given name in the given namespace.
func GetServiceEvents(client *client.Clientset, dsQuery *dataselect.DataSelectQuery, namespace, name string) (*k8scommon.EventList, error) {
	eventList := k8scommon.EventList{
		Events:   make([]k8scommon.Event, 0),
		ListMeta: k8s.ListMeta{TotalItems: 0},
	}

	serviceEvents, err := event.GetEvents(client, namespace, name)
	if err != nil {
		return &eventList, err
	}

	eventList = event.CreateEventList(event.FillEventsType(serviceEvents), dsQuery)
	global.Log.Info(fmt.Sprintf("Found %d events related to %s service in %s namespace", len(eventList.Events), name, namespace))
	return &eventList, nil
}
