// filename redis_pooling.go
// author Yoshiya Ito <myon53@gmail.com>
// 2019-03-18
package middleware

import (
	"github.com/funnyface-inc/go-example/db"
	"github.com/gin-gonic/gin"
)

func RedisPooling() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		redis := db.Redis.Get()
		ctx.Set("redis", redis)
		defer redis.Close()
		ctx.Next()
	}
}
