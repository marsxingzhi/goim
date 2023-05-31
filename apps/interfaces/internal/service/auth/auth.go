package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/marsxingzhi/goim/apps/interfaces/internal/dto/dto_auth"
	"github.com/marsxingzhi/goim/pkg/common/xzhttp"
)

type AuthService interface {
	Register(ctx *gin.Context, req *dto_auth.RegisterReq) *xzhttp.Response
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}
