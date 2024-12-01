package middleware

import (
	"ia-03-backend/pkg/http/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Content-Type", "application/json")

				if commonErr, ok := err.(*resp.ErrorResp); ok {
					ctx.AbortWithStatusJSON(commonErr.StatusCode, commonErr)
					return
				}

				ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp.ErrInternalServer(err.(error)))
			}
		}()

		ctx.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
