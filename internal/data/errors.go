package data

import (
	"errors"
)

var (
	ErrClusterNotFound = errors.New("未找到集群配置, 请联系运维添加")
)
