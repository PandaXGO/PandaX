package storageclass

import (
	"context"
	"fmt"
	"pandax/base/global"

	storage "k8s.io/api/storage/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"pandax/apps/devops/services/k8s/dataselect"
	"pandax/apps/devops/services/k8s/pv"
)

// StorageClassDetail provides the presentation layer view of Storage Class resource.
type StorageClassDetail struct {
	// Extends list item structure.
	StorageClass         `json:",inline"`
	PersistentVolumeList pv.PersistentVolumeList `json:"persistentVolumeList"`
}

// GetStorageClassDetail returns Storage Class resource.
func GetStorageClassDetail(client kubernetes.Interface, name string) (*StorageClassDetail, error) {
	global.Log.Info(fmt.Sprintf("Getting details of %s storage class", name))

	sc, err := client.StorageV1().StorageClasses().Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	persistentVolumeList, err := pv.GetStorageClassPersistentVolumes(client, sc.Name, dataselect.DefaultDataSelect)

	storageClass := toStorageClassDetail(sc, persistentVolumeList)
	return &storageClass, err
}

func toStorageClassDetail(storageClass *storage.StorageClass, persistentVolumeList *pv.PersistentVolumeList) StorageClassDetail {
	return StorageClassDetail{
		StorageClass:         toStorageClass(storageClass),
		PersistentVolumeList: *persistentVolumeList,
	}
}
