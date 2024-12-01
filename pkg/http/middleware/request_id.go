package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type requestIdKeyType string

const requestIdKey requestIdKeyType = "X-Request-ID"

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(
			context.WithValue(c.Request.Context(),
				requestIdKey, uuid.New().String()))

		c.Next()
	}
}
