package router

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/marsxingzhi/goim/apps/interfaces/internal/controller/ctrl_auth"
	"github.com/marsxingzhi/goim/apps/interfaces/internal/service/auth"
)

// auth相关的接口
func registerAuthRouter(group *gin.RouterGroup) {
	authService := auth.NewAuthService()
	authCtrl := ctrl.NewAuthCtrl(authService)

	group.POST("register", authCtrl.Register)
	group.POST("login", authCtrl.Login)
}
