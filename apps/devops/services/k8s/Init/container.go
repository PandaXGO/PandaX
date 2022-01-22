package Init

import (
	"errors"
	"fmt"
	"pandax/base/global"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// GetK8sClient 获取k8s Client
func GetK8sClient(k8sConf string) (*kubernetes.Clientset, error) {

	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(k8sConf))
	// skips the validity check for the server's certificate. This will make your HTTPS connections insecure.
	// config.TLSClientConfig.Insecure = true
	if err != nil {
		global.Log.Error("KubeConfig内容错误", zap.Any("err", err))
		return nil, errors.New("KubeConfig内容错误")
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		global.Log.Error("创建Client失败", zap.Any("err", err))
		return nil, errors.New("创建Client失败！")
	}
	return clientSet, nil
}

// GetRestConf 获取k8s RESTConfig
func GetRestConf(k8sConf string) (restConf *rest.Config, err error) {

	if restConf, err = clientcmd.RESTConfigFromKubeConfig([]byte(k8sConf)); err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	return restConf, nil
}
