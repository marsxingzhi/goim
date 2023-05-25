package server

import (
	"github.com/marsingzhi/goim/apps/interfaces/internal/config"
	"github.com/marsingzhi/goim/apps/interfaces/internal/router"
	"github.com/marsingzhi/goim/pkg/common/xzgin"
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
