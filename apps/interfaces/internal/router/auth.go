package router

import (
	"github.com/gin-gonic/gin"
	ctrl "github.com/marsxingzhi/goim/apps/interfaces/internal/controller"
)

func addRegisterRouter(group *gin.RouterGroup) {
	group.POST("register", ctrl.Register)
}

func addLoginRouter(group *gin.RouterGroup) {
	group.POST("login", ctrl.Login)
}
