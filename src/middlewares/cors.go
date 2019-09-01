package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CORS set the Access-Control headers.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Accept, Accept-Encoding, Origin,Authorization, Content-Length, Content-Type")
		c.Header("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}
