package middlewares

import "github.com/gin-gonic/gin"

func ResponseHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Version", "1.0")
	}
}
