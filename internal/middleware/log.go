package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type key string

var TraceID key

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuidStr := uuid.NewString

		ctx := c.Request.Context()

		ctx = context.WithValue(ctx, TraceID, uuidStr)

	}
}
