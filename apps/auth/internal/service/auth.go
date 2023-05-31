package service

import (
	"context"
	"fmt"
	"github.com/marsxingzhi/goim/pkg/proto/pb_auth"
	"github.com/marsxingzhi/goim/pkg/proto/pb_user"
)

// AuthService 服务
type AuthService interface {
	Register(ctx context.Context, req *pb_auth.RegisterReq) (resp *pb_auth.RegisterResp, err error)
	Login(ctx context.Context, req *pb_auth.LoginReq) (resp *pb_auth.LoginResp, err error)
	Logout(ctx context.Context, req *pb_auth.LogoutReq) (resp *pb_auth.LogoutResp, err error)
	RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (resp *pb_auth.RefreshTokenResp, err error)
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) Register(ctx context.Context, req *pb_auth.RegisterReq) (resp *pb_auth.RegisterResp, err error) {
	resp = new(pb_auth.RegisterResp)
	resp.Msg = "注册成功啦，嘿嘿"
	resp.Code = 0
	resp.UserInfo = new(pb_user.UserInfo)

	fmt.Printf("[测试] register req: %+v\n", req)
	return
}

func (s *authService) Login(ctx context.Context, req *pb_auth.LoginReq) (resp *pb_auth.LoginResp, err error) {
	return
}

func (s *authService) Logout(ctx context.Context, req *pb_auth.LogoutReq) (resp *pb_auth.LogoutResp, err error) {
	return
}
func (s *authService) RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (resp *pb_auth.RefreshTokenResp, err error) {
	return
}
