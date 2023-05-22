package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marsingzhi/goim/pkg/common/xzgin"
)

func main() {
	gs := xzgin.NewGinServer()
	gs.Engine.GET("/ping", func(ctx *gin.Context) {
		ctx.SecureJSON(http.StatusOK, gin.H{
			"ecode": 0,
			"emsg":  "success",
		})
	})

	gs.Run(8081)
}
