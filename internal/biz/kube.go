package biz

import (
	v1 "kube-management-platform/api/kube/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type KubeUsecase struct {
	deploymentRepo KubeDeploymentRepo
	log            *log.Helper
}

func NewKubeUsecase(repo KubeDeploymentRepo, logger log.Logger) *KubeUsecase {
	return &KubeUsecase{deploymentRepo: repo, log: log.NewHelper(logger)}
}
