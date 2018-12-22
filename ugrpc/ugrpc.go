package ugrpc

import (
	"context"
	"time"

	"github.com/rerost/unstable/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	common.Init()
}

func UnstableUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		cfg := common.Config
		if cfg.Interval != 0 && uint64(time.Now().Second())%cfg.Interval == 0 {
			if cfg.SlowResponseOption.GetEnable() {
				time.Sleep(time.Duration(cfg.SlowResponseOption.GetTime()))
			}
			if cfg.ServerErrorOption.GetEnable() {
				return nil, status.Error(codes.Internal, "Fail by unstable")
			}
		}

		return handler(ctx, req)
	}
}
