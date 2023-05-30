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
		Platform:  1,
		Nickname:  "行知",
		Password:  "123456",
		Firstname: "mars",
		Lastname:  "xingzhi",
		Gender:    1,
		Email:     "2277087813@qq.com",
		Mobile:    "123456789",
	}

	// 1. 拨号
	conn, err := grpc.Dial("127.0.0.1:6600", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("[client] faield to Dial auth grpc: ", err)
		return
	}
	defer conn.Close()

	// 2. 创建client
	authClient := pb_auth.NewAuthClient(conn)

	// 3. 客户端调用服务接口函数
	resp, err := authClient.Register(context.Background(), req)
	if err != nil {
		fmt.Println("[client] faield to Register auth grpc: ", err.Error())
		return
	}
	fmt.Println("[client] receive resp msg: ", resp.Msg)

}
