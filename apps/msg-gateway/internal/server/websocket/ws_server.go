package websocket

import (
	"github.com/marsxingzhi/goim/apps/msg-gateway/internal/config"
	"github.com/marsxingzhi/goim/pkg/common/xzgin"
	"github.com/marsxingzhi/goim/pkg/middleware"
)

type WsServer interface {
	Run()
}

type wsServer struct {
	hub  *Hub
	conf *config.Config
	// 由于会使用到中间件，这里直接使用gin，不用原生的http
	ginServer *xzgin.GinServer
}

func NewWsServer(conf *config.Config) WsServer {
	ws := &wsServer{hub: newHub(), conf: conf, ginServer: xzgin.NewGinServer()}
	// 中间件需要加在路由的前面
	ws.ginServer.Use(middleware.JwtAuth())
	ws.addRouter()
	return ws
}

func (ws *wsServer) Run() {
	go ws.hub.run()

	ws.ginServer.Run(ws.conf.WsServer.Port)
	//http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	//	serveWs(ws.hub, w, r)
	//})

	//addr := fmt.Sprintf("0.0.0.0:%d", ws.conf.WsServer.Port)
	//err := http.ListenAndServe(addr, nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
}

func (ws *wsServer) addRouter() {
	ws.ginServer.Engine.GET("ws", ws.hub.Update)
}
