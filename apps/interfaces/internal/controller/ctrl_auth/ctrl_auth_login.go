package ctrl_auth

import (
	"github.com/gin-gonic/gin"
	"github.com/marsxingzhi/goim/pkg/e"
	"net/http"
)

func (ac *AuthCtrl) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		e.ECODE: 0,
		e.EMSG:  "访问了 login api",
	})
}
