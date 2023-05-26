package main

import (
	"github.com/marsxingzhi/goim/apps/interfaces/internal/config"
	"github.com/marsxingzhi/goim/apps/interfaces/internal/server"
)

func main() {
	_ = config.GetConfig()

	s := server.NewServer()
	s.Run()
}
