package auth

import (
	"context"
	"github.com/marsxingzhi/goim/pkg/proto/pb_auth"
)

// 这里不写具体的业务实现
func (s *authServer) Register(ctx context.Context, req *pb_auth.RegisterReq) (resp *pb_auth.RegisterResp, err error) {
	return s.authService.Register(ctx, req)
}

func (s *authServer) Login(ctx context.Context, req *pb_auth.LoginReq) (resp *pb_auth.LoginResp, err error) {
	return s.authService.Login(ctx, req)
}

func (s *authServer) Logout(ctx context.Context, req *pb_auth.LogoutReq) (resp *pb_auth.LogoutResp, err error) {
	return s.authService.Logout(ctx, req)
}
func (s *authServer) RefreshToken(ctx context.Context, req *pb_auth.RefreshTokenReq) (resp *pb_auth.RefreshTokenResp, err error) {
	return s.authService.RefreshToken(ctx, req)
}
