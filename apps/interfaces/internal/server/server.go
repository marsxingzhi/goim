package server

import (
	"github.com/marsxingzhi/goim/apps/interfaces/internal/config"
	"github.com/marsxingzhi/goim/apps/interfaces/internal/router"
	"github.com/marsxingzhi/goim/pkg/common/xzgin"
)

type server struct {
	ginServer *xzgin.GinServer
	conf      *config.Config
}

func NewServer() *server {
	s := &server{
		ginServer: xzgin.NewGinServer(),
		conf:      config.GetConfig(),
	}
	return s
}

func (s *server) Run() {
	router.RegisterRouter(s.ginServer.Engine)
	s.ginServer.Run(s.conf.Port)
}
