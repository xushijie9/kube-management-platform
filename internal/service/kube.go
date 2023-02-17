package service

import (
	"context"

	v1 "kube-management-platform/api/kube/v1"
	"kube-management-platform/internal/biz"
)

type KubeService struct {
	v1.UnsafeKubeServiceServer
	uc *biz.KubeUsecase
}

func NewKubeService(uc *biz.KubeUsecase) *KubeService {
	return &KubeService{uc: uc}
}

// SayHello implements kube.GreeterServer.
func (s *KubeService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	deployment, err := s.uc.GetDeployment(context.TODO(), "local")
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + deployment}, nil
}
