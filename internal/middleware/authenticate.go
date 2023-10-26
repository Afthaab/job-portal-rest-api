package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/afthaab/job-portal/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (m *Mid) Authenticate(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		traceID, ok := ctx.Value(TraceIDKey).(string)
		if !ok {
			log.Error().Msg("trace id not present in the context")

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": http.StatusText(http.StatusInternalServerError),
			})
			return
		}

		authHeader := c.Request.Header.Get("Authorization")

		strSlice := strings.Split(authHeader, "")
		if len(strSlice) != 2 || strings.ToLower(strSlice[0]) != "bearer" {
			err := errors.New(" expected authorization header format: Bearer <token>")
			log.Error().Err(err).Str("trace id", traceID).Send()
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		claims, err := m.auth.ValidateToken(strSlice[1])
		if err != nil {
			log.Error().Err(err).Str("trace id", traceID).Send()
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": http.StatusText(http.StatusInternalServerError),
			})
			return
		}

		// adding claims to the context
		ctx = context.WithValue(ctx, auth.Key, claims)
		c.Request = c.Request.WithContext(ctx)

		next(c)

	}
}
