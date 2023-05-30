package server

type AuthServer interface {
}

type authServer struct {
}

func NewAuthServer() AuthServer {
	return &authServer{}
}
