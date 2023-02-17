package biz

import "context"

type KubeDeploymentRepo interface {
	ListDeployments(ctx context.Context, clusterName string) (string, error)
}

func (kube *KubeUsecase) GetDeployment(ctx context.Context, clusterName string) (string, error) {
	return kube.deploymentRepo.ListDeployments(ctx, clusterName)
}
