package xzgrpc

import (
	"context"
	"github.com/marsxingzhi/goim/pkg/common/xzetcd"
	"github.com/marsxingzhi/goim/pkg/config"
	"google.golang.org/grpc"
	"log"
)

/**
grpc server的封装
*/
func NewServer(c *config.Grpc) *grpc.Server {
	opts := make([]grpc.ServerOption, 0)
	opts = append(opts, grpc.UnaryInterceptor(unaryInterceptor()))
	grpcServer := grpc.NewServer(opts...)
	return grpcServer
}

func GetClientConn(opt *config.GrpcDialOption) *grpc.ClientConn {
	return xzetcd.GetClientConn(opt)
}

func unaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Printf("call %s\n", info.FullMethod)
		resp, err = handler(ctx, req)
		return resp, err
	}
}
