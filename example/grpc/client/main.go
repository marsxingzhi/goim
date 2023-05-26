package main

import (
	"context"
	"fmt"

	"github.com/marsxingzhi/goim/pkg/proto/pb_auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	req := &pb_auth.RegisterReq{
		Nickname: "zhangsan",
	}

	// 1. 拨号
	conn, err := grpc.Dial("127.0.0.1:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("[client] faield to Dial: ", err)
		return
	}
	defer conn.Close()

	// 2. 创建client
	authClient := pb_auth.NewAuthClient(conn)

	// 3. 客户端调用服务接口函数
	resp, err := authClient.Register(context.Background(), req)
	if err != nil {
		fmt.Println("[client] faield to Register: ", err)
		return
	}
	fmt.Println("[client] receive resp: ", resp)

}
