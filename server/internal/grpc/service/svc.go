package hw_service

import (
	"github.com/douyu/jupiter/pkg/server/xgrpc"
)

// TODO 如果有新增service，在这里添加执行注册grpc server service逻辑
func Register(grpcServe *xgrpc.Server) {
	RegisterMissService(grpcServe)
}
