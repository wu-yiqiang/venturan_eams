package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
	"venturan/app/common/response"
	"venturan/global/serviceErrors"
)

// 接口限流
func RateLimitMiddleware(fillInterval time.Duration, cap, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			response.Fail(c, serviceErrors.FrequentRequests)
			c.Abort()
			return
		}
		c.Next()
	}
}
