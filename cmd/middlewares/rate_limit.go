package middlewares

// Rate limit middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"github.com/mcbryan1/resume-builder-backend/cmd/helpers"
)

func RateLimitMiddleware() gin.HandlerFunc {
	// Create a new rate limiter that allows 1 requests per second
	limiter := ratelimit.NewBucketWithRate(1, 1)

	return func(c *gin.Context) {
		if limiter.TakeAvailable(1) < 1 {
			helpers.RespondWithError(c, 429, "Too many requests", "429")
			c.Abort()
			return
		}

		c.Next()
	}
}
