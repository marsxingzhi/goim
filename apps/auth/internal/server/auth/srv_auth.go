package auth

import (
	"context"
	"github.com/marsxingzhi/goim/apps/auth/internal/config"
	"github.com/marsxingzhi/goim/apps/auth/internal/service"
	"github.com/marsxingzhi/goim/pkg/common/xzgrpc"
	"github.com/marsxingzhi/goim/pkg/proto/pb_auth"
)

type AuthServer interface {
	Run()
}

type authServer struct {
	pb_auth.UnimplementedAuthServer
	authService service.AuthService
	conf        *config.Config
}

func NewAuthServer(authService service.AuthService, conf *config.Config) AuthServer {
	return &authServer{authService: authService, conf: conf}
}

func (s *authServer) Run() {
	server := xzgrpc.NewServer(s.conf.GrpcServer)
	pb_auth.RegisterAuthServer(server, s)

	// 启动
	grpcServer := xzgrpc.NewGrpcServer(s.conf.GrpcServer, s.conf.Etcd)
	grpcServer.Run(server)

}

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
