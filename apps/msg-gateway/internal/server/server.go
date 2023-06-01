package server

import "github.com/marsxingzhi/goim/apps/msg-gateway/internal/server/websocket"

type Server struct {
	wsServer websocket.WsServer
}

func NewServer(wsServer websocket.WsServer) *Server {
	return &Server{wsServer: wsServer}
}

func (s *Server) Run() {
	s.wsServer.Run()
}
