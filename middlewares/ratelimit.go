package middlewares

import (
	"bluebell/controllers"
	"net/http"
	"time"

	"github.com/juju/ratelimit"

	"github.com/gin-gonic/gin"
)

// RateLimitMiddleware 令牌桶中间件
func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就返回响应
		if bucket.TakeAvailable(1) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code":    controllers.CodeRateLimit,
				"message": controllers.CodeRateLimit.Msg(),
			})
			c.Abort()
			return
		}
		// 取到令牌则放行
		c.Next()
	}
}
