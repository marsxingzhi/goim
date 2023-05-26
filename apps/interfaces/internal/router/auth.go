package router

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/marsxingzhi/goim/apps/interfaces/internal/controller"
)

func registerAuthRouter(group *gin.RouterGroup) {
	group.POST("sign_in", ctrl.SignIn)
}
