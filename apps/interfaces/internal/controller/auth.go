package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marsxingzhi/goim/pkg/e"
)

func SignIn(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		e.ECODE: 0,
		e.EMSG:  "访问了 auth api",
	})
}
