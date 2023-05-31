package dig

import (
	"github.com/marsxingzhi/goim/apps/auth/internal/config"
	"github.com/marsxingzhi/goim/apps/auth/internal/server"
	"github.com/marsxingzhi/goim/apps/auth/internal/server/auth"
	"github.com/marsxingzhi/goim/apps/auth/internal/service"
	"github.com/marsxingzhi/goim/pkg/common/xzmysql"
)

// TODO-xz 这里不使用dig，先手动创建
func Init() *server.Server {

	// auth服务的config
	authConfig := config.GetConfig()

	xzmysql.NewMysqlClient(authConfig.Mysql)

	authService := service.NewAuthService()
	authServer := auth.NewAuthServer(authService, authConfig)
	return server.NewServer(authServer)
}
