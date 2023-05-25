package main

import (
	"github.com/marsingzhi/goim/apps/interfaces/internal/config"
	"github.com/marsingzhi/goim/apps/interfaces/internal/server"
)

func main() {
	_ = config.GetConfig()

	s := server.NewServer()
	s.Run()
}
