package middlewares

import (
	"github.com/Electra-project/electrapay-api/src/helpers"
	"github.com/gin-gonic/gin"
)

func ResponseHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		version := helpers.GetVersion()
		c.Header("X-Version", version)
	}
}
