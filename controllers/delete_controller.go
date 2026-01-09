package controllers

import (
	"github.com/basitkhan32/crud-api-go-gin/connections"
	"github.com/basitkhan32/crud-api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func DeleteUserController(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	result := connections.DB.Delete(&user, id)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User deleted successfully with id:" + id,
		"user":    user,
	})
}
