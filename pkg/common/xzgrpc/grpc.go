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
	//opts := make([]grpc.ServerOption, 0)
	//opts = append(opts, grpc.UnaryInterceptor(unaryInterceptor()))
	//
	//keepParams := grpc.KeepaliveParams(keepalive.ServerParameters{
	//	MaxConnectionIdle:     time.Duration(c.MaxConnectionIdle) * time.Millisecond,
	//	MaxConnectionAge:      time.Duration(c.MaxConnectionAge) * time.Millisecond,
	//	MaxConnectionAgeGrace: time.Duration(c.MaxConnectionAgeGrace) * time.Millisecond,
	//	Time:                  time.Duration(c.Time) * time.Millisecond,
	//	Timeout:               time.Duration(c.Timeout) * time.Millisecond,
	//})
	//opts = append(opts, keepParams)
	//if c.StreamsLimit > 0 {
	//	// 一个连接中最大并发Stream数
	//	opts = append(opts, grpc.MaxConcurrentStreams(c.StreamsLimit))
	//}
	//if c.MaxRecvMsgSize > 0 {
	//	// 允许接收的最大消息长度
	//	opts = append(opts, grpc.MaxRecvMsgSize(c.MaxRecvMsgSize))
	//}

	grpcServer := grpc.NewServer()
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
