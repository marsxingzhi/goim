package metadata

import (
	"github.com/gin-gonic/gin"
)

// GetString 后续扩展到context.Context
// TODO 参考一下真实项目
func GetInt64(ctx *gin.Context, key string) int64 {
	if value, ok := ctx.Get(key); ok {
		return value.(int64)
	}
	return 0
}

func GetInt8(ctx *gin.Context, key string) int8 {
	if value, ok := ctx.Get(key); ok {
		return value.(int8)
	}
	return 0
}
