// filename errors.go
// author Yoshiya Ito <myon53@gmail.com>
// 2019-03-18
package middleware

import (
	"github.com/gin-gonic/gin"
)

func Handle500() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if 500 <= ctx.Writer.Status() {
			ctx.JSON(500, gin.H{
				"status":  500,
				"message": "internal server error",
			})
		}
	}
}
