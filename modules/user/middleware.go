package user

import "github.com/gin-gonic/gin"

func IsXAPIProvided() gin.HandlerFunc {
	return func(c *gin.Context) {
		xAPIKey := c.GetHeader("X-API-KEY")
		if xAPIKey == "" {
			c.JSON(401, gin.H{
				"error": "X-API-KEY header is required",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
