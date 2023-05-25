package router

import "github.com/gin-gonic/gin"

func RegisterRouter(engine *gin.Engine) {
	// 公开
	openGroup := engine.Group("open")
	registerOpenRouter(openGroup)

	// 私有
	apiGroup := engine.Group("api")
	registerApiRouter(apiGroup)
}

func registerOpenRouter(group *gin.RouterGroup) {
	authGroup := group.Group("auth")
	registerAuthRouter(authGroup)
}

func registerApiRouter(group *gin.RouterGroup) {

}
