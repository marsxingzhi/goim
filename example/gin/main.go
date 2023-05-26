package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marsxingzhi/goim/pkg/common/xzgin"
	"github.com/marsxingzhi/goim/pkg/common/xzjwt"
	"github.com/marsxingzhi/goim/pkg/e"
)

func main() {
	gs := xzgin.NewGinServer()

	gs.Use(JwtAuth(), a())

	gs.Engine.GET("/ping", Ping())

	gs.Run(8081)
}

func Ping() gin.HandlerFunc {
	fmt.Println("ping...") //
	return func(ctx *gin.Context) {
		ctx.SecureJSON(http.StatusOK, gin.H{
			e.ECODE: 0,
			e.EMSG:  "success",
		})
	}
}

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := xzjwt.ParseFromAuthorization(ctx)
		if err != nil {
			ctx.Abort()
			ctx.JSON(http.StatusBadRequest, gin.H{
				e.ECODE: 0,
				e.EMSG:  "faield to varify token",
			})
			fmt.Println("faield to varify token")
		} else {
			fmt.Println("varify token successfully")
		}
	}
}

func a() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("call a middleware...")
	}
}
