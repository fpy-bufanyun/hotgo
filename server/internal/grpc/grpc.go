package grpc

import (
	"context"
	"github.com/douyu/jupiter/pkg/client/grpc/balancer"
	"github.com/douyu/jupiter/pkg/client/grpc/balancer/p2c"
	"github.com/gogf/gf/v2/os/gctx"

	grpcclient "github.com/douyu/jupiter/pkg/client/grpc"
	"github.com/douyu/jupiter/pkg/xlog"
	helloworldv1 "github.com/douyu/proto/gen/go/api/helloworld/v1"
	hwv1 "github.com/douyu/proto/gen/go/api/hw/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Example struct {
	cc grpc.ClientConnInterface
}

func NewExample() ExampleInterface {
	config := grpcclient.StdConfig("hw")
	config.BalancerName = balancer.NameSmoothWeightRoundRobin
	config.BalancerName = p2c.Name
	return &Example{
		cc: config.MustSingleton(),
	}
}

func (s *Example) SayHello(ctx context.Context, req *helloworldv1.SayHelloRequest) (*helloworldv1.SayHelloResponse, error) {

	cc := helloworldv1.NewGreeterServiceClient(s.cc)

	res, err := cc.SayHello(ctx, req)
	if err != nil {
		xlog.L(ctx).Error("sayHello failed", zap.Error(err), zap.Any("res", res), zap.Any("req", req))
		return nil, err
	}

	return res, nil
}

func (s *Example) HwSayHello(ctx context.Context, req *hwv1.SayHelloRequest) (*hwv1.SayHelloResponse, error) {

	cc := hwv1.NewGreeterServiceClient(s.cc)

	res, err := cc.SayHello(ctx, req)
	if err != nil {
		xlog.L(ctx).Error("sayHello failed", zap.Error(err), zap.Any("res", res), zap.Any("req", req))
		return nil, err
	}

	return res, nil
}

// grpc client测试grpc server demo
func test() error {
	ctx := gctx.New()

	example := NewExample()

	resp, err := example.SayHello(ctx, &helloworldv1.SayHelloRequest{
		Name: "bob",
	})
	if err != nil {
		xlog.L(ctx).Error("ExampleGrpc.SayHello failed", zap.Error(err), zap.Any("res", resp), zap.Any("req", nil))
	}
	return nil
}
