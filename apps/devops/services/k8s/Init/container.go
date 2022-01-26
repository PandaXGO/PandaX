package Init

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"k8s.io/client-go/dynamic"
	"os"
	"pandax/base/global"
	"path/filepath"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// GetK8sClient 获取k8s Client
func GetK8sClient(ctx context.Context, kubeConfig string) (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error
	if kubeConfig != "" {
		clientConfig := GetClientConfig(kubeConfig, nil)
		config, err = clientConfig.ClientConfig()
		if err != nil {
			return nil, errors.New("KubeConfig内容错误")
		}
	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, errors.New("KubeConfig内容错误")
		}
	}

	//config, err := clientcmd.RESTConfigFromKubeConfig([]byte(k8sConf))
	// skips the validity check for the server's certificate. This will make your HTTPS connections insecure.
	// config.TLSClientConfig.Insecure = true
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		global.Log.Error("创建Client失败", zap.Any("err", err))
		return nil, errors.New("创建Client失败！")
	}
	_, err = dynamic.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("kubernetes dynamic client create error:%s", err.Error())
	}
	return clientSet, nil
}

func GetClientConfig(kubeConfig string, reader io.Reader) clientcmd.ClientConfig {
	loadingRules := GetLoadingRules(kubeConfig)
	overrides := &clientcmd.ConfigOverrides{ClusterDefaults: clientcmd.ClusterDefaults}
	return clientcmd.NewInteractiveDeferredLoadingClientConfig(loadingRules, overrides, reader)
}

func GetLoadingRules(kubeConfig string) *clientcmd.ClientConfigLoadingRules {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.DefaultClientConfig = &clientcmd.DefaultClientConfig
	if kubeConfig != "" {
		loadingRules.ExplicitPath = kubeConfig
	}

	var otherFiles []string
	homeDir, err := os.UserHomeDir()
	if err == nil {
		otherFiles = append(otherFiles, filepath.Join(homeDir, ".kube", "k3s.yaml"))
	}
	otherFiles = append(otherFiles, "/etc/rancher/k3s/k3s.yaml")
	loadingRules.Precedence = append(loadingRules.Precedence, canRead(otherFiles)...)

	return loadingRules
}

func canRead(files []string) (result []string) {
	for _, f := range files {
		_, err := ioutil.ReadFile(f)
		if err == nil {
			result = append(result, f)
		}
	}
	return
}
