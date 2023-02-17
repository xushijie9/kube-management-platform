package data

import (
	"k8s.io/client-go/tools/clientcmd"
	"kube-management-platform/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"k8s.io/client-go/kubernetes"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewKubeDeploymentRepo)

// Data .
type Data struct {
	KubeClientSet map[string]*kubernetes.Clientset
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, confKube *conf.Kube, logger log.Logger) (*Data, func(), error) {
	_log := log.NewHelper(logger)
	var kubeClientSetMap map[string]*kubernetes.Clientset
	kubeClientSetMap = make(map[string]*kubernetes.Clientset, 0)

	for _, cluster := range confKube.Clusters {
		client := newKubeClientSet(_log, &cluster.Config, cluster.Config)
		kubeClientSetMap[cluster.ClusterName] = client
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{KubeClientSet: kubeClientSetMap}, cleanup, nil
}

func newKubeClientSet(_log *log.Helper, kubeConfig *string, clusterName string) *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err)
	}

	// 初始化 client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		_log.Fatalf("Failed to generate kube client, clusterName: %s", clusterName)
	}

	return clientSet
}
