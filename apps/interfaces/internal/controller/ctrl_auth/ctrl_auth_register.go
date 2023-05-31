package ctrl_auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marsxingzhi/goim/apps/interfaces/internal/dto/dto_auth"
	"github.com/marsxingzhi/goim/pkg/e"
	"net/http"
)

// ctrl pkg 这里作为请求的入口和出口

func (ac *AuthCtrl) Register(ctx *gin.Context) {
	var req dto_auth.RegisterReq

	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println("[controller] failed to bind params: ", err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			e.ECODE: e.ERROR_CODE_REGISTER_PARAMS,
			e.EMSG:  e.ERROR_MSG_REGISTER_PARAMS,
		})

		return
	}

	response := ac.authService.Register(ctx, &req)
	if response == nil {
		// 调用grpc出错了
		ctx.JSON(http.StatusOK, gin.H{
			e.ECODE: e.ERROR_CODE_REGISTER,
			e.EMSG:  e.ERROR_MSG_REGISTER,
		})
		return
	}
	if response.Code > 0 {
		// 错误
		ctx.JSON(http.StatusOK, gin.H{
			e.ECODE: response.Code,
			e.EMSG:  response.Msg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		e.ECODE: response.Code,
		e.EMSG:  response.Msg,
		e.DATA:  response.Data,
	})

	fmt.Printf("[controller] bind success, and req: %+v\n", req)

}
