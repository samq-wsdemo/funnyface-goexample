// filename transaction.go
// author Yoshiya Ito <myon53@gmail.com>
// 2019-03-18
package middleware

import (
	"github.com/funnyface-inc/go-example/db"
	"github.com/gin-gonic/gin"
)

func Transaction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := db.DB.Begin()
		defer func() {
			if 400 <= ctx.Writer.Status() {
				db.Rollback()
				return
			}
			db.Commit()
		}()
		ctx.Set("db", db)
		ctx.Next()
	}
}
