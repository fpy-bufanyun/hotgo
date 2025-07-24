package hw_service

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/douyu/jupiter/pkg/core/metric"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
	"github.com/douyu/jupiter/pkg/util/xerror"
	"github.com/douyu/jupiter/pkg/xlog"
	hwv1 "github.com/douyu/proto/gen/go/api/hw/v1"
	"go.uber.org/zap"
	"hotgo/internal/grpc"
)

// Options wireservice
type Options struct {
	ExampleGrpc grpc.ExampleInterface
}

type Miss struct {
	Options
}

// NewMissService
func NewMissService(options Options) *Miss {
	return &Miss{
		Options: options,
	}
}

func RegisterMissService(grpcServe *xgrpc.Server) {
	exampleInterface := grpc.NewExample()
	options := Options{
		ExampleGrpc: exampleInterface,
	}

	miss := NewMissService(options)

	hwv1.RegisterGreeterServiceServer(grpcServe, miss)

}

func (s *Miss) SayHello(ctx context.Context, req *hwv1.SayHelloRequest) (*hwv1.SayHelloResponse, error) {
	xlog.L(ctx).Info("SayHello started", zap.String("name", req.GetName()))

	if req.GetName() == "" {
		return &hwv1.SayHelloResponse{
			Error: uint32(hwv1.XERROR_ERROR_NAME_EMPTY.GetEcode()),
			Msg:   hwv1.XERROR_ERROR_NAME_EMPTY.GetMsg(),
		}, nil
	}

	err := req.Validate()
	if err != nil {
		return &hwv1.SayHelloResponse{
			Error: uint32(xerror.InvalidArgument.GetEcode()),
			Msg:   err.Error(),
		}, nil
	}

	resp := &hwv1.SayHelloResponse_Data{
		Name: "hello " + req.GetName(),
	}

	if req.Name != "done" {
		resp, err := s.ExampleGrpc.HwSayHello(ctx, &hwv1.SayHelloRequest{
			Name: "done",
		})
		if err != nil {
			xlog.L(ctx).Error("ExampleGrpc.SayHello failed", zap.Error(err), zap.Any("res", resp), zap.Any("req", req))
			// return nil, err
		}
	}

	metric.CustomizedHandleCounter.WithLabelValues("SayHello").Inc()

	return &hwv1.SayHelloResponse{Data: resp}, nil
}

func (s *Miss) SayHi(ctx context.Context, req *hwv1.SayHiRequest) (*hwv1.SayHiResponse, error) {
	err := req.Validate()
	if err != nil {
		return &hwv1.SayHiResponse{
			Error: uint32(xerror.InvalidArgument.GetEcode()),
			Msg:   err.Error(),
		}, nil
	}

	metric.CustomizedHandleCounter.WithLabelValues("SayHi").Inc()

	return &hwv1.SayHiResponse{}, nil
}

func (s *Miss) ProcessConsumer(ctx context.Context, msg *primitive.MessageExt) error {
	metric.CustomizedHandleCounter.WithLabelValues("ProcessConsumer").Inc()

	return nil
}
