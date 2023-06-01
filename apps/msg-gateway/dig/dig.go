package dig

import (
	conf "github.com/marsxingzhi/goim/apps/msg-gateway/internal/config"
	"github.com/marsxingzhi/goim/apps/msg-gateway/internal/server"
	"github.com/marsxingzhi/goim/apps/msg-gateway/internal/server/websocket"
)

func FakeInit() *server.Server {

	config := conf.GetConfig()
	wsServer := websocket.NewWsServer(config)
	return server.NewServer(wsServer)
}
