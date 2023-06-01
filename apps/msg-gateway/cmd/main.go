package main

import (
	"github.com/marsxingzhi/goim/apps/msg-gateway/dig"
)

func main() {

	server := dig.FakeInit()
	server.Run()

	select {}
}
