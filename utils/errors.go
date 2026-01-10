package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrServerCrash(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return true
	}
	return false
}

func ErrBadRequest(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return true
	}
	return false
}
