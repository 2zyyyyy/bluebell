package routers

import (
	"net/http"
	"webapp-scaffold/logger"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	// 使用自定义的中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	return r
}
