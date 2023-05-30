package controller

import (
	"fmt"
	"github.com/marsxingzhi/goim/apps/interfaces/internal/dto/dto_auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marsxingzhi/goim/pkg/e"
)

func Register(ctx *gin.Context) {
	var req dto_auth.RegisterReq

	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println("[controller] failed to bind params: ", err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			e.ECODE: e.REGISTER_PARAMS_ERROR,
			e.EMSG:  "入参错误",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		e.ECODE: 0,
		e.DATA:  req,
	})

	fmt.Printf("[controller] bind success, and req: %+v\n", req)

}

func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		e.ECODE: 0,
		e.EMSG:  "访问了 login api",
	})
}
