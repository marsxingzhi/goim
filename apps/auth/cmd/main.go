package main

import (
	"github.com/marsxingzhi/goim/apps/auth/internal/config"
	"github.com/marsxingzhi/goim/apps/auth/internal/server"
	"github.com/marsxingzhi/goim/apps/auth/internal/server/auth"
	"github.com/marsxingzhi/goim/apps/auth/internal/service"
	"github.com/marsxingzhi/goim/pkg/common/xzmysql"
	"github.com/marsxingzhi/goim/pkg/common/xzredis"
)

func main() {
	server := Init()
	server.Run()

	select {}
}

// TODO 后续可以改成依赖注入
func Init() *server.Server {
	// auth服务的config
	authConfig := config.GetConfig()

	xzmysql.NewMysqlClient(authConfig.Mysql)
	xzredis.NewRedisClient(authConfig.Redis)

	authService := service.NewAuthService(authConfig)
	authServer := auth.NewAuthServer(authService, authConfig)
	return server.NewServer(authServer)
}
