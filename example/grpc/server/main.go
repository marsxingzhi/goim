package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/marsxingzhi/goim/pkg/proto/pb_auth"
	"google.golang.org/grpc"
)

type AuthServer struct {
	pb_auth.UnimplementedAuthServer
}

func NewAuthServer() *AuthServer {
	return &AuthServer{}
}

func (as *AuthServer) Register(ctx context.Context, req *pb_auth.RegisterReq) (resp *pb_auth.RegisterResp, err error) {
	fmt.Println("call Register Server")
	resp = new(pb_auth.RegisterResp)
	resp.Msg = "register成功"
	return
}

func main() {
	// 1. 监听端口
	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Println("[server] failed to Listen: ", err)
		return
	}
	// 2. 创建grpc服务，并注册
	s := grpc.NewServer()
	authServer := NewAuthServer()
	pb_auth.RegisterAuthServer(s, authServer)

	// 3. 启动服务
	log.Fatal(s.Serve(listener))

}
