package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marsxingzhi/goim/apps/interfaces/internal/dto/dto_auth"
	"github.com/marsxingzhi/goim/pkg/common/xzhttp"
	"github.com/marsxingzhi/goim/pkg/proto/pb_auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// rpc调用：1. 拨号； 2. 创建client；3. 调用方法
func (as *authService) Register(ctx *gin.Context, req *dto_auth.RegisterReq) *xzhttp.Response {
	// TODO 这里进行grpc调用，调用auth服务

	// 1. 构建rpc调用的入参
	// 2. 调用rpc方法

	// 将dto_auth.RegisterReq -> pb_auth.RegisterReq
	var transformer func(req *dto_auth.RegisterReq) *pb_auth.RegisterReq

	transformer = func(req *dto_auth.RegisterReq) *pb_auth.RegisterReq {
		var r = new(pb_auth.RegisterReq)
		r.Platform = req.Platform
		r.Nickname = req.Nickname
		r.Password = req.Password
		r.Firstname = req.Firstname
		r.Lastname = req.Lastname
		r.Gender = req.Gender
		r.Email = req.Email
		r.Mobile = req.Mobile
		return r
	}

	conn, err := grpc.Dial("127.0.0.1:6600", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// 这里不做返回
		//ctx.JSON(http.StatusOK, gin.H{
		//	e.ECODE: e.ERROR_CODE_AUTH_CONN,
		//	e.EMSG:  e.ERROR_MSG_AUTH_CONN,
		//})
		fmt.Println("[service-auth] failed to dial: ", err.Error())
		return nil
	}
	defer conn.Close()
	authClient := pb_auth.NewAuthClient(conn)

	resp, err := authClient.Register(ctx, transformer(req))
	if err != nil {
		//ctx.JSON(http.StatusOK, gin.H{
		//	e.ECODE: resp.Code,
		//	e.EMSG:  resp.Msg,
		//})
		fmt.Println("[service-auth] failed to register: ", err.Error())
		return nil
	}

	//ctx.JSON(http.StatusOK, gin.H{
	//	e.ECODE: 0,
	//	e.EMSG:  "请求成功",
	//	e.DATA:  resp,
	//})

	var response = new(xzhttp.Response)
	response.Data = resp.UserInfo
	response.Code = resp.Code
	response.Msg = resp.Msg

	return response
}
