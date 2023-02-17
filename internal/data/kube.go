package data

import (
	"context"
	"fmt"
	"kube-management-platform/internal/biz"

	"github.com/go-kratos/kratos/v2/log"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type kubeDeploymentRepo struct {
	data *Data
	log  *log.Helper
}

func NewKubeDeploymentRepo(data *Data, logger log.Logger) biz.KubeDeploymentRepo {
	return &kubeDeploymentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (kube *kubeDeploymentRepo) ListDeployments(ctx context.Context, clusterName string) (string, error) {
	clientset, ok := kube.data.KubeClientSet[clusterName]
	if !ok {
		kube.log.Errorf("clusterName: %s, 无法操作此集群, 请联系运维配置", clusterName)
		return "", ErrClusterNotFound
	}
	list, err := clientset.AppsV1().Deployments("").List(ctx, meta_v1.ListOptions{})
	if err != nil {
		kube.log.Errorf("clusterName: %s, 获取deployment失败: %v", clusterName, err)
		return "", err
	}
	fmt.Println(len(list.Items))
	return "world", nil
}
