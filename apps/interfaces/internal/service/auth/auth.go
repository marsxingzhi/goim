package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/marsxingzhi/goim/apps/interfaces/internal/dto/dto_auth"
)

type AuthService interface {
	Register(ctx *gin.Context, req *dto_auth.RegisterReq)
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (as *authService) Register(ctx *gin.Context, req *dto_auth.RegisterReq) {
	// TODO 这里进行grpc调用，调用auth服务
}
