package main

import (
	"github.com/marsxingzhi/goim/apps/auth/dig"
)

func main() {
	//config.GetConfig()

	server := dig.Init()
	server.Run()

	select {}
}
