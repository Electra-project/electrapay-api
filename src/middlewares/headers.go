package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ruannelloyd/electrapay-api/src/helpers"
)

func ResponseHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		version := helpers.GetVersion()
		c.Header("X-Version", version)
	}
}
