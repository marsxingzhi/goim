package auth

import (
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
