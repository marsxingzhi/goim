package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marsxingzhi/goim/pkg/common/xzjwt"
	"github.com/marsxingzhi/goim/pkg/common/xzredis"
	"github.com/marsxingzhi/goim/pkg/constant"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. 获取token，检查uid、platform等信息
		// 2. 从redis中获取sessionId，与token中的sessionId进行比较
		// 3. 将uid和platform设置到ctx中

		token, err := xzjwt.ParseFromAuthorization(ctx)
		if err != nil {
			ctx.Abort()
			fmt.Println("[jwtAuth] failed to ParseFromAuthorization: ", err)
			return
		}
		if token.Uid <= 0 {
			ctx.Abort()
			fmt.Println("[jwtAuth] valid uid")
			return
		}
		if token.Platform < 0 {
			ctx.Abort()
			fmt.Println("[jwtAuth] valid platform")
			return
		}
		// 有效期结束后，会自动删除
		if token.SessionID == "" {
			ctx.Abort()
			fmt.Println("[jwtAuth] valid sessionId")
			return
		}
		key := constant.REDIS_KEY_USER_ACCESS_TOKEN_SESSION_ID + strconv.FormatInt(token.Uid, 10) + ":" + strconv.Itoa(int(token.Platform))
		val, err := xzredis.Get(key)
		if err != nil {
			if err == redis.Nil {
				fmt.Println("[jwtAuth] sessionId in cache is nil")
			}
			fmt.Println("[jwtAuth] failed to get sessionId from cache: ", err)
			ctx.Abort()
			return
		}
		if token.SessionID != val {
			fmt.Println("[jwtAuth] sessionId of jwtToken not equals sessionId of cache")
			ctx.Abort()
			return
		}
		ctx.Set(constant.USER_UID, token.Uid)
		ctx.Set(constant.USER_PLATFORM, token.Platform)
	}
}
