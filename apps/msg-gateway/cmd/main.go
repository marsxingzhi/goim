package main

import (
	conf "github.com/marsxingzhi/goim/apps/msg-gateway/internal/config"
	"github.com/marsxingzhi/goim/apps/msg-gateway/internal/server"
	"github.com/marsxingzhi/goim/apps/msg-gateway/internal/server/websocket"
	"github.com/marsxingzhi/goim/pkg/common/xzredis"
)

func main() {

	server := FakeInit()
	server.Run()

	select {}
}

func FakeInit() *server.Server {

	config := conf.GetConfig()

	xzredis.NewRedisClient(config.Redis)

	wsServer := websocket.NewWsServer(config)
	return server.NewServer(wsServer)
}
